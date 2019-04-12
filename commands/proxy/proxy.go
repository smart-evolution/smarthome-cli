package proxy

import (
	"bufio"
	"fmt"
	"github.com/smart-evolution/smarthome-cli/utils"
	"net"
	"os"
	"strings"
)

var comms = map[string]map[string][]string{
	"jeep": map[string][]string{
		"s": []string{"CMD010", "CMD020"},
		"w": []string{"CMD010", "CMD020", "CMD012", "CMD022"},
		"a": []string{"CMD010", "CMD020", "CMD012", "CMD122"},
		"d": []string{"CMD010", "CMD020", "CMD112", "CMD022"},
		"x": []string{"CMD010", "CMD020", "CMD112", "CMD122"},
		"l": []string{"CMDLOK"},
	},
}

func Handler() {
	conn, err := net.Dial("tcp", os.Getenv("SMARTHOME_CLI_SRV"))

	if err != nil {
		fmt.Println("error connecting to the smarthome cli server")
		os.Exit(1)
	}

	var device string
	if len(os.Args) > 2 {
		device = os.Args[2]
	} else {
		fmt.Println("device address required")
		os.Exit(1)
	}

	buff := make([]byte, 512)
	msg := utils.MsgConstructor("proxy", device)
	_, err = conn.Write(msg)
	n, err := conn.Read(buff)

	if err != nil {
		fmt.Println("error reading cli message")
	}

	devType := string(buff[:n])
	fmt.Println("connected to device type '" + devType + "'")
	resBuff := make([]byte, 512)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("CMD: ")
		input, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(input)

		if cmd == "disconnect" {
			conn.Close()
			break
		}

		hardwareComms := comms[devType][cmd]

		for _, c := range hardwareComms {
			_, err = conn.Write([]byte(c))
			if err != nil {
				fmt.Println("RES: sending command failed " + c)
				break
			}

			if c == "CMDLOK" {
				n, err := conn.Read(resBuff)

				if err != nil {
					fmt.Println("RES: error reading message from device")
					break
				}

				response := string(resBuff[:n])
				fmt.Println(response)
			}
		}
	}
}
