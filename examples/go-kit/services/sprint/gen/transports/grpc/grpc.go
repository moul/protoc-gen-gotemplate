package sprint_transportgrpc

import (
	"fmt"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint"
	endpoint "github.com/moul/protoc-gen-gotemplate/examples/go-kit/sprint/gen/endpoints"
	context "golang.org/x/net/context"
)

// avoid import errors
var _ = fmt.Errorf

func MakeGRPCServer(ctx context.Context, endpoints endpoint.Endpoints) pb.SprintServiceServer {
	options := []grpctransport.ServerOption{}
	return &grpcServer{

		addsprint: grpctransport.NewServer(
			ctx,
			endpoints.AddSprintEndpoint,
			decodeAddSprintRequest,
			encodeAddSprintResponse,
			options,
		),

		closesprint: grpctransport.NewServer(
			ctx,
			endpoints.CloseSprintEndpoint,
			decodeCloseSprintRequest,
			encodeCloseSprintResponse,
			options,
		),

		getsprint: grpctransport.NewServer(
			ctx,
			endpoints.GetSprintEndpoint,
			decodeGetSprintRequest,
			encodeGetSprintResponse,
			options,
		),
	}
}

type grpcServer struct {
	addsprint grpctransport.Handler

	closesprint grpctransport.Handler

	getsprint grpctransport.Handler
}

func (s *grpcServer) AddSprint(ctx context.Context, req *pb.AddSprintRequest) (*pb.AddSprintReply, error) {
	_, rep, err := s.addsprint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddSprintReply), nil
}

func decodeAddSprintRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeAddSprintResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.AddSprintReply)
	return resp, nil
}

func (s *grpcServer) CloseSprint(ctx context.Context, req *pb.CloseSprintRequest) (*pb.CloseSprintReply, error) {
	_, rep, err := s.closesprint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CloseSprintReply), nil
}

func decodeCloseSprintRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeCloseSprintResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.CloseSprintReply)
	return resp, nil
}

func (s *grpcServer) GetSprint(ctx context.Context, req *pb.GetSprintRequest) (*pb.GetSprintReply, error) {
	_, rep, err := s.getsprint.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetSprintReply), nil
}

func decodeGetSprintRequest(ctx context.Context, grpcReq interface{}) (interface{}, error) {
	return grpcReq, nil
}

func encodeGetSprintResponse(ctx context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.GetSprintReply)
	return resp, nil
}
