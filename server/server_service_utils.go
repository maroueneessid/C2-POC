package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"log"
	pb "simpleGRPC/proto_defs"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) RegisterAsset(ctx context.Context, assetRegistration *pb.AssetRegistration) (*pb.RegistrationConfirmation, error) {

	// Validate Magic Number
	if assetRegistration.MagicNb != magic {
		refusal := pb.RegistrationConfirmation{Confirmed: false}
		return &refusal, nil
	}

	// Send Notification to the manager through channel
	notif := &pb.Notification{
		SessionId: assetRegistration.SessionId,
		Notif:     "Connection received from ",
	}

	s.notifs <- notif

	// Set Entry in Redis for new Asset
	task := &pb.Task{
		In: &pb.TaskIO{
			Text:   "",
			Binary: []byte{},
		},
		Out: &pb.TaskIO{
			Text:   "",
			Binary: []byte{},
		},
	}

	session := &pb.Session{
		BasicInfo: assetRegistration,
		Task:      task,
		Alive:     true,
	}

	_, err := s.GetAndSetSession(ctx, assetRegistration.SessionId, session)
	if err != nil {
		log.Printf("Failed to set Session: %v", err)
	}

	// Confirm Registration to the Asset
	acceptance := pb.RegistrationConfirmation{Confirmed: true}
	return &acceptance, nil
}

func (s *Server) SendOrder(ctx context.Context, order *pb.ServerOrder) (*emptypb.Empty, error) {

	LogTasks(order.SessionId, "in", order.In.Text)

	session, err := s.GetAndSetSession(ctx, order.SessionId, nil)
	if err != nil {
		log.Printf("Failed to get Session from Redis for SessionId %s: %v", order.SessionId, err)
	}

	session.Task.In = order.In
	if order.In.Text == "exit" {
		session.Alive = false
	}

	_, err = s.GetAndSetSession(ctx, order.SessionId, session)
	if err != nil {
		log.Printf("Failed to update Session: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) SendResponse(ctx context.Context, response *pb.AssetResponse) (*emptypb.Empty, error) {
	if response.Out.Text != "" {
		//fmt.Println(Yellow, "[!] Received output\n", Reset, response.Out.Text)
		LogTasks(response.SessionId, "out", response.Out.Text)
		notif := &pb.Notification{
			SessionId: response.SessionId,
			Notif:     "Returned Task Output from ",
		}
		s.notifs <- notif
	}

	// Get the existing session and update the task with the response
	session, err := s.GetAndSetSession(ctx, response.SessionId, nil)
	if err != nil {
		log.Printf("Failed to get Session from Redis for SessionId %s: %v", response.SessionId, err)
	}

	session.Task.Out = response.Out

	// Set the updated session back to Redis
	_, err = s.GetAndSetSession(ctx, response.SessionId, session)
	if err != nil {
		log.Printf("Failed to update Session: %v", err)
	}

	if session.Task.Out.Text == "download" && len(session.Task.Out.Binary) != 0 {
		// need to extract filename instead of using md5 hash
		err = SaveDownloads(response.SessionId, fmt.Sprintf("%x", md5.Sum(session.Task.Out.Binary)), session.Task.Out.Binary)
		if err != nil {
			log.Printf("%v", err)
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) CheckIn(ctx context.Context, checkin *pb.AssetResponse) (*pb.ServerOrder, error) {
	session := checkin.SessionId

	// Ping function from the manager
	if session == "manager" {
		//log.Printf("Failed to get Session from Redis for SessionId %s", session)
		return &pb.ServerOrder{
			SessionId: session,
			In:        &pb.TaskIO{Text: "Ping", Binary: []byte{}},
		}, nil
	}

	// Get the existing session
	sessionData, err := s.GetAndSetSession(ctx, session, nil)
	if err != nil {
		log.Printf("Failed to get Session from Redis for SessionId %s: %v", session, err)
	}

	if sessionData.Task.In.Text == "" {
		return &pb.ServerOrder{
			SessionId: session,
			In:        &pb.TaskIO{Text: "", Binary: []byte{}},
		}, nil
	}

	// Prepare the response and clear the task input
	toReturn := &pb.ServerOrder{
		SessionId: session,
		In: &pb.TaskIO{
			Text:   sessionData.Task.In.Text,
			Binary: sessionData.Task.In.Binary,
		},
	}

	sessionData.Task.In.Text = ""
	sessionData.Task.In.Binary = []byte{}

	_, err = s.GetAndSetSession(ctx, session, sessionData)
	if err != nil {
		log.Printf("Failed to update Session: %v", err)
	}

	return toReturn, err
}
