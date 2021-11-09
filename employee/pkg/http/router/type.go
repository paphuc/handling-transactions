package router

import "net/http"

type (
	Middleware = func(http.Handler) http.Handler

	Route struct {
		Path        string
		Method      string
		Handler     http.HandlerFunc
		Middlewares []Middleware
	}
)
