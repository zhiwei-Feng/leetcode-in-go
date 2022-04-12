package bytedance0007

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_solve(t *testing.T) {
	ins := []int{1, 8, 3, 6, 5, 4, 7, 2}
	dummy := &ListNode{}
	last := dummy
	for _, v := range ins {
		last.Next = &ListNode{Val: v}
		last = last.Next
	}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{dummy.Next}, dummy.Next},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}else {
				p := tt.want
				for p!=nil {
					fmt.Print(p.Val)
					p = p.Next
				}
				fmt.Println()
			}
		})
	}
}
