package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"

	"github.com/spf13/pflag"

	"github.com/secamp-y3/tutorial.go/02_rpc/protocol"
)

const (
	SERVER_ADDRESS = "localhost:8080"
)

func getEnvOr(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return def
}

// readInput returns the string given via STDIN
func readInput() string {
	fmt.Print("> ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return stdin.Text()
}

func main() {
	// --server SERVER: server address
	server := pflag.StringP("server", "s", getEnvOr("SERVER", SERVER_ADDRESS), "Server address")

	addr, err := net.ResolveTCPAddr("tcp", *server)
	if err != nil {
		log.Fatal(err)
	}

	// main loop
	for {
		msg := readInput()

		conn, err := rpc.Dial("tcp", addr.String())
		if err != nil {
			log.Fatal(err)
		}

		args := protocol.EchoRequestArgs{Payload: msg} // RPC argument
		var reply protocol.EchoRequestReply            // RPC reply holder
		if err := conn.Call("EchoHandler.RequestEcho", args, &reply); err != nil {
			log.Fatal(err)
		}
		fmt.Println(reply.Payload)
	}
}
