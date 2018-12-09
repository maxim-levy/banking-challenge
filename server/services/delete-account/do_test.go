package deleteaccount

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	models.StartBoltDB("delete_account.db")
}

func TestNewDeleteAccount(t *testing.T) {
	s := NewDeleteAccount("TestNewDeleteAccount")
	assert.Equal(t, "TestNewDeleteAccount", s.accountNumber)
	assert.Equal(t, &models.Account{}, s.account)
}

func TestDo(t *testing.T) {
	s := NewDeleteAccount("TestDo")
	err := s.Do()
	assert.Nil(t, err)
	// Should throw validation error
	s = NewDeleteAccount("")
	err = s.Do()
	assert.NotNil(t, err)
}
