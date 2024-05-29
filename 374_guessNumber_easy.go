package leetcode_go

func guess(num int) int {
	return num
}

func guessNumber(n int) int {
	left, right := 0, n
	for left < right {
		mid := left + (right-left)/2
		if guess(mid) <= 0 {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
