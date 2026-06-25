package main

import (
	"fmt"
	"strings"
)

type Command struct {
	Operand string
	Key     string
	Value   string
}

func ParseCommand(command string) Command {
	var com Command

	parsed := strings.Split(command, " ")

	if len(parsed) < 2 {
		return com
	}

	if len(parsed) > 3 {
		return com
	}

	com.Operand = strings.ToLower(parsed[0])
	com.Key = strings.ToLower(parsed[1])

	if com.Operand == "set" {
		com.Value = parsed[2]
	}

	return com
}

func (c *Command) Execute(store Store) {
	switch c.Operand {
	case "set":
		store.Set(c.Key, c.Value)
	case "get":
		fmt.Println(store.Get(c.Key))
	case "delete":
		store.Delete(c.Key)
	default:
		fmt.Println("Invalid Operand")
	}
}

