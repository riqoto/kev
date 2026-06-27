package main

import (
	"bufio"
	"fmt"
	"kev/internal/command"
	"kev/internal/net"
	"kev/internal/store"
	"os"
)

func main() {

	store := store.NewStore(1024)
	store.CleanExpiry()
	scanner := bufio.NewScanner(os.Stdin)
	go net.StartServer(store)
	for {
		fmt.Print("> ")
		ok := scanner.Scan()
		if !ok {
			scanner.Err()
		}
		line := scanner.Text()

		if line == "exit" {
			break
		}

		com := command.ParseCommand(line)
		com.Execute(store, os.Stdout)
	}
}
