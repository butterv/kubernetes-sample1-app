package main

import (
	"fmt"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/butterv/kubernetes-sample1-app/app/domain/model"
	"github.com/butterv/kubernetes-sample1-app/app/domain/repository"
	"github.com/butterv/kubernetes-sample1-app/app/domain/service/health"
	"github.com/butterv/kubernetes-sample1-app/app/domain/service/user"
	healthpb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/health"
	userpb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/user"
	"github.com/butterv/kubernetes-sample1-app/app/infrastructure/persistence"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db := connectDB(dsn)
	repo := persistence.New(db)

	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatal(err)
	}

	s := newGRPCServer(repo)
	reflection.Register(s)
	_ = s.Serve(listenPort)
	s.GracefulStop()
}

func connectDB(connection string) *sqlx.DB {
	db, err := sqlx.Open("mysql", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func newGRPCServer(repo repository.Repository) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	healthpb.RegisterHealthServer(s, health.NewHealthService())
	userpb.RegisterUsersServer(s, user.NewUserService(repo, model.NewDefaultUserIDGenerator()))

	return s
}
