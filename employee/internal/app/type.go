package app

import (
	"context"
	"google.golang.org/grpc"
	"net/http"
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
