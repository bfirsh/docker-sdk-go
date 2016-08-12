package docker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPull(t *testing.T) {
	client, err := FromEnv()
	assert.Nil(t, err)

	image, err := client.Images.Pull("alpine")
	assert.Nil(t, err)

	assert.Equal(t, image.RepoTags[0], "alpine:latest")
}
