package googlehome

import (
	"fmt"

	"github.com/eure/bobo/command"
)

// SayCommand is a command to play voice sound on Google Home.
var SayCommand = command.BasicCommandTemplate{
	Help:           "Say voice on Google Home",
	MentionCommand: "say",
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		if d.TextOther == "" {
			errMessage := fmt.Sprintf("[ERROR] Text is empty")
			task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
			c.Add(task)
			return c
		}

		task, err := NewCastPlayTask(d.TextOther)
		if err != nil {
			errMessage := fmt.Sprintf("[ERROR]\t[NewCastPlayTask]\t`%s`", err.Error())
			task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
			c.Add(task)
			return c
		}

		c.Add(task)
		return c
	},
}
