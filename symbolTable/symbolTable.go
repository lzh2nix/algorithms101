package symbolTable

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
