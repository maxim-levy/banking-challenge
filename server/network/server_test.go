package network

import (
	"client/network"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	// Run the server in Goroutine to stop tests from blocking
	// test execution.
	go func() {
		StartServer()
	}()
}

func TestStartServer(t *testing.T) {
	// Can we connect to the server with no errors.
	_, err := network.NewClient()
	assert.Nil(t, err)
}
