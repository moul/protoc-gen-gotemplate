package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/handlers"
	"google.golang.org/grpc"

	session_svc "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session"
	session_endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/endpoints"
	session_pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/pb"
	session_grpctransport "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/transports/grpc"
	session_httptransport "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/transports/http"

	sprint_svc "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint"
	sprint_endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/endpoints"
	sprint_pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
	sprint_grpctransport "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/transports/grpc"
	sprint_httptransport "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/transports/http"

	user_svc "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user"
	user_endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/endpoints"
	user_pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
	user_grpctransport "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/transports/grpc"
	user_httptransport "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/transports/http"
)

func main() {
	mux := http.NewServeMux()
	ctx := context.Background()
	errc := make(chan error)
	s := grpc.NewServer()
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stdout)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	// initialize services
	{
		svc := session_svc.New()
		endpoints := session_endpoints.MakeEndpoints(svc)
		srv := session_grpctransport.MakeGRPCServer(ctx, endpoints)
		session_pb.RegisterSessionServiceServer(s, srv)
		session_httptransport.RegisterHandlers(ctx, svc, mux, endpoints)
	}
	{
		svc := sprint_svc.New()
		endpoints := sprint_endpoints.MakeEndpoints(svc)
		srv := sprint_grpctransport.MakeGRPCServer(ctx, endpoints)
		sprint_pb.RegisterSprintServiceServer(s, srv)
		sprint_httptransport.RegisterHandlers(ctx, svc, mux, endpoints)
	}
	{
		svc := user_svc.New()
		endpoints := user_endpoints.MakeEndpoints(svc)
		srv := user_grpctransport.MakeGRPCServer(ctx, endpoints)
		user_pb.RegisterUserServiceServer(s, srv)
		user_httptransport.RegisterHandlers(ctx, svc, mux, endpoints)
	}

	// start servers
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger := log.NewContext(logger).With("transport", "HTTP")
		logger.Log("addr", ":8000")
		errc <- http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stderr, mux))
	}()

	go func() {
		logger := log.NewContext(logger).With("transport", "gRPC")
		ln, err := net.Listen("tcp", ":9000")
		if err != nil {
			errc <- err
			return
		}
		logger.Log("addr", ":9000")
		errc <- s.Serve(ln)
	}()

	logger.Log("exit", <-errc)
}
