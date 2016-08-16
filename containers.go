package docker

import (
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"golang.org/x/net/context"
	"io/ioutil"
	"time"
)

// RunOptions is the configuration passed to Containers.Run()
type RunOptions struct {
	Image  string
	Cmd    []string
	Detach bool
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
	// TODO: figure out what to do with contexts
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

// Stop a container
func (container *Container) Stop(timeout *time.Duration) error {
	return container.client.ContainerStop(context.Background(), container.ID, timeout)
}

// Wait waits for a container to finish then returns its exit code
func (container *Container) Wait() (int, error) {
	return container.client.ContainerWait(context.Background(), container.ID)
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
	// TODO: allow setting of context
	resp, err := containers.client.ContainerCreate(context.Background(), createConfig, nil, nil, "")
	if err != nil {
		if client.IsErrImageNotFound(err) {
			images := &ImageCollection{client: containers.client}
			if _, err = images.Pull(options.Image); err != nil {
				return nil, err
			}
			// retry
			resp, err = containers.client.ContainerCreate(context.Background(), createConfig, nil, nil, "")
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	container := &Container{
		ID:     resp.ID,
		client: containers.client,
	}
	err = container.Start()
	if err != nil || options.Detach {
		return container, err
	}
	_, err = container.Wait()
	// TODO: return error on exitCode?
	return container, err
}
