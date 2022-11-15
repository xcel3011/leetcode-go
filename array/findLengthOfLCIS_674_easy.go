package array

func findLengthOfLCIS(nums []int) (ans int) {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	count := 1
	for i := 1; i < len(nums); i++ {

		if nums[i] > nums[i-1] {
			count++
		} else {
			count = 1
		}
		if count > ans {
			ans = count
		}
	}

	return
}

func findLengthOfLCISOfficial(nums []int) (ans int) {
	max := func(nums ...int) int {
		max := nums[0]
		for _, num := range nums {
			if num > max {
				max = num
			}
		}
		return max
	}

	start := 0
	for i, num := range nums {
		if i > 0 && num <= nums[i-1] {
			start = i
		}
		ans = max(ans, i-start+1)
	}
	return
}
