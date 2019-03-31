package connect

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

var comms = map[string]map[string][]string{
	"jeep": map[string][]string{
		"w": []string{"CMD012", "CMD022"},
		"s": []string{"CMD010", "CMD020"},
		"a": []string{"CMD012", "CMD122"},
		"d": []string{"CMD112", "CMD022"},
		"x": []string{"CMD112", "CMD122"},

	},
}

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

	_, err = conn.Write([]byte("CMDWHO"))

	if err != nil {
		fmt.Println("error getting device type")
		os.Exit(1)
	}

	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil {
		fmt.Println("error retrieving device type")
		os.Exit(1)
	}

	devType := string(buff[:n])
	fmt.Println("connected to device type '" + devType + "'")

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
			}
		}
	}
}
