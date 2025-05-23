package main

import (
	"coba/pbuf"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := ":9090"
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	cli := pbuf.NewGreeterClient(conn)

	fmt.Println("doing SayHello")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := cli.SayHello(ctx, &pbuf.HelloRequest{
		Name: "Artorius",
	})

	if err != nil {
		fmt.Println("Error happen")
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("status: %d\n", res.GetHeader().GetStatus())
	fmt.Printf("message: %s\n", res.GetHeader().GetMessage())
	fmt.Printf("result: %s\n", res.GetMessage())

	fmt.Println("============================")

	fmt.Println("doing SayHelloMultiple")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err = cli.SayHelloMultiple(ctx, &pbuf.HelloRequestMultiple{
		Name: []string{
			"Arthur", "Lancelot", "Galahad", "Gawain", "Percival",
		},
	})

	if err != nil {
		fmt.Println("Error happen")
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("status: %d\n", res.GetHeader().GetStatus())
	fmt.Printf("message: %s\n", res.GetHeader().GetMessage())
	fmt.Printf("result: %s\n", res.GetMessage())

	fmt.Println("============================")

	fmt.Println("doing SayHelloError")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err = cli.SayHelloError(ctx, &pbuf.HelloRequest{
		Name: "Artoria",
	})

	if err != nil {
		fmt.Println("Error happen")
		fmt.Printf("error: %v\n", err)
	}

	fmt.Printf("status: %d\n", res.GetHeader().GetStatus())
	fmt.Printf("message: %s\n", res.GetHeader().GetMessage())
	fmt.Printf("result: %s\n", res.GetMessage())
}
