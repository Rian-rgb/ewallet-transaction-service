package test_test

import (
	"ewallet-transaction/test/helper"
	"ewallet-transaction/test/model"
	"os"
	"testing"
)

var PG *model.PostgresContainer

func TestMain(m *testing.M) {
	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")
	var err error

	PG, err = helper.SetupPostgresContainer()
	if err != nil {
		panic(err)
	}

	helper.TestDB = PG.DB

	code := m.Run()

	_ = PG.Close()

	os.Exit(code)
}
