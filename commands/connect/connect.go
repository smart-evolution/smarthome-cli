package connect

import (
	"fmt"
	"net"
)

func Handler(conn net.Conn) {
	fmt.Println("connect")
}
