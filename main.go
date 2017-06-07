package main

import (
	"fmt"

	pb "github.com/iochti/thing-service/proto"
)

func main() {
	psq := PostgresDL{}
	if err := psq.Init(); err != nil {
		fmt.Println(err.Error())
		return
	}

	tng := pb.Thing{
		Name:        "TestThing",
		Description: "This is a description",
	}

	if err := psq.CreateThing(&tng); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(tng)
	fmt.Println(tng.GetCreatedAt().String())
}
