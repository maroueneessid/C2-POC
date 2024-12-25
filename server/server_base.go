package main

import (
	"fmt"
	"log"
	"net"
	pb "simpleGRPC/proto_defs/common"
	pb_man "simpleGRPC/proto_defs/manager"
	"simpleGRPC/utils"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	pb.UnimplementedAssetServiceServer
	pb_man.ManagerAssetServer
	db     *redis.Client
	notifs chan *pb_man.Notification
}

type grpcConfig struct {
	grpcServer   *grpc.Server
	serverConfig *Server
}

var (

	// this initial config will be "inherited"
	// (to keep a common notifications channel between multiple service registration , aka Listeners)
	GlobalConf grpcConfig

	// track and manage listeners
	GlobalListeners map[int]chan bool

	kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: false,           // Allow pings even when there are no active streams
	}

	kasp = keepalive.ServerParameters{

		//MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		//MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 5 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               5 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead

	}
)

const (
	Red          = "\033[31m"
	Green        = "\033[32m"
	Yellow       = "\033[33m"
	Reset        = "\033[0m"
	magic  int32 = 0x45344534
)

func RegisterManagerListener(grpcServer *grpc.Server, serverConfig *Server, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		Error_check("could not listen", err)
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
		if port != 9001 {
			tr := &pb_man.Notification{
				SessionId: "",
				Notif:     fmt.Sprintf(Green + "[+] Listener up" + Reset + "\n"),
			}
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
		tr := &pb_man.Notification{Notif: fmt.Sprintf("%v", Error_check("could not listen", err))}
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
