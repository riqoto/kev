package store

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

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

func (store *Store) SetNumeric(key string, value Number) bool {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, value)

	if err != nil {
		fmt.Println("failed to write numeric value: ", err)
		return false
	}

	if ok := store.SetBytes(key, buf.Bytes()); ok {
		return ok
	}

	return false

}
