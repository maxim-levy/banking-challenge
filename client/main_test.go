package main

import (
	"os"
	"server/models"
	"testing"
)

// Can be used for global test setup.
func TestMain(m *testing.M) {
	// Start DB
	models.StartBoltDB("ledger_test.db")

	code := m.Run()
	os.Exit(code)
}
