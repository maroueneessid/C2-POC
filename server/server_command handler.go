package main

import (
	"fmt"
	"os"
	"strings"
)

func ServerCommandHandler(conf grpcConfig, in string) {

	for {
		if len(strings.Fields(in)) == 0 {
			break
		}
		switch strings.Fields(in)[0] {
		case "exit":
			fmt.Println("Shutting down the server...")
			conf.grpcServer.Stop()
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Please try again.")
			in = ""
		}
	}

}
