package docker

import (
	"github.com/docker/engine-api/client"
)

func FromEnv() (*Client, error) {
	cl, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		APIClient:  cl,
		Containers: &ContainerCollection{client: cl},
	}, nil
}

type Client struct {
	APIClient  *client.Client
	Containers *ContainerCollection
}
