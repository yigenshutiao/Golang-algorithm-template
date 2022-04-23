package test

import "strconv"

func maximalSquare(matrix [][]byte) int {
	var corner int

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			t := min(min(matrix[i-1][j], matrix[i][j-1]), matrix[i-1][j-1])
			if matrix[i][j] > '0' && t >= matrix[i][j] {
				matrix[i][j] = t + 1
			}
			a, _ := strconv.Atoi(string(matrix[i][j]))
			if a > corner {
				corner = a
			}
		}
	}

	return corner * corner
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}
