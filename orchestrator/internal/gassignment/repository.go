package gassignment

import (
	"context"
	"log"
	"time"

	"handling-transactions/orchestrator/internal/gtransaction"
	"handling-transactions/protocol-buffers/assignment"
)

type RepositoryI interface {
	InsertAssignment(context.Context, *assignment.InsertAssignmentRequest) (*assignment.InsertAssignmentResponse, error)
}

func NewRepository(txSrv gtransaction.ServiceI) *Repository {
	return &Repository{
		txSrv: txSrv,
	}
}

func (r *Repository) InsertAssignment(ctx context.Context, p *assignment.InsertAssignmentRequest) (*assignment.InsertAssignmentResponse, error) {
	log.Println("======Start InsertAssignment: ", p)
	tx, err := r.txSrv.GetTxByCorrelationID(p.CorrelationID, p.BeginTxRes.TxRandomID)
	if err != nil {
		return nil, err
	}

	layout := "2006-01-02T15:04:05.000Z"
	startDate, _ := time.Parse(layout, p.StartDate)
	endDate, _ := time.Parse(layout, p.EndDate)
	queryString := `INSERT INTO assignment(id, start_date, end_date) VALUES($1, $2, $3)`
	rs, err := tx.Exec(queryString, p.ID, startDate, endDate)
	if err != nil {
		log.Println("====== InsertAssignment failed: ", err)
		return &assignment.InsertAssignmentResponse{RowAffected: int32(0)}, err
	}
	val, err := rs.RowsAffected()
	if err == nil && val > 0 {
		log.Println("====== InsertAssignment completed: ", val)
		return &assignment.InsertAssignmentResponse{RowAffected: int32(val)}, nil
	}
	log.Println("====== InsertAssignment failed: ", err)
	return &assignment.InsertAssignmentResponse{RowAffected: int32(0)}, err
}
