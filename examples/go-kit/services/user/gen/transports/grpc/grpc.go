package user_grpctransport

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
	context "golang.org/x/net/context"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.UserServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{

		createuser: grpctransport.NewServer(
			ctx,
			endpoints.CreateUserEndpoint,
			decodeCreateUserRequest,
			encodeCreateUserResponse,
			options...,
		),

		getuser: grpctransport.NewServer(
			ctx,
			endpoints.GetUserEndpoint,
			decodeGetUserRequest,
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

func decodeCreateUserRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
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

func decodeGetUserRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeGetUserResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetUserResponse)
	return resp, nil
}
