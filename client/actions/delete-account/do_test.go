package deleteaccount

import (
	"context"
	"server/models"
	"server/network"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	models.StartBoltDB("transfer-funds.db")

	// Run the server in Goroutine to stop tests from blocking
	// test execution.
	go func() {
		network.StartServer()
	}()
}

func TestNewDeleteAccount(t *testing.T) {
	s := NewDeleteAccount("TestNewDeleteAccount")
	assert.Equal(t, "TestNewDeleteAccount", s.accountNumber)
	assert.Equal(t, context.Background(), s.ctx)
}
