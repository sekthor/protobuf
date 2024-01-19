package main

import (
	"fmt"
	"log"

	"github.com/sekthor/protobuf/api"
	"google.golang.org/protobuf/proto"
)

func main() {
	per := &api.Person{
		Name: "sekthor",
		Age:  100,
	}

	data, err := proto.Marshal(per)
	if err != nil {
		log.Fatal("could not marshal person")
	}

	per2 := new(api.Person)
	err = proto.Unmarshal(data, per2)
	if err != nil {
		log.Fatal("could not unmarshal person")
	}

	fmt.Println(per2)
}
