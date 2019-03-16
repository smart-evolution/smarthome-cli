package connect

import (
	"fmt"
	"net"
	"os"
)

func Handler() {
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

	var cmd string
	if len(os.Args) > 3 {
		cmd = os.Args[3]
	} else {
		fmt.Println("device command required")
		os.Exit(1)
	}

	_, err = conn.Write([]byte(cmd))
	if err != nil {
		fmt.Println("error sending command " + cmd)
		os.Exit(1)
	}

	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil {
		fmt.Println("error reading message from device")
	}

	response := string(buff[:n])
	fmt.Println(response)
}
