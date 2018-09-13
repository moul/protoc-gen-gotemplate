package user_httptransport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "moul.io/protoc-gen-gotemplate/examples/go-kit/services/user/gen/endpoints"
	pb "moul.io/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
)

var _ = log.Printf
var _ = gokit_endpoint.Chain
var _ = httptransport.NewClient

func MakeCreateUserHandler(svc pb.UserServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		decodeCreateUserRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeCreateUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func MakeGetUserHandler(svc pb.UserServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		endpoint,
		decodeGetUserRequest,
		encodeResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeGetUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.GetUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(svc pb.UserServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	log.Println("new HTTP endpoint: \"/CreateUser\" (service=User)")
	mux.Handle("/CreateUser", MakeCreateUserHandler(svc, endpoints.CreateUserEndpoint))

	log.Println("new HTTP endpoint: \"/GetUser\" (service=User)")
	mux.Handle("/GetUser", MakeGetUserHandler(svc, endpoints.GetUserEndpoint))

	return nil
}
