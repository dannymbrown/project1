package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	log.Println("Starting server...")

	// listen on port 8000
	ln, _ := net.Listen("tcp", ":4400")

	// accept connection

	// run loop forever (or until ctrl-c)
	for {
		// get message, output
		conn, _ := ln.Accept()
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("\nMessage Received: ", string(message))
		conn.Close()
	}
}
