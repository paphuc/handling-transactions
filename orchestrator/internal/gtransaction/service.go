package gtransaction

import (
	"context"
	"database/sql"
	"errors"
	"handling-transactions/orchestrator/internal/rediscache"
	"handling-transactions/orchestrator/internal/transactioncache"
	"handling-transactions/orchestrator/pkg/constx"

	"github.com/google/uuid"
)

type ServiceI interface {
	GetTxByCorrelationID(correlationID string, txRandomID string) (*sql.Tx, error)
	BeginTx(correlationID string) (bool, string, error)
	Commit(correlationID string, txRandomID string) error
	Rollback(correlationID string, txRandomID string) error
}

func NewTransactionSrv(db *sql.DB, redisService rediscache.ServiceI, txCacheSrv transactioncache.ServiceI) *Service {
	return &Service{
		DB:           db,
		txCacheSrv:   txCacheSrv,
		redisService: redisService,
	}
}

func (t *Service) GetTxByCorrelationID(correlationID string, txRandomID string) (*sql.Tx, error) {
	if val, ok := t.txCacheSrv.Get(correlationID + "_" + txRandomID); ok {
		return val, nil
	}

	return nil, errors.New("Couldn't found the transaction")
}

func (t *Service) BeginTx(correlationID string) (bool, string, error) {
	isRenew := true
	txRandomID := uuid.New().String()

	tx, err := t.DB.Begin()
	if err != nil {
		return isRenew, txRandomID, err
	}

	//Set to local cache and redis cache
	t.txCacheSrv.Set(correlationID+"_"+txRandomID, tx)
	err = t.redisService.Set(context.Background(), correlationID+"_"+txRandomID, constx.False, constx.ExpiredMinutes)
	if err != nil {
		return isRenew, txRandomID, err
	}
	go t.redisService.SubscribeTransaction(context.Background(), correlationID, txRandomID, tx)

	return isRenew, txRandomID, nil
}

func (t *Service) Commit(correlationID string, txRandomID string) error {
	topicKey := correlationID + "_" + txRandomID
	return t.redisService.PublishTransaction(context.Background(), topicKey, constx.Commit)
}

func (t *Service) Rollback(correlationID string, txRandomID string) error {
	topicKey := correlationID + "_" + txRandomID
	return t.redisService.PublishTransaction(context.Background(), topicKey, constx.RollBack)
}
