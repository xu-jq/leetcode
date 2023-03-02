package 数组

// leetcode t704 二分查找

// 直接暴力循环出结果
func search(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return -1
}

// 使用二分法
func search1(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for right >= left {
		mid := (right-left)/2 + left // 防止溢出 等同于(right+left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1 // mid已判断过 舍去mid
		}
		if nums[mid] < target {
			left = mid + 1 // mid已判断过 舍去mid
		}
	}
	return -1
}
