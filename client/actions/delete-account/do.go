package deleteaccount

import (
	"client/network"
	"context"
	"errors"
	"fmt"
	"protos/account"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/rpc"
)

// DeleteAccount service
type DeleteAccount struct {
	accountNumber string
	conn          *rpc.Conn
	ctx           context.Context
}

// NewDeleteAccount instance
func NewDeleteAccount(accountNumber string) *DeleteAccount {
	return &DeleteAccount{
		accountNumber: accountNumber,
		ctx:           context.Background(),
	}
}

// Do steps
func (d *DeleteAccount) Do() (err error) {
	if err = d.validate(); err != nil {
		return err
	}

	if err = d.connectToServer(); err != nil {
		return err
	}

	if err = d.sendCreateRequst(); err != nil {
		return err
	}

	if err = d.printAccountInfo(); err != nil {
		return err
	}

	return nil
}

func (d *DeleteAccount) validate() (err error) {
	if d.accountNumber == "" {
		return errors.New("--account-number flag is required")
	}

	return nil
}

func (d *DeleteAccount) connectToServer() (err error) {
	d.conn, err = network.NewClient()
	if err != nil {
		log.WithError(err).Error("failed to connectToServer")
		return err
	}

	return nil
}

func (d *DeleteAccount) sendCreateRequst() error {
	sess := account.AccountFactory{Client: d.conn.Bootstrap(d.ctx)}
	ca := sess.DeleteAccount(d.ctx, func(p account.AccountFactory_deleteAccount_Params) error {
		err := p.SetAccountNumber(d.accountNumber)
		return err
	})
	resp, err := ca.Struct()
	if err != nil {
		log.WithError(err).Error("CreateAccount call failed")
		return err
	}

	// Check if the deletion was successfull
	if resp.Success() == false {
		err = errors.New("failed to delete account")
		log.Error(err.Error())
		return err
	}

	return nil
}

func (d *DeleteAccount) printAccountInfo() error {
	// Printer
	fmt.Println("------------------------")
	fmt.Println("--- ACCOUNT DELETED! ---")
	fmt.Println("------------------------")
	fmt.Printf("Account number: %s\n", d.accountNumber)
	fmt.Println("------------------------")

	return nil
}
