package array

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
