package _098_validate_binary_search_tree

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tree := util.Ints2TreeNode([]int{2, 1, 3})
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 1",
			args{
				tree,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
