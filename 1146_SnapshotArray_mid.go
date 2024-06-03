package leetcode_go

import (
	"sort"
)

type SnapshotArray struct {
	vals map[int][]struct {
		snapId, val int
	}
	snapId int
}

func Constructor1146(length int) SnapshotArray {
	return SnapshotArray{
		vals: make(map[int][]struct{ snapId, val int }),
	}
}

func (s *SnapshotArray) Set(index int, val int) {
	s.vals[index] = append(s.vals[index], struct{ snapId, val int }{snapId: s.snapId, val: val})
}

func (s *SnapshotArray) Snap() int {
	s.snapId++
	return s.snapId - 1
}

func (s *SnapshotArray) Get(index int, snap_id int) int {
	i := sort.Search(len(s.vals[index]), func(i int) bool {
		return s.vals[index][i].snapId >= snap_id+1
	}) - 1
	if i >= 0 {
		return s.vals[index][i].val
	}
	return 0
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */
