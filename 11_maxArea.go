package leetcode_go

import (
	"math"
)

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0.0

	for right > left {
		currArea := math.Min(float64(height[left]), float64(height[right])) * float64(right-left)
		area = math.Max(currArea, area)

		if height[right] < height[left] {
			right--
		} else {
			left++
		}
	}

	return int(area)
}

func maxArea250331(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func maxArea250415(height []int) int {
	ans := 0
	l, r := 0, len(height)-1
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
