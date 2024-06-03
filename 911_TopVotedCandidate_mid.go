package leetcode_go

import (
	"sort"
)

type TopVotedCandidate []struct {
	val, timestamp int
}

func Constructor911(persons []int, times []int) (votes TopVotedCandidate) {
	top := -1
	voteCount := map[int]int{-1: -1}
	for i := 0; i < len(persons); i++ {
		voteCount[persons[i]]++

		if voteCount[persons[i]] > voteCount[top] {
			top = persons[i]
		}
		votes = append(votes, struct {
			val, timestamp int
		}{val: top, timestamp: times[i]})
	}
	return votes
}

func (votes TopVotedCandidate) Q(t int) int {
	j := sort.Search(len(votes), func(i int) bool {
		return votes[i].timestamp >= t+1
	}) - 1
	if j >= 0 {
		return votes[j].val
	}
	return 0
}

/**
 * Your TopVotedCandidate object will be instantiated and called as such:
 * obj := Constructor(persons, times);
 * param_1 := obj.Q(t);
 */
