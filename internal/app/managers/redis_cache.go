package managers

import (
	"context"
	"encoding/json"
	"fmt"
	"test/internal/app/models"
	"time"

	"github.com/redis/go-redis/v9"
)

func goodsKey(offset, limit int) string {
	return fmt.Sprintf("goods:list:%d:%d", offset, limit)
}

func goodsPattern() string {
	return "goods:list:*"
}

func goodsMetaKey() string {
	return "goods:meta"
}

type RedisCache struct {
	rdb *redis.Client
	exp time.Duration
}

type optionFn func(*RedisCache)

func WithExpiration(seconds int) optionFn {
	return func(c *RedisCache) {
		c.exp = time.Duration(seconds) * time.Second
	}
}

func NewRedisCache(rdb *redis.Client, opts ...optionFn) *RedisCache {
	cache := &RedisCache{
		rdb: rdb,
		exp: 1 * time.Minute,
	}
	for _, apply := range opts {
		apply(cache)
	}
	return cache
}

func (c *RedisCache) CacheGoodsMetadata(ctx context.Context, meta models.PaginationMeta) error {
	return c.cache(ctx, goodsMetaKey(), meta)
}

func (c *RedisCache) GetGoodsMetadata(ctx context.Context) (*models.PaginationMeta, error) {
	value, err := get[models.PaginationMeta](c.rdb, ctx, goodsMetaKey())
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (c *RedisCache) ClearGoodsMetadata(ctx context.Context) error {
	return c.clear(ctx, goodsMetaKey())
}

func (c *RedisCache) CacheGoods(ctx context.Context, offset, limit int, goods []*models.Goods) error {
	return c.cache(ctx, goodsKey(offset, limit), goods)
}

func (c *RedisCache) GetGoods(ctx context.Context, offset, limit int) ([]*models.Goods, error) {
	value, err := get[[]*models.Goods](c.rdb, ctx, goodsKey(offset, limit))
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *RedisCache) ClearGoods(ctx context.Context) error {
	return c.clearByPattern(ctx, goodsPattern())
}

func (c *RedisCache) cache(ctx context.Context, key string, data any) error {
	if data == nil {
		return nil
	}
	metaJSON, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.rdb.Set(ctx, key, metaJSON, c.exp).Err()
}

func (c *RedisCache) clear(ctx context.Context, key string) error {
	_, err := c.rdb.Del(ctx, key).Result()
	return err
}

func (c *RedisCache) clearByPattern(ctx context.Context, pattern string) error {
	iter := c.rdb.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		if err := c.rdb.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	return iter.Err()
}

func get[T any](rdb *redis.Client, ctx context.Context, key string) (T, error) {
	var empty T

	value, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return empty, nil
	}
	if err != nil {
		return empty, err
	}

	var result T
	if err := json.Unmarshal([]byte(value), &result); err != nil {
		return empty, err
	}

	return result, nil
}
