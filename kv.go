package main

import (
	"strings"
	"time"
)

var ttl uint64 = 3600 

type Entry struct {
	Key string
	Value string
	TTL int // second
	ExpireAt time.Time  
}

type Store struct {
	Data map[string]Entry
	Path string
}

func NewStore() Store {
	store := Store {
		Data : make(map[string]Entry),
		Path: "",
	}

	return store
}


func NewEntry() Entry {
	entry := Entry {
		Key: "",
		Value: "",
		TTL: 3600,
		ExpireAt: time.Now().Add(time.Second *  3600),
	}
	
	return entry
}


func (store Store)Get(key string) (string, bool){

	entry, ok := store.Data[key]

	value := entry.Value

	return value,ok

}

func (store Store)Set(key, value string) {
	entry := NewEntry()
	
	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)
	
	entry.Key = key
	entry.Value = value
	
	store.Data[key] = entry
//	Save(store)
}

func (store Store)Delete(key string) string {

	if _,ok := store.Data[key]; ok {
		deleted := store.Data[key].Value
		delete(store.Data, key)
//		Save(store)
		return deleted
	}


	return "key not found"
}

