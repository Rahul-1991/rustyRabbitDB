package commands

import (
	"encoding/json"
	"goredis/store"
)

func RightPush(commandArgList []string) string {
	cacheData, exists := store.HashMap.Load(commandArgList[1])
	var serializedCachedData []string
	if exists {
		err := json.Unmarshal([]byte(cacheData.(store.CacheItem).Value), &serializedCachedData)
		if err != nil {
			return "Invalid data"
		}
	}
	var modifiedArray []string = make([]string, 0, len(serializedCachedData)+len(commandArgList)-2)
	for i := 2; i < len(commandArgList); i++ {
		modifiedArray = append(modifiedArray, commandArgList[i])
	}
	modifiedArray = append(serializedCachedData, modifiedArray...)

	// Serialize the array to JSON and store it as a string
	serialized, _ := json.Marshal(modifiedArray)
	valMap := store.CacheItem{
		Value: string(serialized),
		TTL:   -1,
	}
	store.HashMap.Store(commandArgList[1], valMap)
	return "OK"
}
