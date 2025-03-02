// Code generated by MockGen. DO NOT EDIT.
// Source: snow/engine/common/sender.go

package common

import (
	reflect "reflect"

	ids "github.com/ava-labs/avalanchego/ids"
	snow "github.com/ava-labs/avalanchego/snow"
	gomock "github.com/golang/mock/gomock"
)

// MockSender is a mock of Sender interface.
type MockSender struct {
	ctrl     *gomock.Controller
	recorder *MockSenderMockRecorder
}

// MockSenderMockRecorder is the mock recorder for MockSender.
type MockSenderMockRecorder struct {
	mock *MockSender
}

// NewMockSender creates a new mock instance.
func NewMockSender(ctrl *gomock.Controller) *MockSender {
	mock := &MockSender{ctrl: ctrl}
	mock.recorder = &MockSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSender) EXPECT() *MockSenderMockRecorder {
	return m.recorder
}

// Accept mocks base method.
func (m *MockSender) Accept(ctx *snow.ConsensusContext, containerID ids.ID, container []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept", ctx, containerID, container)
	ret0, _ := ret[0].(error)
	return ret0
}

// Accept indicates an expected call of Accept.
func (mr *MockSenderMockRecorder) Accept(ctx, containerID, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockSender)(nil).Accept), ctx, containerID, container)
}

// SendAccepted mocks base method.
func (m *MockSender) SendAccepted(nodeID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAccepted", nodeID, requestID, containerIDs)
}

// SendAccepted indicates an expected call of SendAccepted.
func (mr *MockSenderMockRecorder) SendAccepted(nodeID, requestID, containerIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAccepted", reflect.TypeOf((*MockSender)(nil).SendAccepted), nodeID, requestID, containerIDs)
}

// SendAcceptedFrontier mocks base method.
func (m *MockSender) SendAcceptedFrontier(nodeID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAcceptedFrontier", nodeID, requestID, containerIDs)
}

// SendAcceptedFrontier indicates an expected call of SendAcceptedFrontier.
func (mr *MockSenderMockRecorder) SendAcceptedFrontier(nodeID, requestID, containerIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAcceptedFrontier", reflect.TypeOf((*MockSender)(nil).SendAcceptedFrontier), nodeID, requestID, containerIDs)
}

// SendAcceptedStateSummary mocks base method.
func (m *MockSender) SendAcceptedStateSummary(nodeID ids.NodeID, requestID uint32, summaryIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAcceptedStateSummary", nodeID, requestID, summaryIDs)
}

// SendAcceptedStateSummary indicates an expected call of SendAcceptedStateSummary.
func (mr *MockSenderMockRecorder) SendAcceptedStateSummary(nodeID, requestID, summaryIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAcceptedStateSummary", reflect.TypeOf((*MockSender)(nil).SendAcceptedStateSummary), nodeID, requestID, summaryIDs)
}

// SendAncestors mocks base method.
func (m *MockSender) SendAncestors(nodeID ids.NodeID, requestID uint32, containers [][]byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAncestors", nodeID, requestID, containers)
}

// SendAncestors indicates an expected call of SendAncestors.
func (mr *MockSenderMockRecorder) SendAncestors(nodeID, requestID, containers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAncestors", reflect.TypeOf((*MockSender)(nil).SendAncestors), nodeID, requestID, containers)
}

