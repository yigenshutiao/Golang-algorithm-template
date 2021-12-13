package _652_find_duplicate_subtrees

import (
	"reflect"
	"testing"
)

func Test_findDuplicateSubtrees(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicateSubtrees(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicateSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
