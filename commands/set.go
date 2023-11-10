package commands

import (
	"goredis/store"
	"strconv"
	"time"
)

func Set(commandArgList []string) string {
	var ttl int64 = -1
	if len(commandArgList) > 3 && (commandArgList[3] == "EX" || commandArgList[3] == "ex") {
		ttl, _ = strconv.ParseInt(commandArgList[4], 10, 64)
		ttl += time.Now().Unix()
	}
	valMap := store.CacheItem{
		Value: commandArgList[2],
		TTL:   ttl,
	}
	store.HashMap.Store(commandArgList[1], valMap)
	return "OK"
}
