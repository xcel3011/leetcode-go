package leetcode_go

func plusOne(digits []int) []int {
	carry := 0
	n := 0
	l := len(digits) - 1
	var res []int
	for ; l >= 0; l-- {
		if l == len(digits)-1 {
			n = digits[l] + 1 + carry
		} else {
			n = digits[l] + carry
		}

		carry = n / 10
		n = n % 10
		res = append([]int{n}, res...)
	}

	if carry != 0 {
		res = append([]int{carry}, res...)
	}
	return res
}
