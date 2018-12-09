package network

import (
	"net"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/rpc"
)

// NewClient tries to connect to the banking server.
// returns a rpc connection if successfull.
func NewClient() (*rpc.Conn, error) {
	addr := "127.0.0.1:8080"

	dial, err := net.Dial("tcp", addr)
	if err != nil {
		log.WithError(err).Errorf("can't connect to server at %s", addr)
		return nil, err
	}

	// Create a rpc connection that we can use to comunicate with the server.
	conn := rpc.NewConn(rpc.StreamTransport(dial))
	// TODO: defer conn.Close()

	return conn, nil
}
