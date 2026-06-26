package repository_test

import (
	"ewallet-transaction/infra"
	"ewallet-transaction/internal/domain/transaction"
	"github.com/Rian-rgb/ewallet-common-lib/database"
	"github.com/Rian-rgb/ewallet-common-lib/logger"
	"gorm.io/gorm"
	"os"
	"testing"
)

var globalTestDB *gorm.DB

func TestMain(m *testing.M) {

	// load log
	infra.InitLogger()

	os.Setenv("TESTCONTAINERS_RYUK_DISABLED", "true")

	testEnv, err := database.SetupPostgresContainer()
	if err != nil {
		logger.Error("failed to turn on postgres test container: %v", err)
	}

	defer testEnv.Container.Terminate(testEnv.Ctx)

	globalTestDB = testEnv.DB

	if err = createTransactionTypes(globalTestDB); err != nil {
		logger.Error("failed to create transaction data types: %v", err)
	}

	err = globalTestDB.AutoMigrate(&transaction.Entity{})
	if err != nil {
		logger.Error("failed to migration database testing: %v", err)
	}

	exitCode := m.Run()

	os.Exit(exitCode)
}
