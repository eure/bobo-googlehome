package googlehome

import (
	"fmt"

	"github.com/eure/bobo/command"
)

// ReloadCommand is a command to reset Google Home client.
var ReloadCommand = command.BasicCommandTemplate{
	Help:           "Reset Google Home client",
	MentionCommand: "reload",
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		task, err := NewReloadClientTask()
		if err != nil {
			errMessage := fmt.Sprintf("[ERROR]\t[NewReloadClientTask]\t`%s`", err.Error())
			task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
			c.Add(task)
			return c
		}

		if err := task.Run(); err != nil {
			errMessage := fmt.Sprintf("[ERROR]\t[NewReloadClientTask.Run]\t`%s`", err.Error())
			task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
			c.Add(task)
			return c
		}
		return c
	},
}
