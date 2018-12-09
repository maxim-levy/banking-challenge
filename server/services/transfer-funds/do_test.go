package transferfunds

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	models.StartBoltDB("transfer_funds.db")
}

func TestNewTransferFunds(t *testing.T) {
	s := NewTransferFunds("sourceAccountID", "destinationAccountID", 1000)
	assert.Equal(t, "sourceAccountID", s.sourceAccountNumber)
	assert.Equal(t, "destinationAccountID", s.destinationAccountNumber)
	assert.Equal(t, uint64(1000), s.fundsAmount)
	assert.Equal(t, &models.Account{}, s.sourceAccount)
	assert.Equal(t, &models.Account{}, s.destinationAccount)
}

func TestDo(t *testing.T) {
	// Create accounts
	accountA := &models.Account{
		ID:      "sourceAccountID",
		Balance: 1000,
	}
	_, err := accountA.Create()
	assert.Nil(t, err)

	accountB := &models.Account{
		ID:      "destinationAccountID",
		Balance: 1000,
	}
	_, err = accountB.Create()
	assert.Nil(t, err)

	// Test transfer
	s := NewTransferFunds("sourceAccountID", "destinationAccountID", 1000)
	err = s.Do()
	assert.Nil(t, err)

	// Should throw errors
	s = NewTransferFunds("", "destinationAccountID", 1000)
	err = s.Do()
	assert.NotNil(t, err)

	s = NewTransferFunds("sourceAccountID", "", 1000)
	err = s.Do()
	assert.NotNil(t, err)

	s = NewTransferFunds("sourceAccountID", "destinationAccountID", 0)
	err = s.Do()
	assert.NotNil(t, err)
}

func TestResult(t *testing.T) {
	// Create accounts
	accountA := &models.Account{
		ID:      "sourceAccountID",
		Balance: 1000,
	}
	_, err := accountA.Create()
	assert.Nil(t, err)

	accountB := &models.Account{
		ID:      "destinationAccountID",
		Balance: 1000,
	}
	_, err = accountB.Create()
	assert.Nil(t, err)

	// Test transfer
	s := NewTransferFunds("sourceAccountID", "destinationAccountID", 1000)
	err = s.Do()
	assert.Nil(t, err)
	sourceAccountNumber,
		sourceAccountBalance,
		destinationAccountNumber,
		destinationAccountBalance := s.Result()
	assert.Equal(t, accountA.ID, sourceAccountNumber)
	assert.Equal(t, int64(0), sourceAccountBalance)
	assert.Equal(t, accountB.ID, destinationAccountNumber)
	assert.Equal(t, int64(2000), destinationAccountBalance)
}
