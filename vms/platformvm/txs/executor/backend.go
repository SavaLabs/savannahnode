// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/SavaLabs/savannahnode/snow"
	"github.com/SavaLabs/savannahnode/snow/uptime"
	"github.com/SavaLabs/savannahnode/utils"
	"github.com/SavaLabs/savannahnode/utils/timer/mockable"
	"github.com/SavaLabs/savannahnode/vms/platformvm/config"
	"github.com/SavaLabs/savannahnode/vms/platformvm/fx"
	"github.com/SavaLabs/savannahnode/vms/platformvm/reward"
	"github.com/SavaLabs/savannahnode/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Manager
	Rewards      reward.Calculator
	Bootstrapped *utils.AtomicBool
}
