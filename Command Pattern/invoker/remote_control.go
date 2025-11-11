package invoker

import "command-pattern/command"

type RemoteControl struct {
	command command.Command
}

func NewRemoteControl(cmd command.Command) *RemoteControl {
	return &RemoteControl{
		command: cmd,
	}
}

func (rc *RemoteControl) PressButton() {
	rc.command.PressButton()
}
