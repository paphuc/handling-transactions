package gemployee

import (
	"context"

	"handling-transactions/protocol-buffers/employee"
)

type GServiceI interface {
	InsertEmployee(ctx context.Context, in *employee.InsertEmployeeRequest) (*employee.InsertEmployeeResponse, error)
	InsertEmployeeDetail(context.Context, *employee.InsertEmployeeDetailRequest) (*employee.InsertEmployeeDetailResponse, error)
}

func NewGService(repos RepositoryI) *GService {
	return &GService{
		repos: repos,
	}
}

func (s *GService) InsertEmployee(ctx context.Context, o *employee.InsertEmployeeRequest) (*employee.InsertEmployeeResponse, error) {
	return s.repos.InsertEmployee(ctx, o)
}

func (s *GService) InsertEmployeeDetail(ctx context.Context, o *employee.InsertEmployeeDetailRequest) (*employee.InsertEmployeeDetailResponse, error) {
	return s.repos.InsertEmployeeDetail(ctx, o)
}
