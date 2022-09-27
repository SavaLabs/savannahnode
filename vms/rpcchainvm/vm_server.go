// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/SavaLabs/savannahnode/api/keystore/gkeystore"
	"github.com/SavaLabs/savannahnode/api/metrics"
	"github.com/SavaLabs/savannahnode/chains/atomic/gsharedmemory"
	"github.com/SavaLabs/savannahnode/database/corruptabledb"
	"github.com/SavaLabs/savannahnode/database/manager"
	"github.com/SavaLabs/savannahnode/database/rpcdb"
	"github.com/SavaLabs/savannahnode/ids"
	"github.com/SavaLabs/savannahnode/ids/galiasreader"
	"github.com/SavaLabs/savannahnode/snow"
	"github.com/SavaLabs/savannahnode/snow/engine/common"
	"github.com/SavaLabs/savannahnode/snow/engine/common/appsender"
	"github.com/SavaLabs/savannahnode/snow/engine/snowman/block"
	"github.com/SavaLabs/savannahnode/utils/logging"
	"github.com/SavaLabs/savannahnode/utils/wrappers"
	"github.com/SavaLabs/savannahnode/version"
	"github.com/SavaLabs/savannahnode/vms/rpcchainvm/ghttp"
	"github.com/SavaLabs/savannahnode/vms/rpcchainvm/grpcutils"
	"github.com/SavaLabs/savannahnode/vms/rpcchainvm/gsubnetlookup"
	"github.com/SavaLabs/savannahnode/vms/rpcchainvm/messenger"

	aliasreaderpb "github.com/SavaLabs/savannahnode/proto/pb/aliasreader"
	appsenderpb "github.com/SavaLabs/savannahnode/proto/pb/appsender"
	httppb "github.com/SavaLabs/savannahnode/proto/pb/http"
	keystorepb "github.com/SavaLabs/savannahnode/proto/pb/keystore"
	messengerpb "github.com/SavaLabs/savannahnode/proto/pb/messenger"
	rpcdbpb "github.com/SavaLabs/savannahnode/proto/pb/rpcdb"
	sharedmemorypb "github.com/SavaLabs/savannahnode/proto/pb/sharedmemory"
	subnetlookuppb "github.com/SavaLabs/savannahnode/proto/pb/subnetlookup"
	vmpb "github.com/SavaLabs/savannahnode/proto/pb/vm"
)

var _ vmpb.VMServer = &VMServer{}

// VMServer is a VM that is managed over RPC.
type VMServer struct {
	vmpb.UnsafeVMServer

	vm   block.ChainVM
	hVM  block.HeightIndexedChainVM
	ssVM block.StateSyncableVM

	processMetrics prometheus.Gatherer
	dbManager      manager.Manager

	serverCloser grpcutils.ServerCloser
	connCloser   wrappers.Closer

	ctx    *snow.Context
	closed chan struct{}
}

// NewServer returns a vm instance connected to a remote vm instance
func NewServer(vm block.ChainVM) *VMServer {
	hVM, _ := vm.(block.HeightIndexedChainVM)
	ssVM, _ := vm.(block.StateSyncableVM)
	return &VMServer{
		vm:   vm,
		hVM:  hVM,
		ssVM: ssVM,
	}
}

