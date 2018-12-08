package main

import (
	"net"
	"os"
	"protos/account"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"zombiezen.com/go/capnproto2/rpc"
	"zombiezen.com/go/capnproto2/server"
)

// accountFactory is a local implementation of AccountFactory.
type accountFactory struct{}

func (af accountFactory) CreateAccount(call account.AccountFactory_createAccount) error {
	server.Ack(call.Options)
	log.Info("CreateAccount called")
	return nil
}

func (af accountFactory) DeleteAccount(call account.AccountFactory_deleteAccount) error {
	log.Info("DeleteAccount called")
	server.Ack(call.Options)

	v, _ := call.Params.SourceAccount()
	log.Infof("%+v", v)
	return nil
}

func startServer(c net.Conn) error {
	// Create a new locally implemented HashFactory.
	srv := account.AccountFactory_ServerToClient(accountFactory{})
	// Listen for calls, using the HashFactory as the bootstrap interface.
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(srv.Client))
	if err := conn.Wait(); err != nil {
		log.WithError(err).Error("conn wait error")
		return err
	}
	return nil
}

func main() {
	log.SetHandler(text.New(os.Stderr))
	log.Info("starting server")
	addr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		log.WithError(err).Fatal("failed")
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.WithError(err).Fatal("failed")
	}

	// Loop to make sure we can keep accepting connections
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.WithError(err).Fatal("failed")
		}

		if err := startServer(conn); err != nil {
			log.WithError(err).Fatal("failed")
		}
	}
}
