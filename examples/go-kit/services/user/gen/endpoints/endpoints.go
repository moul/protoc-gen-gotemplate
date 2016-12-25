package user_endpoints

import (
	"fmt"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
	context "golang.org/x/net/context"
)

var _ = fmt.Errorf

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	CreateUserEndpoint endpoint.Endpoint

	GetUserEndpoint endpoint.Endpoint
}

func (e *Endpoints) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	out, err := e.CreateUserEndpoint(ctx, in)
	if err != nil {
		return &pb.CreateUserResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.CreateUserResponse), err
}

func (e *Endpoints) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	out, err := e.GetUserEndpoint(ctx, in)
	if err != nil {
		return &pb.GetUserResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.GetUserResponse), err
}

func MakeCreateUserEndpoint(svc pb.UserServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CreateUserRequest)
		rep, err := svc.CreateUser(ctx, req)
		if err != nil {
			return &pb.CreateUserResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeGetUserEndpoint(svc pb.UserServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetUserRequest)
		rep, err := svc.GetUser(ctx, req)
		if err != nil {
			return &pb.GetUserResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeEndpoints(svc pb.UserServiceServer) Endpoints {
	return Endpoints{

		CreateUserEndpoint: MakeCreateUserEndpoint(svc),

		GetUserEndpoint: MakeGetUserEndpoint(svc),
	}
}
