package symbolTable_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/lzh2nix/algorithms101/symbolTable"
	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, st symbolTable.SymbolTable, tc string) {
	st.Put("aaa", 0)
	st.Put("bbb", 1)
	st.Put("ccc", 100)
	st.Put("ddd", 9999)
	st.Put("eee", 2423)
	st.Put("fff", 42342)
	st.Put("ggg", 34241)
	assert.True(t, st.Contain("aaa"), tc)
	assert.False(t, st.Contain("kkk"), tc)
	e, v := st.Get("aaa")
	assert.True(t, e, tc)
	assert.Equal(t, 0, v, tc)
	e, _ = st.Get("sssss")
	assert.False(t, e, tc)

	assert.Equal(t, 7, st.Size(), tc)
	st.Put("aaa", 10000)
	e, v = st.Get("aaa")
	assert.True(t, e, tc)
	assert.Equal(t, 10000, v, tc)
	assert.False(t, st.IsEmpty(), tc)

	st.Delete("aaa")
	st.Delete("bbb")
	st.Delete("ccc")
	st.Delete("ddd")
	st.Delete("eee")
	st.Delete("fff")
	st.Delete("ggg")
	assert.True(t, st.IsEmpty(), tc)
	assert.Equal(t, 0, st.Size(), tc)
}

// TestLinkedListST ...
func TestLinkedListST(t *testing.T) {
	st := symbolTable.NewLinkedListST()
	test(t, st, "LinkedListST")
}

func Benchmark1KLinkedListST(b *testing.B) {
	st := symbolTable.NewLinkedListST()
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(k)
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
		}
		st[k] = n
	}
}
