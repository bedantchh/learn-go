package main

import "fmt"

func main() {
	//uninitialised slice is nil
	// var nums []int
	// fmt.Println(nums)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// // nums[0] = 1
	// // fmt.Println(cap(nums))
	// nums = append(nums, 1, 2, 3, 4, 5)
	// fmt.Println(nums)
	// fmt.Println(cap(nums))

	// nums := []int{1, 2, 3, 4}
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))

	// copy function
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)

	//slice operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:1])

	//slice
	// var nums = []int{1, 2}
	// var nums2 = []int{1, 2}

	// fmt.Println(slices.Equal(nums, nums2))

	//2d
	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)

}
