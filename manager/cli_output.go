package main

import (
	"os"
	pb "simpleGRPC/proto_defs/common"
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
	table.SetHeader([]string{"SessionId", "Hostname", "Username", "IPs", "OS", "Last Task"})

	table.Append(infoSlice)

	table.Render()

}

func PrintHelp() string {
	return `
Commands:
  use <sessionId>	: Switch to a specific session
  exec shell <cmd>	: Execute a cmd/shell command on asset
  upload			: Uploads src to dest. "upload path/to/file destination/path/file"
  download			: Downloads specified file to sessions download dir.
  sessions [alive]	: Displays all sessions. "sessions alive" returns only alive sessions
  history			: fetches log file for specific session (for now fetches the whole file)
  listen <port>		: Start listener on specified port
  stop <port>		: Stops listener on specfied port
  kill [all]		: kills a session (only valid when inside a session). "kill all" inside or outside any sessions to kill all active sessions
  clear				: clear terminal
  help				: Display this help message
  exit				: Exit the manager
`
}
