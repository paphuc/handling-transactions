package app

import (
	"database/sql"
	"fmt"
	"handling-transactions/orchestrator/config"
	"handling-transactions/orchestrator/internal/gassignment"
	"handling-transactions/orchestrator/internal/gemployee"
	"handling-transactions/orchestrator/internal/gtransaction"
	"handling-transactions/orchestrator/internal/rediscache"
	"handling-transactions/orchestrator/internal/transactioncache"
	"log"
	"net"

	"handling-transactions/protocol-buffers/assignment"
	"handling-transactions/protocol-buffers/employee"
	"handling-transactions/protocol-buffers/transaction"

	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
)

type ServerI interface {
	Start() error
	Stop()
}

func New(conf *config.Config, db *sql.DB, redisClient *redis.Client) ServerI {
	//Grpc Service
	grpcServer := grpc.NewServer()
	//TransactionCache
	txCacheSrv := transactioncache.NewTransactionCacheSrv()
	//Redis Service
	redisService := rediscache.NewRedisSrv(redisClient, txCacheSrv)
	//Transaction
	txSrv := gtransaction.NewTransactionSrv(db, redisService, txCacheSrv)
	txGSrv := gtransaction.NewGService(txSrv)
	transaction.RegisterTransactionServer(grpcServer, txGSrv)
	//Employee
	employeeRepos := gemployee.NewRepository(txSrv)
	employeeGSrv := gemployee.NewGService(employeeRepos)
	employee.RegisterEmployeeServer(grpcServer, employeeGSrv)
	//Assignment
	assignmentRepos := gassignment.NewRepository(txSrv)
	assignmentGSrv := gassignment.NewGService(assignmentRepos)
	assignment.RegisterAssignmentServer(grpcServer, assignmentGSrv)

	return &Server{
		gServer: grpcServer,
		Config:  conf,
	}

}

func (a *Server) Stop() {
	a.gServer.Stop()
}

func (a *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", a.Config.Server.GPort))
	log.Printf("gRPC Server listens at port %v", a.Config.Server.GPort)
	if err != nil {
		log.Fatalf("gRPC Failed to listen at port %v %s", a.Config.Server.GPort, err)
	}
	return a.gServer.Serve(lis)
}
