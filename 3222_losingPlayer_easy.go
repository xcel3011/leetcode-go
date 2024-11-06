package leetcode_go

import (
	"fmt"
)

func losingPlayer(x int, y int) string {
	cnt := y / 4
	fmt.Println("cnt1", cnt)
	if cnt < x {
		cnt = x
	}
	fmt.Println("cnt2", cnt)
	if cnt%2 == 0 {
		return "Bob"
	} else {
		return "Alice"
	}
}
