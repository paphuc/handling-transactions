package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	"handling-transactions/assignment/config"
	"handling-transactions/assignment/internal/route"
	"handling-transactions/protocol-buffers/assignment"
	"handling-transactions/protocol-buffers/transaction"
)

func New(conf *config.Config) AppI {
	//=====================gRPC==========================
	gRPCClient, err := grpc.Dial(conf.GRPC.Address, grpc.WithInsecure())
	if err != nil {
		log.Panic("Cound't connect to gRPC Server")
	}
	assignmentC := assignment.NewAssignmentClient(gRPCClient)
	txC := transaction.NewTransactionClient(gRPCClient)

	//==================Api Server====================
	r, err := route.NewRouter(assignmentC, txC)
	if err != nil {
		log.Panicf("Failed Init Router")
	}
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.Server.Port),
		Handler: r,
	}
	return &App{
		Server:     httpServer,
		GRPCClient: gRPCClient,
	}
}

func (a *App) StopGRPC() error {
	return a.GRPCClient.Close()
}

func (a *App) Start() error {
	log.Printf("Server is listening at %s", a.Server.Addr)
	return a.Server.ListenAndServe()
}

func (a *App) Stop(ctx context.Context) error {
	return a.Server.Shutdown(ctx)
}
