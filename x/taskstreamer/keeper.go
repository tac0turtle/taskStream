package taskstreamer

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
	taskType "github.com/marbar3778/taskStream/types"
)

type Keeper struct {
	coinKeeper bank.Keeper

	taskStore sdk.StoreKey

	cdc *codec.Codec
}

func NewKeeper(coinKeeper bank.Keeper, taskStore sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper: coinKeeper,
		taskStore:  taskStore,
		cdc:        cdc,
	}
}

func (k Keeper) GetTask(ctx sdk.Context, taskTitle string) taskType.Task {
	store := ctx.KVStore(k.taskStore)
	if !store.Has([]byte(taskTitle)) {
		panic("Naa")
	}
	task := store.Get([]byte(taskTitle))
	var taskDetails taskType.Task
	k.cdc.MustUnmarshalBinaryBare(task, &taskDetails)
	return taskDetails
}

func (k Keeper) SetTask(ctx sdk.Context, taskTitle string, taskData taskType.Task) {
	store := ctx.KVStore(k.taskStore)
	store.Set([]byte(taskTitle), k.cdc.MustMarshalBinaryBare(taskData))
}

func (k Keeper) CreateTask(ctx sdk.Context, taskTitle string, taskDescription string, backers []sdk.AccAddress, value sdk.Coins) {
	task := taskType.CreateTask(taskTitle, taskDescription, backers, value)
	k.SetTask(ctx, taskTitle, task)
}

func (k Keeper) BecomeBacker(ctx sdk.Context, taskTitle string, newBacker sdk.AccAddress, addedValue sdk.Coins) {
	task := k.GetTask(ctx, taskTitle)
	task.Value = task.Value.Add(addedValue)
	task.Backers = append(task.Backers, newBacker)
	k.SetTask(ctx, taskTitle, task)
}

// func (k Keeper) PayoutTask(ctx sdk.Context, taskTitle string, receiver sdk.AccAddress) {

// }
