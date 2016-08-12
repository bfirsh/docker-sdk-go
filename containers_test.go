package docker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleRun(t *testing.T) {
	client, err := FromEnv()
	assert.Nil(t, err)

	container, err := client.Containers.Run(&RunOptions{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	})
	assert.Nil(t, err)

	out, err := container.Logs(nil)
	assert.Nil(t, err)

	assert.Contains(t, string(out), "hello world")
}

func TestDetachedRun(t *testing.T) {
	client, err := FromEnv()
	assert.Nil(t, err)

	container, err := client.Containers.Run(&RunOptions{
		Image:  "bfirsh/reticulate-splines",
		Detach: true,
	})
	assert.Nil(t, err)

	out, err := container.Logs(nil)
	assert.Nil(t, err)

	assert.Contains(t, string(out), "Reticulating splines")

	err = container.Stop(nil)
	assert.Nil(t, err)
}
