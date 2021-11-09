package api

import (
	"handling-transactions/protocol-buffers/assignment"
	"handling-transactions/protocol-buffers/transaction"
)

type (
	Header struct {
		CorrelationID string `json:"correlationID"`
	}

	Assignment struct {
		ID        *int   `json:"id"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	InsertAssignmentRequest struct {
		Header Header     `json:"header"`
		Body   Assignment `json:"body"`
	}

	Handler struct {
		srv ServiceI
	}

	Service struct {
		gAssignmentClient  assignment.AssignmentClient
		gTransactionClient transaction.TransactionClient
	}

	GService struct {
		assignmentClient assignment.AssignmentClient
	}
)
