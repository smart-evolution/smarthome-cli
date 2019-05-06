package proxy

import (
	"bufio"
	"fmt"
	"github.com/smart-evolution/smarthome-cli/cmdapi"
	"github.com/smart-evolution/smarthome-cli/utils"
	"net"
	"os"
	"strings"
)

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

	if _, ok := cmdapi.ApiMap[devType]; !ok {
		fmt.Println("unknown device type '" + devType + "'")
		os.Exit(1)
	}

	fmt.Println("connected to device type '" + devType + "'")
	resBuff := make([]byte, 512)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("CMD: ")
		input, _ := reader.ReadString('\n')
		cmd := strings.TrimSpace(input)

		apiVersion := cmdapi.ApiMap[devType]
		hardwareComms := cmdapi.Comms[apiVersion][cmd]

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
			} else if c == "CMDDIS" {
				conn.Close()
				os.Exit(0)
			}
		}
	}
}
