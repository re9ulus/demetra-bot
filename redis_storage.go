package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
)

type redisStorage struct {
	client *redis.Client
}

func (st *redisStorage) Add(user_id int, val int64) error {
	if val <= 0 {
		return fmt.Errorf("val must be positive")
	}
	current_val := st.Get(user_id)
	str_user_id := strconv.Itoa(user_id)
	err := st.client.Set(str_user_id, current_val+val, 0).Err()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return nil
}

func (st *redisStorage) Spent(user_id int, val int64) error {
	if val <= 0 {
		return fmt.Errorf("val must be positive")
	}
	current_val := st.Get(user_id)
	str_user_id := strconv.Itoa(user_id)
	err := st.client.Set(str_user_id, current_val-val, 0).Err()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return nil
}

// Todo: Add ability to return error from Get
func (st *redisStorage) Get(user_id int) int64 {
	str_user_id := strconv.Itoa(user_id)
	str_val, err := st.client.Get(str_user_id).Result()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		panic(err) // TODO: Return error
	}

	val, err := strconv.ParseInt(str_val, 10, 64)
	if err != nil {
		panic(err) // TODO: Return error
	}

	return val
}

func NewRedisStorage(client *redis.Client) *redisStorage {
	return &redisStorage{client}
}
