package 数组

// leetcode t27 移除元素
func removeElement(nums []int, val int) int {
	left := 0                //左指针
	for _, v := range nums { // v相当于右指针
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}
