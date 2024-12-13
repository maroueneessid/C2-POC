// package server
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	pb "simpleGRPC/proto_defs"
	"strings"
	"sync"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAssetServiceServer
	db *redis.Client
}

const (
	Red          = "\033[31m"
	Green        = "\033[32m"
	Yellow       = "\033[33m"
	Reset        = "\033[0m"
	magic  int32 = 0x45344534
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		Error_check("could not listen", err)
	}

	fmt.Printf(Green + "[!] Started server on LOCALHOST:9001" + Reset + "\n")

	grpcServer := grpc.NewServer()
	pb.RegisterAssetServiceServer(grpcServer, &Server{db: redisClient})

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Printf("could not serve gRPC server: %v", err)
		}
	}()

	for {
		var input string
		fmt.Print("Enter command 'exit' to stop server:\n> ")
		fmt.Scanln(&input)

		switch strings.ToLower(input) {
		case "exit":
			fmt.Println("Shutting down the server...")
			grpcServer.GracefulStop()
			wg.Wait()
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}
}
