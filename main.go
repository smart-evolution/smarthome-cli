//usr/local/go/bin/go run $0 $@; exit $?

package main

import (
    "fmt"
    "github.com/smart-evolution/smarthome-cli/menu"
    "net"
    "os"
)

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
        fmt.Println("please add cmd parameter after script")
        os.Exit(1)
    }

    m := menu.New()
    m.Execute(cmd, conn)
}
