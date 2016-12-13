package usersvc

import (
	"fmt"

	"golang.org/x/net/context"

	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
)

type Service struct{}

func New() pb.UserServiceServer { return &Service{} }

func (svc *Service) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, fmt.Errorf("not implemented")
}
