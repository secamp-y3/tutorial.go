package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"sync"

	"github.com/spf13/pflag"

	"github.com/secamp-y3/tutorial.go/03_state-machine/protocol"
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

// StateMachine manages the state of the server
type StateMachine struct {
	Value int
	mutex sync.Mutex
}

// RequestCommand handles operation commands from clients
func (s *StateMachine) RequestCommand(args protocol.RequestCommandArgs, reply *protocol.RequestCommandReply) error {
	s.mutex.Lock() // start critical section
	oldValue := s.Value
	newValue, err := args.Operator.Apply(oldValue, args.Operand)
	if err != nil {
		s.mutex.Unlock() // end critical section [case: error]
		reply.Value = oldValue
		reply.Error = err
		return nil
	}
	s.Value = newValue
	s.mutex.Unlock() // end critical section [case: normal]

	log.Printf("State update: %d -> %s(%d) -> %d\n", oldValue, args.Operator, args.Operand, newValue)
	reply.Value = newValue
	reply.Error = nil
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
	if err := server.Register(&EchoHandler{}); err != nil { // register EchoHandler
		log.Fatal(err)
	}
	if err := server.Register(&StateMachine{0, sync.Mutex{}}); err != nil { // register StateMachine
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
