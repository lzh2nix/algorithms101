package hashTable_test

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"

	"github.com/lzh2nix/algorithms101/hashTable"
	"github.com/stretchr/testify/assert"
)

type StringKey struct {
	s string
}

func (k *StringKey) Equal(k2 interface{}) bool {
	k22, ok := k2.(*StringKey)
	if !ok {
		return false
	}
	return k.s == k22.s
}
func (k *StringKey) HashCode() int {
	h := 0
	for i := 0; i < len(k.s); i++ {
		h = int(k.s[i]) + 37*h
	}
	return int(math.Abs(float64(h)))
}
func test(t *testing.T, ht hashTable.HashTable, tc string) {
	ht.Put(&StringKey{"a"}, "a")
	ht.Put(&StringKey{"a0"}, "a0")
	ht.Put(&StringKey{"b2"}, "b2")
	ht.Put(&StringKey{"a2"}, "a2")
	ht.Put(&StringKey{"c3"}, "c3")
	ht.Put(&StringKey{"ddd"}, "ddd")
	ht.Put(&StringKey{"dddc"}, "dddc")
	ht.Print()
	v, e := ht.Get(&StringKey{"a"})
	assert.True(t, e, tc)
	assert.Equal(t, v.(string), "a")

	v, e = ht.Get(&StringKey{"ae"})
	assert.False(t, e, tc)

	ht.Put(&StringKey{"a"}, "e")
	v, e = ht.Get(&StringKey{"a"})
	assert.True(t, e, tc)
	assert.NotEqual(t, v.(string), "a")
	assert.Equal(t, v.(string), "e")

	v, e = ht.Get(&StringKey{"a0"})
	assert.True(t, e, tc)
	v, e = ht.Get(&StringKey{"b2"})
	assert.True(t, e, tc)
	v, e = ht.Get(&StringKey{"a2"})
	assert.True(t, e, tc)
	v, e = ht.Get(&StringKey{"c3"})
	assert.True(t, e, tc)
	v, e = ht.Get(&StringKey{"ddd"})
	assert.True(t, e, tc)
	v, e = ht.Get(&StringKey{"dddc"})
	assert.True(t, e, tc)

	ht.Delete(&StringKey{"a"})

	v, e = ht.Get(&StringKey{"a"})
	assert.False(t, e, tc)

	ht.Delete(&StringKey{"a0"})
	ht.Delete(&StringKey{"a"})
	ht.Delete(&StringKey{"b2"})
	ht.Delete(&StringKey{"a2"})
	ht.Delete(&StringKey{"c3"})
	ht.Delete(&StringKey{"ddd"})
	ht.Delete(&StringKey{"dddc"})
}

func TestSeparateChainingHashTable(t *testing.T) {
	ht := hashTable.NewSeparateChainingHashTable(5)
	test(t, ht, "separeatechaining hash table")
}

func TestLinearProbingHashTable(t *testing.T) {
	ht := hashTable.NewLinearProbingHashTable(8)
	test(t, ht, "linearprobing hash table")

	for i := 0; i < 10; i++ {
		ht.Put(&StringKey{strconv.Itoa(i)}, i)
	}
	_, e := ht.Get(&StringKey{"fdafas"})
	assert.False(t, e)
	ht.Delete(&StringKey{"fdafas"})

}

func Benchmark1KSeparateChainingHashTable(b *testing.B) {
	st := hashTable.NewSeparateChainingHashTable(100)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
	}
}
func Benchmark1KLinearProbingHashTable(b *testing.B) {
	st := hashTable.NewLinearProbingHashTable(1000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
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

func Benchmark10KSeparateChainingHashTable(b *testing.B) {
	st := hashTable.NewSeparateChainingHashTable(1000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
	}
}
func Benchmark10KLinearProbingHashTable(b *testing.B) {
	st := hashTable.NewLinearProbingHashTable(10000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(10000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
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

func Benchmark100KSeparateChainingHashTable(b *testing.B) {
	st := hashTable.NewSeparateChainingHashTable(10000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
	}
}
func Benchmark100KLinearProbingHashTable(b *testing.B) {
	st := hashTable.NewLinearProbingHashTable(100000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(100000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
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

func Benchmark1MSeparateChainingHashTable(b *testing.B) {
	st := hashTable.NewSeparateChainingHashTable(100000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
	}
}
func Benchmark1MLinearProbingHashTable(b *testing.B) {
	st := hashTable.NewLinearProbingHashTable(1000000)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000000)
		k := fmt.Sprintf("%d", n)
		if n%3 == 0 {
			st.Delete(&StringKey{k})
			continue
		}
		st.Put(&StringKey{k}, n)
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
