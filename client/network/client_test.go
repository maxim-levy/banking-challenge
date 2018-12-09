package network

import (
	serverNetwork "server/network"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	// Run the server in Goroutine to stop tests from blocking
	// test execution.
	go func() {
		serverNetwork.StartServer()
	}()
}

func TestNewClient(t *testing.T) {
	// Can we create a client with no errors
	_, err := NewClient()
	assert.Nil(t, err)
}
