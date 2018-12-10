package createaccount

import (
	"client/network"
	"context"
	"fmt"
	"protos/account"
	"strconv"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/rpc"
)

// CreateAccount service
type CreateAccount struct {
	initBalance string
	conn        *rpc.Conn
	ctx         context.Context
	account     account.Account
}

// NewCreateAccount instance
func NewCreateAccount(initBalance string) *CreateAccount {
	return &CreateAccount{
		initBalance: initBalance,
		ctx:         context.Background(),
		account:     account.Account{},
	}
}

// Result can be called after successfull Do().
func (c *CreateAccount) Result() account.Account {
	return c.account
}

// Do steps
func (c *CreateAccount) Do() (err error) {
	if err = c.connectToServer(); err != nil {
		return err
	}

	if err = c.sendCreateRequst(); err != nil {
		return err
	}

	if err = c.printAccountInfo(); err != nil {
		return err
	}

	return nil
}

func (c *CreateAccount) connectToServer() (err error) {
	c.conn, err = network.NewClient()
	if err != nil {
		log.WithError(err).Error("failed to connectToServer")
		return err
	}

	return nil
}

func (c *CreateAccount) sendCreateRequst() error {
	sess := account.AccountFactory{Client: c.conn.Bootstrap(c.ctx)}
	ca := sess.CreateAccount(c.ctx, func(p account.AccountFactory_createAccount_Params) error {
		// Set initial balance
		ib, err := strconv.Atoi(c.initBalance)
		if err != nil {
			log.WithError(err).Errorf("failed to convert %s to int", c.initBalance)
			return err
		}
		p.SetInitialBalance(int64(ib))
		return nil
	})
	resp, err := ca.Struct()
	if err != nil {
		log.WithError(err).Error("CreateAccount call failed")
		return err
	}

	c.account, err = resp.Account()
	if err != nil {
		log.WithError(err).Error("failed to decode response")
		return err
	}

	return nil
}

func (c *CreateAccount) printAccountInfo() error {
	accountNumber, err := c.account.AccountNumber()
	if err != nil {
		log.WithError(err).Error("failed to decode SourceAccount field")
		return err
	}
	balance := c.account.Balance()

	// Printer
	fmt.Println("------------------------")
	fmt.Println("--- ACCOUNT CREATED! ---")
	fmt.Println("------------------------")
	fmt.Printf("Account number: %s\n", accountNumber)
	fmt.Printf("Account balance: %d cents\n", balance)
	fmt.Println("------------------------")

	return nil
}
