package corecache

import (
	"bytes"
	"context"
	"encoding/gob"
	"time"

	"github.com/redis/go-redis/v9"
)

const defaultTimeout = 60 * time.Second

type Cache struct {
	client *redis.Client
}

func New(addr string) (*Cache, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Cache{
		client: client,
	}, nil
}

func (c *Cache) Set(ctx context.Context, key string, data interface{}, exp time.Duration) error {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)

	err := enc.Encode(data)
	if err != nil {
		return err
	}

	err = c.client.Set(ctx, key, buf.Bytes(), exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) Get(ctx context.Context, key string, data interface{}) error {
	res := c.client.Get(ctx, key)

	err := res.Err()
	if err != nil {
		return err
	}

	val, err := res.Bytes()
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(bytes.NewReader(val))

	err = dec.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) Close() {
	_ = c.client.Close()
}
