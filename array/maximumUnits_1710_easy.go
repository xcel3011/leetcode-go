package array

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool { return boxTypes[i][1] > boxTypes[j][1] })
	fmt.Println(boxTypes)

	n := truckSize
	res := 0
	for i := 0; i < len(boxTypes) && n > 0; i++ {
		if n > boxTypes[i][0] {
			res += boxTypes[i][0] * boxTypes[i][1]
			n -= boxTypes[i][0]
		} else {
			res += n * boxTypes[i][1]
			n -= n
		}
	}
	return res
}