func (vm *VMServer) Initialize(_ context.Context, req *vmpb.InitializeRequest) (*vmpb.InitializeResponse, error) {
	subnetID, err := ids.ToID(req.SubnetId)
	if err != nil {
		return nil, err
	}
	chainID, err := ids.ToID(req.ChainId)
	if err != nil {
		return nil, err
	}
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}
	xChainID, err := ids.ToID(req.XChainId)
	if err != nil {
		return nil, err
	}
	avaxAssetID, err := ids.ToID(req.AvaxAssetId)
	if err != nil {
		return nil, err
	}

	registerer := prometheus.NewRegistry()

	// Current state of process metrics
	processCollector := collectors.NewProcessCollector(collectors.ProcessCollectorOpts{})
	if err := registerer.Register(processCollector); err != nil {
		return nil, err
	}

	// Go process metrics using debug.GCStats
	goCollector := collectors.NewGoCollector()
	if err := registerer.Register(goCollector); err != nil {
		return nil, err
	}

	// gRPC client metrics
	grpcClientMetrics := grpc_prometheus.NewClientMetrics()
	if err := registerer.Register(grpcClientMetrics); err != nil {
		return nil, err
	}

	// Register metrics for each Go plugin processes
	vm.processMetrics = registerer

	// Dial each database in the request and construct the database manager
	versionedDBs := make([]*manager.VersionedDatabase, len(req.DbServers))
	for i, vDBReq := range req.DbServers {
		version, err := version.Parse(vDBReq.Version)
		if err != nil {
			// Ignore closing errors to return the original error
			_ = vm.connCloser.Close()
			return nil, err
		}

		clientConn, err := grpcutils.Dial(vDBReq.ServerAddr, grpcutils.DialOptsWithMetrics(grpcClientMetrics)...)
		if err != nil {
			// Ignore closing errors to return the original error
			_ = vm.connCloser.Close()
			return nil, err
		}
		vm.connCloser.Add(clientConn)
		db := rpcdb.NewClient(rpcdbpb.NewDatabaseClient(clientConn))
		versionedDBs[i] = &manager.VersionedDatabase{
			Database: corruptabledb.New(db),
			Version:  version,
		}
	}
	dbManager, err := manager.NewManagerFromDBs(versionedDBs)
	if err != nil {
		// Ignore closing errors to return the original error
		_ = vm.connCloser.Close()
		return nil, err
	}
	vm.dbManager = dbManager

	clientConn, err := grpcutils.Dial(req.ServerAddr, grpcutils.DialOptsWithMetrics(grpcClientMetrics)...)
	if err != nil {
		// Ignore closing errors to return the original error
		_ = vm.connCloser.Close()
		return nil, err
	}

	vm.connCloser.Add(clientConn)

	msgClient := messenger.NewClient(messengerpb.NewMessengerClient(clientConn))
	keystoreClient := gkeystore.NewClient(keystorepb.NewKeystoreClient(clientConn))
	sharedMemoryClient := gsharedmemory.NewClient(sharedmemorypb.NewSharedMemoryClient(clientConn))
	bcLookupClient := galiasreader.NewClient(aliasreaderpb.NewAliasReaderClient(clientConn))
	snLookupClient := gsubnetlookup.NewClient(subnetlookuppb.NewSubnetLookupClient(clientConn))
	appSenderClient := appsender.NewClient(appsenderpb.NewAppSenderClient(clientConn))

	toEngine := make(chan common.Message, 1)
	vm.closed = make(chan struct{})
	go func() {
		for {
			select {
			case msg, ok := <-toEngine:
				if !ok {
					return
				}
				// Nothing to do with the error within the goroutine
				_ = msgClient.Notify(msg)
			case <-vm.closed:
				return
			}
		}
	}()

	vm.ctx = &snow.Context{
		NetworkID: req.NetworkId,
		SubnetID:  subnetID,
		ChainID:   chainID,
		NodeID:    nodeID,

		XChainID:    xChainID,
		AVAXAssetID: avaxAssetID,

		Log:          logging.NoLog{},
		Keystore:     keystoreClient,
		SharedMemory: sharedMemoryClient,
		BCLookup:     bcLookupClient,
		SNLookup:     snLookupClient,
		Metrics:      metrics.NewOptionalGatherer(),

		// TODO: support snowman++ fields
	}

	if err := vm.vm.Initialize(vm.ctx, dbManager, req.GenesisBytes, req.UpgradeBytes, req.ConfigBytes, toEngine, nil, appSenderClient); err != nil {
		// Ignore errors closing resources to return the original error
		_ = vm.connCloser.Close()
		close(vm.closed)
		return nil, err
	}

	lastAccepted, err := vm.vm.LastAccepted()
	if err != nil {
		// Ignore errors closing resources to return the original error
		_ = vm.vm.Shutdown()
		_ = vm.connCloser.Close()
		close(vm.closed)
		return nil, err
	}

	blk, err := vm.vm.GetBlock(lastAccepted)
	if err != nil {
		// Ignore errors closing resources to return the original error
		_ = vm.vm.Shutdown()
		_ = vm.connCloser.Close()
		close(vm.closed)
		return nil, err
	}
	parentID := blk.Parent()
	return &vmpb.InitializeResponse{
		LastAcceptedId:       lastAccepted[:],
		LastAcceptedParentId: parentID[:],
		Height:               blk.Height(),
		Bytes:                blk.Bytes(),
		Timestamp:            grpcutils.TimestampFromTime(blk.Timestamp()),
	}, nil
}

