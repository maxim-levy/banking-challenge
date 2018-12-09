package account

import (
	"protos/account"
	deleteaccount "server/services/delete-account"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/server"
)

func (af AccountFactory) DeleteAccount(call account.AccountFactory_deleteAccount) error {
	log.Info("Called: DeleteAccount")
	server.Ack(call.Options)

	// Get params
	accountNumber, err := call.Params.AccountNumber()
	if err != nil {
		log.WithError(err).Error("failed to decode accountNumber")
		return err
	}

	s := deleteaccount.NewDeleteAccount(accountNumber)
	if err := s.Do(); err != nil {
		log.WithError(err).Error("failed to delete account")
		return err
	}

	// Build response
	call.Results.SetSuccess(true)

	return nil
}
