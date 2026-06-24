package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	store := make(map[string]string)
	Set(store, "name", "ahmet")
	name := Get(store, "name")
	fmt.Println(name)
	del := Delete(store, "name")
	fmt.Println("Deleted", del)

	empty := Get(store, "name")
	fmt.Println(empty)

	var i, j, k string

	fmt.Print(">")
	fmt.Scanf("%s %s %s", &i, &j, &k)
	fmt.Println(i, j, k)

	com := ParseCommand(fmt.Sprintf("%s %s %s", i, j, k))

	fmt.Println(com.Operand, com.Key, com.Value)

	com.Execute(store)
	age := Get(store, "age")
	fmt.Println(age)
}

