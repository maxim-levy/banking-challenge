package main

import (
	"fmt"
	"net"
	"os"
	"server/methods"
	"server/models"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func main() {
	log.SetHandler(text.New(os.Stderr))

	// Start DB
	// Will thow with fatal if there is a problem.
	models.StartBoltDB()

	// Get server config
	address := os.Getenv("SERVER_ADDR")
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	// Open ports and start listening for connections.
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.WithError(err).Fatal("failed")
	}
	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.WithError(err).Fatal("failed")
	}
	log.Info("server started")

	// Loop to make sure we can keep accepting connections
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.WithError(err).Fatal("failed")
		}

		// Register method listeners
		if err := methods.StartServer(conn); err != nil {
			log.WithError(err).Fatal("failed")
		}
	}
}
