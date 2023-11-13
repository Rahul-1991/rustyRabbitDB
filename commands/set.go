package commands

import (
	"goredis/store"
	"strconv"
	"strings"
	"time"
)

func Set(commandArgList []string) string {
	var ttl int64 = -1
	if len(commandArgList) > 3 {
		if strings.ToLower(commandArgList[3]) == "ex" {
			ttl, _ = strconv.ParseInt(commandArgList[4], 10, 64)
			ttl += time.Now().Unix()
		} else if strings.ToLower(commandArgList[3]) == "px" {
			ttl, _ = strconv.ParseInt(commandArgList[4], 10, 64)
			ttl = (ttl / 1000) + time.Now().Unix()
		} else if strings.ToLower(commandArgList[3]) == "exat" {
			ttl, _ = strconv.ParseInt(commandArgList[4], 10, 64)
		} else if strings.ToLower(commandArgList[3]) == "pxat" {
			ttl, _ = strconv.ParseInt(commandArgList[4], 10, 64)
			ttl /= 1000
		}
	}
	valMap := store.CacheItem{
		Value: commandArgList[2],
		TTL:   ttl,
	}
	store.HashMap.Store(commandArgList[1], valMap)
	return "OK"
}
