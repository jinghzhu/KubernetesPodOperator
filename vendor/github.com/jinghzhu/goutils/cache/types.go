package cache

import (
	"sync"
	"time"
)

// Map is a go-built in map data structure with a read-write lock.
type Map struct {
	sync.RWMutex
	data map[string]interface{}
}

var (
	once sync.Once
	cMap Map
)

const (
	RedisAddr             = "127.0.0.1:6379"
	RedisPw               = "" // no password set
	RedisDB               = 0  // use default DB
	RedisMaxRetries       = 3  // retry number for each get/put operation
	RedisReConnMaxRetries = 3  // retry number for creating redis client
	// RedisMinRetryBackoff is min backoff between each retry. -1 disables backoff.
	RedisMinRetryBackoff = 128 * time.Millisecond
	// RedisMaxRetryBackoff is max backoff between each retry. -1 disables backoff.
	RedisMaxRetryBackoff = 512 * time.Millisecond
	// RedisDialTimeout is the timeout for establishing new connections.
	RedisDialTimeout = 5 * time.Second
	// RedisReadTimeout Timeout for socket reads.
	RedisReadTimeout = 3 * time.Second
	// RedisWriteTimeout Timeout for socket writes.
	RedisWriteTimeout   = 3 * time.Second
	RedisKeyDuration    = 48 * time.Hour
	RedisReConnInterval = 15 * time.Second
)
