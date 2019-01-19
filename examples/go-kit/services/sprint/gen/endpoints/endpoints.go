package sprint_endpoints

import (
	context "context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	oldcontext "golang.org/x/net/context"
	pb "moul.io/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
)

var _ = endpoint.Chain
var _ = fmt.Errorf
var _ = context.Background

type StreamEndpoint func(server interface{}, req interface{}) (err error)

type Endpoints struct {
	AddSprintEndpoint endpoint.Endpoint

	CloseSprintEndpoint endpoint.Endpoint

	GetSprintEndpoint endpoint.Endpoint
}

func (e *Endpoints) AddSprint(ctx oldcontext.Context, in *pb.AddSprintRequest) (*pb.AddSprintResponse, error) {
	out, err := e.AddSprintEndpoint(ctx, in)
	if err != nil {
		return &pb.AddSprintResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.AddSprintResponse), err
}

func (e *Endpoints) CloseSprint(ctx oldcontext.Context, in *pb.CloseSprintRequest) (*pb.CloseSprintResponse, error) {
	out, err := e.CloseSprintEndpoint(ctx, in)
	if err != nil {
		return &pb.CloseSprintResponse{ErrMsg: err.Error()}, err
	}
	return out.(*pb.CloseSprintResponse), err
}

func (e *Endpoints) GetSprint(ctx oldcontext.Context, in *pb.GetSprintRequest) (*pb.GetSprintResponse, error) {
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
