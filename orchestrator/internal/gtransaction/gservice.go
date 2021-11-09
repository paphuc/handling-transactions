package gtransaction

import (
	"context"
	"log"

	"handling-transactions/protocol-buffers/transaction"
)

type GServiceI interface {
	BeginTx(context.Context, *transaction.BeginTxRequest) (*transaction.BeginTxResponse, error)
	Commit(context.Context, *transaction.CommonTxDoActionRequest) (*transaction.CommonTxResponse, error)
	Rollback(context.Context, *transaction.CommonTxDoActionRequest) (*transaction.CommonTxResponse, error)
}

func NewGService(txSrv ServiceI) *GService {
	return &GService{
		txSrv: txSrv,
	}
}

func (s *GService) BeginTx(ctx context.Context, in *transaction.BeginTxRequest) (*transaction.BeginTxResponse, error) {
	log.Println("======== Begin transaction ======")
	isRenew, txRandomID, err := s.txSrv.BeginTx(in.CorrelationID)
	return &transaction.BeginTxResponse{IsRenew: isRenew, TxRandomID: txRandomID}, err
}

func (s *GService) Commit(ctx context.Context, in *transaction.CommonTxDoActionRequest) (*transaction.CommonTxResponse, error) {
	ok := true
	if in.BeginTxRes.IsRenew {
		err := s.txSrv.Commit(in.CorrelationID, in.BeginTxRes.TxRandomID)
		if err != nil {
			ok = false
			return &transaction.CommonTxResponse{Ok: ok}, err
		}
	}
	return &transaction.CommonTxResponse{Ok: ok}, nil
}

func (s *GService) Rollback(ctx context.Context, in *transaction.CommonTxDoActionRequest) (*transaction.CommonTxResponse, error) {
	ok := true
	if in.BeginTxRes.IsRenew {
		err := s.txSrv.Rollback(in.CorrelationID, in.BeginTxRes.TxRandomID)
		if err != nil {
			ok = false
			return &transaction.CommonTxResponse{Ok: ok}, err
		}
	}
	return &transaction.CommonTxResponse{Ok: ok}, nil
}
