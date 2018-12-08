package main

import (
	"context"
	"net"
	"os"
	"protos/account"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"zombiezen.com/go/capnproto2/rpc"
)

func client(ctx context.Context, c net.Conn) error {
	// Create a connection that we can use to get the HashFactory.
	conn := rpc.NewConn(rpc.StreamTransport(c))
	defer conn.Close()
	// Get the "bootstrap" interface.  This is the capability set with
	// rpc.MainInterface on the remote side.
	af := account.AccountFactory{Client: conn.Bootstrap(ctx)}

	ca := af.CreateAccount(ctx, func(p account.AccountFactory_createAccount_Params) error {
		p.SetInitialBalance(1000)
		return nil
	})
	result1, err := ca.Struct()
	if err != nil {
		log.WithError(err).Error("CreateAccount call failed")
		return err
	}

	log.Infof("%+v", result1.String())

	da := af.DeleteAccount(ctx, func(p account.AccountFactory_deleteAccount_Params) error {
		if err := p.SetSourceAccount("test"); err != nil {
			log.WithError(err).Error("failed to SetSourceAccount")
			return err
		}
		return nil
	})
	result2, err := da.Struct()
	if err != nil {
		log.WithError(err).Error("DeleteAccount call failed")
		return err
	}

	log.Infof("%+v", result2.String())

	return nil
}

func main() {
	log.SetHandler(text.New(os.Stderr))
	log.Info("starting client")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.WithError(err).Fatal("can't connect to server")
	}

	if err := client(context.Background(), conn); err != nil {
		log.WithError(err).Fatal("client has failed")
	}
}
