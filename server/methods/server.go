package methods

import (
	"net"
	"protos/account"
	"protos/transfer"
	accountMethods "server/methods/account"
	transferMethods "server/methods/transfer"

	"zombiezen.com/go/capnproto2/rpc"
)

// StartServer starts listening for calls.
func StartServer(c net.Conn) error {
	// Create a new locally implemented accountFactory.
	accountSrv := account.AccountFactory_ServerToClient(accountMethods.AccountFactory{})
	// Create a new locally implemented transferFactory.
	transferSrv := transfer.TransferFactory_ServerToClient(transferMethods.TransferFactory{})

	// Listen for calls, using the accountFactory as the bootstrap interface.
	conn := rpc.NewConn(
		rpc.StreamTransport(c),
		rpc.MainInterface(accountSrv.Client),
		rpc.MainInterface(transferSrv.Client),
	)
	defer conn.Wait()

	return nil
}
