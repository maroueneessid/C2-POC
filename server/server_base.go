package main

import (
	"fmt"
	"log"
	"net"
	pb "simpleGRPC/proto_defs/common"
	pb_man "simpleGRPC/proto_defs/manager"
	"simpleGRPC/utils"
	"sync"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func RegisterManagerListener(grpcServer *grpc.Server, serverConfig *Server, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("could not listen: %v", err)
	}

	pb_man.RegisterManagerAssetServer(grpcServer, serverConfig)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Printf("Starting Remote manager failed: %v", err)
		}
	}()

	return nil

}

func RegisterAssetListener(grpcServer *grpc.Server, serverConfig *Server, port int) error {

	// Only append non used ports
	if !ListenerPresence(GlobalListeners, port) {
		GlobalListeners[port] = make(chan bool)
		// no sending listener setup notification for the original port
		if port != GLOBAL_PORT {
			tr := &pb_man.Notification{
				SessionId: "",
				Notif:     fmt.Sprintf(Green + "[+] Listener up" + Reset + "\n"),
			}
			UpdatePortPersistenceConfig()
			serverConfig.notifs <- tr
		}

	} else {
		tr := &pb_man.Notification{
			SessionId: "",
			Notif:     fmt.Sprintf(Red + "[!] Listener already established" + Reset + "\n"),
		}
		serverConfig.notifs <- tr
		return nil
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		tr := &pb_man.Notification{Notif: fmt.Sprintf("%v", err)}
		GlobalConf.serverConfig.notifs <- tr
		return nil
	}

	pb.RegisterAssetServiceServer(grpcServer, serverConfig)

	go func() {

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Printf("Asset Listener Failed: %v", err)
		}
	}()

	<-GlobalListeners[port]
	grpcServer.GracefulStop()

	return nil
}

func InitGrpcConfig(notifs chan *pb_man.Notification) grpcConfig {

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
		notifs = make(chan *pb_man.Notification, 1000)

	}

	grpcServer := grpc.NewServer(grpc.Creds(tlsCred), grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))

	serverConfig := &Server{db: redisClient, notifs: notifs}

	tr := grpcConfig{grpcServer, serverConfig}

	return tr

}
