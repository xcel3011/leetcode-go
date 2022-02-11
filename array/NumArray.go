package array

type NumArray struct {
	prevSum []int
}

func Constructor(nums []int) NumArray {
	array := NumArray{
		prevSum: make([]int, len(nums)+1),
	}

	for i := 1; i <= len(nums); i++ {
		array.prevSum[i] = array.prevSum[i-1] + nums[i-1]
	}

	return array
}

func (array *NumArray) SumRange(left int, right int) int {
	return array.prevSum[right+1] - array.prevSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
