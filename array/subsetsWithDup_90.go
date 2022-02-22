package array

import (
	"sort"
	"strconv"
	"strings"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	m := make(map[string]struct{})
	toStr := func(nums []int) string {
		b := new(strings.Builder)
		for _, num := range nums {
			b.WriteString(strconv.Itoa(num))
		}
		return b.String()
	}

	res = append(res, []int{})

	for i := 0; i < len(nums); i++ {
		all := len(res)
		for j := 0; j < all; j++ {
			temp := append([]int{}, append(res[j], nums[i])...)
			tempStr := toStr(temp)
			if _, ok := m[tempStr]; ok {
				continue
			}

			m[tempStr] = struct{}{}
			res = append(res, temp)
		}
	}

	return res
}
