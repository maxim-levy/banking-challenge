package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartBoltDB(t *testing.T) {
	// Will throw fatal on error
	StartBoltDB("ledger_test.db")
}

func TestGetDB(t *testing.T) {
	db := GetDB()
	assert.NotNil(t, db)
	assert.Equal(t, db.Path(), "ledger_test.db")
}
