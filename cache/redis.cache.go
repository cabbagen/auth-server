package cache

import "github.com/go-redis/redis/v7"

type RedisCache struct {
	Type         string
	Client       *redis.Client
}

var defaultRedisCache *RedisCache

func NewRedisCache() *RedisCache {
	redisCache := RedisCache { "redis", nil }
	defaultRedisCache = &redisCache

	return &redisCache
}

func GetRedisCacheInstance() *RedisCache {
	if defaultRedisCache.Type != "" {
		return defaultRedisCache
	}
	return NewRedisCache()
}

func (rc *RedisCache) Connect() {
	rc.Client = redis.NewClient(&redis.Options {
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

func (rc *RedisCache) Destroy() {
	rc.Client.Close()
}

func (rc *RedisCache) GetList(key string, start, stop int64) ([]string, error) {
	isExist, error := rc.Client.Exists(key).Result()

	if error != nil {
		return []string{}, error
	}
	if isExist == 0 {
		return []string{}, nil
	}
	return rc.Client.LRange(key, start, stop).Result()
}

func (rc *RedisCache) PushList(key string, values ...string) (int64, error) {
	return rc.Client.RPush(key, values).Result()
}

func (rc *RedisCache) UnShiftList(key string, values ...string) (int64, error) {
	return rc.Client.LPush(key, values).Result()
}

func (rc *RedisCache) GetSet(key string) ([]string, error) {
	return rc.Client.SMembers(key).Result()
}

func (rc *RedisCache) PushSet(key string, values ...string) (int64, error) {
	return rc.Client.SAdd(key, values).Result()
}

