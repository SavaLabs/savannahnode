// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"github.com/SavaLabs/savannahnode/ids"
	"github.com/SavaLabs/savannahnode/snow"
	"github.com/SavaLabs/savannahnode/vms"
)

var (
	_ vms.Factory = &Factory{}

	// ID that this Fx uses when labeled
	ID = ids.ID{'p', 'r', 'o', 'p', 'e', 'r', 't', 'y', 'f', 'x'}
)

type Factory struct{}

func (f *Factory) New(*snow.Context) (interface{}, error) { return &Fx{}, nil }
