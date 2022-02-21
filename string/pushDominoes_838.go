package string

func pushDominoes(dominoes string) string {
	s := []byte(dominoes)
	i, n, left := 0, len(dominoes), byte('L')

	for i < n {
		j := i
		// 找到一段连续没有被推动的骨牌
		for j < n && s[j] == '.' {
			j++
		}

		right := byte('R')
		if j < n {
			right = s[j]
		}

		if left == right { // 方向相同，那么这些竖立骨牌也会倒向同一方向
			for i < j {
				s[i] = left
				i++
			}
		} else if left == 'R' && right == 'L' { // 方向相对，那么就从两侧向中间倒
			k := j - 1
			for i < k {
				s[i] = 'R'
				s[k] = 'L'
				i++
				k--
			}
		}

		left = right
		i = j + 1
	}
	return string(s)
}
