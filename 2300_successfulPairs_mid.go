package leetcode_go

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	ans := make([]int, len(spells))
	for i := 0; i < len(spells); i++ {
		ans[i] = len(potions) - sort.SearchInts(potions, int(success-1)/spells[i]+1)
	}
	return ans
}
