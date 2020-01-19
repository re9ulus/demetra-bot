package main

import (
	"fmt"
)

type Storage interface {
	Add(user_id int, val int64) error
	Get(user_id int) int64
	Spent(user_id int, val int64) error
}

type InMemoryStorage struct {
	amounts map[int]int64
}

func (st *InMemoryStorage) Add(user_id int, val int64) error {
	if val <= 0 {
		return fmt.Errorf("val must be positive")
	}
	st.amounts[user_id] += val
	return nil
}

func (st *InMemoryStorage) Spent(user_id int, val int64) error {
	if val <= 0 {
		return fmt.Errorf("val must be positive")
	}
	st.amounts[user_id] -= val
	return nil
}

func (st *InMemoryStorage) Get(user_id int) int64 {
	return st.amounts[user_id]
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		amounts: make(map[int]int64),
	}
}
