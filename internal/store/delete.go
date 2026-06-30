package store

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
