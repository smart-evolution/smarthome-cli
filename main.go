package main

import (
	"fmt"
    "net"
)

const (
	SERVICE_ADDR = "93.180.188.173:3333"
)

func main() {
    conn, err := net.Dial("tcp", SERVICE_ADDR)

    if err != nil {
        fmt.Println(err)
    }
    // fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    // status, err := bufio.NewReader(conn).ReadString('\n')

    msgBytes := []byte("----- some message")
    l, err := conn.Write(msgBytes)

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(l)
}
