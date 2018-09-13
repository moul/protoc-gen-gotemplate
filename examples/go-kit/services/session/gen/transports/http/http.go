package session_httptransport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "moul.io/protoc-gen-gotemplate/examples/go-kit/services/session/gen/endpoints"
	pb "moul.io/protoc-gen-gotemplate/examples/go-kit/services/session/gen/pb"
)

var _ = log.Printf
var _ = gokit_endpoint.Chain
var _ = httptransport.NewClient

func MakeLoginHandler(svc pb.SessionServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
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

func RegisterHandlers(svc pb.SessionServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	log.Println("new HTTP endpoint: \"/Login\" (service=Session)")
	mux.Handle("/Login", MakeLoginHandler(svc, endpoints.LoginEndpoint))

	return nil
}
