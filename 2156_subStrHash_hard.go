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
		//hash*power：将当前窗口的哈希值乘以 power，这是为了给新加入窗口的字符腾出位置。
		//在哈希计算中，就像在进制转换里，每一位向左移动一位时要乘以基数一样。
		//val(s[i])：将新加入窗口的字符 s[i] 转换为对应的数值，并加到哈希值中。
		//pk*val(s[i+k])%modulo：计算窗口最右边字符 s[i+k] 对哈希值的贡献，并从当前哈希值中减去。
		//因为窗口要向左滑动，最右边的字符会离开窗口，所以要去掉它的影响。
		//+ modulo：在减去最右边字符的贡献时，可能会得到负数。加上 modulo 是为了确保结果为非负数。
		//% modulo：最后对结果取模，保证哈希值在 [0, modulo-1] 的范围内。
		hash = (hash*power + val(s[i]) - pk*val(s[i+k])%modulo + modulo) % modulo
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}

	return ans
}

//给定整数 p 和 m ，一个长度为 k 且下标从 0 开始的字符串 s 的哈希值按照如下函数计算：
//
//hash(s, p, m) = (val(s[0]) * p0 + val(s[1]) * p1 + ... + val(s[k-1]) * pk-1) mod m.
//其中 val(s[i]) 表示 s[i] 在字母表中的下标，从 val('a') = 1 到 val('z') = 26 。
//
//给你一个字符串 s 和整数 power，modulo，k 和 hashValue 。请你返回 s 中 第一个 长度为 k 的 子串 sub ，满足 hash(sub, power, modulo) == hashValue 。
//
//测试数据保证一定 存在 至少一个这样的子串。
//
//子串 定义为一个字符串中连续非空字符组成的序列。
//
//
//
//示例 1：
//
//输入：s = "leetcode", power = 7, modulo = 20, k = 2, hashValue = 0
//输出："ee"
//解释："ee" 的哈希值为 hash("ee", 7, 20) = (5 * 1 + 5 * 7) mod 20 = 40 mod 20 = 0 。
//"ee" 是长度为 2 的第一个哈希值为 0 的子串，所以我们返回 "ee" 。
//示例 2：
//
//输入：s = "fbxzaad", power = 31, modulo = 100, k = 3, hashValue = 32
//输出："fbx"
//解释："fbx" 的哈希值为 hash("fbx", 31, 100) = (6 * 1 + 2 * 31 + 24 * 312) mod 100 = 23132 mod 100 = 32 。
//"bxz" 的哈希值为 hash("bxz", 31, 100) = (2 * 1 + 24 * 31 + 26 * 312) mod 100 = 25732 mod 100 = 32 。
//"fbx" 是长度为 3 的第一个哈希值为 32 的子串，所以我们返回 "fbx" 。
//注意，"bxz" 的哈希值也为 32 ，但是它在字符串中比 "fbx" 更晚出现。
//
//
//提示：
//
//1 <= k <= s.length <= 2 * 104
//1 <= power, modulo <= 109
//0 <= hashValue < modulo
//s 只包含小写英文字母。
//测试数据保证一定 存在 满足条件的子串。

func subStrHash250324(s string, power int, modulo int, k int, hashValue int) string {
	val := func(c byte) int { return int(c - 'a' + 1) }
	// 用秦九韶算法计算 s[n-k:] 的哈希值
	hash, pk := 0, 1
	n := len(s)
	for i := n - 1; i >= n-k; i-- {
		hash = (hash*power + val(s[i])) % modulo
		pk = pk * power % modulo
	}

	var ans string
	if hash == hashValue {
		ans = s[n-k:]
	}

	// 向左滑动窗口
	for i := n - k - 1; i >= 0; i-- {
		hash = (hash*power + val(s[i]) - pk*val(s[i+k])%modulo + modulo) % modulo
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}

	return ans
}
