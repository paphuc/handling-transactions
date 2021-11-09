package rediscache

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"handling-transactions/orchestrator/internal/transactioncache"
	"handling-transactions/orchestrator/pkg/constx"
	"handling-transactions/orchestrator/pkg/msgx"

	"github.com/go-redis/redis/v8"
)

type ServiceI interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Del(ctx context.Context, keys ...string) error
	Keys(ctx context.Context, pattern string) ([]string, error)
	SubscribeTransaction(ctx context.Context, correlationID string, txRandomID string, tx *sql.Tx)
	PublishTransaction(ctx context.Context, topic string, action string) error
}

func NewRedisSrv(redisClient *redis.Client, txCacheSrv transactioncache.ServiceI) *Service {
	return &Service{
		redisClient: redisClient,
		txCacheSrv:  txCacheSrv,
	}
}

func (s *Service) SubscribeTransaction(ctx context.Context, correlationID string, txRandomID string, tx *sql.Tx) {
	topicKey := correlationID + "_" + txRandomID
	expirationTime := constx.ExpiredMinutes

	// Listen message Commit, Rollback local transaction
	tTx := s.redisClient.Subscribe(ctx, correlationID)
	go func() {
		for {
			//TODO read time alive of topic from config file
			ctx, cancel := context.WithTimeout(context.Background(), expirationTime)
			defer cancel()
			msg, err := tTx.ReceiveMessage(ctx)
			if err != nil {
				s.txCacheSrv.Remove(topicKey)
				s.redisClient.Del(context.Background(), topicKey)
				tx.Rollback()
				tTx.Close()
				return
			}
			var txInfo msgx.TransactionInfo
			if err := json.Unmarshal([]byte(msg.Payload), &txInfo); err != nil {
				s.txCacheSrv.Remove(topicKey)
				s.redisClient.Del(context.Background(), topicKey)
				tx.Rollback()
				tTx.Close()
				return
			}
			switch txInfo.Action {
			case constx.Commit:
				s.txCacheSrv.Remove(topicKey)
				s.redisClient.Del(context.Background(), topicKey)
				tx.Commit()
				tTx.Close()
				return
			case constx.RollBack:
				s.txCacheSrv.Remove(topicKey)
				s.redisClient.Del(context.Background(), topicKey)
				tx.Rollback()
				tTx.Close()
				return
			}
		}
	}()
	// Listen message change state on redis cache
	tRd := s.redisClient.Subscribe(ctx, topicKey)
	go func() {
		for {
			_, cancel := context.WithTimeout(context.Background(), expirationTime)
			defer cancel()
			msg, err := tRd.ReceiveMessage(ctx)
			if err != nil {
				tRd.Close()
				s.PublishTransaction(context.Background(), correlationID, constx.RollBack)
				return
			}
			var txInfo msgx.TransactionInfo
			if err := json.Unmarshal([]byte(msg.Payload), &txInfo); err != nil {
				tRd.Close()
				s.PublishTransaction(context.Background(), correlationID, constx.RollBack)
				return
			}
			switch txInfo.Action {
			case constx.Commit:
				err := s.Set(context.Background(), topicKey, constx.True, expirationTime)
				if err != nil {
					tRd.Close()
					s.PublishTransaction(context.Background(), correlationID, constx.RollBack)
					return
				}
				rs, err := s.Keys(context.Background(), fmt.Sprintf("%s_*", correlationID))
				if err != nil {
					tRd.Close()
					s.PublishTransaction(context.Background(), correlationID, constx.RollBack)
					return
				}
				count := 0
				for _, k := range rs {
					v, err := s.Get(context.Background(), k)
					if err != nil {
						tRd.Close()
						s.PublishTransaction(context.Background(), correlationID, constx.RollBack)
						return
					}
					if v == constx.True {
						count++
					}
				}
				if count == len(rs) {
					s.PublishTransaction(context.Background(), correlationID, constx.Commit)
				}
				tRd.Close()
				return
			case constx.RollBack:
				s.PublishTransaction(context.Background(), correlationID, constx.RollBack)
				tRd.Close()
				return
			}
		}
	}()
}

func (s *Service) PublishTransaction(ctx context.Context, topic string, action string) error {
	txInfo := &msgx.TransactionInfo{CorrelationID: topic, Action: action}
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return err
	}
	//TODO publish retry should be implement
	if err := s.redisClient.Publish(ctx, topic, txInfoBytes).Err(); err != nil {
		//TODO implement dead letter queue
		return err
	}
	return nil
}
