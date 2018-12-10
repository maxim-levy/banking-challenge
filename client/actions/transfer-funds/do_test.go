package transferfunds

import (
	createaccount "client/actions/create-account"
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

func TestDo(t *testing.T) {
	// Create accounts
	accountA := createaccount.NewCreateAccount("1000")
	err := accountA.Do()
	assert.Nil(t, err)
	accountAResp := accountA.Result()
	accountANumber, err := accountAResp.AccountNumber()
	assert.Nil(t, err)

	accountB := createaccount.NewCreateAccount("1000")
	err = accountB.Do()
	assert.Nil(t, err)
	accountBResp := accountB.Result()
	accountBNumber, err := accountBResp.AccountNumber()
	assert.Nil(t, err)

	s := NewTransferFunds(accountANumber, accountBNumber, "1000")
	err = s.Do()
	assert.Nil(t, err)
	// Check result
	record := s.Result()
	sourceAccount, err := record.SourceAccount()
	assert.Nil(t, err)
	assert.Equal(t, accountANumber, sourceAccount)
	destinationAccount, err := record.DestinationAccount()
	assert.Nil(t, err)
	assert.Equal(t, accountBNumber, destinationAccount)
	assert.Equal(t, uint64(1000), record.Amount())
	assert.Equal(t, int64(0), record.SourceBalance())
	assert.Equal(t, int64(2000), record.DestinationBalance())
}
