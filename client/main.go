package main

import (
	createaccount "client/actions/create-account"
	deleteaccount "client/actions/delete-account"
	transferfunds "client/actions/transfer-funds"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/urfave/cli"
)

func main() {
	// Configure loging
	log.SetHandler(text.New(os.Stderr))

	// Setup CLI
	app := cli.NewApp()
	app.Name = "crypto-banking"
	app.Usage = "Crypto-banking is a system that consist of a client and server (bank) that allows you to do basic finintial operations."
	app.Version = "1.0.0"
	app.Commands = []cli.Command{
		{
			Name:    "create-account",
			Aliases: []string{"ca"},
			Usage:   "Register an new account",
			Action: func(c *cli.Context) error {
				s := createaccount.NewCreateAccount(
					c.String("initial-balance"),
				)
				return s.Do()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "initial-balance, ib",
					Value: "0",
					Usage: "Initial balance for the new account, in cents",
				},
			},
		},
		{
			Name:    "delete-account",
			Aliases: []string{"da"},
			Usage:   "Delete an existing account",
			Action: func(c *cli.Context) error {
				s := deleteaccount.NewDeleteAccount(
					c.String("account-number"),
				)
				return s.Do()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "account-number, an",
					Value: "",
					Usage: "Number of account to be deleted",
				},
			},
		},
		{
			Name:    "transfer-funds",
			Aliases: []string{"tf"},
			Usage:   "Transfer funds from one account to another",
			Action: func(c *cli.Context) error {
				s := transferfunds.NewTransferFunds(
					c.String("source"),
					c.String("destination"),
					c.String("amount"),
				)
				return s.Do()
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "source",
					Value: "",
					Usage: "Source account number",
				},
				cli.StringFlag{
					Name:  "destination",
					Value: "",
					Usage: "Destination account number",
				},
				cli.StringFlag{
					Name:  "amount",
					Value: "0",
					Usage: "Amount to transfer from source to destination in cents",
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Sorry, Something went wrong while running crypto-banking application")
	}
}
