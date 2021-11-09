package api

import (
	"handling-transactions/protocol-buffers/employee"
	"handling-transactions/protocol-buffers/transaction"
)

const (
	assignmentAPI = "http://localhost:8082/api/v1/assignment"
)

type (
	Header struct {
		CorrelationID string `json:"correlationID"`
	}

	EmployeeDetail struct {
		AssignmentID int32  `json:"assignmentID"`
		Salary       string `json:"salary"`
		HomeAddress  string `json:"homeAddress"`
		Title        string `json:"title"`
	}

	Employee struct {
		ID              *string           `json:"id"`
		AssignmentID    int32             `json:"assignmentID"`
		FirstName       string            `json:"firstName"`
		LastName        string            `json:"lastName"`
		StartDate       string            `json:"startDate"`
		EndDate         string            `json:"endDate"`
		EmployeeDetails []*EmployeeDetail `json:"employeeDetails"`
	}

	EmployeeRequest struct {
		Header Header   `json:"header"`
		Body   Employee `json:"body"`
	}

	AssignmentRequest struct {
		Header Header     `json:"header"`
		Body   Assignment `json:"body"`
	}

	Assignment struct {
		ID        int32  `json:"id"`
		StartDate string `json:"startDate"`
		EndDate   string `json:"endDate"`
	}

	Handler struct {
		srv ServiceI
	}

	Service struct {
		gEmployeeClient    employee.EmployeeClient
		gTransactionClient transaction.TransactionClient
	}

	GService struct {
		employeeClient employee.EmployeeClient
	}
)
