package ds

func NewLinkedList[T any](initialData ...[]T) *LinkedList[T] {
	ll := &LinkedList[T]{
		head: nil,
		tail: nil,
	}

	for _, data := range initialData {
		for _, i := range data {
			ll.Append(i)
		}
	}

	return ll
}

type LinkedList[T any] struct {
	head  *node[T]
	tail  *node[T]
	count int
}

func (ll *LinkedList[T]) Len() int {
	return ll.count
}

func (ll *LinkedList[T]) Append(value T) {
	ll.count++
	newNode := ll.newNode(value)
	if ll.head == nil {
		// list is empty
		ll.head = newNode
		return
	}
	if ll.tail == nil {
		// list has one node
		newNode.prev = ll.head
		ll.head.next = newNode
		ll.tail = newNode
		return
	}

	newNode.prev = ll.tail
	ll.tail.next = newNode
	ll.tail = newNode
}

func (ll *LinkedList[T]) Get(index int) (T, bool) {
	n := ll.head

	if index >= ll.count {
		var zero T
		return zero, false
	}

	for i := 0; i < index; i++ {
		n = n.next
	}
	return n.val, true
}

func (ll *LinkedList[T]) Iter() <-chan T {
	retChan := make(chan T)
	go func() {
		defer close(retChan)
		n := ll.head
		for n != nil {
			val := n.val
			n = n.next
			retChan <- val
		}
	}()
	return retChan
}

func (ll LinkedList[T]) newNode(val T) *node[T] {
	return &node[T]{
		val: val,
	}
}

type node[T any] struct {
	val  T
	next *node[T]
	prev *node[T]
}
