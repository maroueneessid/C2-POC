package main

import (
	"context"
	"log"
	"os"
	pb "simpleGRPC/proto_defs/common"
	"simpleGRPC/utils"
	"time"

	"flag"

	"google.golang.org/grpc"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

const magic int32 = 0x45344534

var session string = ""

func main() {

	connUrl := flag.String("host", "localhost:8080", "host to connect to. In format host:port")

	flag.Parse()
	flag.Usage()

	tlsCred, err := utils.SimpleClientTLS()
	if err != nil {
		log.Fatalf("[-] Error loading SSL Cert: %v", err)
	}
	conn, err := grpc.NewClient(*connUrl, grpc.WithTransportCredentials(tlsCred))

	if err != nil {
		Error_check("could not connect to GRPCServer", err)
	}

	defer conn.Close()

	c := pb.NewAssetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	registration, err := RegisterForm()

	Error_check("[-] Error generating registration form", err)

	r, err := c.RegisterAsset(ctx, registration)

	if err != nil {
		log.Fatal("[-] No response to the form was received")
	}

	if !r.Confirmed {
		os.Exit(0)
	}

	for {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

		checkIn := &pb.AssetResponse{
			SessionId: session,
			Out: &pb.TaskIO{
				Text:   "",
				Binary: []byte{},
			},
		}
		order, _ := c.CheckIn(ctx, checkIn)

		textOut, binOut := OrderHandler(order)

		time.Sleep(time.Duration(5) * time.Second)

		msg := pb.AssetResponse{
			SessionId: session,
			Out: &pb.TaskIO{
				Text:   textOut,
				Binary: binOut,
			},
		}

		c.SendResponse(ctx, &msg)

		cancel()
	}

}
