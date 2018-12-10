package transferfunds

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

func TestNewTransferFunds(t *testing.T) {
	s := NewTransferFunds("sourceAccountID", "destinationAccountID", "1000")
	assert.Equal(t, "sourceAccountID", s.sourceAccount)
	assert.Equal(t, "destinationAccountID", s.destinationAccount)
	assert.Equal(t, "1000", s.fundsAmount)
	assert.Equal(t, context.Background(), s.ctx)
}
