package docker

import (
	"encoding/json"
	"fmt"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"golang.org/x/net/context"
	"io"
)

// Image represents an image
type Image struct {
	types.ImageInspect
	client *client.Client
	ID     string
}

// ImageCollection represents all possible images
type ImageCollection struct {
	client *client.Client
}

// Pull an image and return it
func (images *ImageCollection) Pull(name string) (*Image, error) {
	resp, err := images.client.ImagePull(context.Background(), name, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}
	var dec = json.NewDecoder(resp)
	for {
		var jm jsonmessage.JSONMessage
		if err := dec.Decode(&jm); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if jm.Error != nil {
			return nil, jm.Error
		} else if jm.ErrorMessage != "" {
			return nil, fmt.Errorf(jm.ErrorMessage)
		}
	}
	return images.Get(name)
}

// Get an image and return it
func (images *ImageCollection) Get(name string) (*Image, error) {
	image, _, err := images.client.ImageInspectWithRaw(context.Background(), name)
	if err != nil {
		return nil, err
	}
	return &Image{
		client:       images.client,
		ImageInspect: image,
	}, nil
}
