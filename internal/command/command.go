package command

import (
	"fmt"
	"io"
	"kev/internal/store"
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

func (c *Command) Execute(store *store.Store, w io.Writer) {

	switch c.Operand {
	case "set":
		store.Set(c.Key, c.Value)
		fmt.Fprintln(w, "ok")
	case "get":
		value, ok := store.Get(c.Key)
		if ok {
			fmt.Fprintln(w, value)
		} else {
			fmt.Fprintln(w, "nil")
		}

	case "delete":
		store.Delete(c.Key)
		fmt.Fprintln(w, "ok")
	default:
		fmt.Fprintln(w, "Invalid Operand")

	}
}
