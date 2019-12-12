package googlehome

import (
	"errors"
	"sync"

	"github.com/evalphobia/google-home-client-go/googlehome"
)

type castPlayTask struct {
	client *googlehome.Client
	text   string
}

// NewCastPlayTask is a task to play voice on Google Home.
func NewCastPlayTask(text string) (castPlayTask, error) {
	cli, err := getGlobalClient()
	return NewCastPlayTaskWithClient(cli, text), err
}

// NewCastPlayTaskWithClient is a task to play voice on Google Home.
func NewCastPlayTaskWithClient(client *googlehome.Client, text string) castPlayTask {
	return castPlayTask{
		client: client,
		text:   text,
	}
}

func (castPlayTask) GetName() string {
	return "cast_play_task"
}

func (t castPlayTask) Run() error {
	if t.client == nil {
		return errors.New("Set Google Home Client")
	}

	return t.client.Notify(t.text)
}

var ghOnce sync.Once
var globalClient *googlehome.Client

func getGlobalClient() (*googlehome.Client, error) {
	var err error
	ghOnce.Do(func() {
		globalClient, err = googlehome.NewClient()
	})
	return globalClient, err
}
