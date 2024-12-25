// package server
package main

import (
	"fmt"
	"strings"
)

func main() {

	GlobalConf = InitGrpcConfig(nil)
	RegisterManagerListener(GlobalConf.grpcServer, GlobalConf.serverConfig, 9001)

	GlobalListeners = make(map[int]chan bool)

	fmt.Printf(Green + "[!] Started server on LOCALHOST:9001" + Reset + "\n")

	for {

		var input string
		fmt.Print("Enter command 'exit' to stop server:\n> ")
		fmt.Scanln(&input)

		ciInput := strings.ToLower(input)

		ServerCommandHandler(GlobalConf, ciInput)

	}
}