func (vm *VMServer) SetState(_ context.Context, stateReq *vmpb.SetStateRequest) (*vmpb.SetStateResponse, error) {
	err := vm.vm.SetState(snow.State(stateReq.State))
	if err != nil {
		return nil, err
	}

	lastAccepted, err := vm.vm.LastAccepted()
	if err != nil {
		return nil, err
	}

	blk, err := vm.vm.GetBlock(lastAccepted)
	if err != nil {
		return nil, err
	}

	parentID := blk.Parent()
	return &vmpb.SetStateResponse{
		LastAcceptedId:       lastAccepted[:],
		LastAcceptedParentId: parentID[:],
		Height:               blk.Height(),
		Bytes:                blk.Bytes(),
		Timestamp:            grpcutils.TimestampFromTime(blk.Timestamp()),
	}, nil
}

func (vm *VMServer) Shutdown(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	if vm.closed == nil {
		return &emptypb.Empty{}, nil
	}
	errs := wrappers.Errs{}
	errs.Add(vm.vm.Shutdown())
	close(vm.closed)
	vm.serverCloser.Stop()
	errs.Add(vm.connCloser.Close())
	return &emptypb.Empty{}, errs.Err
}

func (vm *VMServer) CreateHandlers(context.Context, *emptypb.Empty) (*vmpb.CreateHandlersResponse, error) {
	handlers, err := vm.vm.CreateHandlers()
	if err != nil {
		return nil, err
	}
	resp := &vmpb.CreateHandlersResponse{}
	for prefix, h := range handlers {
		handler := h

		serverListener, err := grpcutils.NewListener()
		if err != nil {
			return nil, err
		}
		serverAddr := serverListener.Addr().String()

		// Start the gRPC server which serves the HTTP service
		go grpcutils.Serve(serverListener, func(opts []grpc.ServerOption) *grpc.Server {
			if len(opts) == 0 {
				opts = append(opts, grpcutils.DefaultServerOptions...)
			}
			server := grpc.NewServer(opts...)
			vm.serverCloser.Add(server)
			httppb.RegisterHTTPServer(server, ghttp.NewServer(handler.Handler))
			return server
		})

		resp.Handlers = append(resp.Handlers, &vmpb.Handler{
			Prefix:      prefix,
			LockOptions: uint32(handler.LockOptions),
			ServerAddr:  serverAddr,
		})
	}
	return resp, nil
}

func (vm *VMServer) CreateStaticHandlers(context.Context, *emptypb.Empty) (*vmpb.CreateStaticHandlersResponse, error) {
	handlers, err := vm.vm.CreateStaticHandlers()
	if err != nil {
		return nil, err
	}
	resp := &vmpb.CreateStaticHandlersResponse{}
	for prefix, h := range handlers {
		handler := h

		serverListener, err := grpcutils.NewListener()
		if err != nil {
			return nil, err
		}
		serverAddr := serverListener.Addr().String()

		// Start the gRPC server which serves the HTTP service
		go grpcutils.Serve(serverListener, func(opts []grpc.ServerOption) *grpc.Server {
			if len(opts) == 0 {
				opts = append(opts, grpcutils.DefaultServerOptions...)
			}
			server := grpc.NewServer(opts...)
			vm.serverCloser.Add(server)
			httppb.RegisterHTTPServer(server, ghttp.NewServer(handler.Handler))
			return server
		})

		resp.Handlers = append(resp.Handlers, &vmpb.Handler{
			Prefix:      prefix,
			LockOptions: uint32(handler.LockOptions),
			ServerAddr:  serverAddr,
		})
	}
	return resp, nil
}

func (vm *VMServer) Connected(_ context.Context, req *vmpb.ConnectedRequest) (*emptypb.Empty, error) {
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}

	peerVersion, err := version.ParseApplication(req.Version)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, vm.vm.Connected(nodeID, peerVersion)
}

func (vm *VMServer) Disconnected(_ context.Context, req *vmpb.DisconnectedRequest) (*emptypb.Empty, error) {
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, vm.vm.Disconnected(nodeID)
}

