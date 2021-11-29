package n1229

import (
	"reflect"
	"testing"
)

func Test_minAvailableDuration(t *testing.T) {
	type args struct {
		slots1   [][]int
		slots2   [][]int
		duration int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"case1", args{[][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 8}, []int{60, 68}},
		{"case2", args{[][]int{{10, 50}, {60, 120}, {140, 210}}, [][]int{{0, 15}, {60, 70}}, 12}, []int{}},
		{"case3", args{[][]int{{10, 60}}, [][]int{{12, 17}, {21, 50}}, 8}, []int{21, 29}},
		{"case4", args{[][]int{{216397070, 363167701}, {98730764, 158208909}, {441003187, 466254040}, {558239978, 678368334}, {683942980, 717766451}}, [][]int{{50490609, 222653186}, {512711631, 670791418}, {730229023, 802410205}, {812553104, 891266775}, {230032010, 399152578}}, 456085}, []int{98730764, 99186849}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAvailableDuration(tt.args.slots1, tt.args.slots2, tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minAvailableDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
