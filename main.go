package main

import (
	"bufio"
	"fmt"
	"kev/internal/command"
	"kev/internal/server"
	"kev/internal/store"
	"os"
	//	"weak"
	// "fmt"
	// "kev/internal/query"
)

func main() {

	store := store.NewStore(1024)
	store.CleanExpiry()
	scanner := bufio.NewScanner(os.Stdin)
	go server.StartServer(store)
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

	// var uq query.UnparsedQuery = "set age 25 -ttl=60 -type=int"
	// q := uq.Parse()
	// fmt.Println(q)
}
