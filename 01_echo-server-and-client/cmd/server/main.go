package main

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/pflag"
)

const (
	DEFAULT_HOST    = "localhost"
	DEFAULT_PORT    = 8080
	BUFFER_CAPACITY = 1024
)

// provides an echo-server
func handle(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, BUFFER_CAPACITY)
	n, _ := conn.Read(buf)
	input := buf[:n]

	log.Printf("[%s] %s\n", conn.RemoteAddr().String(), input)

	conn.Write(input)
}

func main() {
	// --host HOST: hostname
	host := pflag.StringP("host", "h", DEFAULT_HOST, "Hostname")
	// --port PORT: port number
	port := pflag.IntP("port", "p", DEFAULT_PORT, "Port number to listen")
	pflag.Parse()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}

	socket, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()
	log.Print(fmt.Sprintf("Listning: %s", addr.String()))

	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal("Connection failed")
		}
		go handle(conn)
	}
}
