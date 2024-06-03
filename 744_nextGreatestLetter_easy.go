package leetcode_go

func nextGreatestLetter(letters []byte, target byte) byte {
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}

func nextGreatestLetter1(letters []byte, target byte) byte {
	left, right := 0, len(letters)
	for left < right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if len(letters) == left || letters[left] <= target {
		return letters[0]
	}
	return letters[left]
}
