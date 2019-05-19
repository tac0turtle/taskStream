package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Task struct {
	TaskTitle       string           `json:"task_title"`
	TaskDescription string           `json:"task_description"`
	Backers         []sdk.AccAddress `json:"backers"`
	Value           sdk.Coins        `json:"value"`
	IsCompleted     bool             `json:"is_completed"`
}

func CreateTask(taskTitle string, taskDescription string, backers []sdk.AccAddress, value sdk.Coins, completed bool) Task {
	return Task{
		TaskTitle:       taskTitle,
		TaskDescription: taskDescription,
		Backers:         backers,
		Value:           value,
		IsCompleted:     completed,
	}
}

func (t Task) GetTitle() string             { return t.TaskTitle }
func (t Task) GetDescription() string       { return t.TaskDescription }
func (t Task) GetBackers() []sdk.AccAddress { return t.Backers }

func (t Task) String() string {
	return fmt.Sprintf(`
	Task Title: %s,
	Task Description: %s,
	Backers: %s,
	Value: %v`, t.TaskTitle, t.TaskDescription, t.Backers, t.Value)
}
