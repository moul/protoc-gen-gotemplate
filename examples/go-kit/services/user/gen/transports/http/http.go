package user_httptransport

import (
	"encoding/json"
	context "golang.org/x/net/context"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/endpoints"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/user/gen/pb"
)

func MakeCreateUserHandler(ctx context.Context, svc pb.UserServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
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

func MakeGetUserHandler(ctx context.Context, svc pb.UserServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
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

func RegisterHandlers(ctx context.Context, svc pb.UserServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	log.Println("new HTTP endpoint: \"/CreateUser\" (service=User)")
	mux.Handle("/CreateUser", MakeCreateUserHandler(ctx, svc, endpoints.CreateUserEndpoint))

	log.Println("new HTTP endpoint: \"/GetUser\" (service=User)")
	mux.Handle("/GetUser", MakeGetUserHandler(ctx, svc, endpoints.GetUserEndpoint))

	return nil
}
