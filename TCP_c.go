package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Errorf("Reading error")
		}
		if strings.TrimSpace(text) == "exit" {
			fmt.Println("TCP client exiting...")
			return
		} else {
			fmt.Fprintf(conn, text+"\n")
		}
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Errorf("Connection problem")
		}
		fmt.Print("Message from server: " + message)
	}
}
