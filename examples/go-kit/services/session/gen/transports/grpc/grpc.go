package session_grpctransport

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"

	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/pb"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.SessionServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{

		login: grpctransport.NewServer(
			ctx,
			endpoints.LoginEndpoint,
			decodeRequest,
			encodeLoginResponse,
			options...,
		),
	}
}

type grpcServer struct {
	login grpctransport.Handler
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginResponse), nil
}

func encodeLoginResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginResponse)
	return resp, nil
}

func decodeRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

type streamHandler interface {
	Do(server interface{}, req interface{}) (err error)
}

type server struct {
	e endpoints.StreamEndpoint
}

func (s server) Do(server interface{}, req interface{}) (err error) {
	if err := s.e(server, req); err != nil {
		return err
	}
	return nil
}
