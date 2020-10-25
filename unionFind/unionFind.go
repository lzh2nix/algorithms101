package unionFind

import "fmt"

type DynamicConnect interface {
	// connect p, and q
	Union(p, q int)
	// check p and q's connectivity
	Connected(p, q int) bool
}

// p, q is connect if p and q have same id
/*       0 1 2 3 4 5 6 7 8 9
   id[]  0 1 1 8 8 0 0 1 8 8

0,5,6 is connected
1,2,7 is connected
3,4,8,9 is connected
*/
type QuickFind struct {
	ID []int
}

// create a quick find  instance
func NewQuickFind(N int) DynamicConnect {
	ids := make([]int, N)
	for i := 0; i < N; i++ {
		ids[i] = i
	}
	return &QuickFind{ids}
}

// implement DynamicConnet Union method
func (qf *QuickFind) Union(p, q int) {

	qId := qf.ID[q]
	for i := 0; i < len(qf.ID); i++ {
		if qf.ID[i] == qId {
			qf.ID[i] = qf.ID[p]
		}
	}
	fmt.Println(qf.ID)
}

// implement DynamicConnet Connected method
func (qf *QuickFind) Connected(p, q int) bool {
	return qf.ID[p] == qf.ID[q]
}

// p, q is connect if p and q have root id
/*       0 1 2 3 4 5 6 7 8 9
   id[]  5 2 7 4 9 6 6 7 3 9

                6   7    9
               /   /    /
              5   2    4
             /   /    /
            0   7    3
                    /
                   8
0,5,6 is connected
1,2,7 is connected
3,4,8,9 is connected
*/
type QuickUnion struct {
	ID []int
}

// create a quick find  instance
func NewQuickUnion(N int) DynamicConnect {
	ids := make([]int, N)
	for i := 0; i < N; i++ {
		ids[i] = i
	}
	return &QuickUnion{ids}
}

// implement DynamicConnet Union method
func (qu *QuickUnion) Union(p, q int) {
	qu.ID[p] = q
}
func (qu *QuickUnion) root(p int) int {
	for {
		if p == qu.ID[p] {
			break
		}
		p = qu.ID[p]
	}
	return p
}

// implement DynamicConnet Connected method
func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.root(qu.ID[p]) == qu.root(qu.ID[q])
}

// same as quickunion but connect the root when union action
type QuickUnionLazy struct {
	ID []int
}

// create a quick find  instance
func NewQuickUnionLazy(N int) DynamicConnect {
	ids := make([]int, N)
	for i := 0; i < N; i++ {
		ids[i] = i
	}
	return &QuickUnionLazy{ids}
}

// implement DynamicConnet Union method
func (qu *QuickUnionLazy) Union(p, q int) {
	rootP := qu.root(p)
	rootQ := qu.root(q)
	qu.ID[rootP] = rootQ
}
func (qu *QuickUnionLazy) root(p int) int {
	for {
		if p == qu.ID[p] {
			break
		}
		p = qu.ID[p]
	}
	return p
}

// implement DynamicConnet Connected method
func (qu *QuickUnionLazy) Connected(p, q int) bool {
	return qu.root(qu.ID[p]) == qu.root(qu.ID[q])
}

// same as quickunion but balance by linking root of smaller tree to root of large tree
type QuickUnionWeight struct {
	ID []int
	SZ []int
}

// create a quick find  instance
func NewQuickUnionWeight(N int) DynamicConnect {
	ids := make([]int, N)
	sz := make([]int, N)
	for i := 0; i < N; i++ {
		ids[i] = i
		sz[i] = 1
	}
	return &QuickUnionWeight{ids,sz}
}

// implement DynamicConnet Union method
func (qu *QuickUnionWeight) Union(p, q int) {
	rootP := qu.root(p)
	rootQ := qu.root(q)
	if rootP == rootQ {
		return
	}
	if qu.SZ[rootP] > qu.SZ[rootQ] {
		qu.SZ[rootP] += qu.SZ[rootQ]
		qu.ID[rootQ] = qu.ID[rootP]
	} else {
		qu.SZ[rootQ] += qu.SZ[rootP]
		qu.ID[rootP] = qu.ID[rootQ]
	}
}
func (qu *QuickUnionWeight) root(p int) int {
	for {
		if p == qu.ID[p] {
			break
		}
		qu.ID[p] = qu.ID[qu.ID[p]]
		p = qu.ID[p]
	}
	return p
}

// implement DynamicConnet Connected method
func (qu *QuickUnionWeight) Connected(p, q int) bool {
	return qu.root(qu.ID[p]) == qu.root(qu.ID[q])
}
