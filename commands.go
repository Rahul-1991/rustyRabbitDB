package main

//func Get(commandArgList []string) string {
//	if len(commandArgList) < 2 {
//		return "(nil)"
//	}
//	cacheData, exists := HashMap.Load(commandArgList[1])
//	if exists {
//		if cacheData.(store.CacheItem).TTL < time.Now().Unix() && cacheData.(store.CacheItem).TTL != -1 {
//			fmt.Println("Key expired so deleting")
//			HashMap.Delete(commandArgList[1])
//			return "(nil)"
//		}
//		return cacheData.(store.CacheItem).Value
//	} else {
//		return "(nil)"
//	}
//}

//func Set(commandArgList []string) string {
//	var ttl int64 = -1
//	if len(commandArgList) > 3 && (commandArgList[3] == "EX" || commandArgList[3] == "ex") {
//		ttl, _ = strconv.ParseInt(commandArgList[4], 10, 64)
//		ttl += time.Now().Unix()
//	}
//	valMap := store.CacheItem{
//		Value: commandArgList[2],
//		TTL:   ttl,
//	}
//	HashMap.Store(commandArgList[1], valMap)
//	return "OK"
//}

//func Exists(commandArgList []string) string {
//	var count int = 0
//	for i := 1; i < len(commandArgList); i++ {
//		_, exists := HashMap.Load(commandArgList[i])
//		if exists {
//			count++
//		}
//	}
//	return strconv.FormatInt(int64(count), 10)
//}

//func Delete(commandArgList []string) string {
//	var count int = 0
//	for i := 1; i < len(commandArgList); i++ {
//		_, exists := HashMap.Load(commandArgList[i])
//		if exists {
//			HashMap.Delete(commandArgList[1])
//			count++
//		}
//	}
//	return strconv.FormatInt(int64(count), 10)
//}

//func AtomicIncrement(commandArgList []string) string {
//	cachedData, exists := HashMap.Load(commandArgList[1])
//	valueAsString := "0"
//	if exists {
//		valueAsString = cachedData.(CacheItem).Value
//	}
//	valueAsInt, _ := strconv.ParseInt(valueAsString, 10, 64)
//	updatedValue := atomic.AddInt64(&valueAsInt, 1)
//	updatedValueAsString := strconv.FormatInt(updatedValue, 10)
//	valMap := CacheItem{
//		Value: updatedValueAsString,
//		TTL:   cachedData.(CacheItem).TTL,
//	}
//	HashMap.Store(commandArgList[1], valMap)
//	return updatedValueAsString
//}

//func AtomicDecrement(commandArgList []string) string {
//	cachedData, exists := HashMap.Load(commandArgList[1])
//	valueAsString := "0"
//	if exists {
//		valueAsString = cachedData.(CacheItem).Value
//	}
//	valueAsInt, _ := strconv.ParseInt(valueAsString, 10, 64)
//	updatedValue := atomic.AddInt64(&valueAsInt, -1)
//	updatedValueAsString := strconv.FormatInt(updatedValue, 10)
//	valMap := CacheItem{
//		Value: updatedValueAsString,
//		TTL:   cachedData.(CacheItem).TTL,
//	}
//	HashMap.Store(commandArgList[1], valMap)
//	return updatedValueAsString
//}
