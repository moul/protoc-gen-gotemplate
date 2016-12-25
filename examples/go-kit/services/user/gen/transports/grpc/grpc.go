package user_grpctransport

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"

	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.UserServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{

		createuser: grpctransport.NewServer(
			ctx,
			endpoints.CreateUserEndpoint,
			decodeRequest,
			encodeCreateUserResponse,
			options...,
		),

		getuser: grpctransport.NewServer(
			ctx,
			endpoints.GetUserEndpoint,
			decodeRequest,
			encodeGetUserResponse,
			options...,
		),
	}
}

type grpcServer struct {
	createuser grpctransport.Handler

	getuser grpctransport.Handler
}

func (s *grpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, rep, err := s.createuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateUserResponse), nil
}

func encodeCreateUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CreateUserResponse)
	return resp, nil
}

func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, rep, err := s.getuser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetUserResponse), nil
}

func encodeGetUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetUserResponse)
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
