package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path"
	pb_man "simpleGRPC/proto_defs/manager"
)

func KillListenerHelper(c map[int]chan bool, target int) {
	for port, sigChan := range c {
		if port == target {
			sigChan <- false
			close(sigChan)
			delete(c, port)
		}
	}
	UpdatePortPersistenceConfig()
}

func ListenerPresence(c map[int]chan bool, target int) bool {
	_, exists := c[target]
	return exists
}

func KillAll(c map[int]chan bool) {
	for port, sigChan := range c {
		sigChan <- false
		close(sigChan)
		delete(c, port)

	}
	UpdatePortPersistenceConfig()
}

func InitPortPersistenceConfig(s *Server) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %v", err)
	}

	last_run := "last_run.ini"

	logDir := path.Join(homeDir, ".customC2")

	last_run_full_path := path.Join(logDir, last_run)

	f, err := os.OpenFile(last_run_full_path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		encoder := json.NewEncoder(f)
		if err := encoder.Encode(PersistentListeners); err != nil {
			return fmt.Errorf("failed to write default config to file: %v", err)
		}
	}

	jsonParser := json.NewDecoder(f)
	jsonParser.Decode(&PersistentListeners)

	fmt.Printf(Yellow+"[!] Found %d Listener from last run\n"+Reset, len(PersistentListeners.Ports))

	for _, port := range PersistentListeners.Ports {
		fmt.Printf(Yellow+"On %d\n"+Reset, port)
		s.StartNewListener(context.Background(), &pb_man.Listener{Port: uint32(port)})
	}

	return nil
}

func UpdatePortPersistenceConfig() error {

	logPath, err := GetLogPath("", "l")
	if err != nil {
		return fmt.Errorf("failed to get port persistence logs: %v", err)
	}
	f, err := os.OpenFile(logPath, os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to get port persistence logs: %v", err)
	}

	encoder := json.NewEncoder(f)

	var updated []int

	for k := range GlobalListeners {
		updated = append(updated, k)
	}

	PersistentListeners.Ports = updated
	err = encoder.Encode(PersistentListeners)
	if err != nil {
		return fmt.Errorf("failed to get update persistence log: %v", err)
	}

	defer f.Close()

	return nil
}
