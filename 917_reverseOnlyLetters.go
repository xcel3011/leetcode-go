package leetcode_go

func reverseOnlyLetters(s string) string {
	ss := []byte(s)
	slow, fast := 0, len(ss)-1
	for fast > slow {
		// fmt.Println("slow:", slow, "ss[slow]:", string(ss[slow]), "fast:", fast, "ss[fast]:", string(ss[fast]))
		if !((ss[slow] >= 'A' && ss[slow] <= 'Z') || (ss[slow] >= 'a' && ss[slow] <= 'z')) {
			slow++
			// fmt.Println("slow continue")
			continue
		}

		if !((ss[fast] >= 'A' && ss[fast] <= 'Z') || (ss[fast] >= 'a' && ss[fast] <= 'z')) {
			fast--
			// fmt.Println("fast continue")
			continue
		}

		ss[slow], ss[fast] = ss[fast], ss[slow]
		slow++
		fast--
	}
	return string(ss)
}
