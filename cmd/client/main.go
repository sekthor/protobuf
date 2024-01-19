package main

import (
	"context"
	"fmt"
	"log"

	"github.com/sekthor/protobuf/api"
	"google.golang.org/grpc"
)

func main() {
	//var opt grpc.DialOption

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not call grpc server")
	}
	defer conn.Close()

	client := api.NewPersonServiceClient(conn)

	// this is the request we send to the server
	request := api.PersonRequest{
		Id: 1,
	}

	person, err := client.GetPerson(context.Background(), &request)
	if err != nil {
		log.Fatal("could not fetch person:", err.Error())
	}

	fmt.Println(person)
}
