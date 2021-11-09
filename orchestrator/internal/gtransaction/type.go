package gtransaction

import (
	"database/sql"
	"handling-transactions/orchestrator/internal/rediscache"
	"handling-transactions/orchestrator/internal/transactioncache"
)

type (
	Service struct {
		MapTx        map[string]*sql.Tx
		DB           *sql.DB
		redisService rediscache.ServiceI
		txCacheSrv   transactioncache.ServiceI
	}

	GService struct {
		txSrv ServiceI
	}
)