// SendAppGossip mocks base method.
func (m *MockSender) SendAppGossip(appGossipBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppGossip", appGossipBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppGossip indicates an expected call of SendAppGossip.
func (mr *MockSenderMockRecorder) SendAppGossip(appGossipBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppGossip", reflect.TypeOf((*MockSender)(nil).SendAppGossip), appGossipBytes)
}

// SendAppGossipSpecific mocks base method.
func (m *MockSender) SendAppGossipSpecific(nodeIDs ids.NodeIDSet, appGossipBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppGossipSpecific", nodeIDs, appGossipBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppGossipSpecific indicates an expected call of SendAppGossipSpecific.
func (mr *MockSenderMockRecorder) SendAppGossipSpecific(nodeIDs, appGossipBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppGossipSpecific", reflect.TypeOf((*MockSender)(nil).SendAppGossipSpecific), nodeIDs, appGossipBytes)
}

// SendAppRequest mocks base method.
func (m *MockSender) SendAppRequest(nodeIDs ids.NodeIDSet, requestID uint32, appRequestBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppRequest", nodeIDs, requestID, appRequestBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppRequest indicates an expected call of SendAppRequest.
func (mr *MockSenderMockRecorder) SendAppRequest(nodeIDs, requestID, appRequestBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppRequest", reflect.TypeOf((*MockSender)(nil).SendAppRequest), nodeIDs, requestID, appRequestBytes)
}

// SendAppResponse mocks base method.
func (m *MockSender) SendAppResponse(nodeID ids.NodeID, requestID uint32, appResponseBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppResponse", nodeID, requestID, appResponseBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppResponse indicates an expected call of SendAppResponse.
func (mr *MockSenderMockRecorder) SendAppResponse(nodeID, requestID, appResponseBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppResponse", reflect.TypeOf((*MockSender)(nil).SendAppResponse), nodeID, requestID, appResponseBytes)
}

// SendChits mocks base method.
func (m *MockSender) SendChits(nodeID ids.NodeID, requestID uint32, votes []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendChits", nodeID, requestID, votes)
}

// SendChits indicates an expected call of SendChits.
func (mr *MockSenderMockRecorder) SendChits(nodeID, requestID, votes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChits", reflect.TypeOf((*MockSender)(nil).SendChits), nodeID, requestID, votes)
}

// SendGet mocks base method.
func (m *MockSender) SendGet(nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGet", nodeID, requestID, containerID)
}

// SendGet indicates an expected call of SendGet.
func (mr *MockSenderMockRecorder) SendGet(nodeID, requestID, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGet", reflect.TypeOf((*MockSender)(nil).SendGet), nodeID, requestID, containerID)
}

// SendGetAccepted mocks base method.
func (m *MockSender) SendGetAccepted(nodeIDs ids.NodeIDSet, requestID uint32, containerIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAccepted", nodeIDs, requestID, containerIDs)
}

// SendGetAccepted indicates an expected call of SendGetAccepted.
func (mr *MockSenderMockRecorder) SendGetAccepted(nodeIDs, requestID, containerIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAccepted", reflect.TypeOf((*MockSender)(nil).SendGetAccepted), nodeIDs, requestID, containerIDs)
}

// SendGetAcceptedFrontier mocks base method.
func (m *MockSender) SendGetAcceptedFrontier(nodeIDs ids.NodeIDSet, requestID uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAcceptedFrontier", nodeIDs, requestID)
}

// SendGetAcceptedFrontier indicates an expected call of SendGetAcceptedFrontier.
func (mr *MockSenderMockRecorder) SendGetAcceptedFrontier(nodeIDs, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAcceptedFrontier", reflect.TypeOf((*MockSender)(nil).SendGetAcceptedFrontier), nodeIDs, requestID)
}

// SendGetAcceptedStateSummary mocks base method.
func (m *MockSender) SendGetAcceptedStateSummary(nodeIDs ids.NodeIDSet, requestID uint32, heights []uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAcceptedStateSummary", nodeIDs, requestID, heights)
}

// SendGetAcceptedStateSummary indicates an expected call of SendGetAcceptedStateSummary.
func (mr *MockSenderMockRecorder) SendGetAcceptedStateSummary(nodeIDs, requestID, heights interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAcceptedStateSummary", reflect.TypeOf((*MockSender)(nil).SendGetAcceptedStateSummary), nodeIDs, requestID, heights)
}

// SendGetAncestors mocks base method.
func (m *MockSender) SendGetAncestors(nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAncestors", nodeID, requestID, containerID)
}

// SendGetAncestors indicates an expected call of SendGetAncestors.
func (mr *MockSenderMockRecorder) SendGetAncestors(nodeID, requestID, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAncestors", reflect.TypeOf((*MockSender)(nil).SendGetAncestors), nodeID, requestID, containerID)
}

// SendGetStateSummaryFrontier mocks base method.
func (m *MockSender) SendGetStateSummaryFrontier(nodeIDs ids.NodeIDSet, requestID uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetStateSummaryFrontier", nodeIDs, requestID)
}

// SendGetStateSummaryFrontier indicates an expected call of SendGetStateSummaryFrontier.
func (mr *MockSenderMockRecorder) SendGetStateSummaryFrontier(nodeIDs, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetStateSummaryFrontier", reflect.TypeOf((*MockSender)(nil).SendGetStateSummaryFrontier), nodeIDs, requestID)
}

// SendGossip mocks base method.
func (m *MockSender) SendGossip(container []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGossip", container)
}

// SendGossip indicates an expected call of SendGossip.
func (mr *MockSenderMockRecorder) SendGossip(container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGossip", reflect.TypeOf((*MockSender)(nil).SendGossip), container)
}

// SendPullQuery mocks base method.
func (m *MockSender) SendPullQuery(nodeIDs ids.NodeIDSet, requestID uint32, containerID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPullQuery", nodeIDs, requestID, containerID)
}

// SendPullQuery indicates an expected call of SendPullQuery.
func (mr *MockSenderMockRecorder) SendPullQuery(nodeIDs, requestID, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPullQuery", reflect.TypeOf((*MockSender)(nil).SendPullQuery), nodeIDs, requestID, containerID)
}

// SendPushQuery mocks base method.
func (m *MockSender) SendPushQuery(nodeIDs ids.NodeIDSet, requestID uint32, container []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPushQuery", nodeIDs, requestID, container)
}

// SendPushQuery indicates an expected call of SendPushQuery.
func (mr *MockSenderMockRecorder) SendPushQuery(nodeIDs, requestID, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPushQuery", reflect.TypeOf((*MockSender)(nil).SendPushQuery), nodeIDs, requestID, container)
}

// SendPut mocks base method.
func (m *MockSender) SendPut(nodeID ids.NodeID, requestID uint32, container []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPut", nodeID, requestID, container)
}

// SendPut indicates an expected call of SendPut.
func (mr *MockSenderMockRecorder) SendPut(nodeID, requestID, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPut", reflect.TypeOf((*MockSender)(nil).SendPut), nodeID, requestID, container)
}

// SendStateSummaryFrontier mocks base method.
func (m *MockSender) SendStateSummaryFrontier(nodeID ids.NodeID, requestID uint32, summary []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendStateSummaryFrontier", nodeID, requestID, summary)
}

// SendStateSummaryFrontier indicates an expected call of SendStateSummaryFrontier.
func (mr *MockSenderMockRecorder) SendStateSummaryFrontier(nodeID, requestID, summary interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendStateSummaryFrontier", reflect.TypeOf((*MockSender)(nil).SendStateSummaryFrontier), nodeID, requestID, summary)
}

// MockStateSummarySender is a mock of StateSummarySender interface.
type MockStateSummarySender struct {
	ctrl     *gomock.Controller
	recorder *MockStateSummarySenderMockRecorder
}

// MockStateSummarySenderMockRecorder is the mock recorder for MockStateSummarySender.
type MockStateSummarySenderMockRecorder struct {
	mock *MockStateSummarySender
}

// NewMockStateSummarySender creates a new mock instance.
func NewMockStateSummarySender(ctrl *gomock.Controller) *MockStateSummarySender {
	mock := &MockStateSummarySender{ctrl: ctrl}
	mock.recorder = &MockStateSummarySenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateSummarySender) EXPECT() *MockStateSummarySenderMockRecorder {
	return m.recorder
}

// SendGetStateSummaryFrontier mocks base method.
func (m *MockStateSummarySender) SendGetStateSummaryFrontier(nodeIDs ids.NodeIDSet, requestID uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetStateSummaryFrontier", nodeIDs, requestID)
}

// SendGetStateSummaryFrontier indicates an expected call of SendGetStateSummaryFrontier.
func (mr *MockStateSummarySenderMockRecorder) SendGetStateSummaryFrontier(nodeIDs, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetStateSummaryFrontier", reflect.TypeOf((*MockStateSummarySender)(nil).SendGetStateSummaryFrontier), nodeIDs, requestID)
}

// SendStateSummaryFrontier mocks base method.
func (m *MockStateSummarySender) SendStateSummaryFrontier(nodeID ids.NodeID, requestID uint32, summary []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendStateSummaryFrontier", nodeID, requestID, summary)
}

// SendStateSummaryFrontier indicates an expected call of SendStateSummaryFrontier.
func (mr *MockStateSummarySenderMockRecorder) SendStateSummaryFrontier(nodeID, requestID, summary interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendStateSummaryFrontier", reflect.TypeOf((*MockStateSummarySender)(nil).SendStateSummaryFrontier), nodeID, requestID, summary)
}

// MockAcceptedStateSummarySender is a mock of AcceptedStateSummarySender interface.
type MockAcceptedStateSummarySender struct {
	ctrl     *gomock.Controller
	recorder *MockAcceptedStateSummarySenderMockRecorder
}

// MockAcceptedStateSummarySenderMockRecorder is the mock recorder for MockAcceptedStateSummarySender.
type MockAcceptedStateSummarySenderMockRecorder struct {
	mock *MockAcceptedStateSummarySender
}

// NewMockAcceptedStateSummarySender creates a new mock instance.
func NewMockAcceptedStateSummarySender(ctrl *gomock.Controller) *MockAcceptedStateSummarySender {
	mock := &MockAcceptedStateSummarySender{ctrl: ctrl}
	mock.recorder = &MockAcceptedStateSummarySenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAcceptedStateSummarySender) EXPECT() *MockAcceptedStateSummarySenderMockRecorder {
	return m.recorder
}

// SendAcceptedStateSummary mocks base method.
func (m *MockAcceptedStateSummarySender) SendAcceptedStateSummary(nodeID ids.NodeID, requestID uint32, summaryIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAcceptedStateSummary", nodeID, requestID, summaryIDs)
}

// SendAcceptedStateSummary indicates an expected call of SendAcceptedStateSummary.
func (mr *MockAcceptedStateSummarySenderMockRecorder) SendAcceptedStateSummary(nodeID, requestID, summaryIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAcceptedStateSummary", reflect.TypeOf((*MockAcceptedStateSummarySender)(nil).SendAcceptedStateSummary), nodeID, requestID, summaryIDs)
}

// SendGetAcceptedStateSummary mocks base method.
func (m *MockAcceptedStateSummarySender) SendGetAcceptedStateSummary(nodeIDs ids.NodeIDSet, requestID uint32, heights []uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAcceptedStateSummary", nodeIDs, requestID, heights)
}

// SendGetAcceptedStateSummary indicates an expected call of SendGetAcceptedStateSummary.
func (mr *MockAcceptedStateSummarySenderMockRecorder) SendGetAcceptedStateSummary(nodeIDs, requestID, heights interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAcceptedStateSummary", reflect.TypeOf((*MockAcceptedStateSummarySender)(nil).SendGetAcceptedStateSummary), nodeIDs, requestID, heights)
}

// MockFrontierSender is a mock of FrontierSender interface.
type MockFrontierSender struct {
	ctrl     *gomock.Controller
	recorder *MockFrontierSenderMockRecorder
}

// MockFrontierSenderMockRecorder is the mock recorder for MockFrontierSender.
type MockFrontierSenderMockRecorder struct {
	mock *MockFrontierSender
}

// NewMockFrontierSender creates a new mock instance.
func NewMockFrontierSender(ctrl *gomock.Controller) *MockFrontierSender {
	mock := &MockFrontierSender{ctrl: ctrl}
	mock.recorder = &MockFrontierSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFrontierSender) EXPECT() *MockFrontierSenderMockRecorder {
	return m.recorder
}

// SendAcceptedFrontier mocks base method.
func (m *MockFrontierSender) SendAcceptedFrontier(nodeID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAcceptedFrontier", nodeID, requestID, containerIDs)
}

// SendAcceptedFrontier indicates an expected call of SendAcceptedFrontier.
func (mr *MockFrontierSenderMockRecorder) SendAcceptedFrontier(nodeID, requestID, containerIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAcceptedFrontier", reflect.TypeOf((*MockFrontierSender)(nil).SendAcceptedFrontier), nodeID, requestID, containerIDs)
}

// SendGetAcceptedFrontier mocks base method.
func (m *MockFrontierSender) SendGetAcceptedFrontier(nodeIDs ids.NodeIDSet, requestID uint32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAcceptedFrontier", nodeIDs, requestID)
}

// SendGetAcceptedFrontier indicates an expected call of SendGetAcceptedFrontier.
func (mr *MockFrontierSenderMockRecorder) SendGetAcceptedFrontier(nodeIDs, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAcceptedFrontier", reflect.TypeOf((*MockFrontierSender)(nil).SendGetAcceptedFrontier), nodeIDs, requestID)
}

// MockAcceptedSender is a mock of AcceptedSender interface.
type MockAcceptedSender struct {
	ctrl     *gomock.Controller
	recorder *MockAcceptedSenderMockRecorder
}

// MockAcceptedSenderMockRecorder is the mock recorder for MockAcceptedSender.
type MockAcceptedSenderMockRecorder struct {
	mock *MockAcceptedSender
}

// NewMockAcceptedSender creates a new mock instance.
func NewMockAcceptedSender(ctrl *gomock.Controller) *MockAcceptedSender {
	mock := &MockAcceptedSender{ctrl: ctrl}
	mock.recorder = &MockAcceptedSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAcceptedSender) EXPECT() *MockAcceptedSenderMockRecorder {
	return m.recorder
}

// SendAccepted mocks base method.
func (m *MockAcceptedSender) SendAccepted(nodeID ids.NodeID, requestID uint32, containerIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAccepted", nodeID, requestID, containerIDs)
}

// SendAccepted indicates an expected call of SendAccepted.
func (mr *MockAcceptedSenderMockRecorder) SendAccepted(nodeID, requestID, containerIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAccepted", reflect.TypeOf((*MockAcceptedSender)(nil).SendAccepted), nodeID, requestID, containerIDs)
}

// SendGetAccepted mocks base method.
func (m *MockAcceptedSender) SendGetAccepted(nodeIDs ids.NodeIDSet, requestID uint32, containerIDs []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAccepted", nodeIDs, requestID, containerIDs)
}

// SendGetAccepted indicates an expected call of SendGetAccepted.
func (mr *MockAcceptedSenderMockRecorder) SendGetAccepted(nodeIDs, requestID, containerIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAccepted", reflect.TypeOf((*MockAcceptedSender)(nil).SendGetAccepted), nodeIDs, requestID, containerIDs)
}

// MockFetchSender is a mock of FetchSender interface.
type MockFetchSender struct {
	ctrl     *gomock.Controller
	recorder *MockFetchSenderMockRecorder
}

// MockFetchSenderMockRecorder is the mock recorder for MockFetchSender.
type MockFetchSenderMockRecorder struct {
	mock *MockFetchSender
}

// NewMockFetchSender creates a new mock instance.
func NewMockFetchSender(ctrl *gomock.Controller) *MockFetchSender {
	mock := &MockFetchSender{ctrl: ctrl}
	mock.recorder = &MockFetchSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFetchSender) EXPECT() *MockFetchSenderMockRecorder {
	return m.recorder
}

// SendAncestors mocks base method.
func (m *MockFetchSender) SendAncestors(nodeID ids.NodeID, requestID uint32, containers [][]byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendAncestors", nodeID, requestID, containers)
}

// SendAncestors indicates an expected call of SendAncestors.
func (mr *MockFetchSenderMockRecorder) SendAncestors(nodeID, requestID, containers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAncestors", reflect.TypeOf((*MockFetchSender)(nil).SendAncestors), nodeID, requestID, containers)
}

// SendGet mocks base method.
func (m *MockFetchSender) SendGet(nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGet", nodeID, requestID, containerID)
}

// SendGet indicates an expected call of SendGet.
func (mr *MockFetchSenderMockRecorder) SendGet(nodeID, requestID, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGet", reflect.TypeOf((*MockFetchSender)(nil).SendGet), nodeID, requestID, containerID)
}

// SendGetAncestors mocks base method.
func (m *MockFetchSender) SendGetAncestors(nodeID ids.NodeID, requestID uint32, containerID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGetAncestors", nodeID, requestID, containerID)
}

// SendGetAncestors indicates an expected call of SendGetAncestors.
func (mr *MockFetchSenderMockRecorder) SendGetAncestors(nodeID, requestID, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGetAncestors", reflect.TypeOf((*MockFetchSender)(nil).SendGetAncestors), nodeID, requestID, containerID)
}

// SendPut mocks base method.
func (m *MockFetchSender) SendPut(nodeID ids.NodeID, requestID uint32, container []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPut", nodeID, requestID, container)
}

// SendPut indicates an expected call of SendPut.
func (mr *MockFetchSenderMockRecorder) SendPut(nodeID, requestID, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPut", reflect.TypeOf((*MockFetchSender)(nil).SendPut), nodeID, requestID, container)
}

// MockQuerySender is a mock of QuerySender interface.
type MockQuerySender struct {
	ctrl     *gomock.Controller
	recorder *MockQuerySenderMockRecorder
}

// MockQuerySenderMockRecorder is the mock recorder for MockQuerySender.
type MockQuerySenderMockRecorder struct {
	mock *MockQuerySender
}

// NewMockQuerySender creates a new mock instance.
func NewMockQuerySender(ctrl *gomock.Controller) *MockQuerySender {
	mock := &MockQuerySender{ctrl: ctrl}
	mock.recorder = &MockQuerySenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerySender) EXPECT() *MockQuerySenderMockRecorder {
	return m.recorder
}

// SendChits mocks base method.
func (m *MockQuerySender) SendChits(nodeID ids.NodeID, requestID uint32, votes []ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendChits", nodeID, requestID, votes)
}

// SendChits indicates an expected call of SendChits.
func (mr *MockQuerySenderMockRecorder) SendChits(nodeID, requestID, votes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChits", reflect.TypeOf((*MockQuerySender)(nil).SendChits), nodeID, requestID, votes)
}

// SendPullQuery mocks base method.
func (m *MockQuerySender) SendPullQuery(nodeIDs ids.NodeIDSet, requestID uint32, containerID ids.ID) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPullQuery", nodeIDs, requestID, containerID)
}

// SendPullQuery indicates an expected call of SendPullQuery.
func (mr *MockQuerySenderMockRecorder) SendPullQuery(nodeIDs, requestID, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPullQuery", reflect.TypeOf((*MockQuerySender)(nil).SendPullQuery), nodeIDs, requestID, containerID)
}

// SendPushQuery mocks base method.
func (m *MockQuerySender) SendPushQuery(nodeIDs ids.NodeIDSet, requestID uint32, container []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendPushQuery", nodeIDs, requestID, container)
}

// SendPushQuery indicates an expected call of SendPushQuery.
func (mr *MockQuerySenderMockRecorder) SendPushQuery(nodeIDs, requestID, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPushQuery", reflect.TypeOf((*MockQuerySender)(nil).SendPushQuery), nodeIDs, requestID, container)
}

// MockGossiper is a mock of Gossiper interface.
type MockGossiper struct {
	ctrl     *gomock.Controller
	recorder *MockGossiperMockRecorder
}

// MockGossiperMockRecorder is the mock recorder for MockGossiper.
type MockGossiperMockRecorder struct {
	mock *MockGossiper
}

// NewMockGossiper creates a new mock instance.
func NewMockGossiper(ctrl *gomock.Controller) *MockGossiper {
	mock := &MockGossiper{ctrl: ctrl}
	mock.recorder = &MockGossiperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGossiper) EXPECT() *MockGossiperMockRecorder {
	return m.recorder
}

// SendGossip mocks base method.
func (m *MockGossiper) SendGossip(container []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendGossip", container)
}

// SendGossip indicates an expected call of SendGossip.
func (mr *MockGossiperMockRecorder) SendGossip(container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendGossip", reflect.TypeOf((*MockGossiper)(nil).SendGossip), container)
}

// MockAppSender is a mock of AppSender interface.
type MockAppSender struct {
	ctrl     *gomock.Controller
	recorder *MockAppSenderMockRecorder
}

// MockAppSenderMockRecorder is the mock recorder for MockAppSender.
type MockAppSenderMockRecorder struct {
	mock *MockAppSender
}

// NewMockAppSender creates a new mock instance.
func NewMockAppSender(ctrl *gomock.Controller) *MockAppSender {
	mock := &MockAppSender{ctrl: ctrl}
	mock.recorder = &MockAppSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppSender) EXPECT() *MockAppSenderMockRecorder {
	return m.recorder
}

// SendAppGossip mocks base method.
func (m *MockAppSender) SendAppGossip(appGossipBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppGossip", appGossipBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppGossip indicates an expected call of SendAppGossip.
func (mr *MockAppSenderMockRecorder) SendAppGossip(appGossipBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppGossip", reflect.TypeOf((*MockAppSender)(nil).SendAppGossip), appGossipBytes)
}

// SendAppGossipSpecific mocks base method.
func (m *MockAppSender) SendAppGossipSpecific(nodeIDs ids.NodeIDSet, appGossipBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppGossipSpecific", nodeIDs, appGossipBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppGossipSpecific indicates an expected call of SendAppGossipSpecific.
func (mr *MockAppSenderMockRecorder) SendAppGossipSpecific(nodeIDs, appGossipBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppGossipSpecific", reflect.TypeOf((*MockAppSender)(nil).SendAppGossipSpecific), nodeIDs, appGossipBytes)
}

// SendAppRequest mocks base method.
func (m *MockAppSender) SendAppRequest(nodeIDs ids.NodeIDSet, requestID uint32, appRequestBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppRequest", nodeIDs, requestID, appRequestBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppRequest indicates an expected call of SendAppRequest.
func (mr *MockAppSenderMockRecorder) SendAppRequest(nodeIDs, requestID, appRequestBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppRequest", reflect.TypeOf((*MockAppSender)(nil).SendAppRequest), nodeIDs, requestID, appRequestBytes)
}

// SendAppResponse mocks base method.
func (m *MockAppSender) SendAppResponse(nodeID ids.NodeID, requestID uint32, appResponseBytes []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAppResponse", nodeID, requestID, appResponseBytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAppResponse indicates an expected call of SendAppResponse.
func (mr *MockAppSenderMockRecorder) SendAppResponse(nodeID, requestID, appResponseBytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAppResponse", reflect.TypeOf((*MockAppSender)(nil).SendAppResponse), nodeID, requestID, appResponseBytes)
}
