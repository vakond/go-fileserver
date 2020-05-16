package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"vakond/fileserver/api"

	"google.golang.org/grpc"
)

const (
	service = "Fileserver"

	rpcAddress = "localhost:8090"

	timeout = 60 * time.Second
)

// Rpc represents a gRPC client connection.
type Rpc struct {
	name       string
	connection *grpc.ClientConn
	caller     api.FileserverClient
}

// newRpc creates new instance of Rpc.
func newRpc() (*Rpc, error) {
	var err error
	rpc := &Rpc{name: service}
	rpc.connection, err = grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("rpc connect: %w", err)
	}
	rpc.caller = api.NewFileserverClient(rpc.connection)
	return rpc, nil
}

// Close closes the rpc connection.
func (r *Rpc) Close() {
	err := r.connection.Close()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}

// Versions makes RPC call 'Versions'.
func (r *Rpc) Versions() (*api.VersionsResponse, error) {
	r.name = service + ".Versions"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return r.caller.Versions(ctx, &api.Request{})
}

// GetDownloadStream starts RPC call 'Download'.
func (r *Rpc) GetDownloadStream(ver string) (api.Fileserver_DownloadClient, context.CancelFunc, error) {
	r.name = service + ".Download"
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	stream, err := r.caller.Download(ctx, &api.Request{Version: ver})
	if err != nil {
		cancel()
		return nil, nil, err
	}
	return stream, cancel, nil
}

// Error generates new error with some context info.
func (r *Rpc) Error(e error) error {
	return fmt.Errorf("%s: %w", r.name, e)
}
