package n1396

import (
	"fmt"
	"strconv"
	"strings"
)

// 要想计算AB两站之间的平均时间，必须记录每个A->B
// 使用<A_B>作为key，存储每个A->B记录的时间
// 要知道AB之间的时间，可以以到站为节点计算，需要记录入站的时间
// 以<id>为key，<stationName_t>为value

type UndergroundSystem struct {
	TimeRecords   map[string][]int
	PersonRecords map[int]string
}

func Constructor() UndergroundSystem {
	system := UndergroundSystem{
		TimeRecords:   make(map[string][]int),
		PersonRecords: make(map[int]string),
	}
	return system
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	val := fmt.Sprintf("%v_%v", stationName, t)
	this.PersonRecords[id] = val
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	inRecord := this.PersonRecords[id]
	inVal := strings.Split(inRecord, "_")
	inStationName := inVal[0]
	inT, _ := strconv.Atoi(inVal[1])
	key := fmt.Sprintf("%v_%v", inStationName, stationName)
	if _, ok := this.TimeRecords[key]; !ok {
		this.TimeRecords[key] = make([]int, 0)
	}
	this.TimeRecords[key] = append(this.TimeRecords[key], t-inT)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := fmt.Sprintf("%v_%v", startStation, endStation)
	records := this.TimeRecords[key]
	var ans float64

	for _, v := range records {
		ans += float64(v)
	}
	ans = ans / float64(len(records))
	return ans
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
