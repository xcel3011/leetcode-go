package leetcode_go

func numOfSubarrays(arr []int, k int, threshold int) int {
	sum, count := 0, 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum/k >= threshold {
		count++
	}
	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		if sum/k >= threshold {
			count++
		}
	}
	return count
}
