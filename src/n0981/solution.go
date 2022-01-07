package n0981

import (
	"sort"
)

type StampVal struct {
	TimeStamp int
	Val       string
}

type TimeMap struct {
	M map[string][]StampVal
}

func Constructor() TimeMap {
	m := TimeMap{
		M: make(map[string][]StampVal),
	}
	return m
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.M[key] = append(this.M[key], StampVal{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list := this.M[key]
	index := sort.Search(len(list), func(i int) bool {
		return list[i].TimeStamp > timestamp
	})
	if index == 0 {
		return ""
	} 
	return list[index-1].Val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
