package commands

import (
	"fmt"
	"goredis/store"
	utils "goredis/utility"
)

func SaveOnDisk() string {
	mapCopy := utils.CopyMap(store.HashMap)
	err := utils.WriteToRDB("backup.rdb", mapCopy)
	if err != nil {
		return "Error saving backup"
	}
	fmt.Println("Backup completed")
	return "OK"
}
