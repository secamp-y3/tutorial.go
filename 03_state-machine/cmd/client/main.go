package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"

	"github.com/secamp-y3/tutorial.go/03_state-machine/protocol"
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
func readInput() []string {
	fmt.Print("> ")
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	return strings.Split(stdin.Text(), " ")
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
		input := readInput()

		conn, err := rpc.Dial("tcp", addr.String())
		if err != nil {
			log.Fatal(err)
		}

		/*
		 * callEcho と callOperation の戻り値とエラーを別々に受け取って
		 * 個別に画面に表示する処理を書くのが面倒なので，どちらともチャネルを介して受け取ってみる．
		 */
		success := make(chan string, 1)
		failure := make(chan error, 1)
		if input[0] == "echo" {
			go callEcho(conn, strings.Join(input[1:], " "), success, failure)
		} else {
			go callOperation(conn, input, success, failure)
		}

		/*
		 * success か failure のどちらか一方に値が渡されるまで待機
		 */
		select {
		case reply := <-success:
			fmt.Println(reply)
		case err := <-failure:
			fmt.Printf("[ERROR] %s\n", err)
		}
	}
}

func callEcho(conn *rpc.Client, msg string, s chan<- string, f chan<- error) {
	defer conn.Close()

	args := protocol.EchoRequestArgs{Payload: msg} // RPC argument
	var reply protocol.EchoRequestReply            // RPC reply holder
	if err := conn.Call("EchoHandler.RequestEcho", args, &reply); err != nil {
		f <- err
		return
	}
	s <- reply.Payload
}

func callOperation(conn *rpc.Client, input []string, s chan<- string, f chan<- error) {
	defer conn.Close()

	if len(input) < 2 {
		f <- fmt.Errorf("Invalid form of command: %v", input)
		return
	}

	operator, err := protocol.Translate(input[0])
	if err != nil {
		f <- err
		return
	}

	operand, err := strconv.Atoi(input[1])
	if err != nil {
		f <- err
		return
	}

	var reply protocol.RequestCommandReply
	if err := conn.Call("StateMachine.RequestCommand", protocol.RequestCommandArgs{Operator: operator, Operand: operand}, &reply); err != nil {
		f <- err
		return
	}

	if reply.Error != nil {
		f <- err
		return
	}
	s <- fmt.Sprintf("Updated: %d", reply.Value)
}
