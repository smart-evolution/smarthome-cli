package status

import (
	"fmt"
	"github.com/smart-evolution/smarthome-cli/utils"
	"net"
)

func Handler(conn net.Conn) {
	buff := make([]byte, 512)
	msg := utils.MsgConstructor("status")
	_, err := conn.Write(msg)
	n, err := conn.Read(buff)

	if err != nil {
		fmt.Println("error reading cli message")
	}

	response := string(buff[:n])
	fmt.Println(response)
}
