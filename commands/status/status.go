package status

import (
	"fmt"
	"github.com/smart-evolution/smarthome-cli/utils"
	"net"
	"os"
)

func Handler() {
	conn, err := net.Dial("tcp", os.Getenv("SMARTHOME_CLI_SRV"))

	if err != nil {
		fmt.Println("error connecting to the smarthome cli server")
		os.Exit(1)
	}

	buff := make([]byte, 512)
	msg := utils.MsgConstructor("status")
	_, err = conn.Write(msg)
	n, err := conn.Read(buff)

	if err != nil {
		fmt.Println("error reading cli message")
	}

	response := string(buff[:n])
	fmt.Println(response)
}
