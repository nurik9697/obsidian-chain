package keeper

import (
	"obsidian/x/notary/types"
)

var _ types.QueryServer = Keeper{}
