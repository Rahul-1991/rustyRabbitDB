package utility

import (
	"compress/gzip"
	"encoding/gob"
	"goredis/store"
	"os"
	"sync"
)

func CopyMap(originalMap sync.Map) map[string]store.CacheItem {
	copiedMap := make(map[string]store.CacheItem)
	originalMap.Range(func(key, value interface{}) bool {
		copiedMap[key.(string)] = value.(store.CacheItem)
		return true
	})
	return copiedMap
}

func WriteToRDB(filename string, data map[string]store.CacheItem) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a gzip writer
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	encoder := gob.NewEncoder(gzipWriter)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func ReadFromRDB(filename string) (sync.Map, error) {
	file, err := os.Open(filename)
	if err != nil {
		return sync.Map{}, err
	}
	defer file.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return sync.Map{}, err
	}
	defer gzipReader.Close()

	decoder := gob.NewDecoder(gzipReader)
	var copiedMap = make(map[string]store.CacheItem)
	err = decoder.Decode(&copiedMap)
	if err != nil {
		return sync.Map{}, err
	}

	result := sync.Map{}
	for key, value := range copiedMap {
		result.Store(key, value)
	}

	return result, nil
}
