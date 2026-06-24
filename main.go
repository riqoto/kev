package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	store := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := scanner.Text()

		if line == "exit" {
			break
		}

		com := ParseCommand(line)
		com.Execute(store)
	}
}


