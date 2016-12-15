package sprint_grpctransport

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
	context "golang.org/x/net/context"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.SprintServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{

		addsprint: grpctransport.NewServer(
			ctx,
			endpoints.AddSprintEndpoint,
			decodeAddSprintRequest,
			encodeAddSprintResponse,
			options...,
		),

		closesprint: grpctransport.NewServer(
			ctx,
			endpoints.CloseSprintEndpoint,
			decodeCloseSprintRequest,
			encodeCloseSprintResponse,
			options...,
		),

		getsprint: grpctransport.NewServer(
			ctx,
			endpoints.GetSprintEndpoint,
			decodeGetSprintRequest,
			encodeGetSprintResponse,
			options...,
		),
	}
}

type grpcServer struct {
	addsprint grpctransport.Handler

	closesprint grpctransport.Handler

	getsprint grpctransport.Handler
}

func (s *grpcServer) AddSprint(ctx context.Context, req *pb.AddSprintRequest) (*pb.AddSprintResponse, error) {
	_, rep, err := s.addsprint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddSprintResponse), nil
}

func decodeAddSprintRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeAddSprintResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.AddSprintResponse)
	return resp, nil
}

func (s *grpcServer) CloseSprint(ctx context.Context, req *pb.CloseSprintRequest) (*pb.CloseSprintResponse, error) {
	_, rep, err := s.closesprint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CloseSprintResponse), nil
}

func decodeCloseSprintRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeCloseSprintResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CloseSprintResponse)
	return resp, nil
}

func (s *grpcServer) GetSprint(ctx context.Context, req *pb.GetSprintRequest) (*pb.GetSprintResponse, error) {
	_, rep, err := s.getsprint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetSprintResponse), nil
}

func decodeGetSprintRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeGetSprintResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetSprintResponse)
	return resp, nil
}
