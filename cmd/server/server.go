package server

import (
	"log"
	"net"
	"users/config"
	"users/generated/users"
	"users/service"
	"users/storage/postgres"

	"google.golang.org/grpc"
)

func ServerRun(userRepo *postgres.UserRepo, cfg *config.Config) {
	listener, err := net.Listen("tcp", cfg.GRPC_PORT)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	users.RegisterAuthServiceServer(s, service.NewUserServer(userRepo))

	log.Printf("Server is running on %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
