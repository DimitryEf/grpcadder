package main

import (
	"context"
	"flag"
	"github.com/DimitryEf/grpcadder/pkg/api"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main(){
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("where arguments?")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())
}