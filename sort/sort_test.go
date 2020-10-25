package sort_test

import (
	"fmt"
	"math/rand"
	stdSort "sort"
	"testing"
	"time"

	"github.com/lzh2nix/algorithms101/sort"
	"github.com/stretchr/testify/assert"
)

func testSort(t *testing.T, sort func([]int), tc string) {
	rand.Seed(time.Now().UnixNano())
	a := []int{}
	for i := 0; i < 20; i++ {
		a = append(a, rand.Intn(1000))
	}
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	fmt.Println(tc+" before test = ", a)
	sort(a)
	fmt.Println(tc+" before test = ", a)
	assert.True(t, isSorted(a), tc)

}
func isSorted(a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i+1] < a[i] {
			return false
		}
	}
	return true
}

func TestSelectSort(t *testing.T) {
	testSort(t, sort.SelectSort, "select sort")
}
func TestInsertSort(t *testing.T) {
	testSort(t, sort.InsertSort, "insert sort")
}

func TestShellSort(t *testing.T) {
	testSort(t, sort.ShellSort, "shell sort")
}

func TestMergeSort(t *testing.T) {
	testSort(t, sort.MergeSort, "merge sort")
}
func randomArray(n int) []int {
	a := []int{}
	for i := 0; i < n; i++ {
		a = append(a, rand.Intn(1000000))
	}
	return a
}
func Benchmark1kSeletSort(b *testing.B) {
	a := randomArray(1000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.SelectSort(d)
	}
}
func Benchmark1kInsertSort(b *testing.B) {
	a := randomArray(1000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.InsertSort(d)
	}
}
func Benchmark1kShellSort(b *testing.B) {
	a := randomArray(1000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.InsertSort(d)
	}
}

func Benchmark1kMergeSort(b *testing.B) {
	a := randomArray(1000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.MergeSort(d)
	}
}
func Benchmark1kStdSort(b *testing.B) {
	a := randomArray(1000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		stdSort.Ints(d)
	}
}

func Benchmark10kSelectSort(b *testing.B) {
	a := randomArray(10000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.SelectSort(d)
	}
}
func Benchmark10kInsertSort(b *testing.B) {
	a := randomArray(10000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.SelectSort(d)
	}
}
func Benchmark10kShellSort(b *testing.B) {
	a := randomArray(10000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.SelectSort(d)
	}
}
func Benchmark10kMergeSort(b *testing.B) {
	a := randomArray(10000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.MergeSort(d)
	}
}
func Benchmark10kStdSort(b *testing.B) {
	a := randomArray(10000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		stdSort.Ints(d)
	}
}

func Benchmark100kSelectSort(b *testing.B) {
	a := randomArray(100000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.SelectSort(d)
	}
}

func Benchmark100kInsertSort(b *testing.B) {
	a := randomArray(100000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.SelectSort(d)
	}
}

func Benchmark100kShellSort(b *testing.B) {
	a := randomArray(100000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.SelectSort(d)
	}
}

func Benchmark100kMergeSort(b *testing.B) {
	a := randomArray(100000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.MergeSort(d)
	}
}
func Benchmark100kStdSort(b *testing.B) {
	a := randomArray(100000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		stdSort.Ints(d)
	}
}
