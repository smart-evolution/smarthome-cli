package main

import (
	"fmt"
    "net"
    "os"
)

func msgConstructor(cmd string) []byte {
    msgString :=  `{
        "cmd": "` + cmd + `"
    }`

    return []byte(msgString)
}

func main() {
    conn, err := net.Dial("tcp", os.Getenv("SMARTHOME_CLI_SRV"))

    if err != nil {
        fmt.Println("error connecting to the smarthome cli server")
        os.Exit(1)
    }

    var cmd string
    if len(os.Args) > 1 {
        cmd = os.Args[1]
    } else {
        cmd = ""
    }

    var response string
    buff := make([]byte, 512)

    switch cmd {
    case "status":
        msg := msgConstructor("status")
        _, err = conn.Write(msg)
        n, err := conn.Read(buff)

        if err != nil {
            fmt.Println("error reading cli message")

        }

        response = string(buff[:n])
    default:
        response = "Invalid command"
    }


    fmt.Println(response)
}
