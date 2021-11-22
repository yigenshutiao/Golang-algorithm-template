package _076_minimum_window_substring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{
				"ADOBECODEBANC",
				"ABC",
			},
			"ACB",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindows(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func minWindows(s string, t string) string {
	source, target := map[byte]int{}, map[byte]int{}
	for k := range t {
		source[t[k]]++
	}

	lens := 99999
	start, end := -1, -1

	for l, r := 0, 0; r < len(s); r++ {
		target[s[r]] += 1
		for l <= r && checks(source, target) {
			if r-l+1 < lens {
				lens = r - l + 1
				start, end = l, r+1
			}
			target[s[l]]--
			l++
		}
	}

	if start == -1 {
		return ""
	}
	return s[start:end]
}

func checks(source, target map[byte]int) bool {
	for k, v := range source {
		if target[k] < v {
			return false
		}
	}
	return true
}
