package gassignment

import (
	"handling-transactions/orchestrator/internal/gtransaction"
)

type (
	GService struct {
		repos RepositoryI
	}

	Repository struct {
		txSrv gtransaction.ServiceI
	}
)
