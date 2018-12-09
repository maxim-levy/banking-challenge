package account

import (
	"protos/account"
	createaccount "server/services/create-account"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/server"
)

func (af accountFactory) CreateAccount(call account.AccountFactory_createAccount) error {
	log.Info("Called: CreateAccount")
	server.Ack(call.Options)

	s := createaccount.NewCreateAccount(call.Params.InitialBalance())
	if err := s.Do(); err != nil {
		log.WithError(err).Error("failed to create new account")
		return err
	}

	// Get response
	accountNumber, balance := s.Result()

	// Build response
	account, err := call.Results.NewAccount()
	if err != nil {
		log.WithError(err).Error("failed to get Account struct")
		return err
	}
	account.SetAccountNumber(accountNumber)
	account.SetBalance(balance)
	call.Results.SetAccount(account)

	return nil
}
