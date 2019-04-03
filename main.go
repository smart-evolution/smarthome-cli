//usr/local/go/bin/go run $0 $@; exit $?

package main

import (
	"fmt"
	"github.com/smart-evolution/smarthome-cli/commands/connect"
	"github.com/smart-evolution/smarthome-cli/commands/default"
	"github.com/smart-evolution/smarthome-cli/commands/scan"
	"github.com/smart-evolution/smarthome-cli/commands/send"
	"github.com/smart-evolution/smarthome-cli/commands/status"
	"github.com/smart-evolution/smarthome-cli/commands/version"
	"os"
)

//go:generate bash ./scripts/version.sh ./scripts/version_tpl.txt ./utils/version.go

func main() {
	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	} else {
		fmt.Println("please add cmd parameter after script")
		os.Exit(1)
	}

	switch cmd {
	case "connect":
		connect.Handler()
	case "send":
		send.Handler()
	case "status":
		status.Handler()
	case "scan":
		scan.Handler()
	case "version":
		version.Handler()
	default:
		_default.Handler()
	}
}
