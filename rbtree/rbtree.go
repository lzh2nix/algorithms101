package rbtree

import "fmt"

const RED = 0
const BLACK = 1

type RBNode struct {
	key         string
	value       int
	left, right *RBNode
	count       int
	color       int
}

func size(x *RBNode) int {
	if x == nil {
		return 0
	}
	return x.count
}

type RBT struct {
	root *RBNode
}

func NewRBT() *RBT {
	return &RBT{}
}
func (rbt *RBT) Put(key string, v int) {
	rbt.root = put(rbt.root, key, v)
	//	rbt.root.color = BLACK
}
func isRed(x *RBNode) bool {
	if x == nil {
		return false
	}
	return x.color == RED
}

func rotateLeft(h *RBNode) *RBNode {
	x := h.right
	h.right = x.left
	x.left = h
	x.count = h.count
	h.count = size(h.left) + size(h.right) + 1
	x.color = h.color
	h.color = RED
	return x
}

func rotateRight(h *RBNode) *RBNode {
	x := h.left
	h.left = x.right
	x.right = h
	x.count = h.count
	h.count = size(h.left) + size(h.right) + 1
	x.color = h.color
	h.color = RED
	return x
}
func flip(x int) int {
	if x == RED {
		return BLACK
	}
	return RED
}
func flipColors(h *RBNode) {
	h.color = flip(h.color)
	h.left.color = flip(h.left.color)
	h.right.color = flip(h.right.color)
}

func put(x *RBNode, key string, v int) *RBNode {
	if x == nil {
		return &RBNode{key: key, value: v, count: 1, color: RED}
	}
	if x.key > key {
		x.left = put(x.left, key, v)
	} else if x.key < key {
		x.right = put(x.right, key, v)
	} else {
		x.value = v
	}
	if isRed(x.right) && !isRed(x.left) {
		x = rotateLeft(x)
	}
	if isRed(x.left) && isRed(x.left.left) {
		x = rotateRight(x)
	}
	if isRed(x.left) && isRed(x.right) {
		flipColors(x)
	}
	x.count = size(x.left) + size(x.right) + 1
	return x
}
func (rbt *RBT) Get(key string) (bool, int) {
	x := rbt.root
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

func moveRedLeft(h *RBNode) *RBNode {
	flipColors(h)
	if isRed(h.right.left) {
		h.right = rotateRight(h.right)
		h = rotateLeft(h)
		flipColors(h)
	}
	return h
}
func moveRedRight(h *RBNode) *RBNode {
	flipColors(h)
	if isRed(h.left.left) {
		h = rotateRight(h)
		flipColors(h)
	}
	return h
}

// Delete ...
func (rbt *RBT) Delete(key string) {
	rbt.root = delete(rbt.root, key)
	//	if rbt.root != nil {
	//rbt.root.color = BLACK
	//}
}

func delete(h *RBNode, key string) *RBNode {
	if h == nil {
		return nil
	}
	if h.key > key {
		if h.left != nil && !isRed(h.left) && !isRed(h.left.left) {
			h = moveRedLeft(h)
		}
		h.left = delete(h.left, key)

	} else {
		if isRed(h.left) {
			h = rotateRight(h)
		}
		if h.key == key && h.right == nil {
			return nil
		}
		if h.right != nil && !isRed(h.right) && !isRed(h.right.left) {
			h = moveRedRight(h)
		}
		if h.key == key {
			min := min(h.right)
			h.value = min.value
			h.key = min.key
			h.right = deleteMin(h.right)
		} else {
			h.right = delete(h.right, key)
		}
	}
	h = fixUp(h)
	h.count = size(h.left) + size(h.right) + 1
	return h

}

// Contain ...
func (rbt *RBT) Contain(key string) bool {
	e, _ := rbt.Get(key)
	return e == true
}

// Size of RBT
func (rbt *RBT) Size() int {
	if rbt.root == nil {
		return 0
	}
	return rbt.root.count
}

func (rbt *RBT) Min() (bool, string) {
	min := min(rbt.root)
	if min == nil {
		return false, ""
	}
	return true, min.key
}

func min(x *RBNode) *RBNode {
	for x != nil {
		if x.left == nil {
			return x
		}
		x = x.left
	}
	return nil
}

func fixUp(h *RBNode) *RBNode {
	if isRed(h.right) {
		h = rotateLeft(h)
	}
	if isRed(h.left) && isRed(h.left.left) {
		h = rotateRight(h)
	}
	if isRed(h.right) && isRed(h.left) {
		flipColors(h)
	}
	return h
}
func deleteMin(x *RBNode) *RBNode {
	if x.left == nil {
		return nil
	}
	if x.left != nil && !isRed(x.left) && !isRed(x.left.left) {
		x = moveRedLeft(x)
	}
	x.left = deleteMin(x.left)
	x = fixUp(x)
	x.count = size(x.left) + size(x.right) + 1
	return x
}
func (rbt *RBT) IsEmpty() bool {
	return rbt.root == nil
}

func (rbt *RBT) Print() {
	print(rbt.root)
	fmt.Println()
}

func print(x *RBNode) {
	if x == nil {
		return
	}
	print(x.left)
	fmt.Print(x.key, x.color, x.count, ",")
	print(x.right)
}
