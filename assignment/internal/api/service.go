package api

import (
	"context"
	"log"

	"handling-transactions/protocol-buffers/assignment"
	"handling-transactions/protocol-buffers/transaction"
)

type (
	ServiceI interface {
		InsertAssignment(ctx context.Context, a InsertAssignmentRequest) (*int32, error)
	}
)

func NewService(assignmentClient assignment.AssignmentClient, txClient transaction.TransactionClient) Service {
	return Service{
		gAssignmentClient:  assignmentClient,
		gTransactionClient: txClient,
	}
}

func (s *Service) InsertAssignment(ctx context.Context, a InsertAssignmentRequest) (*int32, error) {
	//1. Begin Transaction
	txRequest := &transaction.BeginTxRequest{CorrelationID: a.Header.CorrelationID}
	txRes, err := s.gTransactionClient.BeginTx(ctx, txRequest)
	if err != nil {
		return nil, err
	}
	assignmentRequest := assignment.InsertAssignmentRequest{
		CorrelationID: a.Header.CorrelationID,
		ID:            int32(*a.Body.ID),
		StartDate:     a.Body.StartDate,
		EndDate:       a.Body.EndDate,
		BeginTxRes:    txRes,
	}
	//2. Insert To Assignment Table
	assignmentRes, err := s.gAssignmentClient.InsertAssignment(ctx, &assignmentRequest)

	txInfo := &transaction.CommonTxDoActionRequest{
		CorrelationID: a.Header.CorrelationID,
		BeginTxRes:    txRes,
	}
	if err != nil {
		s.gTransactionClient.Rollback(ctx, txInfo)
		log.Println("======Rollback DB======")
		return nil, err
	}
	s.gTransactionClient.Commit(ctx, txInfo)
	return &assignmentRes.RowAffected, nil
}
