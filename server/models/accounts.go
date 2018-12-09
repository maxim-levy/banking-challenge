package models

import (
	"fmt"
	"strconv"

	"github.com/apex/log"
	"github.com/boltdb/bolt"
)

// AccountsBucketName name of bolt db accounts bucket.
const AccountsBucketName = "Accounts"

// Account model
type Account struct {
	ID      string
	Balance int64
}

// Create new account
func (a *Account) Create() (*Account, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(AccountsBucketName))
		return b.Put([]byte(a.ID), []byte(fmt.Sprintf("%d", a.Balance)))
	})
	if err != nil {
		log.WithError(err).Error("failed to Create account")
		return a, err
	}

	return a, nil
}

// GetByAccountNumber returns a signgle account marching the number provided.
// If no account is found it will return an empty Account struct.
func (a *Account) GetByAccountNumber(accountNumber string) (*Account, error) {
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(AccountsBucketName))
		v := b.Get([]byte(accountNumber))
		if string(v) != "" {
			a.ID = accountNumber
			balance, err := strconv.Atoi(string(v))
			if err != nil {
				log.WithError(err).Errorf("not a number: %s", v)
				return err
			}
			a.Balance = int64(balance)
		}
		return nil
	})
	if err != nil {
		log.WithError(err).Error("failed to GetByAccountNumber")
		return a, err
	}

	return a, nil
}

// Delete account
func (a *Account) Delete(accountNumber string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(AccountsBucketName))
		err := b.Delete([]byte(accountNumber))
		return err
	})
	if err != nil {
		log.WithError(err).Error("failed to Delete account")
		return err
	}

	return nil
}

// TransferFunds between two accounts
func (a *Account) TransferFunds(
	sourceAccountNumber string,
	sourceAccountBalance int64,
	destinationAccountNumber string,
	destinationAccountBalance int64,
) error {
	// Start a writable transaction.
	tx, err := db.Begin(true)
	if err != nil {
		log.WithError(err).Error("failed to start transaction")
		return err
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte(AccountsBucketName))

	// Wite source changes
	err = b.Put([]byte(sourceAccountNumber), []byte(fmt.Sprintf("%d", sourceAccountBalance)))
	if err != nil {
		log.WithError(err).Error("failed to update account")
		return err
	}

	// Wite destination changes
	err = b.Put([]byte(destinationAccountNumber), []byte(fmt.Sprintf("%d", destinationAccountBalance)))
	if err != nil {
		log.WithError(err).Error("failed to update account")
		return err
	}

	// Commit the transaction and check for error.
	if err := tx.Commit(); err != nil {
		log.WithError(err).Error("failed to commit the transaction")
		return err
	}

	return nil
}
