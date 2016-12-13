package session_endpoints

import (
	"fmt"
	"github.com/go-kit/kit/endpoint"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session"
	context "golang.org/x/net/context"
)

var _ = fmt.Errorf

type Endpoints struct {
	LoginEndpoint endpoint.Endpoint

	LogoutEndpoint endpoint.Endpoint
}

func (e *Endpoints) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	out, err := e.LoginEndpoint(ctx, in)
	if err != nil {
		return &pb.LoginResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.LoginReply), err
}

func (e *Endpoints) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	out, err := e.LogoutEndpoint(ctx, in)
	if err != nil {
		return &pb.LogoutResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.LogoutReply), err
}

func MakeLoginEndpoint(svc pb.SessionServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.LoginRequest)
		rep, err := svc.Login(ctx, req)
		if err != nil {
			return &pb.LoginReply{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeLogoutEndpoint(svc pb.SessionServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.LogoutRequest)
		rep, err := svc.Logout(ctx, req)
		if err != nil {
			return &pb.LogoutReply{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeEndpoints(svc pb.SessionServiceServer) Endpoints {
	return Endpoints{

		LoginEndpoint: MakeLoginEndpoint(svc),

		LogoutEndpoint: MakeLogoutEndpoint(svc),
	}
}
