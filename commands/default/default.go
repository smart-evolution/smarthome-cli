package _default

import (
	"fmt"
	"net"
)

func Handler(conn net.Conn) {
	fmt.Println("Invalid command")
}
