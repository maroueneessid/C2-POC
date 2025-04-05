package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path"
	pb "simpleGRPC/proto_defs/common"
	pb_man "simpleGRPC/proto_defs/manager"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

func GetLogPath(sessionId string, logType string) (string, error) {

	// f for sessions tasks.log and d for sessions download dir

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}

	logDir := path.Join(homeDir, ".customC2")

	listenerLog := path.Join(logDir, ConfigFileName)

	sessionLog := path.Join(logDir, sessionId)

	sessionDownloads := path.Join(sessionLog, "downloads")

	logFile := path.Join(sessionLog, "tasks.log")

	tr := ""

	switch logType {
	case "f":
		tr = logFile
	case "d":
		tr = sessionDownloads
	case "l":
		tr = listenerLog

	}

	return tr, nil
}

func (s *Server) GetHistory(ctx context.Context, query *pb_man.HistoryQuery) (*pb_man.HistoryQuery, error) {

	log, err := GetSessionLog(query.SessionId)
	tr := &pb_man.HistoryQuery{
		SessionId: query.SessionId,
		History:   log,
	}

	return tr, err

}

func LogTasks(sessionId string, taskType string, taskIO string, operatorToken string) error {
	taskInSign := "<-"
	taskOutSign := "->"

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	logDir := path.Join(homeDir, ".customC2")
	// Ensure the directory exists
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %v", err)
	}

	sessionLog := path.Join(logDir, sessionId)
	if err := os.MkdirAll(sessionLog, 0755); err != nil {
		return fmt.Errorf("failed to create Session directory: %v", err)
	}

	sessionDownloads := path.Join(sessionLog, "downloads")
	if err := os.MkdirAll(sessionDownloads, 0755); err != nil {
		return fmt.Errorf("failed to create Session directory: %v", err)
	}

	logFile := path.Join(sessionLog, "tasks.log")

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("error opening log file %v: %v", logFile, err)
	}
	defer f.Close()

	log.SetOutput(f)

	operator := ""
	for name, tok := range OpConf.Operators {
		if bytes.Equal([]byte(operatorToken), []byte(tok)) {
			operator = name
		}
	}
	if taskType == "in" {
		log.Println("<" + operator + "> :\n" + taskInSign + " " + taskIO)
	} else if taskType == "out" {
		log.Println(taskOutSign + " " + taskIO)
	}

	return nil
}

func SaveDownloads(sessionId string, filename string, data []byte) error {
	downloadDir, err := GetLogPath(sessionId, "d")
	if err != nil {
		return fmt.Errorf("failed to get session download directory: %v", err)
	}

	downloadedFile := path.Join(downloadDir, filename)

	err = os.WriteFile(downloadedFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil

}

func GetSessionLog(sessionId string) (string, error) {

	logFile, err := GetLogPath(sessionId, "f")
	if err != nil {
		return "", fmt.Errorf("failed to read log file: %v", err)
	}

	fileBytes, err := os.ReadFile(logFile)
	if err != nil {
		return "", fmt.Errorf("failed to read log file: %v", err)
	}

	return string(fileBytes), nil
}

func (s *Server) GetAndSetSession(ctx context.Context, sessionId string, newSession *pb.Session) (*pb.Session, error) {

	serializedSession, err := s.db.Get(ctx, sessionId).Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("failed to get Session from Redis for SessionId %s: %v", sessionId, err)
	}

	var session pb.Session
	if err == nil {
		err = proto.Unmarshal([]byte(serializedSession), &session)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal Session: %v", err)
		}
	}

	if newSession != nil {
		newSerializedSession, err := proto.Marshal(newSession)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize Session: %v", err)
		}
		s.db.Set(ctx, sessionId, newSerializedSession, 0)
	}

	return &session, nil
}

func (s *Server) CheckSession(empty *pb.None, stream pb_man.ManagerAsset_CheckSessionServer) error {

	keys, _ := s.db.Keys(stream.Context(), "*").Result()

	for _, key := range keys {

		tosend, _ := s.GetAndSetSession(stream.Context(), key, nil)

		stream.Send(tosend)

	}

	return nil
}
