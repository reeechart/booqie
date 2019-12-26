package server

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/reeechart/booql/logger"
	"google.golang.org/grpc"
)

type server struct {
	logger.UnimplementedLoggerServer
	host string
	port int
}

func NewServer(host string, port int) *server {
	return &server{
		host: host,
		port: port,
	}
}

func (s *server) LogEvent(ctx context.Context, in *logger.Event) (*empty.Empty, error) {
	fmt.Printf("User %d: %s\n", in.GetUserId(), in.GetActivity())
	return new(empty.Empty), nil
}

func (s *server) Run() {
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	logger.RegisterLoggerServer(grpcServer, s)
	log.Printf("Audit server is listening at %s\n", addr)
	log.Fatal(grpcServer.Serve(listener))
}
