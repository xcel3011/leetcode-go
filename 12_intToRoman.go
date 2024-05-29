package leetcode_go

func intToRoman(num int) string {
	// 1 <= num <= 3999
	I := [10]string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	X := [10]string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	C := [10]string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	M := [4]string{"", "M", "MM", "MMM"}

	return M[num/1000] + C[num%1000/100] + X[num%100/10] + I[num%10]
}
