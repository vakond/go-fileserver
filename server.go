package main

//go:generate protoc --proto_path=proto --go_out=plugins=grpc:api fileserver.proto

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"vakond/fileserver/api"
)

var server *grpc.Server

func init() {
	server = grpc.NewServer()
}

// serve listens gRPC server.
func serve() error {
	api.RegisterFileserverServer(server, &fileserver{})

	if config.verbose {
		log.Printf("Listening gRPC server on port %d...\n", config.port)
	}

	if config.port <= 0 {
		return fmt.Errorf("invalid port %d", config.port)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.port))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to listen: %w", err))
	}

	return server.Serve(listener)
}
