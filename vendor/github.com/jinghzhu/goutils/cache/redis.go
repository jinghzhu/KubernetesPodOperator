package cache

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-redis/redis"

	logger "github.com/jinghzhu/goutils/logger"
)

var (
	redisClient      *redis.Client
	redisConfig      *redis.Options
	redisOnce        sync.Once
	initaialized     uint64
	reset            uint64
	_reset           uint64
	_Reset           uint64
	mutexGetClient   sync.Mutex
	mutexResetClient sync.Mutex
)

// GetRedisClient returns a global instance of redis.Client in singleton model.
func GetRedisClient() *redis.Client {
	if atomic.LoadUint64(&initaialized) == 1 {
		return redisClient
	}

	mutexGetClient.Lock()
	defer mutexGetClient.Unlock()

	if initaialized == 0 {
		redisClient = redis.NewClient(GetRedisConfig())
		atomic.StoreUint64(&initaialized, 1)
	}

	return redisClient
}

// ResetRedisClient tries to reset the connection to Redis.
func ResetRedisClient() {
	// No need to let many goroutines to drive this method. Only one is enough.
	if atomic.LoadUint64(&_Reset) == 1 {
		return
	}
	atomic.StoreUint64(&_Reset, 1)

	logger.Info("Start to try to reset Redis connection")
	// The lock and following _reset are to prevent many goroutines to wait to reset the connection. This
	// may help them release the resource in their earlist time.
	mutexResetClient.Lock()
	defer mutexResetClient.Unlock()
	if _reset == 0 {
		go resetRedisClient()
	}
}

// resetRedisClient palys the main role to reset the connection to Redis. It will retry it with an interval.
// If it fails, an alert mail will be sent.
func resetRedisClient() {
	atomic.StoreUint64(&_reset, 1)
	defer func() {
		atomic.StoreUint64(&_Reset, 0)
		atomic.StoreUint64(&_reset, 0)
	}()

	var i int
	for i = 0; i < RedisReConnMaxRetries; i++ {
		logger.Info(fmt.Sprintf("Reset Redis connection in round#%d\n", i))
		atomic.StoreUint64(&initaialized, 0)
		GetRedisClient()
		if IsConnected() {
			logger.Info("Successfully reset Redis connection")
			break
		}
		time.Sleep(RedisReConnInterval)
	}
	if i >= RedisReConnMaxRetries {
		logger.Error("Fail to reset Redis connection")
	}
}

// IsConnected calls redis.Client.Ping() to verify the connection to Redis.
func IsConnected() bool {
	c := GetRedisClient()
	_, err := c.Ping().Result()
	if err != nil {
		logger.Info("Redis is disconnected")
		return false
	}
	return true
}

// GetRedisConfig returns the global configurationo of Redis in singleton model.
func GetRedisConfig() *redis.Options {
	redisOnce.Do(func() {
		redisConfig = &redis.Options{
			Addr:       RedisAddr,
			Password:   RedisPw,
			DB:         RedisDB,
			MaxRetries: RedisMaxRetries,
		}
		logger.Info(
			fmt.Sprintf("Init Redis configuration. address = %s, db = %d", RedisAddr, RedisDB),
		)
	})

	return redisConfig
}

// Get is a wrapper for redis.Client.Get. It supports asynchronously try to reset the connection to Redis
// if it fails to get data. If the key doesn't exist in Redis, it returns an empty byte array together with
// a redis.Nil error
func Get(key string) ([]byte, error) {
	client := GetRedisClient()
	val, err := client.Get(key).Bytes()
	if err == redis.Nil {
		logger.Info(
			fmt.Sprintf("Key(%s) doesn't exist", key),
		)
		return []byte{}, err
	}
	if err != nil {
		errMsg := "Fail to get value from Redis. Try to reset the connection."
		logger.Error(errMsg)
		go ResetRedisClient()
		return val, fmt.Errorf("%s: %v", errMsg, err)
	}

	return val, nil
}

// Set is a wrapper for redis.Client.Set. It supports asynchronously try to reset the connection to Redis
// if it fails to set data.
func Set(key string, val interface{}, keyDuration time.Duration) error {
	client := GetRedisClient()

	err := client.Set(key, val, keyDuration).Err()
	if err != nil {
		errMsg := "Error to put key values into Redis. Try to reset the connection."
		logger.Error(errMsg)
		go ResetRedisClient()
		return fmt.Errorf("%s: %v", errMsg, err)
	}

	return nil
}

func SetBytes(key string, val []byte, keyDuration time.Duration) error {
	client := GetRedisClient()

	err := client.Set(key, val, keyDuration).Err()
	if err != nil {
		errMsg := "Error to put key values into Redis. Try to reset the connection."
		logger.Error(errMsg)
		go ResetRedisClient()
		return fmt.Errorf("%s: %v", errMsg, err)
	}

	return nil
}

// Del is a wrapper for redis.Client.Del. It supports asynchronously try to reset the connection to Redis
// if it fails to delete data.
func Del(key string) error {
	client := GetRedisClient()
	err := client.Del(key).Err()
	if err != nil {
		errMsg := "Error to delete data from Redis. Try to reset connection."
		logger.Error(errMsg)
		go ResetRedisClient()
		return fmt.Errorf("%s: %v", errMsg, err)
	}

	logger.Info(fmt.Sprintf("Successfully delete data(%s) from Redis", key))
	return nil
}
