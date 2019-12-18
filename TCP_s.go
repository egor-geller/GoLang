package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Launching server...")
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Errorf("Connection problem")
	}
	conn, err1 := ln.Accept()
	if err1 != nil {
		fmt.Errorf("Connection problem")
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Errorf("Connection problem")
		}
		num, err := strconv.Atoi(message)
		if err == nil {
			fmt.Printf("Multiply by 2: %d\n", num*2)
		} else {
			newMessage := strings.ToUpper(message)
			conn.Write([]byte(newMessage + "\n"))
		}
	}
}
