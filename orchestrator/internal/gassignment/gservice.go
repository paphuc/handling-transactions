package gassignment

import (
	"context"

	"handling-transactions/protocol-buffers/assignment"
)

type GServiceI interface {
	UpdateAssignment(context.Context, *assignment.InsertAssignmentRequest) (*assignment.InsertAssignmentResponse, error)
}

func NewGService(repos RepositoryI) *GService {
	return &GService{
		repos: repos,
	}
}

func (s *GService) InsertAssignment(ctx context.Context, p *assignment.InsertAssignmentRequest) (*assignment.InsertAssignmentResponse, error) {
	return s.repos.InsertAssignment(ctx, p)
}
