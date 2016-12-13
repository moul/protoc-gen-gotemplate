package sessionsvc

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/pb"
)

type Service struct{}

func New() pb.SessionServiceServer {
	return &Service{}
}

func (svc *Service) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Logout(ctx context.Context, in *pb.LogoutRequest) (*pb.LogoutResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
