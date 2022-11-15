package main

import "fmt"

func BubbleSort(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		var flag bool
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				flag = true
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
		if flag == false {
			break
		}
	}

	return nums
}

func main() {
	fmt.Println(BubbleSort([]int{9, 12, 1, 5, 2, 5, 67}))
}
