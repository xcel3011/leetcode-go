package string

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		buf := []byte(str)
		sort.Slice(buf, func(i, j int) bool { return buf[i] < buf[j] })
		s := string(buf)
		m[s] = append(m[s], str)
	}

	var ans [][]string
	for _, ss := range m {
		ans = append(ans, ss)
	}
	return ans
}
