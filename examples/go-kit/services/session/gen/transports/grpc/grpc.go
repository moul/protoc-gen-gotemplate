package session_transportgrpc

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session"
	endpoint "github.com/moul/protoc-gen-gotemplate/examples/go-kit/session/gen/endpoints"
	context "golang.org/x/net/context"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) pb.SessionServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{

		login: grpctransport.NewServer(
			ctx,
			endpoints.LoginEndpoint,
			decodeLoginRequest,
			encodeLoginResponse,
			options,
		),

		logout: grpctransport.NewServer(
			ctx,
			endpoints.LogoutEndpoint,
			decodeLogoutRequest,
			encodeLogoutResponse,
			options,
		),
	}
}

type grpcServer struct {
	login grpctransport.Handler

	logout grpctransport.Handler
}

func (s *grpcServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	_, rep, err := s.login.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LoginReply), nil
}

func decodeLoginRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeLoginResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LoginReply)
	return resp, nil
}

func (s *grpcServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	_, rep, err := s.logout.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.LogoutReply), nil
}

func decodeLogoutRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeLogoutResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.LogoutReply)
	return resp, nil
}
