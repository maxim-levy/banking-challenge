package createaccount

import (
	"context"
	"protos/account"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCreateAccount(t *testing.T) {
	s := NewCreateAccount("1000")
	assert.Equal(t, "1000", s.initBalance)
	assert.Equal(t, context.Background(), s.ctx)
	assert.Equal(t, account.Account{}, s.account)
}
