package list

// Element is a list element.
type Element[V any] struct {
	next, prev *Element[V]
	Value      V
}

// NewElement creates a list element.
func NewElement[V any](v V) *Element[V] {
	e := &Element[V]{
		Value: v,
	}
	e.next = e
	e.prev = e
	return e
}

// Next returns the next element or itself.
func (e *Element[V]) Next() *Element[V] {
	return e.next
}

// Prev returns the previous element or itself.
func (e *Element[V]) Prev() *Element[V] {
	return e.prev
}

// link inserts an element after this element.
func (e *Element[V]) link(s *Element[V]) {
	n := e.next
	e.next = s
	s.prev = e
	n.prev = s
	s.next = n
}

// unlink unlinks this element.
func (e *Element[V]) unlink() {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = e
	e.prev = e
}
