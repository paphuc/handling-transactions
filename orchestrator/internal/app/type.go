package app

import (
	"handling-transactions/orchestrator/config"

	grpc "google.golang.org/grpc"
)

type (
	Server struct {
		gServer *grpc.Server
		Config  *config.Config
	}
)
