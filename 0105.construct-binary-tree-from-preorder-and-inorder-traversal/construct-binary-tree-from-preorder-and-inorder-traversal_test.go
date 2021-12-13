package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"fmt"
	"testing"
)

func Test_buildTree(t *testing.T) {
	// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	res := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(res)
}
