package main

import (
	"os"
	"server/models"
	"server/network"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

func main() {
	log.SetHandler(text.New(os.Stderr))

	// Start DB
	// Will thow with fatal if there is a problem.
	models.StartBoltDB("ledger.db")

	// Start the server and listen for connections.
	// The server will also prevent the program from exiting.
	network.StartServer()
}
