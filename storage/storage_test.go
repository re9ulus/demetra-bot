package storage

import (
	"testing"
)

func TestInMemroyStorage(t *testing.T) {
	s := NewInMemoryStorage()
	user_id := 123
	val := s.Get(user_id)
	if val != 0 {
		t.Errorf("Empty storage `Get` expected 0, get %v", val)
	}
	s.Add(user_id, 42)
	val = s.Get(user_id)
	if val != 42 {
		t.Errorf("Storage `Get` expected %v, get %v", 42, val)
	}
	s.Spent(user_id, 13)
	val = s.Get(user_id)
	if val != 29 {
		t.Errorf("Stoarge `Get` expected %v, get %v", 29, val)
	}
}
