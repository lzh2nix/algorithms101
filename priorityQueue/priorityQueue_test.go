package priorityQueue_test

import (
	"math/rand"
	"testing"

	"github.com/lzh2nix/algorithms101/priorityQueue"
	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, pq priorityQueue.MaxPQ, tc string) {
	pq.Insert(1000)
	pq.Insert(100)
	pq.Insert(1100)
	pq.Insert(1120)
	pq.Insert(1001)
	pq.Insert(100000)
	pq.Insert(1009)
	assert.False(t, pq.IsEmpty(), tc)
	assert.Equal(t, 100000, pq.DelMax(), tc)
	assert.Equal(t, 1120, pq.DelMax(), tc)
	assert.Equal(t, 1100, pq.DelMax(), tc)
	assert.Equal(t, 1009, pq.DelMax(), tc)
	assert.Equal(t, 1001, pq.DelMax(), tc)
	assert.Equal(t, 1000, pq.DelMax(), tc)
	assert.Equal(t, 100, pq.DelMax(), tc)
	assert.True(t, pq.IsEmpty(), tc)
	assert.Equal(t, 0, pq.DelMax(), tc)

}

func TestUnOrderMaxPQ(t *testing.T) {
	pq := priorityQueue.NewUnOrderMaxPQ(5)
	test(t, pq, "UnOrderMaxPQ")
}

func TestOrderMaxPQ(t *testing.T) {
	pq := priorityQueue.NewOrderMaxPQ(5)
	test(t, pq, "OrderMaxPQ")
}

func TestBinaryHeapMaxPQ(t *testing.T) {
	pq := priorityQueue.NewBinaryHeapMaxPQ(5)
	test(t, pq, "binaryHeap")
}

func Benchmark1KUnOrder(b *testing.B) {
	pq := priorityQueue.NewUnOrderMaxPQ(1000)
	for i := 0; i < 1000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}
func Benchmark10KUnOrder(b *testing.B) {
	pq := priorityQueue.NewUnOrderMaxPQ(10000)
	for i := 0; i < 10000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}

func Benchmark100KUnOrder(b *testing.B) {
	pq := priorityQueue.NewUnOrderMaxPQ(100000)
	for i := 0; i < 100000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}

func Benchmark1KOrder(b *testing.B) {
	pq := priorityQueue.NewOrderMaxPQ(1000)
	for i := 0; i < 1000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}
func Benchmark10KOrder(b *testing.B) {
	pq := priorityQueue.NewOrderMaxPQ(10000)
	for i := 0; i < 10000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}

func Benchmark100KOrder(b *testing.B) {
	pq := priorityQueue.NewOrderMaxPQ(100000)
	for i := 0; i < 100000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}

func Benchmark1KBinaryHeap(b *testing.B) {
	pq := priorityQueue.NewBinaryHeapMaxPQ(1000)
	for i := 0; i < 1000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}
func Benchmark10KBinaryHeap(b *testing.B) {
	pq := priorityQueue.NewBinaryHeapMaxPQ(10000)
	for i := 0; i < 10000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}

func Benchmark100KBinaryHeap(b *testing.B) {
	pq := priorityQueue.NewBinaryHeapMaxPQ(100000)
	for i := 0; i < 100000; i++ {
		pq.Insert(rand.Intn(1000000))
	}
	for i := 0; i < b.N; i++ {
		pq.Insert(rand.Intn(1000000))
		pq.DelMax()
	}
}
