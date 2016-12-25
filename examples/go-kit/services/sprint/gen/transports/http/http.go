package sprint_httptransport

import (
	"encoding/json"
	context "golang.org/x/net/context"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/sprint/gen/pb"
)

func MakeAddSprintHandler(ctx context.Context, svc pb.SprintServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeAddSprintRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeAddSprintRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.AddSprintRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func MakeCloseSprintHandler(ctx context.Context, svc pb.SprintServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeCloseSprintRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeCloseSprintRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.CloseSprintRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func MakeGetSprintHandler(ctx context.Context, svc pb.SprintServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeGetSprintRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeGetSprintRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.GetSprintRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(ctx context.Context, svc pb.SprintServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	log.Println("new HTTP endpoint: \"/AddSprint\" (service=Sprint)")
	mux.Handle("/AddSprint", MakeAddSprintHandler(ctx, svc, endpoints.AddSprintEndpoint))

	log.Println("new HTTP endpoint: \"/CloseSprint\" (service=Sprint)")
	mux.Handle("/CloseSprint", MakeCloseSprintHandler(ctx, svc, endpoints.CloseSprintEndpoint))

	log.Println("new HTTP endpoint: \"/GetSprint\" (service=Sprint)")
	mux.Handle("/GetSprint", MakeGetSprintHandler(ctx, svc, endpoints.GetSprintEndpoint))

	return nil
}
