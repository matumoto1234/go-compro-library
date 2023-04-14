package data_structure

type CumulativeSum struct {
	Sum []int
}

func NewCumulativeSum(a []int) *CumulativeSum {
	sum := make([]int, len(a)+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	return &CumulativeSum{Sum: sum}
}

// [l, r)
func (cs *CumulativeSum) Query(l, r int) int {
	return cs.Sum[r] - cs.Sum[l]
}
