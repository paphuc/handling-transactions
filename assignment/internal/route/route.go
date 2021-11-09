package route

import (
	"net/http"

	"handling-transactions/assignment/internal/api"
	"handling-transactions/assignment/pkg/http/middleware"
	"handling-transactions/assignment/pkg/http/router"
	"handling-transactions/protocol-buffers/assignment"
	"handling-transactions/protocol-buffers/transaction"

	"github.com/gorilla/mux"
)

func NewRouter(assignmentC assignment.AssignmentClient, txC transaction.TransactionClient) (http.Handler, error) {
	r := mux.NewRouter()
	newAPIHandler, err := newAPIHandler(assignmentC, txC)
	if err != nil {
		return nil, err
	}

	routes := []router.Route{}
	routes = append(routes, newAPIHandler.Routes()...)

	//Routes
	for _, rt := range routes {
		var h http.Handler
		h = rt.Handler
		for i := len(rt.Middlewares) - 1; i >= 0; i-- {
			h = rt.Middlewares[i](h)
		}
		r.Path(rt.Path).Methods(rt.Method).Handler(h)
	}

	return middleware.CORS(r), nil
}

func newAPIHandler(assignmentC assignment.AssignmentClient, txC transaction.TransactionClient) (*api.Handler, error) {
	srv := api.NewService(assignmentC, txC)
	handler := api.NewHandler(&srv)
	return handler, nil
}
