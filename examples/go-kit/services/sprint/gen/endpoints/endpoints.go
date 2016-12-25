package sprint_endpoints

import (
	"fmt"

	"github.com/go-kit/kit/endpoint"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
	context "golang.org/x/net/context"
)

var _ = fmt.Errorf

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	AddSprintEndpoint endpoint.Endpoint

	CloseSprintEndpoint endpoint.Endpoint

	GetSprintEndpoint endpoint.Endpoint
}

func (e *Endpoints) AddSprint(ctx context.Context, in *pb.AddSprintRequest) (*pb.AddSprintResponse, error) {
	out, err := e.AddSprintEndpoint(ctx, in)
	if err != nil {
		return &pb.AddSprintResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.AddSprintResponse), err
}

func (e *Endpoints) CloseSprint(ctx context.Context, in *pb.CloseSprintRequest) (*pb.CloseSprintResponse, error) {
	out, err := e.CloseSprintEndpoint(ctx, in)
	if err != nil {
		return &pb.CloseSprintResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.CloseSprintResponse), err
}

func (e *Endpoints) GetSprint(ctx context.Context, in *pb.GetSprintRequest) (*pb.GetSprintResponse, error) {
	out, err := e.GetSprintEndpoint(ctx, in)
	if err != nil {
		return &pb.GetSprintResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.GetSprintResponse), err
}

func MakeAddSprintEndpoint(svc pb.SprintServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.AddSprintRequest)
		rep, err := svc.AddSprint(ctx, req)
		if err != nil {
			return &pb.AddSprintResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeCloseSprintEndpoint(svc pb.SprintServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.CloseSprintRequest)
		rep, err := svc.CloseSprint(ctx, req)
		if err != nil {
			return &pb.CloseSprintResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeGetSprintEndpoint(svc pb.SprintServiceServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*pb.GetSprintRequest)
		rep, err := svc.GetSprint(ctx, req)
		if err != nil {
			return &pb.GetSprintResponse{ErrMsg: err.Error()}, err
		}
		return rep, nil
	}
}

func MakeEndpoints(svc pb.SprintServiceServer) Endpoints {
	return Endpoints{

		AddSprintEndpoint: MakeAddSprintEndpoint(svc),

		CloseSprintEndpoint: MakeCloseSprintEndpoint(svc),

		GetSprintEndpoint: MakeGetSprintEndpoint(svc),
	}
}
