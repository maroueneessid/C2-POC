package main

import (
	"fmt"
	"log"
	"net"
	"os"
	pb "simpleGRPC/proto_defs"
	"simpleGRPC/utils"
	"strings"
	"sync"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAssetServiceServer
	db     *redis.Client
	notifs chan *pb.Notification
}

type grpcConfig struct {
	grpcServer   *grpc.Server
	serverConfig *Server
}

const (
	Red          = "\033[31m"
	Green        = "\033[32m"
	Yellow       = "\033[33m"
	Reset        = "\033[0m"
	magic  int32 = 0x45344534
)

func RegisterListener(grpcServer *grpc.Server, serverConfig *Server, port int) error {

	// Anon function to check if port already listening
	checkPresence := func(arr []int, target int) bool {
		for _, v := range arr {
			if v == target {
				return true
			}
		}
		return false
	}

	// Only append non used ports
	if !checkPresence(GlobalListeners, port) {
		GlobalListeners = append(GlobalListeners, port)
		// no sending listener setup notification for the original port
		if port != 9001 {
			tr := &pb.Notification{
				SessionId: "",
				Notif:     fmt.Sprintf(Green + "[+] Listener up" + Reset + "\n"),
			}
			serverConfig.notifs <- tr
		}

	} else {
		tr := &pb.Notification{
			SessionId: "",
			Notif:     fmt.Sprintf(Red + "[!] Listener already established" + Reset + "\n"),
		}
		serverConfig.notifs <- tr
		return nil
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		Error_check("could not listen", err)
	}

	pb.RegisterAssetServiceServer(grpcServer, serverConfig)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Printf("could not serve gRPC server: %v", err)
		}
	}()

	return nil
}

func InitGrpcConfig(notifs chan *pb.Notification) grpcConfig {

	tlsCred, err := utils.SimpleServerTLS()

	if err != nil {
		log.Fatalf("[-] Error loading SSL Cert: %v", err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	if notifs == nil {
		notifs = make(chan *pb.Notification, 1000)

	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCred))

	serverConfig := &Server{db: redisClient, notifs: notifs}

	tr := grpcConfig{grpcServer, serverConfig}

	return tr

}

func ServerCommandHandler(conf grpcConfig, in string) {

	for {
		if len(strings.Fields(in)) == 0 {
			break
		}
		switch strings.Fields(in)[0] {
		case "exit":
			fmt.Println("Shutting down the server...")
			conf.grpcServer.GracefulStop()
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Please try again.")
		}
	}

}
