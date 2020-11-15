package symbolTable_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/lzh2nix/algorithms101/rbtree"
	"github.com/lzh2nix/algorithms101/symbolTable"
	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, st symbolTable.SymbolTable, tc string) {
	st.Put("ddd", 9999)
	st.Put("aaa", 0)
	st.Put("eee", 2423)
	st.Put("bbb", 1)
	st.Put("fff", 42342)
	st.Put("ggg", 34241)
	st.Put("ccc", 100)
	assert.True(t, st.Contain("aaa"), tc)
	assert.False(t, st.Contain("kkk"), tc)
	e, v := st.Get("aaa")
	assert.True(t, e, tc)
	assert.Equal(t, 0, v, tc)
	e, _ = st.Get("sssss")
	assert.False(t, e, tc)

	e, v = st.Get("bbb")
	assert.True(t, e, tc)
	assert.Equal(t, 1, v, tc)

	e, v = st.Get("ddd")
	assert.True(t, e, tc)
	assert.Equal(t, 9999, v, tc)

	e, v = st.Get("ccc")
	assert.True(t, e, tc)
	assert.Equal(t, 100, v, tc)

	assert.Equal(t, 7, st.Size(), tc)
	st.Put("aaa", 10000)
	e, v = st.Get("aaa")
	assert.True(t, e, tc)
	assert.Equal(t, 10000, v, tc)
	assert.False(t, st.IsEmpty(), tc)
	st.Print()
	e2, v2 := st.Min()
	assert.True(t, e2, tc)
	assert.Equal(t, "aaa", v2, tc)
	st.Delete("ddd")
	st.Delete("aaa")
	st.Delete("bbb")
	st.Delete("ccc")
	st.Delete("fff")
	st.Delete("ggg")
	e2, v2 = st.Min()
	assert.True(t, e2, tc)
	assert.Equal(t, "eee", v2, tc)
	st.Delete("eee")
	assert.True(t, st.IsEmpty(), tc)
	assert.Equal(t, 0, st.Size(), tc)
	e, _ = st.Get("eee")
	assert.False(t, e)
	e2, _ = st.Min()
	assert.False(t, e, tc)
}

// TestLinkedListST ...
func TestLinkedListST(t *testing.T) {
	st := symbolTable.NewLinkedListST()
	test(t, st, "LinkedListST")
}

// TestLinkedListST ...
func TestBSTST(t *testing.T) {
	st := symbolTable.NewBST()
	test(t, st, "BST version symbolTable")
	for i := 0; i < 100; i++ {
		n := rand.Intn(100)
		k := fmt.Sprintf("%d", n)
		if i%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
	st.Print()
}

func Benchmark1KLinkedListST(b *testing.B) {
	st := symbolTable.NewLinkedListST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark1KBSTST(b *testing.B) {
	st := symbolTable.NewBST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark1kRTB(b *testing.B) {
	st := rbtree.NewRBT()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark1KStdMap(b *testing.B) {
	st := map[string]int{}
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			delete(st, k)
			continue
		}
		st[k] = n
	}
}

func Benchmark10KLinkedListST(b *testing.B) {
	st := symbolTable.NewLinkedListST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark10KBSTST(b *testing.B) {
	st := symbolTable.NewBST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark10KRBT(b *testing.B) {
	st := rbtree.NewRBT()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark10KStdMap(b *testing.B) {
	st := map[string]int{}
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			delete(st, k)
			continue
		}
		st[k] = n
	}
}

func Benchmark100KLinkedListST(b *testing.B) {
	st := symbolTable.NewLinkedListST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}

func Benchmark100KBSTST(b *testing.B) {
	st := symbolTable.NewBST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark100RBT(b *testing.B) {
	st := rbtree.NewRBT()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark100KStdMap(b *testing.B) {
	st := map[string]int{}
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			delete(st, k)
			continue
		}
		st[k] = n
	}
}

func Benchmark1MLinkedListST(b *testing.B) {
	st := symbolTable.NewLinkedListST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark1MBSTST(b *testing.B) {
	st := symbolTable.NewBST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark1MRBT(b *testing.B) {
	st := rbtree.NewRBT()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
			continue
		}
		st.Put(k, n)
	}
}
func Benchmark1MStdMap(b *testing.B) {
	st := map[string]int{}
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			delete(st, k)
			continue
		}
		st[k] = n
	}
}
