package main

import (
	"bufio"
	"log"
	"os"
	pb "simpleGRPC/proto_defs/common"
	pb_man "simpleGRPC/proto_defs/manager"
	"simpleGRPC/utils"
	"time"

	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Constants for colored output
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
	Clear  = "\033[H\033[J"
)

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             5 * time.Second,  // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

var session string = ""
var cli string = "Enter 'help' for more commands:\n" + Yellow + "[!] session: %s\n" + Reset + ">>"

func main() {

	host := flag.String("host", "localhost:9001", "Server to connect to. In format host:port")
	flag.Parse()
	flag.Usage()

	// load TLS and establish conn

	tlsCred, err := utils.SimpleClientTLS()

	if err != nil {
		log.Fatalf("[-] Error loading SSL Cert: %v", err)
	}
	conn, err := grpc.NewClient(*host,
		grpc.WithTransportCredentials(tlsCred),
		grpc.WithKeepaliveParams(kacp),
	)

	ErrorCheck("Could not connect to gRPC Server", err)
	defer conn.Close()

	client := pb_man.NewManagerAssetClient(conn)
	pingClient := pb.NewAssetServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		CommandHandling(pingClient, client, scanner)
	}

}
