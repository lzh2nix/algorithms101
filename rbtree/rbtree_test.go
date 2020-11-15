package rbtree_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/lzh2nix/algorithms101/rbtree"
	"github.com/stretchr/testify/assert"
)

func test(t *testing.T, st *rbtree.RBT, tc string) {
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
func TestRBT(t *testing.T) {
	st := rbtree.NewRBT()
	test(t, st, "RBT version symbolTable")
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
