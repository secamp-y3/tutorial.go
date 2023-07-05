package main

import (
	"fmt"
	// "io"
	// "log"
	// "net"
	"bufio"
	"os"
	// "github.com/spf13/pflag"
)

var (
	count = 0
)

func input() string {
	// buf := []byte{}
	fmt.Printf("%d > ", count)
	stdin := bufio.NewScanner(os.Stdin)
	if stdin.Scan() {
		return stdin.Text()
	}
	return ""
}

func main() {
	// server := pflag.StringP("server", "s", "127.0.0.1:8080", "Server address")

	// addr, err := net.ResolveTCPAddr("tcp", *server)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// conn, err := net.Dial("tcp", addr.String())
	for {
		count += 1
		msg := input()
		fmt.Printf("[%d] %s\n", count, msg)
	}

}
