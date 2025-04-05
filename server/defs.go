package main

import (
	pb "simpleGRPC/proto_defs/common"
	pb_man "simpleGRPC/proto_defs/manager"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	ConfigFileName = "conf.ini"
)

const (
	GLOBAL_PORT       = 9001
	Red               = "\033[31m"
	Green             = "\033[32m"
	Yellow            = "\033[33m"
	Reset             = "\033[0m"
	magic       int32 = 0x45344534
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

// struct for persistent listeners config file
type PersistentListenersConf struct {
	Ports []int `json:"ports"`
}

type OperatorsConf struct {
	Operators map[string]string `json:"operators"`
}

var (

	// this initial config will be "inherited"
	// (to keep a common notifications channel between multiple service registration , aka Listeners)
	// need to make the chanell multi read to allow all managers to receive notif without race condition

	GlobalConf grpcConfig

	// track and manage listeners
	GlobalListeners map[int]chan bool

	PersistentListeners = &PersistentListenersConf{Ports: []int{}}

	OpConf = &OperatorsConf{Operators: map[string]string{}}

	kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: false,           // Allow pings even when there are no active streams
	}

	kasp = keepalive.ServerParameters{
		MaxConnectionAgeGrace: 5 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  5 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               5 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead

	}
)
