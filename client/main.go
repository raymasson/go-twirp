package main

import (
	"context"
	"fmt"
	"net/http"

	pb "github.com/raymasson/go-twirp/users"
)

func main() {
	client := pb.NewEmailBossProtobufClient("http://localhost:8000", &http.Client{})

	resp, err := client.UpdateEmail(context.Background(), &pb.UpdateEmailRequest{UserId: 1234, NewEmail: "newemail@gmail.com"})
	if err == nil {
		fmt.Println(resp.CleanedEmail) // prints "newemail@gmail.com"
	} else {
		fmt.Println(err)
	}
}
