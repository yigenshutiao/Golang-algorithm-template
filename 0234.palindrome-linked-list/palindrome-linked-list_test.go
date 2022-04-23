package _234_palindrome_linked_list

import (
	"github.com/yigenshutiao/Golang-algorithm-template/util"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	head := util.Ints2List([]int{1, 2, 2, 1})
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 1",
			args{
				head,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
