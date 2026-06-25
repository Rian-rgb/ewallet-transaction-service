package helper

import (
	"context"
	"ewallet-transaction/internal/domain/transaction"
	"ewallet-transaction/test/model"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var TestDB *gorm.DB

func SetupPostgresContainer() (*model.PostgresContainer, error) {
	ctx := context.Background()
	container, err := tcpostgres.Run(
		ctx,
		"postgres:16-alpine",
		tcpostgres.WithDatabase("testdb"),
		tcpostgres.WithUsername("test"),
		tcpostgres.WithPassword("12345"),
		tcpostgres.BasicWaitStrategies(),
	)

	if err != nil {
		return nil, err
	}

	dsn, err := container.ConnectionString(
		ctx,
		"sslmode=disable",
	)

	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(
		gormpostgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	if err = CreateTransactionTypes(db); err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(
		&transaction.Entity{},
	); err != nil {
		return nil, err
	}

	return &model.PostgresContainer{
		Ctx:       ctx,
		Container: container,
		DB:        db,
	}, nil
}
