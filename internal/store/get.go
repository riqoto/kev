package store

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

func (store *Store) Get(key string) (string, bool) {

	if bytes, ok := store.GetBytes(key); ok {
		return string(bytes), ok
	}

	return "", false
}
