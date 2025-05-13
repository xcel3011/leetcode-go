package leetcode_go

import (
	"slices"
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

func successfulPairs250417(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	ans := make([]int, len(spells))
	for i, spell := range spells {
		// 找到spell*potion >= success的最小potion
		index, _ := slices.BinarySearch(potions, int(success-1)/spell+1)
		//if !ok {
		//	ans[i] = 0
		//} else {
		ans[i] = len(potions) - index
		//}
	}
	return ans
}
