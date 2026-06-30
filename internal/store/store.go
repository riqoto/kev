package store

import (
	"kev/internal/list"
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
