package store

import (
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *Store {
	return &Store{rdb: rdb}
}

func (s *Store) SaveData(ctx context.Context, data *BufferData) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error while marshalling", err)
		return err
	}

	if err = s.rdb.Set(ctx, "record:"+data.Value, jsonData, 0).Err(); err != nil {
		log.Println("Error while setting into redis, err: ", err)
		return err
	}

	return nil
}
