package docker

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"golang.org/x/net/context"
	"io/ioutil"
)

// RunOptions is the configuration passed to Containers.Run()
type RunOptions struct {
	Image string
	Cmd   []string
}

// Container represents a container
type Container struct {
	client *client.Client
	ID     string
}

// Logs returns logs for a container
func (container *Container) Logs(options *types.ContainerLogsOptions) ([]byte, error) {
	if options == nil {
		options = &types.ContainerLogsOptions{ShowStdout: true}
	}
	out, err := container.client.ContainerLogs(context.Background(), container.ID, *options)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(out)
}

// Start starts a container
func (container *Container) Start() error {
	options := types.ContainerStartOptions{}
	return container.client.ContainerStart(context.Background(), container.ID, options)
}

// ContainerCollection represents all possible containers
type ContainerCollection struct {
	client *client.Client
}

// Run a container
func (containers *ContainerCollection) Run(options *RunOptions) (*Container, error) {
	createConfig := &container.Config{
		Image: options.Image,
		Cmd:   options.Cmd,
	}
	resp, err := containers.client.ContainerCreate(context.Background(), createConfig, nil, nil, "")
	if err != nil {
		return nil, err
	}
	container := &Container{
		ID:     resp.ID,
		client: containers.client,
	}
	err = container.Start()
	return container, err
}
