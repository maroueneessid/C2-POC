package main

import (
	"bufio"
	"os"
	pb "simpleGRPC/proto_defs"

	"flag"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Constants for colored output
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
	Clear  = "\033[H\033[J"
)

var session string = ""
var cli string = "Enter 'help' for more commands:\n" + Yellow + "[!] session: %s\n" + Reset + ">>"

func main() {

	host := flag.String("host", "localhost:9001", "Server to connect to. In format host:port")
	flag.Parse()
	flag.Usage()

	// Establish connection to gRPC server
	conn, err := grpc.NewClient(*host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	ErrorCheck("Could not connect to gRPC Server", err)
	defer conn.Close()

	client := pb.NewAssetServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		CommandHandling(client, scanner)
	}

}
