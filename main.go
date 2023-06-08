package main

import  (
	"fmt"
	"net"
	"os"
	"time"
)

func raw_connect(host string, ports []string) {
    for _, port := range ports {
        timeout := time.Second
        conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
        if err != nil {
            fmt.Println("Connecting error:", err)
        }
        if conn != nil {
            defer conn.Close()
            fmt.Println("Opened", net.JoinHostPort(host, port))
        }
    }
}

func main() {
	host := os.Args[1]
	ports := os.Args[2:]
	raw_connect(host, ports)
}
