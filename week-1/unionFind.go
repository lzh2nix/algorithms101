package week1

import "fmt"

type DynamicConnect interface {
	// connect p, and q
	Union(p, q int)
	// check p and q's connectivity
	Connected(p, q int) bool
}

// p, q is connect if p and q have same id
/*       0 1 2 3 4 5 6 7 8 9
   id[] 0 1 1 8 8 0 0 1 8 8

0,5,6 is connected
1,2,7 is connected
3,4,8,9 is connected
*/
type QuickFind struct {
	ID []int
}

func NewQuickFind(N int) DynamicConnect {
	ids := make([]int, N)
	for i := 0; i < N; i++ {
		ids[i] = i
	}
	return &QuickFind{ids}
}
func (qf *QuickFind) Union(p, q int) {

	qId := qf.ID[q]
	for i := 0; i < len(qf.ID); i++ {
		if qf.ID[i] == qId {
			qf.ID[i] = qf.ID[p]
		}
	}
	fmt.Println(qf.ID)
}

func (qf *QuickFind) Connected(p, q int) bool {
	return qf.ID[p] == qf.ID[q]
}
