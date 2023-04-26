package 数组

// 暴力两层for
func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	res := 999
	sum := 0
	for i := 0; i < length; i++ {
		sum = 0
		for j := i; j < length; j++ {
			sum += nums[j]
			if sum >= target {
				a := j - i + 1
				if res > a {
					res = a
				}
				break
			}
		}
	}
	if res == 999 {
		return 0
	}
	return res
}

// 滑动窗口
func minSubArrayLen1(target int, nums []int) int {
	l := len(nums)
	j := 0
	sum := 0
	res := 100001
	for i := 0; i < l; i++ {
		sum += nums[i]
		for sum >= target {
			subLength := i - j + 1
			if subLength < res {
				res = subLength
			}
			sum -= nums[j]
			j++
		}
	}
	if res == 100001 {
		return 0
	}
	return res
}
