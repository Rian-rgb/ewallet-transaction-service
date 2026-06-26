package infra

import (
	"ewallet-transaction/infra/grpc"

	"github.com/Rian-rgb/ewallet-common-lib/redis"
	"gorm.io/gorm"
)

type AppDependencies struct {
	PostgresDB   *gorm.DB
	RedisRepo    *redis.RedisRepository
	GrpcRegistry *grpc.ConnRegistry
}
