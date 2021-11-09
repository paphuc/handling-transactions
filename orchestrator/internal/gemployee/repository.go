package gemployee

import (
	"context"
	"log"

	"handling-transactions/orchestrator/internal/gtransaction"
	"handling-transactions/protocol-buffers/employee"

	"github.com/google/uuid"
)

type RepositoryI interface {
	InsertEmployee(ctx context.Context, in *employee.InsertEmployeeRequest) (*employee.InsertEmployeeResponse, error)
	InsertEmployeeDetail(context.Context, *employee.InsertEmployeeDetailRequest) (*employee.InsertEmployeeDetailResponse, error)
}

func NewRepository(txSrv gtransaction.ServiceI) *Repository {
	return &Repository{
		txSrv: txSrv,
	}
}

func (r *Repository) InsertEmployee(ctx context.Context, o *employee.InsertEmployeeRequest) (*employee.InsertEmployeeResponse, error) {
	log.Println("======Start InsertEmployee: ", o)
	tx, err := r.txSrv.GetTxByCorrelationID(o.CorrelationID, o.BeginTxRes.TxRandomID)
	if err != nil {
		log.Println("GetTxByCorrelationID Err: ", err)
		return nil, err
	}
	id := uuid.New().String()
	queryString := `INSERT INTO employee(id, first_name, last_name) VALUES($1, $2, $3)`
	rs, err := tx.Exec(queryString, id, o.FirstName, o.LastName)
	if val, err := rs.RowsAffected(); err == nil && val > 0 {
		log.Println("====== InsertEmployee completed ======: ", id)
		return &employee.InsertEmployeeResponse{Id: id}, nil
	}
	return nil, err
}

func (r *Repository) InsertEmployeeDetail(ctx context.Context, o *employee.InsertEmployeeDetailRequest) (*employee.InsertEmployeeDetailResponse, error) {
	log.Println("======Start InsertEmployeeDetail: ", o)
	tx, err := r.txSrv.GetTxByCorrelationID(o.CorrelationID, o.BeginTxRes.TxRandomID)
	if err != nil {
		return nil, err
	}
	numRowAffected := int64(0)
	for _, od := range o.EmployeeDetails {
		queryString := `INSERT INTO employee_detail(employee_id, assignment_id, salary, home_address, title) VALUES($1, $2, $3, $4, $5)`
		rs, err := tx.Exec(queryString, od.EmployeeID, od.AssignmentID, od.Salary, od.HomeAddress, od.Title)
		if err != nil {
			return nil, err
		}
		rowAffected, err := rs.RowsAffected()
		if err != nil {
			return nil, err
		}
		numRowAffected += rowAffected
	}
	log.Println("====== InsertEmployeeDetail Completed: ", numRowAffected)
	return &employee.InsertEmployeeDetailResponse{RowAffected: int32(numRowAffected)}, nil
}
