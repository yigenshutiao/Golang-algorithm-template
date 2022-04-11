package _094_binary_tree_inorder_traversal

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}

	root := util.PreIn2Tree([]int{1, 2, 5}, []int{5, 1, 2})
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test1",
			args{root},
			[]int{2, 1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
