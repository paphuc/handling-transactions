package rediscache

import (
	"handling-transactions/orchestrator/internal/transactioncache"

	redis "github.com/go-redis/redis/v8"
)

type (
	Service struct {
		redisClient *redis.Client
		txCacheSrv  transactioncache.ServiceI
	}
)
