//usr/local/go/bin/go run $0 $@; exit $?

package main

import (
	"fmt"
	"github.com/smart-evolution/smarthome-cli/commands/connect"
	"github.com/smart-evolution/smarthome-cli/commands/default"
	"github.com/smart-evolution/smarthome-cli/commands/status"
	"os"
)

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
	case "status":
		status.Handler()
	default:
		_default.Handler()
	}
}
