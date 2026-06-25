package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	store := NewStore(1024)
	store.CleanExpiry()
	scanner := bufio.NewScanner(os.Stdin)
	go StartServer(store)	
	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()

		if line == "exit" {
			break
		}

		com := ParseCommand(line)
		com.Execute(store, os.Stdout)
	}
}


