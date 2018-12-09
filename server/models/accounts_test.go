package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	StartBoltDB("ledger_account_test.db")
}

func TestCreate(t *testing.T) {
	a := &Account{
		ID:      "TestCreate",
		Balance: 1000,
	}
	b, err := a.Create()
	assert.Nil(t, err)
	assert.Equal(t, b.ID, a.ID)
	assert.Equal(t, b.Balance, a.Balance)
}

func TestGetByAccountNumber(t *testing.T) {
	// Create record
	account := &Account{
		ID:      "TestGetByAccountNumber",
		Balance: 1000,
	}
	_, err := account.Create()
	assert.Nil(t, err)

	// Fetch the record
	a, err := new(Account).GetByAccountNumber(account.ID)
	assert.Nil(t, err)
	assert.Equal(t, a.ID, account.ID)
	assert.Equal(t, a.Balance, account.Balance)
}

func TestDelete(t *testing.T) {
	// Create record
	account := &Account{
		ID:      "TestDelete",
		Balance: 1000,
	}
	_, err := account.Create()
	assert.Nil(t, err)

	// Delete record
	err = account.Delete(account.ID)
	assert.Nil(t, err)
}

func TestTransferFunds(t *testing.T) {
	// Create records
	accountA := &Account{
		ID:      "TestTransferFundsAccountA",
		Balance: 1000,
	}
	_, err := accountA.Create()
	assert.Nil(t, err)

	accountB := &Account{
		ID:      "TestTransferFundsAccountB",
		Balance: 1000,
	}
	_, err = accountB.Create()
	assert.Nil(t, err)

	// Transfer from A to B
	err = new(Account).TransferFunds(
		accountA.ID,
		500,
		accountB.ID,
		1500,
	)
	assert.Nil(t, err)

	// Fetch accounts to verify writes
	aa, err := new(Account).GetByAccountNumber(accountA.ID)
	assert.Nil(t, err)
	assert.Equal(t, aa.Balance, int64(500))

	bb, err := new(Account).GetByAccountNumber(accountB.ID)
	assert.Nil(t, err)
	assert.Equal(t, bb.Balance, int64(1500))
}
