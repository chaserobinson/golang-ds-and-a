package list

type Node struct {
	next  *Node
	value interface{}
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}

// trimmed removes references and returns node
func (n *Node) trimmed() *Node {
	n.next = nil
	return n
}

type SinglyLinked struct {
	head   *Node
	tail   *Node
	length int
}

func NewSinglyLinked(values ...interface{}) *SinglyLinked {
	sl := &SinglyLinked{}
	for _, value := range values {
		sl.Push(value)
	}
	return sl
}

const indexOutOfRange = "Index out of range."

// SliceValues will return a slice of all values within list
func (sl *SinglyLinked) SliceValues() []interface{} {
	var slice []interface{}
	if sl.head != nil {
		slice = append(slice, sl.head.value)
		curr := sl.head
		for curr.next != nil {
			slice = append(slice, curr.next.value)
			curr = curr.next
		}
	}
	return slice
}

// Push a value onto the end of the list
func (sl *SinglyLinked) Push(value interface{}) *Node {
	node := NewNode(value)

	if sl.head == nil {
		sl.head = node
	} else {
		sl.tail.next = node
	}
	sl.tail = node
	sl.length++
	return node
}

// Pop the last value from the list
func (sl *SinglyLinked) Pop() *Node {
	if sl.head == nil {
		return nil
	}

	last := sl.tail
	sl.length--

	if sl.length == 0 {
		sl.head = nil
		sl.tail = nil
		return last
	}

	curr := sl.head
	for curr.next != sl.tail {
		curr = curr.next
	}
	sl.tail = curr
	sl.tail.next = nil
	return last
}

// Shift removes and returns the first (head) node
func (sl *SinglyLinked) Shift() *Node {
	if sl.head == nil {
		return nil
	}

	node := sl.head
	sl.head = node.next
	sl.length--
	if sl.length == 0 {
		sl.tail = nil
	}
	return node.trimmed()
}

// Unshift adds a node to the beginning
func (sl *SinglyLinked) Unshift(value interface{}) *Node {
	node := NewNode(value)
	if sl.head == nil {
		sl.tail = node
	} else {
		node.next = sl.head
	}

	sl.head = node
	sl.length++
	return node
}

// Get returns node at index
func (sl *SinglyLinked) Get(i int) *Node {
	if sl.head == nil || i < 0 || i >= sl.length {
		return nil
	}

	node := sl.head
	for j := 0; j < i; j++ {
		node = node.next
	}
	return node
}

// Set value at index
func (sl *SinglyLinked) Set(i int, value interface{}) *Node {
	node := sl.Get(i)
	if node != nil {
		node.value = value
	}
	return node
}

// Insert a new node with specified value
func (sl *SinglyLinked) Insert(i int, value interface{}) *Node {
	if i < 0 || i > sl.length {
		panic(indexOutOfRange)
	}
	if i == 0 {
		return sl.Unshift(value)
	}
	if i == sl.length {
		return sl.Push(value)
	}
	node := NewNode(value)
	prev := sl.Get(i - 1)
	node.next = prev.next
	prev.next = node
	sl.length++
	return node
}

// Remove node at index
func (sl *SinglyLinked) Remove(i int) *Node {
	if i < 0 || i > sl.length-1 {
		panic("Index out of range.")
	}
	if i == 0 {
		return sl.Shift()
	}
	if i == sl.length-1 {
		return sl.Pop()
	}
	prev := sl.Get(i - 1)
	node := prev.next
	prev.next = node.next
	sl.length--
	return node.trimmed()
}

// Reverse inverts the order of elements within list
func (sl *SinglyLinked) Reverse() *SinglyLinked {
	if sl.length <= 1 {
		return sl
	}

	node := sl.head
	sl.head = sl.tail
	sl.tail = node

	var prev *Node = nil
	for node != nil {
		curr := node
		node = node.next
		curr.next = prev
		prev = curr
	}
	return sl
}
