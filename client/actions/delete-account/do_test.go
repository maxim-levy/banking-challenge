package deleteaccount

import (
	createaccount "client/actions/create-account"
	"context"
	"server/models"
	"server/network"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	models.StartBoltDB("delete-account.db")

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

func TestDo(t *testing.T) {
	s := createaccount.NewCreateAccount("1000")
	err := s.Do()
	assert.Nil(t, err)
	resp := s.Result()
	an, err := resp.AccountNumber()
	assert.Nil(t, err)

	// Delete account
	d := NewDeleteAccount(an)
	err = d.Do()
	assert.Nil(t, err)
}