func (vm *VMServer) BuildBlock(context.Context, *emptypb.Empty) (*vmpb.BuildBlockResponse, error) {
	blk, err := vm.vm.BuildBlock()
	if err != nil {
		return nil, err
	}
	blkID := blk.ID()
	parentID := blk.Parent()
	return &vmpb.BuildBlockResponse{
		Id:        blkID[:],
		ParentId:  parentID[:],
		Bytes:     blk.Bytes(),
		Height:    blk.Height(),
		Timestamp: grpcutils.TimestampFromTime(blk.Timestamp()),
	}, nil
}

func (vm *VMServer) ParseBlock(_ context.Context, req *vmpb.ParseBlockRequest) (*vmpb.ParseBlockResponse, error) {
	blk, err := vm.vm.ParseBlock(req.Bytes)
	if err != nil {
		return nil, err
	}
	blkID := blk.ID()
	parentID := blk.Parent()
	return &vmpb.ParseBlockResponse{
		Id:        blkID[:],
		ParentId:  parentID[:],
		Status:    uint32(blk.Status()),
		Height:    blk.Height(),
		Timestamp: grpcutils.TimestampFromTime(blk.Timestamp()),
	}, nil
}

func (vm *VMServer) GetBlock(_ context.Context, req *vmpb.GetBlockRequest) (*vmpb.GetBlockResponse, error) {
	id, err := ids.ToID(req.Id)
	if err != nil {
		return nil, err
	}
	blk, err := vm.vm.GetBlock(id)
	if err != nil {
		return &vmpb.GetBlockResponse{
			Err: errorToErrCode[err],
		}, errorToRPCError(err)
	}

	parentID := blk.Parent()
	return &vmpb.GetBlockResponse{
		ParentId:  parentID[:],
		Bytes:     blk.Bytes(),
		Status:    uint32(blk.Status()),
		Height:    blk.Height(),
		Timestamp: grpcutils.TimestampFromTime(blk.Timestamp()),
	}, nil
}

func (vm *VMServer) SetPreference(_ context.Context, req *vmpb.SetPreferenceRequest) (*emptypb.Empty, error) {
	id, err := ids.ToID(req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, vm.vm.SetPreference(id)
}

func (vm *VMServer) Health(ctx context.Context, req *emptypb.Empty) (*vmpb.HealthResponse, error) {
	vmHealth, err := vm.vm.HealthCheck()
	if err != nil {
		return &vmpb.HealthResponse{}, err
	}
	dbHealth, err := vm.dbHealthChecks()
	if err != nil {
		return &vmpb.HealthResponse{}, err
	}
	report := map[string]interface{}{
		"database": dbHealth,
		"health":   vmHealth,
	}

	details, err := json.Marshal(report)
	return &vmpb.HealthResponse{
		Details: details,
	}, err
}

func (vm *VMServer) dbHealthChecks() (interface{}, error) {
	details := make(map[string]interface{}, len(vm.dbManager.GetDatabases()))

	// Check Database health
	for _, client := range vm.dbManager.GetDatabases() {
		// Shared gRPC client don't close
		health, err := client.Database.HealthCheck()
		if err != nil {
			return nil, fmt.Errorf("failed to check db health %q: %w", client.Version.String(), err)
		}
		details[client.Version.String()] = health
	}

	return details, nil
}

func (vm *VMServer) Version(context.Context, *emptypb.Empty) (*vmpb.VersionResponse, error) {
	version, err := vm.vm.Version()
	return &vmpb.VersionResponse{
		Version: version,
	}, err
}

func (vm *VMServer) AppRequest(_ context.Context, req *vmpb.AppRequestMsg) (*emptypb.Empty, error) {
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}
	deadline, err := grpcutils.TimestampAsTime(req.Deadline)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, vm.vm.AppRequest(nodeID, req.RequestId, deadline, req.Request)
}

func (vm *VMServer) AppRequestFailed(_ context.Context, req *vmpb.AppRequestFailedMsg) (*emptypb.Empty, error) {
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, vm.vm.AppRequestFailed(nodeID, req.RequestId)
}

func (vm *VMServer) AppResponse(_ context.Context, req *vmpb.AppResponseMsg) (*emptypb.Empty, error) {
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, vm.vm.AppResponse(nodeID, req.RequestId, req.Response)
}

