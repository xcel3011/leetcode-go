package leetcode_go

import (
	"math"
)

func minSpeedOnTime(dist []int, hour float64) int {
	h100 := int(math.Round(hour * 100))
	n := len(dist)
	// 除去最后一趟车，前面的每趟车最少花费一小时
	// 所以hour必须大于n-1
	if h100 <= (n-1)*100 {
		return -1
	}

	// 因为1<=dist[i]<=1e5,且小数点最多保留两位
	// 最后一趟车最大速度dist[n-1]=1e5/0.01=1e7
	// 所以速度取值范围为[1,1e7)
	l, r := 1, int(1e7)
	for l < r {
		mid := l + (r-l)/2
		sum := 0
		// 前n-1段耗时
		for _, d := range dist[:n-1] {
			sum += (d-1)/mid + 1
		}
		// 最后一段贡献的时间： dist[n-1] / mid
		// 避免float
		sum *= mid
		sum += dist[n-1]
		if sum*100 <= h100*mid {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
