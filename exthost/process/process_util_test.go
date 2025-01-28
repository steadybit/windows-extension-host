package stopprocess

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStopProcesses(t *testing.T) {
	command := exec.Command("tail", "-f", "/dev/null", "&")
	err := command.Start()
	assert.NoError(t, err)
	ids := FindProcessIds("tail")
	assert.Equal(t, 1, len(ids))
	assert.Equal(t, command.Process.Pid, ids[0])
	err = StopProcesses(ids, true)
	assert.NoError(t, err)
}
