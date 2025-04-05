package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
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

	logDir := path.Join(homeDir, ".customC2")
	lastRunPath := path.Join(logDir, ConfigFileName)

	f, err := os.OpenFile(lastRunPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file: %v", err)
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file info: %v", err)
	}

	if fileInfo.Size() == 0 {
		// First run: save empty/default config
		initialData := struct {
			PersistentListeners *PersistentListenersConf `json:"persistentListeners"`
			OperatorsConf       *OperatorsConf           `json:"operatorsConf"`
		}{
			PersistentListeners: PersistentListeners,
			OperatorsConf:       OpConf,
		}

		encoder := json.NewEncoder(f)
		if err := encoder.Encode(initialData); err != nil {
			return fmt.Errorf("failed to write default config to file: %v", err)
		}
		fmt.Println("[*] Created default config file")
		return fmt.Errorf("initial config created; please edit and restart")
	}

	// Decode existing config
	decoder := json.NewDecoder(f)
	var rawConf struct {
		PersistentListeners *PersistentListenersConf `json:"persistentListeners"`
		OperatorsConf       *OperatorsConf           `json:"operatorsConf"`
	}
	if err := decoder.Decode(&rawConf); err != nil {
		return fmt.Errorf("failed to decode config: %v", err)
	}

	// Apply config
	if rawConf.PersistentListeners != nil {
		PersistentListeners = rawConf.PersistentListeners
	} else {
		return fmt.Errorf("missing persistentListeners section in config")
	}

	if rawConf.OperatorsConf != nil && len(rawConf.OperatorsConf.Operators) > 0 {
		OpConf = rawConf.OperatorsConf
	} else {
		log.Fatal("[X] Invalid or missing operator configuration. Must specify at least one operator.")
	}

	fmt.Printf(Yellow+"[!] Found %d Listener(s) from last run\n"+Reset, len(PersistentListeners.Ports))

	for _, port := range PersistentListeners.Ports {
		fmt.Printf(Yellow+"    * Port %d\n"+Reset, port)
		s.StartNewListener(context.Background(), &pb_man.Listener{Port: uint32(port)})
	}

	return nil
}

func UpdatePortPersistenceConfig() error {
	logPath, err := GetLogPath("", "l")
	if err != nil {
		return fmt.Errorf("failed to get port persistence logs: %v", err)
	}

	// Step 1: Read existing config
	fileData, err := os.ReadFile(logPath)
	if err != nil {
		return fmt.Errorf("failed to read config: %v", err)
	}

	var conf struct {
		PersistentListeners *PersistentListenersConf `json:"persistentListeners"`
		OperatorsConf       *OperatorsConf           `json:"operatorsConf"`
	}

	err = json.Unmarshal(fileData, &conf)
	if err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	var updated []int
	for k := range GlobalListeners {
		updated = append(updated, k)
	}
	conf.PersistentListeners = &PersistentListenersConf{Ports: updated}

	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config for writing: %v", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&conf); err != nil {
		return fmt.Errorf("failed to write updated config: %v", err)
	}

	return nil
}
