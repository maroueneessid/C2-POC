package main

import (
	"os"
	pb "simpleGRPC/proto_defs"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func tableWriter(session *pb.Session) {

	sessionId := session.BasicInfo.SessionId
	hostname := session.BasicInfo.Hostname
	username := session.BasicInfo.Username
	ip := strings.Join(session.BasicInfo.IP, "\n")
	operatingSystem := session.BasicInfo.OS
	lastTask := "[!] Use 'history' command"

	infoSlice := []string{
		sessionId,
		hostname,
		username,
		ip,
		operatingSystem,
		lastTask,
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"SessionId", "Hostname", "Username", "OS", "IPs", "Last Task"})

	table.Append(infoSlice)

	table.Render()

}

func PrintHelp() string {
	return `
Commands:
  use <sessionId>	: Switch to a specific session
  exec <command>	: Execute a command on the server (only 'shell' for now)
  upload			: Uploads src to dest. "upload path/to/file destination/path/file"
  download			: Downloads specified file to sessions download dir.
  sessions [alive]	: Displays all sessions. "sessiions alive" returns only alive sessions
  history			: fetches log file for specific session (for now fetches the whole file)
  kill				: kills a session (only valid when inside a session)
  clear				: clear terminal
  help				: Display this help message
  exit				: Exit the manager
`
}
