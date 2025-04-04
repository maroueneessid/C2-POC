package main

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "simpleGRPC/proto_defs/common"
	pb_man "simpleGRPC/proto_defs/manager"
)

func GetNotified(client pb_man.ManagerAssetClient) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := client.Subscribe(ctx, &pb_man.Notification{})

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

		fmt.Printf("\n%s%s %s%s", Yellow, notif.Notif, notif.SessionId, Reset)
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

func StartListenerOrder(client pb_man.ManagerAssetClient, port uint32) {

	newListener := &pb_man.Listener{
		Port: port,
	}
	client.StartNewListener(context.Background(), newListener)

}
