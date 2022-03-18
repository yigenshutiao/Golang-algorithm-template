package _082_remove_duplicates_from_sorted_list_ii

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}

	input := util.Ints2List([]int{1, 2, 3, 3, 4, 4, 5})
	got := util.Ints2List([]int{1, 2, 5})
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"test 1",
			args{
				input,
			},
			got,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicate(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
