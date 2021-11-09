package transactioncache

import (
	"database/sql"
	"sync"
)

type ServiceI interface {
	Set(key string, value *sql.Tx)
	Get(key string) (*sql.Tx, bool)
	Remove(key string)
}

func NewTransactionCacheSrv() *Service {
	var mapTx = sync.Map{}
	return &Service{
		mapTx: mapTx,
	}
}

//TODO Implement expiration time for a value

func (s *Service) Set(key string, value *sql.Tx) {
	s.mapTx.Store(key, value)
}

func (s *Service) Get(key string) (*sql.Tx, bool) {
	value, ok := s.mapTx.Load(key)
	if !ok {
		return nil, ok
	}
	return value.(*sql.Tx), ok
}

func (s *Service) Remove(key string) {
	s.mapTx.Delete(key)
}
