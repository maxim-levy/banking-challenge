package methods

import (
	"net"
	"protos/account"
	accountMethods "server/methods/account"

	"zombiezen.com/go/capnproto2/rpc"
)

// Register methods and start listening for calls.
func Register(c net.Conn) error {
	// Create a new locally implemented accountFactory.
	accountSrv := account.AccountFactory_ServerToClient(accountMethods.AccountFactory{})

	// Listen for calls, using the accountFactory as the bootstrap interface.
	conn := rpc.NewConn(
		rpc.StreamTransport(c),
		rpc.MainInterface(accountSrv.Client),
	)
	defer conn.Wait()

	return nil
}
