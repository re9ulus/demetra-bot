package main

import (
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"testing"
)

func TestRedisStorage(t *testing.T) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer mr.Close()
	s := NewRedisStorage(
		redis.NewClient(
			&redis.Options{Addr: mr.Addr()},
		),
	)
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

	user_id2 := 456
	mr.Set("456", "100")
	val = s.Get(user_id2)
	if val != 100 {
		t.Errorf("Storage `Get` expected %v, get %v", 100, val)
	}

}
