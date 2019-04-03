package scan

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

var
(
	addrs []string
)

func Handler() {
	for i := 1; i <= 255; i++ {
		fmt.Print("\033[H\033[2J")
		addr := "192.168.1." + strconv.Itoa(i)
		fmt.Print("checking " + strconv.Itoa(i) + " address out of 255")

		d := net.Dialer{Timeout: time.Duration(500) * time.Millisecond}
		conn, err := d.Dial("tcp", addr + ":81")
		if err != nil {
			continue
		}

		_, err = conn.Write([]byte("CMDWHO"))

		if err != nil {
			continue
		}

		buff := make([]byte, 512)
		n, err := conn.Read(buff)

		if err != nil {
			continue
		}

		devType := string(buff[:n])
		addrs = append(addrs, addr + "::" + devType)
	}

	fmt.Println(strconv.Itoa(len(addrs)) + " compliant devices found")
	for _, dev := range addrs {
		fmt.Println("- " + dev)
	}
}
