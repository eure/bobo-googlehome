package googlehome

import (
	"errors"

	"github.com/evalphobia/google-home-client-go/googlehome"
)

type castVolumeTask struct {
	client *googlehome.Client
	volume float64
}

// NewCastVolumeTask is a task to set volume on Google Home.
func NewCastVolumeTask(volume float64) (castVolumeTask, error) {
	cli, err := getGlobalClient()
	return NewCastVolumeTaskWithClient(cli, volume), err
}

// NewCastVolumeTaskWithClient is a task to set volume on Google Home.
func NewCastVolumeTaskWithClient(client *googlehome.Client, volume float64) castVolumeTask {
	return castVolumeTask{
		client: client,
		volume: volume,
	}
}

func (castVolumeTask) GetName() string {
	return "cast_volume_task"
}

func (t castVolumeTask) Run() error {
	if t.client == nil {
		return errors.New("Set Google Home Client")
	}

	return t.client.SetVolume(t.volume)
}
