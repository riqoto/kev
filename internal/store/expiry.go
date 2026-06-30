package store

import "time"

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
