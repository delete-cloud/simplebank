package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/delete-cloud/simplebank/gapi"
	"github.com/delete-cloud/simplebank/pb"
	"github.com/delete-cloud/simplebank/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/delete-cloud/simplebank/api"
	db "github.com/delete-cloud/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listen, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server on %s", listen.Addr().String())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
