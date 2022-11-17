package string

// 方法一：分桶
func numMatchingSubseq(s string, words []string) int {
	m := [26][]string{}
	for _, word := range words {
		m[word[0]-'a'] = append(m[word[0]-'a'], word)
	}
	var ans int
	for _, c := range s {
		p := m[c-'a']
		m[c-'a'] = nil
		for _, t := range p {
			if len(t) == 1 {
				ans++
			} else {
				m[t[1]-'a'] = append(m[t[1]-'a'], t[1:])
			}
		}
	}
	return ans
}
