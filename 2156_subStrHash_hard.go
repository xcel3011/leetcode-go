package leetcode_go

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	// 定义字符转下标函数
	val := func(c byte) int {
		return int(c - 'a' + 1)
	}

	n := len(s)
	hash, pk := 0, 1
	for i := n - 1; i >= n-k; i-- {
		hash = (hash*power + val(s[i])) % modulo
		pk = pk * power % modulo
	}

	var ans string
	if hash == hashValue {
		ans = s[n-k:]
	}

	for i := n - k - 1; i >= 0; i-- {
		hash = (hash*power + val(s[i]) - pk*val(s[i+k])%modulo + modulo) % modulo
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}

	return ans
}
