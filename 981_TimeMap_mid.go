package leetcode_go

import (
	"sort"
)

type TimeMap map[string][]struct {
	value     string
	timestamp int
}

func Constructor981() TimeMap {
	return make(TimeMap)
}

func (t TimeMap) Set(key string, value string, timestamp int) {
	t[key] = append(t[key], struct {
		value     string
		timestamp int
	}{value: value, timestamp: timestamp})
}

func (t TimeMap) Get(key string, timestamp int) string {
	i := sort.Search(len(t[key]), func(i int) bool {
		return t[key][i].timestamp > timestamp
	})
	if i > 0 {
		return t[key][i-1].value
	}
	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
