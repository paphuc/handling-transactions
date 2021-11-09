package route

import (
	"handling-transactions/protocol-buffers/employee"
	"handling-transactions/protocol-buffers/transaction"
	"net/http"

	"handling-transactions/employee/internal/api"
	"handling-transactions/employee/pkg/http/middleware"
	"handling-transactions/employee/pkg/http/router"

	"github.com/gorilla/mux"
)

func NewRouter(employeeC employee.EmployeeClient, txC transaction.TransactionClient) (http.Handler, error) {
	r := mux.NewRouter()
	apiHandler, err := newAPIHandler(employeeC, txC)
	if err != nil {
		return nil, err
	}

	routes := []router.Route{}
	routes = append(routes, apiHandler.Routes()...)

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

func newAPIHandler(employeeC employee.EmployeeClient, txC transaction.TransactionClient) (*api.Handler, error) {
	srv := api.NewService(employeeC, txC)
	handler := api.NewHandler(&srv)
	return handler, nil
}
