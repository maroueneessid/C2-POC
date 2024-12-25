// package server
package main

import (
	"fmt"
	"strings"
)

func main() {

	GlobalConf = InitGrpcConfig(nil)
	RegisterManagerListener(GlobalConf.grpcServer, GlobalConf.serverConfig, GLOBAL_PORT)
	GlobalListeners = make(map[int]chan bool)
	fmt.Printf(Green+"[!] Started server on LOCALHOST:%d"+Reset+"\n", GLOBAL_PORT)

	InitPortPersistenceConfig(GlobalConf.serverConfig)

	for {

		var input string
		fmt.Print("Enter command 'exit' to stop server:\n> ")
		fmt.Scanln(&input)

		ciInput := strings.ToLower(input)

		ServerCommandHandler(GlobalConf, ciInput)

	}
}
