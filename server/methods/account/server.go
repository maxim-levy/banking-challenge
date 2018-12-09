package account

import (
	"net"
	"protos/account"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/rpc"
)

// accountFactory is a local implementation of AccountFactory.
type accountFactory struct{}

// StartServer starts listening for calls.
func StartServer(c net.Conn) error {
	// Create a new locally implemented accountFactory.
	srv := account.AccountFactory_ServerToClient(accountFactory{})
	// Listen for calls, using the accountFactory as the bootstrap interface.
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(srv.Client))
	if err := conn.Wait(); err != nil {
		log.WithError(err).Error("conn wait error")
		return err
	}

	return nil
}
