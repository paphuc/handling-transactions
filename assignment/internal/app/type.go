package app

import (
	"context"
	"net/http"

	"google.golang.org/grpc"
)

type (
	App struct {
		Server     *http.Server
		GRPCClient *grpc.ClientConn
	}

	AppI interface {
		Start() error
		Stop(ctx context.Context) error
		StopGRPC() error
	}
)
