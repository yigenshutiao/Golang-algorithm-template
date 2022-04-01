package _020_valid_parentheses

import "testing"

func Test_isValids(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 1",
			args{
				"()",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValids(tt.args.s); got != tt.want {
				t.Errorf("isValids() = %v, want %v", got, tt.want)
			}
		})
	}
}
