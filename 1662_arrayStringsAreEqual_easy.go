package leetcode_go

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func arrayStringsAreEqual1(word1 []string, word2 []string) bool {
	var w1, w2, i, j int
	for w1 < len(word1) && w2 < len(word2) {
		if word1[w1][i] != word2[w2][j] {
			return false
		}

		i++
		if i == len(word1[w1]) {
			w1++
			i = 0
		}

		j++
		if j == len(word2[w2]) {
			w2++
			j = 0
		}
	}
	return w1 == len(word1) && w2 == len(word2)
}
