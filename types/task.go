package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type Task struct {
	TaskTitle       string           `json:"task_title"`
	TaskDescription string           `json:"task_description"`
	Backers         []sdk.AccAddress `json:"backers"`
	Value           sdk.Coins        `json:"value`
}

func CreateTask(taskTitle string, taskDescription string, backers []sdk.AccAddress, value sdk.Coins) Task {
	return Task{
		TaskTitle:       taskTitle,
		TaskDescription: taskDescription,
		Backers:         backers,
		Value:           value,
	}
}