func (vm *VMServer) AppGossip(_ context.Context, req *vmpb.AppGossipMsg) (*emptypb.Empty, error) {
	nodeID, err := ids.ToNodeID(req.NodeId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, vm.vm.AppGossip(nodeID, req.Msg)
}

func (vm *VMServer) Gather(context.Context, *emptypb.Empty) (*vmpb.GatherResponse, error) {
	// Gather metrics registered to snow context Gatherer. These
	// metrics are defined by the underlying vm implementation.
	mfs, err := vm.ctx.Metrics.Gather()
	if err != nil {
		return nil, err
	}

	// Gather metrics registered by rpcchainvm server Gatherer. These
	// metrics are collected for each Go plugin process.
	pluginMetrics, err := vm.processMetrics.Gather()
	if err != nil {
		return nil, err
	}
	mfs = append(mfs, pluginMetrics...)

	return &vmpb.GatherResponse{MetricFamilies: mfs}, err
}

func (vm *VMServer) GetAncestors(_ context.Context, req *vmpb.GetAncestorsRequest) (*vmpb.GetAncestorsResponse, error) {
	blkID, err := ids.ToID(req.BlkId)
	if err != nil {
		return nil, err
	}
	maxBlksNum := int(req.MaxBlocksNum)
	maxBlksSize := int(req.MaxBlocksSize)
	maxBlocksRetrivalTime := time.Duration(req.MaxBlocksRetrivalTime)

	blocks, err := block.GetAncestors(
		vm.vm,
		blkID,
		maxBlksNum,
		maxBlksSize,
		maxBlocksRetrivalTime,
	)
	return &vmpb.GetAncestorsResponse{
		BlksBytes: blocks,
	}, err
}

func (vm *VMServer) BatchedParseBlock(
	ctx context.Context,
	req *vmpb.BatchedParseBlockRequest,
) (*vmpb.BatchedParseBlockResponse, error) {
	blocks := make([]*vmpb.ParseBlockResponse, len(req.Request))
	for i, blockBytes := range req.Request {
		block, err := vm.ParseBlock(ctx, &vmpb.ParseBlockRequest{
			Bytes: blockBytes,
		})
		if err != nil {
			return nil, err
		}
		blocks[i] = block
	}
	return &vmpb.BatchedParseBlockResponse{
		Response: blocks,
	}, nil
}

func (vm *VMServer) VerifyHeightIndex(context.Context, *emptypb.Empty) (*vmpb.VerifyHeightIndexResponse, error) {
	var err error
	if vm.hVM != nil {
		err = vm.hVM.VerifyHeightIndex()
	} else {
		err = block.ErrHeightIndexedVMNotImplemented
	}

	return &vmpb.VerifyHeightIndexResponse{
		Err: errorToErrCode[err],
	}, errorToRPCError(err)
}

func (vm *VMServer) GetBlockIDAtHeight(ctx context.Context, req *vmpb.GetBlockIDAtHeightRequest) (*vmpb.GetBlockIDAtHeightResponse, error) {
	var (
		blkID ids.ID
		err   error
	)
	if vm.hVM != nil {
		blkID, err = vm.hVM.GetBlockIDAtHeight(req.Height)
	} else {
		err = block.ErrHeightIndexedVMNotImplemented
	}

	return &vmpb.GetBlockIDAtHeightResponse{
		BlkId: blkID[:],
		Err:   errorToErrCode[err],
	}, errorToRPCError(err)
}

func (vm *VMServer) StateSyncEnabled(context.Context, *emptypb.Empty) (*vmpb.StateSyncEnabledResponse, error) {
	var (
		enabled bool
		err     error
	)
	if vm.ssVM != nil {
		enabled, err = vm.ssVM.StateSyncEnabled()
	}

	return &vmpb.StateSyncEnabledResponse{
		Enabled: enabled,
		Err:     errorToErrCode[err],
	}, errorToRPCError(err)
}

func (vm *VMServer) GetOngoingSyncStateSummary(
	context.Context,
	*emptypb.Empty,
) (*vmpb.GetOngoingSyncStateSummaryResponse, error) {
	var (
		summary block.StateSummary
		err     error
	)
	if vm.ssVM != nil {
		summary, err = vm.ssVM.GetOngoingSyncStateSummary()
	} else {
		err = block.ErrStateSyncableVMNotImplemented
	}

	if err != nil {
		return &vmpb.GetOngoingSyncStateSummaryResponse{
			Err: errorToErrCode[err],
		}, errorToRPCError(err)
	}

	summaryID := summary.ID()
	return &vmpb.GetOngoingSyncStateSummaryResponse{
		Id:     summaryID[:],
		Height: summary.Height(),
		Bytes:  summary.Bytes(),
	}, nil
}

func (vm *VMServer) GetLastStateSummary(
	ctx context.Context,
	empty *emptypb.Empty,
) (*vmpb.GetLastStateSummaryResponse, error) {
	var (
		summary block.StateSummary
		err     error
	)
	if vm.ssVM != nil {
		summary, err = vm.ssVM.GetLastStateSummary()
	} else {
		err = block.ErrStateSyncableVMNotImplemented
	}

	if err != nil {
		return &vmpb.GetLastStateSummaryResponse{
			Err: errorToErrCode[err],
		}, errorToRPCError(err)
	}

	summaryID := summary.ID()
	return &vmpb.GetLastStateSummaryResponse{
		Id:     summaryID[:],
		Height: summary.Height(),
		Bytes:  summary.Bytes(),
	}, nil
}

func (vm *VMServer) ParseStateSummary(
	ctx context.Context,
	req *vmpb.ParseStateSummaryRequest,
) (*vmpb.ParseStateSummaryResponse, error) {
	var (
		summary block.StateSummary
		err     error
	)
	if vm.ssVM != nil {
		summary, err = vm.ssVM.ParseStateSummary(req.Bytes)
	} else {
		err = block.ErrStateSyncableVMNotImplemented
	}

	if err != nil {
		return &vmpb.ParseStateSummaryResponse{
			Err: errorToErrCode[err],
		}, errorToRPCError(err)
	}

	summaryID := summary.ID()
	return &vmpb.ParseStateSummaryResponse{
		Id:     summaryID[:],
		Height: summary.Height(),
	}, nil
}

func (vm *VMServer) GetStateSummary(
	ctx context.Context,
	req *vmpb.GetStateSummaryRequest,
) (*vmpb.GetStateSummaryResponse, error) {
	var (
		summary block.StateSummary
		err     error
	)
	if vm.ssVM != nil {
		summary, err = vm.ssVM.GetStateSummary(req.Height)
	} else {
		err = block.ErrStateSyncableVMNotImplemented
	}

	if err != nil {
		return &vmpb.GetStateSummaryResponse{
			Err: errorToErrCode[err],
		}, errorToRPCError(err)
	}

	summaryID := summary.ID()
	return &vmpb.GetStateSummaryResponse{
		Id:    summaryID[:],
		Bytes: summary.Bytes(),
	}, nil
}

func (vm *VMServer) BlockVerify(_ context.Context, req *vmpb.BlockVerifyRequest) (*vmpb.BlockVerifyResponse, error) {
	blk, err := vm.vm.ParseBlock(req.Bytes)
	if err != nil {
		return nil, err
	}
	if err := blk.Verify(); err != nil {
		return nil, err
	}
	return &vmpb.BlockVerifyResponse{
		Timestamp: grpcutils.TimestampFromTime(blk.Timestamp()),
	}, nil
}

func (vm *VMServer) BlockAccept(_ context.Context, req *vmpb.BlockAcceptRequest) (*emptypb.Empty, error) {
	id, err := ids.ToID(req.Id)
	if err != nil {
		return nil, err
	}
	blk, err := vm.vm.GetBlock(id)
	if err != nil {
		return nil, err
	}
	if err := blk.Accept(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (vm *VMServer) BlockReject(_ context.Context, req *vmpb.BlockRejectRequest) (*emptypb.Empty, error) {
	id, err := ids.ToID(req.Id)
	if err != nil {
		return nil, err
	}
	blk, err := vm.vm.GetBlock(id)
	if err != nil {
		return nil, err
	}
	if err := blk.Reject(); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (vm *VMServer) StateSummaryAccept(
	_ context.Context,
	req *vmpb.StateSummaryAcceptRequest,
) (*vmpb.StateSummaryAcceptResponse, error) {
	var (
		accepted bool
		err      error
	)
	if vm.ssVM != nil {
		var summary block.StateSummary
		summary, err = vm.ssVM.ParseStateSummary(req.Bytes)
		if err == nil {
			accepted, err = summary.Accept()
		}
	} else {
		err = block.ErrStateSyncableVMNotImplemented
	}

	return &vmpb.StateSummaryAcceptResponse{
		Accepted: accepted,
		Err:      errorToErrCode[err],
	}, errorToRPCError(err)
}
