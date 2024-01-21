package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net/http"

	pb "Test-Task2/proto"
)

func main() {
	log := logrus.New()
	log.Debug("Client starts")
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	grpcMux := runtime.NewServeMux()
	err = pb.RegisterUserServiceHandler(context.Background(), grpcMux, conn)
	err = pb.RegisterAuthHandler(context.Background(), grpcMux, conn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("client starts to work")
	log.Fatal(http.ListenAndServe("localhost:8082", grpcMux))
}
