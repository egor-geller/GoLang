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
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		num, err := strconv.Atoi(message)
		if err == nil {
			fmt.Printf("Multiply by 2: %d\n", num*2)
		} else {
			newMessage := strings.ToUpper(message)
			conn.Write([]byte(newMessage + "\n"))
		}
	}
}
