package symbolTable

import "fmt"

// string:int based symbol table interface
type SymbolTable interface {
	// put key-value pair to symbol table a[key] = val(if key already exist, update it)
	Put(key string, val int)
	// value pair with key(false if not exist)
	Get(key string) (bool, int)
	// remove key form table
	Delete(key string)
	// is there a value paired with key
	Contain(key string) bool
	// is the table empty
	IsEmpty() bool
	// number of key-value pair in symbol table
	Size() int
	// print all element
	Print()
	// min in SymbolTable
	Min() (bool, string)
}

// linked list based symbol table
//    search  insert Delete
//      N       N      N
type LinkedListST struct {
	Node
	size int
}

// linked list node
type Node struct {
	val  int
	key  string
	next *Node
}

// new a linkedlist based symbol table
func NewLinkedListST() SymbolTable {
	return &LinkedListST{}
}

// implement SymbolTable's Put
func (st *LinkedListST) Put(key string, val int) {
	n := st.next
	for n != nil {
		if n.key == key {
			n.val = val
			return
		}
		n = n.next
	}
	n = &Node{key: key, val: val, next: st.next}
	st.next = n
	st.size++
}

// implement SymbolTable's Get
func (st *LinkedListST) Get(key string) (bool, int) {
	n := st.next
	for n != nil {
		if n.key == key {
			return true, n.val
		}
		n = n.next
	}
	return false, 0
}

// implement SymbolTable's Delete
func (st *LinkedListST) Delete(key string) {
	pre := &st.Node
	n := st.next
	for n != nil {
		if n.key == key {
			pre.next = n.next
			n.next = nil
			st.size--
			break
		}
		pre = n
		n = n.next
	}
}

// implement SymbolTable's Contain
func (st *LinkedListST) Contain(key string) bool {
	t, _ := st.Get(key)
	return t == true
}

// implement SymbolTable's IsEmpty
func (st *LinkedListST) IsEmpty() bool {
	return st.size == 0
}

// implement SybmolTable's Size
func (st *LinkedListST) Size() int {
	return st.size
}

func (st *LinkedListST) Print() {
	v := st.Node.next
	for v != nil {
		fmt.Print(v.key, ",")
		v = v.next
	}
	fmt.Println()
}

func (st *LinkedListST) Min() (bool, string) {
	v := st.Node.next
	min := v
	if min == nil {
		return false, ""
	}
	for v != nil {
		if v.key < min.key {
			min = v
		}
		v = v.next
	}
	return true, min.key
}

type BSTNode struct {
	key         string
	value       int
	left, right *BSTNode
	count       int
}

func size(x *BSTNode) int {
	if x == nil {
		return 0
	}
	return x.count
}

type BST struct {
	root *BSTNode
}

func NewBST() SymbolTable {
	return &BST{}
}
func (bst *BST) Put(key string, v int) {
	bst.root = put(bst.root, key, v)
}
func put(x *BSTNode, key string, v int) *BSTNode {
	if x == nil {
		return &BSTNode{key: key, value: v, count: 1}
	}
	if x.key > key {
		x.left = put(x.left, key, v)
	} else if x.key < key {
		x.right = put(x.right, key, v)
	} else {
		x.value = v
	}
	x.count = size(x.left) + size(x.right) + 1
	return x
}
func (bst *BST) Get(key string) (bool, int) {
	x := bst.root
	for x != nil {
		if x.key == key {
			return true, x.value
		} else if x.key > key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return false, 0
}

// Delete ...
func (bst *BST) Delete(key string) {
	bst.root = delete(bst.root, key)
}

func delete(x *BSTNode, key string) *BSTNode {
	if x == nil {
		return nil
	}
	if x.key > key {
		x.left = delete(x.left, key)
	} else if x.key < key {
		x.right = delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		t := x
		x = min(t.right)
		x.right = deleteMin(t.right)
		x.left = t.left
		t = nil
	}
	x.count = size(x.left) + size(x.right) + 1
	return x
}

// Contain ...
func (bst *BST) Contain(key string) bool {
	e, _ := bst.Get(key)
	return e == true
}

// Size of BST
func (bst *BST) Size() int {
	if bst.root == nil {
		return 0
	}
	return bst.root.count
}

func (bst *BST) Min() (bool, string) {
	min := min(bst.root)
	if min == nil {
		return false, ""
	}
	return true, min.key
}

func min(x *BSTNode) *BSTNode {
	for x != nil {
		if x.left == nil {
			return x
		}
		x = x.left
	}
	return nil
}

func deleteMin(x *BSTNode) *BSTNode {
	if x.left == nil {
		return x.left
	}
	x.left = deleteMin(x.left)
	x.count = size(x.left) + size(x.right) + 1
	return x
}
func (bst *BST) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BST) Print() {
	print(bst.root)
	fmt.Println()
}

func print(x *BSTNode) {
	if x == nil {
		return
	}
	print(x.left)
	fmt.Print(x.key, ",")
	print(x.right)
}
