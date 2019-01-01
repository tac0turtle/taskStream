package taskstream

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	coinKeeper bank.Keeper

	ownersKey   sdk.StoreKey
	valueKey    sdk.StoreKey
	taskKey     sdk.StoreKey
	finishedKey sdk.StoreKey

	cdc *codec.Codec
}

func (k Keeper) createTask(ctx sdk.Context, taskName string, value string) {
	store := ctx.KVStore(k.taskKey)
	store.Set([]byte(taskName), []byte(value))
}

func (k Keeper) solveTask(ctx sdk.Context, taskName string, value string, solver sdk.AccAddress) {
	store := ctx.KVStore(k.finishedKey)
	store.Set([]byte(taskName), solver)
}

func (k Keeper) setOwner(ctx sdk.Context, taskName string, taker sdk.AccAddress) {
	store := ctx.KVStore(k.ownersKey)
	store.Set([]byte(taskName), taker)
}
