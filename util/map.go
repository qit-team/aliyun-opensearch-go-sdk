package util

import "sort"

type Util struct {
}

func (u *Util) KSort(data map[string]interface{}) map[string]interface{} {
	var keys []string
	newMap := map[string]interface{}{}
	for k, _ := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys {
		newMap[v] = data[v]
	}
	return newMap
}


func (u *Util) Empty(data map[string]interface{}, key string) bool {
	if data == nil || len(data) == 0 {
		return true
	}
	if _, ok := data[key]; ok {
		return false
	}
	return true
}
