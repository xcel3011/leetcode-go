package array

type NumArray struct {
	prevSum []int
}

func NewNumArray(nums []int) NumArray {
	array := NumArray{
		prevSum: make([]int, len(nums)+1),
	}

	for i, num := range nums {
		array.prevSum[i+1] = array.prevSum[i] + num
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
