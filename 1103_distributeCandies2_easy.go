package leetcode_go

import (
	"fmt"
)

func distributeCandies(candies int, num_people int) []int {
	ans := make([]int, num_people)
	n := 0
	for candies > 0 {
		for i := 0; i < num_people; i++ {
			count := n*num_people + i
			if count > candies {
				count = candies
			}
			candies -= count
			fmt.Println(count, candies)
			ans[i] += count
			if candies == 0 {
				break
			}
		}
		n++
	}
	return ans
}
