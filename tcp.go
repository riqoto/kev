package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
//	"os"
)


// create TCPConfig or use default if exist for port and conn type
func StartServer(store *Store) {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("failed to connect: ", err)
	}
	fmt.Println("Connected")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept: ", err)
		}

		go handleConnection(conn, store)
	}
}

func handleConnection(conn net.Conn, store *Store) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	
	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("failed to read: ", err)
			return
		}

		command := ParseCommand(strings.ToLower(strings.TrimSpace(message)))
		command.Execute(store, conn)
	}
}


