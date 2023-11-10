package commands

import (
	store "goredis/store"
	"strconv"
)

func Exists(commandArgList []string) string {
	var count int = 0
	for i := 1; i < len(commandArgList); i++ {
		_, exists := store.HashMap.Load(commandArgList[i])
		if exists {
			count++
		}
	}
	return strconv.FormatInt(int64(count), 10)
}
