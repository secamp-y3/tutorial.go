package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/spf13/pflag"
)

const (
	SERVER_ADDRESS  = "localhost:8080"
	BUFFER_CAPACITY = 1024
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

// send transmits a messege and receives a reply via the given connection
func send(conn net.Conn, msg string, reply *string) error {
	defer conn.Close()

	// send msg as an array of bytes
	conn.Write([]byte(msg))

	// receive reply from server as an array of bytes
	buf := make([]byte, BUFFER_CAPACITY)
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}

	*reply = string(buf[:n])
	return nil
}

func main() {
	// --server SERVER: server address
	server := pflag.StringP("server", "s", getEnvOr("SERVER", SERVER_ADDRESS), "Server address")

	addr, err := net.ResolveTCPAddr("tcp", *server)
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg := readInput()

		conn, err := net.Dial("tcp", addr.String())
		if err != nil {
			log.Fatal(err)
		}

		var reply string
		if send(conn, msg, &reply) != nil {
			log.Fatal(err)
		}
		fmt.Println(reply)
	}
}
