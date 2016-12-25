package sprint_clientgrpc

import (
	jwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"

	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
)

func New(conn *grpc.ClientConn, logger log.Logger) pb.SprintServiceServer {

	var addsprintEndpoint endpoint.Endpoint
	{
		addsprintEndpoint = grpctransport.NewClient(
			conn,
			"sprint.SprintService",
			"AddSprint",
			EncodeAddSprintRequest,
			DecodeAddSprintResponse,
			pb.AddSprintResponse{},
			append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.FromGRPCContext()))...,
		).Endpoint()
	}

	var closesprintEndpoint endpoint.Endpoint
	{
		closesprintEndpoint = grpctransport.NewClient(
			conn,
			"sprint.SprintService",
			"CloseSprint",
			EncodeCloseSprintRequest,
			DecodeCloseSprintResponse,
			pb.CloseSprintResponse{},
			append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.FromGRPCContext()))...,
		).Endpoint()
	}

	var getsprintEndpoint endpoint.Endpoint
	{
		getsprintEndpoint = grpctransport.NewClient(
			conn,
			"sprint.SprintService",
			"GetSprint",
			EncodeGetSprintRequest,
			DecodeGetSprintResponse,
			pb.GetSprintResponse{},
			append([]grpctransport.ClientOption{}, grpctransport.ClientBefore(jwt.FromGRPCContext()))...,
		).Endpoint()
	}

	return &endpoints.Endpoints{

		AddSprintEndpoint: addsprintEndpoint,

		CloseSprintEndpoint: closesprintEndpoint,

		GetSprintEndpoint: getsprintEndpoint,
	}
}

func EncodeAddSprintRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddSprintRequest)
	return req, nil
}

func DecodeAddSprintResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.AddSprintResponse)
	return response, nil
}

func EncodeCloseSprintRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CloseSprintRequest)
	return req, nil
}

func DecodeCloseSprintResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.CloseSprintResponse)
	return response, nil
}

func EncodeGetSprintRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetSprintRequest)
	return req, nil
}

func DecodeGetSprintResponse(_ context.Context, grpcResponse interface{}) (interface{}, error) {
	response := grpcResponse.(*pb.GetSprintResponse)
	return response, nil
}
