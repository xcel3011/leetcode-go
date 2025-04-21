package leetcode_go

// d0=a1-a0
// a1=a0+d0
// a2=a1+d1=a0+d0+d1
// a3=a2+d2=a0+d0+d1+d2
// a[i]=a0+d0+d1+...+di
// lower<=ai<=upper
// lower<=a0+si<=upper
// lower-si<=a0<=upper-si
func numberOfArrays(differences []int, lower int, upper int) int {
	var s, mins, maxs int
	for _, d := range differences {
		s += d
		mins = min(mins, s)
		maxs = max(maxs, s)
	}
	return max(0, upper-lower-maxs+mins+1)
}
