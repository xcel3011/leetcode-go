package array

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	for i := 0; i < len(arr); i++ {
		if m[arr[i]] == 1 {
			k--
			if k == 0 {
				return arr[i]
			}
		}
	}
	return ""
}
