package taskstream

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	coinKeeper bank.Keeper

	authorsKey sdk.StoreKey
	valueKey   sdk.StoreKey
	taskKey    sdk.StoreKey

	cdc *codec.Codec
}

func (k Keeper) createTask(ctx sdk.Context, taskName string, value string) {
	store := ctx.KVStore(k.taskKey)
	store.Set([]byte(taskName), []byte(value))
}

func (k Keeper) solveTask(ctx sdk.Context, taskName string) string {
	store := ctx.KVStore(k.taskKey)
	bz := store.Get([]byte(taskName))
	return string(bz)
}
