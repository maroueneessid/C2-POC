// package server
package main

import (
	"fmt"
	"strings"
)

// this initial config will be "inherited"
// (to keep a common notifications channel between multiple service registration , aka Listeners)
var GlobalConf grpcConfig

// keep track of listeners
var GlobalListeners []int

func main() {

	GlobalConf = InitGrpcConfig(nil)
	RegisterListener(GlobalConf.grpcServer, GlobalConf.serverConfig, 9001)
	fmt.Printf(Green + "[!] Started server on LOCALHOST:9001" + Reset + "\n")

	for {

		var input string
		fmt.Print("Enter command 'exit' to stop server:\n> ")
		fmt.Scanln(&input)

		ciInput := strings.ToLower(input)

		ServerCommandHandler(GlobalConf, ciInput)

	}
}
