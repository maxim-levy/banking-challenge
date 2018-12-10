package transferfunds

import (
	"client/network"
	"context"
	"errors"
	"fmt"
	"protos/account"
	"strconv"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/rpc"
)

// TransferFunds service
type TransferFunds struct {
	sourceAccount      string
	destinationAccount string
	fundsAmount        string
	conn               *rpc.Conn
	ctx                context.Context
	record             account.TransactionalRecord
}

// NewTransferFunds instance
func NewTransferFunds(sourceAccount, destinationAccount, fundsAmount string) *TransferFunds {
	return &TransferFunds{
		sourceAccount:      sourceAccount,
		destinationAccount: destinationAccount,
		fundsAmount:        fundsAmount,
		ctx:                context.Background(),
	}
}

// Result can be called after successfull Do().
func (t *TransferFunds) Result() account.TransactionalRecord {
	// Close conection not to block the thread.
	t.conn.Close()
	return t.record
}

// Do steps
func (t *TransferFunds) Do() (err error) {
	if err = t.validate(); err != nil {
		return err
	}

	if err = t.connectToServer(); err != nil {
		return err
	}

	if err = t.sendCreateRequst(); err != nil {
		return err
	}

	if err = t.printAccountInfo(); err != nil {
		return err
	}

	return nil
}

func (t *TransferFunds) validate() (err error) {
	if t.sourceAccount == "" {
		return errors.New("--source flag is required")
	}

	if t.destinationAccount == "" {
		return errors.New("--destination flag is required")
	}

	amount, err := strconv.Atoi(t.fundsAmount)
	if err != nil {
		return err
	}
	if amount <= 0 {
		return errors.New("--amount must be greater than 0")
	}

	return nil
}

func (t *TransferFunds) connectToServer() (err error) {
	t.conn, err = network.NewClient()
	if err != nil {
		log.WithError(err).Error("failed to connectToServer")
		return err
	}

	return nil
}

func (t *TransferFunds) sendCreateRequst() error {
	sess := account.AccountFactory{Client: t.conn.Bootstrap(t.ctx)}
	ca := sess.TransferFunds(t.ctx, func(p account.AccountFactory_transferFunds_Params) (err error) {
		err = p.SetSourceAccount(t.sourceAccount)
		if err != nil {
			log.WithError(err).Error("failed to SetSourceAccount")
			return err
		}
		err = p.SetDestinationAccount(t.destinationAccount)
		if err != nil {
			log.WithError(err).Error("failed to SetDestinationAccount")
			return err
		}
		amount, err := strconv.Atoi(t.fundsAmount)
		if err != nil {
			return err
		}
		p.SetAmount(uint64(amount))
		return err
	})
	resp, err := ca.Struct()
	if err != nil {
		log.WithError(err).Error("TransferFunds call failed")
		return err
	}

	t.record, err = resp.Record()
	if err != nil {
		log.WithError(err).Error("failed to decode response")
		return err
	}

	return nil
}

func (t *TransferFunds) printAccountInfo() error {
	sourceAccount, err := t.record.SourceAccount()
	if err != nil {
		log.WithError(err).Error("failed to decode SourceAccount field")
		return err
	}

	destinationAccount, err := t.record.DestinationAccount()
	if err != nil {
		log.WithError(err).Error("failed to decode SourceAccount field")
		return err
	}

	// Printer
	fmt.Println("--------------------------")
	fmt.Println("--- FUNDS TRANSFERRED! ---")
	fmt.Println("--------------------------")
	fmt.Printf("Source account number: %s\n", sourceAccount)
	fmt.Printf("Source account new balance: %d cents\n", t.record.SourceBalance())
	fmt.Printf("TRANSFERRED: %d cents\n", t.record.Amount())
	fmt.Printf("Destination account number: %s\n", destinationAccount)
	fmt.Printf("Destination account new balance: %d cents\n", t.record.DestinationBalance())
	fmt.Println("--------------------------")

	return nil
}
