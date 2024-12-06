package main

import  (
	"fmt"
	"net"
	"os"
	"time"
	"flag"
)

func connectAll(host string, ports []string) {
    for _, port := range ports {
	    socket := net.JoinHostPort(host, port)
	    connect(socket)
    }
}

func connect(socket string) {
	timeout := 15 * time.Second
	fmt.Printf("Is %s open...  ", socket)
	conn, err := net.DialTimeout("tcp", socket, timeout)
	if err != nil {
	    fmt.Println("Connecting error:", err)
	}
	if conn != nil {
	    defer conn.Close()
	    fmt.Printf("... Opened %s successfully\n", socket)
	}
}

func main() {
	socketFlag := flag.String("socket", "", "remote host and port to test. format 'host:port'")
	hostFlag := flag.String("host", "", "remote host to test")
	portsFlag := flag.String("ports", "", "ports to test. accepts single port or multiple ports comma separated, or in a range")
	flag.Parse()

	if *socketFlag != "" {
		connect(*socketFlag)
		os.Exit(0)
	}

	if *hostFlag != "" && *portsFlag != "" {
		connectAll(*hostFlag, []string{*portsFlag})
		os.Exit(0)
	}
	fmt.Println("no data passed in")
	os.Exit(-1)
}
