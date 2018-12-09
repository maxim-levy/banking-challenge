package createaccount

import (
	"errors"
	"server/models"
	"strings"

	"github.com/apex/log"
	"github.com/google/uuid"
)

// CreateAccount service
type CreateAccount struct {
	initBalance   int64
	accountNumber string
	account       *models.Account
}

// NewCreateAccount instance
func NewCreateAccount(initBalance int64) *CreateAccount {
	return &CreateAccount{
		initBalance: initBalance,
		account:     new(models.Account),
	}
}

// Result can be called after successfull Do().
func (c *CreateAccount) Result() (accountNumber string, balance int64) {
	accountNumber = c.accountNumber
	balance = c.initBalance
	return
}

// Do steps
func (c *CreateAccount) Do() (err error) {
	if err := c.validate(); err != nil {
		return err
	}

	if err := c.generateAccountNumber(); err != nil {
		return err
	}

	if err := c.saveAccount(); err != nil {
		return err
	}

	return nil
}

func (c *CreateAccount) validate() (err error) {
	if c.initBalance < 0 {
		err = errors.New("can not create a new account with a negative balance")
		log.Warn(err.Error())
		return err
	}

	return nil
}

func (c *CreateAccount) generateAccountNumber() error {
	id, err := uuid.NewRandom()
	if err != nil {
		log.WithError(err).Error("failed to generate account number")
		return err
	}

	// Remove hyphens to make the string 32 bytes.
	c.accountNumber = strings.Replace(id.String(), "-", "", -1)

	// Check if there is already an account with this accountNumber
	c.account, err = c.account.GetByAccountNumber(c.accountNumber)
	if err != nil {
		log.WithError(err).Error("failed to GetByAccountNumber")
		return err
	}

	if c.account.Balance != 0 {
		err = errors.New("account with this id already exists")
		log.Error(err.Error())
		return err
	}

	return nil
}

func (c *CreateAccount) saveAccount() (err error) {
	c.account.ID = c.accountNumber
	c.account.Balance = c.initBalance
	c.account, err = c.account.Create()
	if err != nil {
		log.WithError(err).Error("failed to save account")
		return err
	}

	return nil
}
