package main

import (
	"kev/list"
	"sync"
	"time"
)

type Entry struct {
	Key   string
	Value string
	TTL   time.Time
}

type Store struct {
	mu       sync.RWMutex
	Data     map[string]*list.Element[Entry]
	lru      *list.List[Entry]
	capacity int
}

func NewStore(capacity int) *Store {
	return &Store{
		Data:     make(map[string]*list.Element[Entry]),
		lru:      list.New[Entry](),
		capacity: capacity,
	}
}


func (store *Store) CleanExpiry() {
	go func() {
		for {
			time.Sleep(time.Second)
			store.mu.Lock()
			for key, element := range store.Data {
				if time.Now().After(element.Value.TTL) {
					delete(store.Data, key)
					store.lru.Remove(element)
				}
			}
			store.mu.Unlock()
		}
	}()
}
func (store *Store) Get(key string) (string, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()
	element, ok := store.Data[key]
	if !ok {
		return "", false
	}
	store.lru.MoveToFront(element)
	return element.Value.Value, true
}

func (store *Store) Set(key, value string) {
	store.mu.Lock()
	defer store.mu.Unlock()

	if element, ok := store.Data[key]; ok {
		element.Value.Value = value
		store.lru.MoveToFront(element)
		return
	}

	if len(store.Data) >= store.capacity {
		back := store.lru.Back()
		if back != nil {
			delete(store.Data, back.Value.Key)
			store.lru.Remove(back)
		}
	}

	entry := Entry{Key: key, Value: value, TTL: time.Now().Add(time.Second * 60)}
	element := store.lru.PushFront(entry)
	store.Data[key] = element
}
func (store *Store) Delete(key string) string {
	store.mu.Lock()
	defer store.mu.Unlock()
	if element , ok := store.Data[key]; ok {
		deleted := element.Value.Value
		delete(store.Data, key)
		store.lru.Remove(element)

		return deleted
	}
	return "key not found"
}


