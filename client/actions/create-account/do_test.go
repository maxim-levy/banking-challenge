package createaccount

import (
	"context"
	"protos/account"
	"server/models"
	"server/network"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	models.StartBoltDB("create-account.db")

	// Run the server in Goroutine to stop tests from blocking
	// test execution.
	go func() {
		network.StartServer()
	}()
}

func TestNewCreateAccount(t *testing.T) {
	s := NewCreateAccount("1000")
	assert.Equal(t, "1000", s.initBalance)
	assert.Equal(t, context.Background(), s.ctx)
	assert.Equal(t, account.Account{}, s.account)
}

func TestDo(t *testing.T) {
	s := NewCreateAccount("1000")
	err := s.Do()
	assert.Nil(t, err)
	resp := s.Result()
	an, err := resp.AccountNumber()
	assert.Nil(t, err)
	assert.Equal(t, 32, len(an))
}
