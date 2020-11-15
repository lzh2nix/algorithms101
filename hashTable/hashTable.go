package hashTable

import (
	"fmt"
	"strconv"
)

// generic type for hash Key
type Key interface {
	// a.Equal(b)
	// a == b
	Equal(interface{}) bool
	// hashcode of Key
	HashCode() int
}

type HashTable interface {
	// put a element to HashTable
	Put(k Key, e interface{})
	// get from HashTable
	Get(k Key) (interface{}, bool)
	// Delete for HashTable
	Delete(k Key)
	// print the hash table
	Print()
}

type Node struct {
	key   Key
	value interface{}
	next  *Node
}

type SeparateChainingHashTable struct {
	// number of chains
	m     int
	nodes []*Node
}

func NewSeparateChainingHashTable(n int) HashTable {
	nodes := make([]*Node, n, n)
	return &SeparateChainingHashTable{n, nodes}
}

// impl HashTables Put
func (ht *SeparateChainingHashTable) Put(key Key, value interface{}) {
	h := key.HashCode()
	idx := h % ht.m

	for o := ht.nodes[idx]; o != nil; o = o.next {
		if o.key.Equal(key) {
			o.value = value
			return
		}
	}
	ht.nodes[idx] = &Node{key, value, ht.nodes[idx]}
}

// impl HashTables Get
func (ht *SeparateChainingHashTable) Get(key Key) (interface{}, bool) {
	h := key.HashCode()
	idx := h % ht.m
	for o := ht.nodes[idx]; o != nil; o = o.next {
		if o.key.Equal(key) {
			return o.value, true
		}
	}
	return nil, false
}

// impl HashTables Delete
func (ht *SeparateChainingHashTable) Delete(key Key) {
	h := key.HashCode()
	idx := h % ht.m
	if ht.nodes[idx] == nil {
		return
	}
	if ht.nodes[idx].key.Equal(key) {
		ht.nodes[idx] = ht.nodes[idx].next
		return
	}
	prev := ht.nodes[idx]
	for o := ht.nodes[idx].next; o != nil; o = o.next {
		if o.key.Equal(key) {
			prev.next = o.next
			o.next = nil
			break
		}
		prev = o
	}
}

// impl HashTables Print
func (ht *SeparateChainingHashTable) Print() {
	for i := 0; i < ht.m; i++ {
		fmt.Print("[" + strconv.Itoa(i) + "]-->")
		for o := ht.nodes[i]; o != nil; o = o.next {
			fmt.Print(o.key, "-->")
		}
		fmt.Println("")
	}
}

type item struct {
	key   Key
	value interface{}
}
type LinearProbingHashTable struct {
	// number of chains
	m     int
	nodes []*item
}

func NewLinearProbingHashTable(n int) HashTable {
	nodes := make([]*item, n, n)

	return &LinearProbingHashTable{n, nodes}
}

// impl HashTables Put
func (ht *LinearProbingHashTable) Put(key Key, value interface{}) {
	h := key.HashCode()
	idx := h % ht.m
	count := 0
	for ; ht.nodes[idx] != nil; idx = (idx + 1) % ht.m {
		if ht.nodes[idx].key.Equal(key) {
			break
		}
		if count > ht.m {
			return
		}
		count++
	}
	ht.nodes[idx] = &item{key, value}
}

// impl HashTables Get
func (ht *LinearProbingHashTable) Get(key Key) (interface{}, bool) {
	h := key.HashCode()
	idx := h % ht.m
	count := 0
	for ; ht.nodes[idx] != nil; idx = (idx + 1) % ht.m {
		if ht.nodes[idx].key.Equal(key) {
			return ht.nodes[idx].value, true
		}
		if count > ht.m {
			break
		}
		count++
	}
	return nil, false
}

// impl HashTables Delete
func (ht *LinearProbingHashTable) Delete(key Key) {
	h := key.HashCode()
	idx := h % ht.m
	count := 0
	for ; ht.nodes[idx] != nil; idx = (idx + 1) % ht.m {
		if ht.nodes[idx].key.Equal(key) {
			ht.nodes[idx] = nil
			return
		}
		if count > ht.m {
			break
		}
		count++
	}
}

// impl HashTables Print
func (ht *LinearProbingHashTable) Print() {
	for i := 0; i < ht.m; i++ {
		fmt.Println(ht.nodes[i])
	}
}
