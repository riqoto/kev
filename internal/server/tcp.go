package server

import (
	"bufio"
	"fmt"
	"kev/internal/command"
	"kev/internal/store"
	"net"
	"strings"
)

// create TCPConfig or use default if exist for port and conn type
func StartServer(store *store.Store) {
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

func handleConnection(conn net.Conn, store *store.Store) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("failed to read: ", err)
			return
		}

		query := command.ParseCommand(strings.ToLower(strings.TrimSpace(message)))
		query.Execute(store, conn)
	}
}
