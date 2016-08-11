package docker

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestSimpleRun(t *testing.T) {
	client, err := FromEnv()
	assert.Nil(t, err)

	opts := &RunOptions{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world"},
	}

	container, err := client.Containers.Run(opts)
	assert.Nil(t, err)

	out, err := container.Logs(nil)
	assert.Nil(t, err)

	s, err := ioutil.ReadAll(out)
	assert.Nil(t, err)

	assert.Contains(t, string(s), "hello world")
}
