package session_httptransport

import (
	"encoding/json"
	context "golang.org/x/net/context"
	"log"
	"net/http"

	gokit_endpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	pb "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session"
	endpoints "github.com/moul/protoc-gen-gotemplate/examples/go-kit/services/session/gen/endpoints"
)

func MakeLoginHandler(ctx context.Context, svc pb.SessionServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeLoginRequest,
		encodeLoginResponse,
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

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func MakeLogoutHandler(ctx context.Context, svc pb.SessionServiceServer, endpoint gokit_endpoint.Endpoint) *httptransport.Server {
	return httptransport.NewServer(
		ctx,
		endpoint,
		decodeLogoutRequest,
		encodeLogoutResponse,
		[]httptransport.ServerOption{}...,
	)
}

func decodeLogoutRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req pb.LogoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return &req, nil
}

func encodeLogoutResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func RegisterHandlers(ctx context.Context, svc pb.SessionServiceServer, mux *http.ServeMux, endpoints endpoints.Endpoints) error {

	log.Println("new HTTP endpoint: \"/Login\" (service=Session)")
	mux.Handle("/Login", MakeLoginHandler(ctx, svc, endpoints.LoginEndpoint))

	log.Println("new HTTP endpoint: \"/Logout\" (service=Session)")
	mux.Handle("/Logout", MakeLogoutHandler(ctx, svc, endpoints.LogoutEndpoint))

	return nil
}
