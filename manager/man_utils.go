package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "simpleGRPC/proto_defs"
	"time"
)

func Ping(client pb.AssetServiceClient) bool {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	checkIn := &pb.AssetResponse{
		SessionId: "manager",
		Out: &pb.TaskIO{
			Text:   "",
			Binary: []byte{},
		},
	}
	defer cancel()

	pinged, _ := client.CheckIn(ctx, checkIn)

	select {
	case <-ctx.Done():
		log.Fatalf(Red+"[-] Server shutdown: %v"+Reset, ctx.Err())
		return false
	default:
		if pinged.In.Text == "Ping" {
			return true
		}

	}

	return false
}

func GetNotified(client pb.AssetServiceClient) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Subscribe(ctx, &pb.Notification{})

	if err != nil {
		log.Fatalf("error subscribing: %v", err)
	}

	for {
		notif, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving notifs: %v", err)

		}

		fmt.Printf("\n[!] %s%s %s%s", Yellow, notif.Notif, notif.SessionId, Reset)
	}

}

func ErrorCheck(context string, err error) {
	if err != nil {
		log.Fatalf("%s : %v", context, err)
	}
}

func BuildOrder(cmd string, binContent []byte) *pb.ServerOrder {

	if binContent == nil || len(binContent) == 0 {
		return &pb.ServerOrder{
			SessionId: session,
			In: &pb.TaskIO{
				Text:   cmd,
				Binary: []byte{},
			},
		}
	}

	return &pb.ServerOrder{
		SessionId: session,
		In: &pb.TaskIO{
			Text:   cmd,
			Binary: binContent,
		},
	}

}
