package network

import (
	"fmt"
	"net"
	"os"

	"github.com/apex/log"
	"zombiezen.com/go/capnproto2/rpc"
)

// NewClient tries to connect to the banking server.
// returns a rpc connection if successfull.
func NewClient() (*rpc.Conn, error) {
	addr := os.Getenv("CLIENT_SERVER_ADDR")
	port := os.Getenv("CLIENT_SERVER_PORT")
	// Set defaults
	if addr == "" {
		addr = "127.0.0.1"
	}
	if port == "" {
		port = "8080"
	}

	dial, err := net.Dial("tcp", fmt.Sprintf("%s:%s", addr, port))
	if err != nil {
		log.WithError(err).Errorf("can't connect to server at %s", fmt.Sprintf("%s:%s", addr, port))
		return nil, err
	}

	// Create a rpc connection that we can use to comunicate with the server.
	conn := rpc.NewConn(rpc.StreamTransport(dial))
	// TODO: defer conn.Close()

	return conn, nil
}
