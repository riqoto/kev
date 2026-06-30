package command

import (
	"fmt"
	"io"
	"kev/internal/store"
)

func (c *Command) Execute(store *store.Store, w io.Writer) {

	switch c.Operand {
	case "set":
		store.SetBytes(c.Key, []byte(c.Value))
		fmt.Fprintln(w, "ok")
	case "get":
		value, ok := store.GetBytes(c.Key)
		if ok {
			fmt.Fprintln(w, string(value))
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
