package priorityQueue

import (
	"sort"
)

type MaxPQ interface {
	// insert a element to priority queue
	Insert(k int)
	// delete the largest one from queue
	DelMax() int
	// check empty of pq
	IsEmpty() bool
}

type UnOrderMaxPQ struct {
	items []int
	n     int
	cap   int
}

func NewUnOrderMaxPQ(cap int) MaxPQ {
	pq := make([]int, cap)
	return &UnOrderMaxPQ{pq, 0, cap}
}

// implement of MaxPQ Insert method
func (pq *UnOrderMaxPQ) Insert(k int) {
	pq.items = append(pq.items[:pq.n], k)
	pq.n++
}

// implement of MaxPQ DelMax method
func (pq *UnOrderMaxPQ) DelMax() int {
	if pq.IsEmpty() {
		return 0
	}
	maxI := 0
	for i := 0; i < pq.n; i++ {
		if pq.items[i] > pq.items[maxI] {
			maxI = i
		}
	}
	max := pq.items[maxI]
	pq.items[maxI] = pq.items[pq.n-1]
	pq.n--
	return max
}

// implement of MaxPQ IsEmpty method
func (pq *UnOrderMaxPQ) IsEmpty() bool {
	return pq.n == 0
}

type OrderMaxPQ struct {
	items []int
	cap   int
	n     int
}

func NewOrderMaxPQ(cap int) MaxPQ {
	pq := make([]int, cap)
	return &OrderMaxPQ{pq, cap, 0}
}

// implement of MaxPQ Insert method
func (pq *OrderMaxPQ) Insert(k int) {
	i := sort.SearchInts(pq.items[:pq.n], k)
	pq.items = append(pq.items[:pq.n], 0)
	copy(pq.items[i+1:], pq.items[i:])
	pq.items[i] = k
	pq.n++
}

// implement of MaxPQ DelMax method
func (pq *OrderMaxPQ) DelMax() int {
	if pq.IsEmpty() {
		return 0
	}
	max := pq.items[len(pq.items)-1]
	pq.items = pq.items[:len(pq.items)-1]
	pq.n--
	return max
}

// implement of MaxPQ IsEmpty method
func (pq *OrderMaxPQ) IsEmpty() bool {
	return len(pq.items) == 0
}

// binary heap
type BinaryHeapMaxPQ struct {
	items []int
	cap   int
	n     int
}

// make a new binary heap based priority queue
func NewBinaryHeapMaxPQ(cap int) MaxPQ {
	pq := make([]int, cap+1)
	return &BinaryHeapMaxPQ{pq, cap, 1}
}

// implement of MaxPQ Insert method
func (pq *BinaryHeapMaxPQ) Insert(k int) {
	pq.items = append(pq.items[:pq.n], k)
	pq.swim(pq.n)
	pq.n++
}

// implement of MaxPQ DelMax method
func (pq *BinaryHeapMaxPQ) DelMax() int {
	if pq.n == 1 {
		return 0
	}
	max := pq.items[1]
	pq.items[pq.n-1], pq.items[1] = pq.items[1], pq.items[pq.n-1]
	pq.n--
	pq.items = pq.items[:pq.n]
	pq.slink(1)
	return max
}

// implement of MaxPQ IsEmpty method
func (pq *BinaryHeapMaxPQ) IsEmpty() bool {
	return pq.n == 1
}

func (pq *BinaryHeapMaxPQ) swim(k int) {
	for k > 1 && pq.items[k/2] < pq.items[k] {
		pq.items[k/2], pq.items[k] = pq.items[k], pq.items[k/2]
		k = k / 2
	}
}
func (pq *BinaryHeapMaxPQ) slink(k int) {
	for 2*k < pq.n {
		j := 2 * k
		if j+1 < pq.n && pq.items[j+1] > pq.items[j] {
			j++
		}
		if pq.items[k] > pq.items[j] {
			break
		}
		pq.items[k], pq.items[j] = pq.items[j], pq.items[k]
		k = j
	}
}
