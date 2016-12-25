package session_httptransport

import (
	"encoding/json"
	context "golang.org/x/net/context"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/pb"
)

func MakeLoginHandler(ctx context.Context, svc pb.SessionServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeLoginRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(ctx context.Context, svc pb.SessionServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	log.Println("new HTTP endpoint: \"/Login\" (service=Session)")
	mux.Handle("/Login", MakeLoginHandler(ctx, svc, endpoints.LoginEndpoint))

	return nil
}
