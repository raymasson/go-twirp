package main

import (
	"context"
	"net/http"

	pb "github.com/raymasson/go-twirp/users"
)

type EmailBossServer struct{}

func (s *EmailBossServer) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateEmailResponse, error) {
	return &pb.UpdateEmailResponse{
		CleanedEmail: req.NewEmail,
	}, nil
}

// Run the implementation in a local server
func main() {
	twirpHandler := pb.NewEmailBossServer(&EmailBossServer{}, nil)
	// You can use any mux you like - NewHelloWorldServer gives you an http.Handler.
	mux := http.NewServeMux()
	// The generated code includes a const, <ServiceName>PathPrefix, which
	// can be used to mount your service on a mux.
	mux.Handle(pb.EmailBossPathPrefix, twirpHandler)
	http.ListenAndServe(":8000", mux)
}
