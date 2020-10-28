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
	for i := 0; i < 50; i++ {
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

func TestQuickSort(t *testing.T) {
	testSort(t, sort.QuickSort, "quick sort")
	//sort.QuickSort([]int{194, 924, 238, 372, 148, 217, 773, 773, 687, 654, 789, 477, 565, 230, 715, 773, 285, 63, 343, 248, 949, 829, 592, 997, 195, 629, 768, 600, 167, 18, 701, 937, 599, 542, 700, 918, 72, 810, 287, 925, 192, 801, 131, 300, 310, 56, 155, 92, 680, 605})
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
func Benchmark1kQuickSort(b *testing.B) {
	a := randomArray(1000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.QuickSort(d)
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
func Benchmark10kQuickSort(b *testing.B) {
	a := randomArray(10000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.QuickSort(d)
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
func Benchmark100kQuickSort(b *testing.B) {
	a := randomArray(100000)
	for i := 0; i < b.N; i++ {
		d := make([]int, len(a))
		copy(d, a)
		sort.QuickSort(d)
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
