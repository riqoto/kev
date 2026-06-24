package main

import "strings"

func Get(store map[string]string, key string) (string, bool){
	value, ok := store[key]

	return value,ok

}

func Set(store map[string]string, key, value string) {
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	store[key] = value
}

func Delete(store map[string]string, key string) string {
	if _,ok := store[key]; ok {
		deleted := store[key]
		delete(store, key)
		return deleted
	}

	return "key not found"
}

