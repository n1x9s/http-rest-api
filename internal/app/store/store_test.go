package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost user=postgres password=15931 dbname=goapitest sslmode=disable"
	}

	os.Exit(m.Run())
}
