package main

import (
	"context"
	"fmt"
	"io"
	pb "simpleGRPC/proto_defs"
	"time"
)

func SendOrderFromManager(client pb.AssetServiceClient, order *pb.ServerOrder) error {

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	_, err := client.SendOrder(ctx, order)
	if err != nil {
		fmt.Println(Red+"Error executing command:"+Reset, err)
		cancel()
		return err
	}

	cancel()

	return nil
}

func GetAllSessions(client pb.AssetServiceClient, onlyAlive bool) {

	req := pb.None{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	stream, err := client.CheckSession(ctx, &req)
	ErrorCheck("[-] Error getting stream of Session", err)

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

		} else {
			ErrorCheck("[-] Error reading some message in stream", err)
		}

		if msg.Alive {
			tableWriter(msg)
		} else {
			if onlyAlive {
				continue
			}
			fmt.Print(Red)
			tableWriter(msg)
			fmt.Print(Reset + "\n")

		}

	}
}

func KillAll(client pb.AssetServiceClient) {
	req := pb.None{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	stream, err := client.CheckSession(ctx, &req)
	ErrorCheck("[-] Error getting stream of Session", err)

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

		} else {
			ErrorCheck("[-] Error reading some message in stream", err)
		}

		if msg.Alive {

			order := &pb.ServerOrder{
				SessionId: msg.BasicInfo.SessionId,
				In: &pb.TaskIO{
					Text:   "exit",
					Binary: []byte{},
				},
			}
			session = ""
			rt := SendOrderFromManager(client, order)
			if rt != nil {
				continue
			}
		}

	}
}

func FetchSessionHistory(client pb.AssetServiceClient, sessionId string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	topass := &pb.HistoryQuery{
		SessionId: sessionId,
		History:   "",
	}

	tr, err := client.GetHistory(ctx, topass)
	if err != nil {
		return fmt.Sprintf("[-] %v", err)
	}

	return tr.History

}
