// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"github.com/SavaLabs/savannahnode/ids"
	"github.com/SavaLabs/savannahnode/snow"
	"github.com/SavaLabs/savannahnode/vms/components/avax"
	"github.com/SavaLabs/savannahnode/vms/secp256k1fx"
)

// UnsignedTx is an unsigned transaction
type UnsignedTx interface {
	// TODO: Remove this initialization pattern from both the platformvm and the
	// avm.
	snow.ContextInitializable
	secp256k1fx.UnsignedTx
	Initialize(unsignedBytes []byte)

	// InputIDs returns the set of inputs this transaction consumes
	InputIDs() ids.Set

	Outputs() []*avax.TransferableOutput

	// Attempts to verify this transaction without any provided state.
	SyntacticVerify(ctx *snow.Context) error

	// Visit calls [visitor] with this transaction's concrete type
	Visit(visitor Visitor) error
}
