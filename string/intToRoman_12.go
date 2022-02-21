package string

func intToRoman(num int) string {
	m := map[int]string{1: "I", 4: "IV", 5: "V", 9: "IX", 10: "X", 40: "XL", 50: "L", 90: "XC", 100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M"}
	s := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	res := ""

	for _, v := range s {
		for num >= v {
			num -= v
			res += m[v]
		}
		if num == 0 {
			break
		}
	}
	return res
}
