package sprint_grpctransport

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"

	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoints.Endpoints) pb.SprintServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{

		addsprint: grpctransport.NewServer(
			ctx,
			endpoints.AddSprintEndpoint,
			decodeRequest,
			encodeAddSprintResponse,
			options...,
		),

		closesprint: grpctransport.NewServer(
			ctx,
			endpoints.CloseSprintEndpoint,
			decodeRequest,
			encodeCloseSprintResponse,
			options...,
		),

		getsprint: grpctransport.NewServer(
			ctx,
			endpoints.GetSprintEndpoint,
			decodeRequest,
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

func encodeGetSprintResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetSprintResponse)
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
