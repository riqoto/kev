package main

import (
	"strings"
	"sync"
	"time"
)

type Entry struct {
	Key   string
	Value string
	TTL   time.Time
}

type Store struct {
	mu   sync.RWMutex
	Data map[string]Entry
}

func NewStore() *Store {
	return &Store{
		Data: make(map[string]Entry),
	}
}

func NewEntry() Entry {
	entry := Entry{
		Key:   "",
		Value: "",
		TTL:   time.Now().Add(time.Second * 3600),
	}

	return entry
}

func (store *Store) Get(key string) (string, bool) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	entry, ok := store.Data[key]

	value := entry.Value

	return value, ok

}

func (store *Store) Set(key, value string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	entry := NewEntry()

	key = strings.TrimSpace(key)
	value = strings.TrimSpace(value)

	entry.Key = key
	entry.Value = value

	store.Data[key] = entry

}

func (store Store) Delete(key string) string {
	store.mu.Lock()
	defer store.mu.Unlock()
	if _, ok := store.Data[key]; ok {
		deleted := store.Data[key].Value
		delete(store.Data, key)

		return deleted
	}
	return "key not found"
}

