package main

import "fmt"

func InsertSort(nums []int) []int {
	if len(nums) < 1 {
		return nil
	}

	for i := 1; i < len(nums); i++ {
		val := nums[i]

		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = val
	}

	return nums
}

func insert(nums []int) []int {

	if len(nums) < 1 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		//
		val := nums[i]

		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}

		nums[j+1] = val
	}

	return nums
}

func main1234() {
	fmt.Println(InsertSort([]int{4, 5, 6, 1, 3, 2}))
}
