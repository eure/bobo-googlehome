package googlehome

import (
	"sync"

	"github.com/evalphobia/google-home-client-go/googlehome"
)

type reloadClientTask struct {
	client *googlehome.Client
}

// NewReloadClientTask is a task to play voice on Google Home.
func NewReloadClientTask() (reloadClientTask, error) {
	cli, err := getGlobalClient()
	return NewReloadClientTaskWithClient(cli), err
}

// NewReloadClientTaskWithClient is a task to play voice on Google Home.
func NewReloadClientTaskWithClient(client *googlehome.Client) reloadClientTask {
	return reloadClientTask{
		client: client,
	}
}

func (reloadClientTask) GetName() string {
	return "reload_client_task"
}

func (t reloadClientTask) Run() error {
	var err error
	globalClient, err = googlehome.NewClient()
	return err
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
