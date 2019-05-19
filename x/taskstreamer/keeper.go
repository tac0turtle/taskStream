package taskstreamer

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	bank bank.Keeper

	taskStore sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(bank bank.Keeper, taskStore sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		bank:      bank,
		taskStore: taskStore,
		cdc:       cdc,
	}
}

// Get a individual task
// Get all tasks
// CreateTask
// Add value to task - add new value giver as backer
// Give proof of compeltion
// Vote if the prof is accepted
// Payout the person who completed it
