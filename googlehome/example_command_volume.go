package googlehome

import (
	"fmt"
	"strconv"

	"github.com/eure/bobo/command"
)

// VolumeCommand is a command to set volume of Google Home.
var VolumeCommand = command.BasicCommandTemplate{
	Help:           "Set Google Home volume",
	MentionCommand: "volume",
	GenerateFn: func(d command.CommandData) command.Command {
		c := command.Command{}

		volText := d.TextOther
		volume, err := strconv.ParseFloat(volText, 64)
		switch {
		case err != nil,
			volume < 0.0,
			volume > 1.0:
			errMessage := fmt.Sprintf("[ERROR]\tVolume must be from [0.0] to [1.0]")
			task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
			c.Add(task)
			return c
		}

		task, err := NewCastVolumeTask(volume)
		if err != nil {
			errMessage := fmt.Sprintf("[ERROR]\t[NewCastVolumeTask]\t`%s`", err.Error())
			task := command.NewReplyEngineTask(d.Engine, d.Channel, errMessage)
			c.Add(task)
			return c
		}

		c.Add(task)
		return c
	},
}
