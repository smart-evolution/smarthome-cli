package scan

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

var addrs []string

func scan(wg *sync.WaitGroup, ip string) {
	defer wg.Done()

	d := net.Dialer{Timeout: time.Duration(1000) * time.Millisecond}
	conn, err := d.Dial("tcp", ip+":81")
	if err != nil {
		return
	}

	_, err = conn.Write([]byte("CMDWHO"))

	if err != nil {
		return
	}

	buff := make([]byte, 512)
	n, err := conn.Read(buff)

	if err != nil {
		return
	}

	devType := string(buff[:n])
	addrs = append(addrs, ip+"::"+devType)
}

func Handler() {
	var wg sync.WaitGroup

	for i := 1; i <= 255; i++ {
		ip := "192.168.1." + strconv.Itoa(i)
		wg.Add(1)
		go scan(&wg, ip)
	}
	wg.Wait()

	fmt.Println(strconv.Itoa(len(addrs)) + " compliant devices found")
	for _, dev := range addrs {
		fmt.Println("- " + dev)
	}
}
