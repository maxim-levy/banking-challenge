package createaccount

import (
	"server/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	models.StartBoltDB("create_account.db")
}

func TestNewCreateAccount(t *testing.T) {
	s := NewCreateAccount(1000)
	assert.Equal(t, int64(1000), s.initBalance)
	assert.Equal(t, &models.Account{}, s.account)
}

func TestDo(t *testing.T) {
	s := NewCreateAccount(1000)
	err := s.Do()
	assert.Nil(t, err)
	// Should throw validation error
	s = NewCreateAccount(-100)
	err = s.Do()
	assert.NotNil(t, err)
}

func TestResult(t *testing.T) {
	s := NewCreateAccount(1000)
	err := s.Do()
	assert.Nil(t, err)
	accountNumber, balance := s.Result()
	assert.Equal(t, 32, len(accountNumber))
	assert.Equal(t, int64(1000), balance)
}

var benchmarkPackageErr error

func BenchmarkCreateAccount(b *testing.B) {
	var err error
	// run the Do function b.N times
	for n := 0; n < b.N; n++ {
		s := NewCreateAccount(1000)
		// always record the result of Fib to prevent
		// the compiler eliminating the function call.
		err = s.Do()
	}
	// always store the result to a package level variable
	// so the compiler cannot eliminate the Benchmark itself.
	benchmarkPackageErr = err
}
