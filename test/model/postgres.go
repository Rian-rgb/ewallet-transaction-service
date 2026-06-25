package model

import (
	"context"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/gorm"
)

type PostgresContainer struct {
	Ctx       context.Context
	Container *tcpostgres.PostgresContainer
	DB        *gorm.DB
}

func (p *PostgresContainer) Close() error {
	return p.Container.Terminate(p.Ctx)
}
