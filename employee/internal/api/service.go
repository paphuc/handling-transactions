package api

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"handling-transactions/employee/pkg/http/request"
	"handling-transactions/protocol-buffers/employee"
	"handling-transactions/protocol-buffers/transaction"
)

type (
	ServiceI interface {
		InsertEmployee(ctx context.Context, e EmployeeRequest) (*string, error)
	}
)

func NewService(employeeClient employee.EmployeeClient, txClient transaction.TransactionClient) Service {
	return Service{
		gEmployeeClient:    employeeClient,
		gTransactionClient: txClient,
	}
}

func (s *Service) InsertEmployee(ctx context.Context, e EmployeeRequest) (*string, error) {
	//1. Begin Transaction
	txRequest := &transaction.BeginTxRequest{CorrelationID: e.Header.CorrelationID}
	txRes, err := s.gTransactionClient.BeginTx(ctx, txRequest)
	if err != nil {
		log.Println("Failed to begin transaction: ", err)
		return nil, err
	}

	log.Println("==========Tx res: ", txRes)
	employeeRequest := employee.InsertEmployeeRequest{
		CorrelationID: e.Header.CorrelationID,
		FirstName:     e.Body.FirstName,
		LastName:      e.Body.LastName,
		BeginTxRes:    txRes,
	}
	txInfo := &transaction.CommonTxDoActionRequest{
		CorrelationID: e.Header.CorrelationID,
		BeginTxRes:    txRes,
	}
	//2. Insert into Employee Table
	employeeRes, err := s.gEmployeeClient.InsertEmployee(ctx, &employeeRequest)
	if err != nil {
		s.gTransactionClient.Rollback(ctx, txInfo)
		return nil, err
	}

	//3. Make a request to assignment service to add assignment
	pReq := AssignmentRequest{
		Header: Header{
			CorrelationID: e.Header.CorrelationID,
		},
		Body: Assignment{
			ID:        e.Body.AssignmentID,
			StartDate: e.Body.StartDate,
			EndDate:   e.Body.EndDate,
		},
	}
	pBytes, err := json.Marshal(pReq)
	if err != nil {
		return nil, err
	}
	resP, err := request.Post(assignmentAPI, pBytes)
	if err != nil {
		s.gTransactionClient.Rollback(ctx, txInfo)
		log.Println("==== Add assignment failed, rollback DB")
		log.Println("Error: ", err)
		return nil, err
	}

	if resP.ID == nil {
		s.gTransactionClient.Rollback(ctx, txInfo)
		log.Println("==== Add assignment failed, rollback DB")
		return nil, errors.New("can not add assignment")
	}

	//4 Insert into Employee Detail Table
	employeeDetails := make([]*employee.EmployeeDetail, 0)
	for _, od := range e.Body.EmployeeDetails {
		employeeDetails = append(employeeDetails, &employee.EmployeeDetail{
			AssignmentID: e.Body.AssignmentID,
			Salary:       od.Salary,
			HomeAddress:  od.HomeAddress,
			Title:        od.Title,
			EmployeeID:   employeeRes.Id,
		})
	}
	_, err = s.gEmployeeClient.InsertEmployeeDetail(context.Background(), &employee.InsertEmployeeDetailRequest{
		CorrelationID:   e.Header.CorrelationID,
		EmployeeDetails: employeeDetails,
		BeginTxRes:      txRes,
	})
	if err != nil {
		s.gTransactionClient.Rollback(ctx, txInfo)
		return nil, err
	}
	s.gTransactionClient.Commit(ctx, txInfo)
	return &employeeRes.Id, nil
}
