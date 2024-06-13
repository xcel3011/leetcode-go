package leetcode_go

import (
	"fmt"
)

// A->1
// AA->27
func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber != 0 {
		columnNumber--
		ans = string('A'+byte(columnNumber%26)) + ans
		columnNumber /= 26
		fmt.Println("ans:", ans)
	}
	return ans
}
