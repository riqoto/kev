package main

import "testing"
func lruKeys(s *Store) []string {
	keys := []string{}

	node := s.lru.Front()
	for node != nil {
		keys = append(keys, node.Value.Key)
		node = node.Next()
	}

	return keys
}
func assertLRU(t *testing.T, s *Store, expected []string) {
	got := lruKeys(s)

	if len(got) != len(expected) {
		t.Fatalf("expected %v got %v", expected, got)
	}

	for i := range expected {
		if got[i] != expected[i] {
			t.Fatalf("expected %v got %v", expected, got)
		}
	}
}
func TestLRUOrder(t *testing.T) {
	s := NewStore(3)

	s.Set("A", "1")
	assertLRU(t, s, []string{"A"})

	s.Set("B", "2")
	assertLRU(t, s, []string{"B", "A"})

	s.Set("C", "3")
	assertLRU(t, s, []string{"C", "B", "A"})

	s.Get("A")
	assertLRU(t, s, []string{"A", "C", "B"})

	s.Set("D", "4")
	assertLRU(t, s, []string{"D", "A", "C"})

	s.Get("C")
	assertLRU(t, s, []string{"C", "D", "A"})
}
func TestLRUEviction(t *testing.T) {

    s := NewStore(3)

    s.Set("A", "1")
    s.Set("B", "2")
    s.Set("C", "3")

    s.Get("A")
    s.Set("D", "4")

    if _, ok := s.Get("B"); ok {
        t.Fatal("B should have been evicted")
    }

    if _, ok := s.Get("A"); !ok {
        t.Fatal("A should still exist")
    }
}
