package store

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"kev/list"
	"sync"
	"time"
)

type Entry struct {
	Key   string
	Value []byte
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
func (store *Store) GetBytes(key string) ([]byte, bool) {
	store.mu.Lock()
	defer store.mu.Unlock()
	element, ok := store.Data[key]
	if !ok {
		return nil, false
	}
	store.lru.MoveToFront(element)
	return element.Value.Value, true
}

func (store *Store) SetBytes(key string, value []byte) bool {
	store.mu.Lock()

	defer store.mu.Unlock()

	if element, ok := store.Data[key]; ok {
		element.Value.Value = value
		store.lru.MoveToFront(element)
		return true
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
	return true
}

func (store *Store) Get(key string) (string, bool) {

	if bytes, ok := store.GetBytes(key); ok {
		return string(bytes), ok
	}

	return "", false
}

func (store *Store) SetNumeric(key string, value Number) bool {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, value)

	if err != nil {
		fmt.Println("failed to write numeric value: ", err)
		return false
	}

	if ok := store.SetBytes(key, []byte(buf)); ok {
		return ok
	}

	return false

}

func (store *Store) Delete(key string) string {
	store.mu.Lock()
	defer store.mu.Unlock()
	if element, ok := store.Data[key]; ok {
		deleted := element.Value.Value
		delete(store.Data, key)
		store.lru.Remove(element)

		return string(deleted)
	}
	return "key not found"
}
