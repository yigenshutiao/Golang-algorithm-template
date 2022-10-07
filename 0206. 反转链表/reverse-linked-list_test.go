package _206_reverse_linked_list

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"reflect"
	"testing"
)

func Test_reverseListRecursion(t *testing.T) {
	nodes := util.Ints2List([]int{1, 2, 3, 4, 5})
	wants := util.Ints2List([]int{5, 4, 3, 2, 1})
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test1",
			args{
				nodes,
			},
			wants,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListRecursion(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}
