package main

import (
	"context"

	"github.com/sekthor/protobuf/api"
)

type Server struct {
	api.UnsafePersonServiceServer
}

func (s Server) GetPerson(context.Context, *api.PersonRequest) (*api.Person, error) {
	return &api.Person{
		Name: "sekthor",
		Age:  100,
	}, nil
}
