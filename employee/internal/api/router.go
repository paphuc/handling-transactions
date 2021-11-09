package api

import (
	"net/http"

	"handling-transactions/employee/pkg/http/router"
)

func (h *Handler) Routes() []router.Route {
	return []router.Route{
		{
			Path:    "/api/v1/employee",
			Method:  http.MethodPost,
			Handler: h.InsertEmployee,
		},
	}
}
