package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	healthpb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/health"
	userpb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	mux := runtime.NewServeMux()

	if err := registerHandlers(ctx, mux); err != nil {
		return err
	}
	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8080")
	return gwServer.ListenAndServe()
}

func registerHandlers(ctx context.Context, mux *runtime.ServeMux) error {
	opts := grpcDialOptions()

	if err := healthpb.RegisterHealthHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}
	if err := userpb.RegisterUsersHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	return nil
}

func grpcDialOptions() []grpc.DialOption {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	return opts
}
