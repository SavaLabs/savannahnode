// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/SavaLabs/savannahnode/snow"
	"github.com/SavaLabs/savannahnode/snow/consensus/snowball"
	"github.com/SavaLabs/savannahnode/snow/consensus/snowman"
	"github.com/SavaLabs/savannahnode/snow/engine/common"
	"github.com/SavaLabs/savannahnode/snow/engine/snowman/block"
	"github.com/SavaLabs/savannahnode/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx        *snow.ConsensusContext
	VM         block.ChainVM
	Sender     common.Sender
	Validators validators.Set
	Params     snowball.Parameters
	Consensus  snowman.Consensus
}
