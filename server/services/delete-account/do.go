package deleteaccount

import (
	"errors"
	"server/models"

	"github.com/apex/log"
)

// DeleteAccount service
type DeleteAccount struct {
	accountNumber string
	account       *models.Account
}

// NewDeleteAccount instance
func NewDeleteAccount(accountNumber string) *DeleteAccount {
	return &DeleteAccount{
		accountNumber: accountNumber,
		account:       new(models.Account),
	}
}

// Do steps
func (d *DeleteAccount) Do() (err error) {
	if err := d.validate(); err != nil {
		return err
	}

	if err := d.saveAccount(); err != nil {
		return err
	}

	return nil
}

func (d *DeleteAccount) validate() (err error) {
	if d.accountNumber == "" {
		err = errors.New("account-number is required")
		log.Warn(err.Error())
		return err
	}

	return nil
}

func (d *DeleteAccount) saveAccount() (err error) {
	err = d.account.Delete(d.accountNumber)
	if err != nil {
		log.WithError(err).Error("failed to delete account")
		return err
	}

	return nil
}
