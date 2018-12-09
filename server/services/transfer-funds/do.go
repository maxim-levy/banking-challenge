package transferfunds

import (
	"errors"
	"fmt"
	"server/models"

	"github.com/apex/log"
)

// TransferFunds service
type TransferFunds struct {
	sourceAccountNumber      string
	destinationAccountNumber string
	fundsAmount              uint64
	sourceAccount            *models.Account
	destinationAccount       *models.Account
}

// NewTransferFunds instance
func NewTransferFunds(sourceAccount, destinationAccount string, fundsAmount uint64) *TransferFunds {
	return &TransferFunds{
		sourceAccountNumber:      sourceAccount,
		destinationAccountNumber: destinationAccount,
		fundsAmount:              fundsAmount,
		sourceAccount:            new(models.Account),
		destinationAccount:       new(models.Account),
	}
}

// Result can be called after successfull Do().
func (t *TransferFunds) Result() (
	sourceAccountNumber string,
	sourceAccountBalance int64,
	destinationAccountNumber string,
	destinationAccountBalance int64,
) {
	sourceAccountNumber = t.sourceAccountNumber
	sourceAccountBalance = t.sourceAccount.Balance
	destinationAccountNumber = t.destinationAccountNumber
	destinationAccountBalance = t.destinationAccount.Balance
	return
}

// Do steps
func (t *TransferFunds) Do() (err error) {
	if err := t.validate(); err != nil {
		return err
	}

	if err := t.getSourceAccount(); err != nil {
		return err
	}

	if err := t.getDestinationAccount(); err != nil {
		return err
	}

	if err := t.checkBalances(); err != nil {
		return err
	}

	if err := t.transfer(); err != nil {
		return err
	}

	return nil
}

func (t *TransferFunds) validate() (err error) {
	if t.sourceAccountNumber == "" {
		err = errors.New("source-account is required")
		log.Warn(err.Error())
		return err
	}

	if t.destinationAccountNumber == "" {
		err = errors.New("destination-account is required")
		log.Warn(err.Error())
		return err
	}

	if t.fundsAmount <= 0 {
		err = errors.New("funds-amount must be greater than 0")
		log.Warn(err.Error())
		return err
	}

	return nil
}

func (t *TransferFunds) getSourceAccount() (err error) {
	t.sourceAccount, err = t.sourceAccount.GetByAccountNumber(t.sourceAccountNumber)
	if err != nil {
		log.WithError(err).Error("failed to get source account")
		return err
	}

	// Check if response is empty
	if t.sourceAccount.ID == "" {
		err = fmt.Errorf("no source account with account number (%s) found", t.sourceAccountNumber)
		log.Warn(err.Error())
		return err
	}

	return nil
}

func (t *TransferFunds) getDestinationAccount() (err error) {
	t.destinationAccount, err = t.destinationAccount.GetByAccountNumber(t.destinationAccountNumber)
	if err != nil {
		log.WithError(err).Error("failed to get destination account")
		return err
	}

	// Check if response is empty
	if t.destinationAccount.ID == "" {
		err = fmt.Errorf("no destination account with account number (%s) found", t.destinationAccountNumber)
		log.Warn(err.Error())
		return err
	}

	return nil
}

func (t *TransferFunds) checkBalances() error {
	if int64(t.fundsAmount) > t.sourceAccount.Balance {
		err := fmt.Errorf(
			"insufficient funds, requested to transfer %d out of %d balance",
			t.fundsAmount,
			t.sourceAccount.Balance,
		)
		log.Warn(err.Error())
		return err
	}

	return nil
}

func (t *TransferFunds) transfer() error {
	// Calculate new balances
	t.sourceAccount.Balance = t.sourceAccount.Balance - int64(t.fundsAmount)
	t.destinationAccount.Balance = t.destinationAccount.Balance + int64(t.fundsAmount)

	// Save changes
	err := new(models.Account).TransferFunds(
		t.sourceAccountNumber,
		t.sourceAccount.Balance,
		t.destinationAccountNumber,
		t.destinationAccount.Balance,
	)
	if err != nil {
		log.WithError(err).Error("failed to transfer funds")
		return err
	}

	return nil
}
