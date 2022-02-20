package _301_remove_invalid_parentheses

import (
	"reflect"
	"testing"
)

func Test_removeInvalidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []string
	}{
		{
			"test 1",
			args{
				"(a)())()",
			},
			[]string{"(())()", "()()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := removeInvalidParentheses(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("removeInvalidParentheses() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
