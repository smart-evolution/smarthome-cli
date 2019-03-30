package connect

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)


type CmdMap	map[string][]string

func Handler() {
	var jeepComm = make(CmdMap)

	jeepComm["forward"] = []string{"CMD012", "CMD022"}
	jeepComm["stop"] = []string{"CMD010", "CMD020"}

	var device string
	if len(os.Args) > 2 {
		device = os.Args[2]
	} else {
		fmt.Println("device address required")
		os.Exit(1)
	}

	conn, err := net.Dial("tcp", device)

	if err != nil {
		fmt.Println("error connecting device " + device)
		os.Exit(1)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("CMD: ")
		input, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(input)

		if cmd == "disconnect" {
			break
		}

		tC := jeepComm[cmd]

		for _, c := range tC {
			_, err = conn.Write([]byte(c))
			if err != nil {
				fmt.Println("RES: sending command failed " + c)
			}
		}
	}
}
