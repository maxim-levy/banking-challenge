package transfer

import (
	"protos/transfer"
	transferfunds "server/services/transfer-funds"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/server"
)

func (tf TransferFactory) TransferFunds(call transfer.TransferFactory_transferFunds) error {
	log.Info("Called: TransferFunds")
	server.Ack(call.Options)

	// Get params
	sourceAccount, err := call.Params.SourceAccount()
	if err != nil {
		log.WithError(err).Error("failed to decode sourceAccount")
		return err
	}

	destinationAccount, err := call.Params.DestinationAccount()
	if err != nil {
		log.WithError(err).Error("failed to decode destinationAccount")
		return err
	}

	s := transferfunds.NewTransferFunds(
		sourceAccount,
		destinationAccount,
		call.Params.Amount(),
	)
	if err := s.Do(); err != nil {
		log.WithError(err).Error("failed to transfer funds")
		return err
	}

	// Get response
	sourceAccountNumber,
		sourceAccountBalance,
		destinationAccountNumber,
		destinationAccountBalance := s.Result()

	// Build response
	record, err := call.Results.NewRecord()
	if err != nil {
		log.WithError(err).Error("failed to get Record struct")
		return err
	}

	record.SetSourceAccount(sourceAccountNumber)
	record.SetSourceBalance(sourceAccountBalance)
	record.SetDestinationAccount(destinationAccountNumber)
	record.SetDestinationBalance(destinationAccountBalance)
	record.SetAmount(call.Params.Amount())
	call.Results.SetRecord(record)

	return nil
}
