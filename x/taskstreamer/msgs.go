// package taskstreamer

// import (
// 	"encoding/json"
// 	"fmt"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	taskType "github.com/marbar3778/taskStream/types"
// )

// type MsgCreateTask struct {
// 	TaskTitle       string
// 	TaskDescription string
// 	Backer          sdk.AccAddress
// 	Value           sdk.Coins
// }

// func NewMsgCreateTask(taskTitle string, taskDescription string, backer sdk.AccAddress, value sdk.Coins) taskType.Task {
// 	return MsgCreateTask{
// 		TaskTitle:       taskTitle,
// 		TaskDescription: taskDescription,
// 		Backer:          backer,
// 		Value:           value,
// 	}
// }

// func (msg MsgCreateTask) Route() string { return "tasker" }

// func (msg MsgCreateTask) Type() string { return "create_task" }

// func (msg MsgCreateTask) Validatebasic() sdk.Error {
// 	if len(msg.TaskTitle) == 0 || len(msg.TaskDescription) == 0 {
// 		return sdk.ErrUnknownRequest("There is no task title and/or task description")
// 	}
// 	if msg.Backer.Empty() {
// 		return sdk.ErrInvalidAddress(fmt.Sprintf("By creating a task you must be the first backer, %s", msg.Backer.String()))
// 	}
// 	if !msg.Value.IsAllPositive() {
// 		return sdk.ErrInsufficientCoins("You must provide a value to the task")
// 	}
// 	return nil
// }

// func (msg MsgCreateTask) GetSignBytes() []byte {
// 	b, err := json.Marshal(msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return sdk.MustSortJSON(b)
// }

// func (msg MsgCreateTask) GetSigners() []sdk.AccAddress {
// 	return []sdk.AccAddress{msg.Backer}
// }
