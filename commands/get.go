package commands

import (
	"fmt"
	store "goredis/store"
	"time"
)

func Get(commandArgList []string) string {
	if len(commandArgList) < 2 {
		return "(nil)"
	}
	cacheData, exists := store.HashMap.Load(commandArgList[1])
	if exists {
		if cacheData.(store.CacheItem).TTL < time.Now().Unix() && cacheData.(store.CacheItem).TTL != -1 {
			fmt.Println("Key expired so deleting")
			store.HashMap.Delete(commandArgList[1])
			return "(nil)"
		}
		return cacheData.(store.CacheItem).Value
	} else {
		return "(nil)"
	}
}
