package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/spf13/pflag"

	"github.com/secamp-y3/tutorial.go/02_rpc/protocol"
)

const (
	DEFAULT_HOST = "localhost"
	DEFAULT_PORT = 8080
)

// EchoHandler is an RPC handler that echoes back the given payload
type EchoHandler struct{}

// RequestEcho echoes back the given payload
func (s *EchoHandler) RequestEcho(args protocol.EchoRequestArgs, reply *protocol.EchoRequestReply) error {
	log.Printf(">>> %s\n", args.Payload)

	// echo back
	reply.Payload = args.Payload

	return nil
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

	// prepare RPC server with EchoHandler
	server := rpc.NewServer()
	if err := server.Register(&EchoHandler{}); err != nil {
		log.Fatal(err)
	}

	// main loop
	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal("Connection failed")
		}
		go server.ServeConn(conn)
	}
}
