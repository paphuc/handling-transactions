package api

import (
	"net/http"

	"handling-transactions/assignment/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/assignment",
			Method:  http.MethodPost,
			Handler: h.AddAssignment,
		},
	}
}
