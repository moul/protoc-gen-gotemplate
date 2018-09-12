package user_grpctransport

import (
	context "context"
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"

	endpoints "moul.io/protoc-gen-gotemplate/examples/go-kit/services/user/gen/endpoints"
	pb "moul.io/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(endpoints endpoints.Endpoints) pb.UserServiceServer {
	var options []grpctransport.ServerOption
	_ = options
	return &grpcServer{

		createuser: grpctransport.NewServer(
			endpoints.CreateUserEndpoint,
			decodeRequest,
			encodeCreateUserResponse,
			options...,
		),

		getuser: grpctransport.NewServer(
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

func (s *grpcServer) CreateUser(ctx oldcontext.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
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

func (s *grpcServer) GetUser(ctx oldcontext.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
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
