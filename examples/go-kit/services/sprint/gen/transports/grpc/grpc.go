package sprint_grpctransport

import (
	context "context"
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"

	endpoints "moul.io/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/endpoints"
	pb "moul.io/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(endpoints endpoints.Endpoints) pb.SprintServiceServer {
	var options []grpctransport.ServerOption
	_ = options
	return &grpcServer{

		addsprint: grpctransport.NewServer(
			endpoints.AddSprintEndpoint,
			decodeRequest,
			encodeAddSprintResponse,
			options...,
		),

		closesprint: grpctransport.NewServer(
			endpoints.CloseSprintEndpoint,
			decodeRequest,
			encodeCloseSprintResponse,
			options...,
		),

		getsprint: grpctransport.NewServer(
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

func (s *grpcServer) AddSprint(ctx oldcontext.Context, req *pb.AddSprintRequest) (*pb.AddSprintResponse, error) {
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

func (s *grpcServer) CloseSprint(ctx oldcontext.Context, req *pb.CloseSprintRequest) (*pb.CloseSprintResponse, error) {
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

func (s *grpcServer) GetSprint(ctx oldcontext.Context, req *pb.GetSprintRequest) (*pb.GetSprintResponse, error) {
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
