package _240_search_a_2d_matrix_ii

// searchMatrix 方法比较巧妙，从第一行最右边开始遍历，是二分属性
// 第一行左边的值都小于x ，第一行下面的值都大于x
func searchMatrix(matrix [][]int, target int) bool {

	i, j := len(matrix[0])-1, 0
	for i >= 0 && j < len(matrix) {
		if matrix[j][i] == target {
			return true
		} else if matrix[j][i] > target {
			i--
		} else if matrix[j][i] < target {
			j++
		}
	}

	return false
}
