package 数组

import "sort"

// leetcode T977.有序数组的平方

// 先平方再排序
func sortedSquares(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		nums[i] = nums[i] * nums[i]
	}
	sort.Ints(nums)
	return nums
}

// 双指针后再反转
func sortedSquares2(nums []int) []int {
	left := 0
	right := len(nums) - 1
	res := make([]int, 0)
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res = append(res, nums[right]*nums[right])
			right--
		} else {
			res = append(res, nums[left]*nums[left])
			left++
		}
	}
	left = 0
	right = len(res) - 1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	return res
}

// 纯双指针
func sortedSquares3(nums []int) []int {
	left := 0
	right := len(nums) - 1
	n := len(nums) - 1
	res := make([]int, len(nums))
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[n] = nums[right] * nums[right]
			right--
		} else {
			res[n] = nums[left] * nums[left]
			left++
		}
		n--
	}
	return res
}
