package cache

import (
	"fmt"
	"time"
)

func init() {
	GetCacheMap()
}

// GetCacheMap is a singleton methond and return the global instance of map cache.
func GetCacheMap() *Map {
	once.Do(func() {
		cMap = Map{data: make(map[string]interface{})}
	})
	return &cMap
}

// StartClean starts a goroutine to clean the data of cache in some interval.
func (instance *Map) StartClean(interval time.Duration, maxDuration float64) {
	go instance.clean(interval, maxDuration)
}

func (instance *Map) clean(interval time.Duration, maxDuration float64) {
	for {
		func() {
			time.Sleep(interval)
			instance.Lock()
			defer instance.Unlock()
			fmt.Println("Start clean at " + time.Now().String())
			fmt.Println("Before clean, data cache size is " + fmt.Sprintf("%d", len(instance.data)))
			for k, v := range instance.data {
				t, ok := v.(time.Time)
				if !ok || time.Now().Sub(t).Seconds() > maxDuration {
					delete(instance.data, k)
				}
			}
			fmt.Println("After clean, data cache size is " + fmt.Sprintf("%d", len(instance.data)))
			fmt.Println("End clean at " + time.Now().String())
		}()
	}
}

// Add is to insert data into map cache.
func (instance *Map) Add(k string, v interface{}) {
	instance.Lock()
	defer instance.Unlock()
	instance.data[k] = v
}

// Contains returns if the map cache contains the key.
func (instance *Map) Contains(k string) bool {
	instance.RLock()
	defer instance.RUnlock()
	if _, ok := instance.data[k]; ok {
		return true
	}

	return false
}
