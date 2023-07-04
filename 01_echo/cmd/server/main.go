package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/spf13/pflag"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	var buf []byte
	conn.Read(buf)
	io.WriteString(conn, string(buf))
}

func main() {
	port := pflag.IntP("port", "p", 8080, "Port number to listen")
	pflag.Parse()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", *port))
	log.Print(fmt.Sprintf("Listning: %s", addr))
	check(err)
	socket, err := net.ListenTCP("tcp", addr)
	check(err)
	for {
		conn, err := socket.Accept()
		check(err)
		go handle(conn)
	}
}
