package commands

import (
	store "goredis/store"
	"strconv"
	"sync/atomic"
)

func AtomicIncrement(commandArgList []string) string {
	cachedData, exists := store.HashMap.Load(commandArgList[1])
	valueAsString := "0"
	if exists {
		valueAsString = cachedData.(store.CacheItem).Value
	}
	valueAsInt, _ := strconv.ParseInt(valueAsString, 10, 64)
	updatedValue := atomic.AddInt64(&valueAsInt, 1)
	updatedValueAsString := strconv.FormatInt(updatedValue, 10)
	valMap := store.CacheItem{
		Value: updatedValueAsString,
		TTL:   cachedData.(store.CacheItem).TTL,
	}
	store.HashMap.Store(commandArgList[1], valMap)
	return updatedValueAsString
}
