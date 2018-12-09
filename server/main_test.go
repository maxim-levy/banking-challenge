package main

import (
	"os"
	"testing"
)

// Can be used for global test setup.
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}
