package leetcode_go

func vowelStrings(words []string, queries [][]int) []int {
	cnt := make([]int, len(words)+1)
	sum := 0
	for i := 0; i < len(words); i++ {
		if (words[i][0] == 'a' ||
			words[i][0] == 'e' ||
			words[i][0] == 'i' ||
			words[i][0] == 'o' ||
			words[i][0] == 'u') &&
			(words[i][len(words[i])-1] == 'a' ||
				words[i][len(words[i])-1] == 'e' ||
				words[i][len(words[i])-1] == 'i' ||
				words[i][len(words[i])-1] == 'o' ||
				words[i][len(words[i])-1] == 'u') {
			sum++
		}
		cnt[i+1] = sum
	}

	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = cnt[queries[i][1]+1] - cnt[queries[i][0]]
	}
	return res
}
