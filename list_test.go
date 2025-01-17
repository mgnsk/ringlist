package ringlist_test

import (
	"reflect"
	"testing"

	"github.com/mgnsk/ringlist"
)

func TestPushFront(t *testing.T) {
	var l ringlist.List[int]

	l.PushFront(0)
	assertEqual(t, l.Len(), 1)

	l.PushFront(1)
	assertEqual(t, l.Len(), 2)

	expectValidRing(t, &l)
}

func TestPushBack(t *testing.T) {
	var l ringlist.List[int]

	l.PushFront(0)
	assertEqual(t, l.Len(), 1)

	l.PushFront(1)
	assertEqual(t, l.Len(), 2)

	expectValidRing(t, &l)
}

func TestMoveToFront(t *testing.T) {
	t.Run("moving the back element", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.MoveToFront(l.Back())

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "two")
		assertEqual(t, l.Back().Value, "one")
	})

	t.Run("moving the middle element", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.PushBack("three")
		l.MoveToFront(l.Front().Next())

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "two")
		assertEqual(t, l.Back().Value, "three")
	})
}

func TestMoveToBack(t *testing.T) {
	t.Run("moving the front element", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.MoveToBack(l.Front())

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "two")
		assertEqual(t, l.Back().Value, "one")
	})

	t.Run("moving the middle element", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.PushBack("three")
		l.MoveToBack(l.Front().Next())

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "one")
		assertEqual(t, l.Back().Value, "two")
	})
}

func TestMoveBefore(t *testing.T) {
	t.Run("before itself", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.PushBack("three")
		expectValidRing(t, &l)

		one := l.Front()
		two := l.Front().Next()
		three := l.Front().Next().Next()

		assertEqual(t, one.Value, "one")
		assertEqual(t, two.Value, "two")
		assertEqual(t, three.Value, "three")

		l.MoveToFront(one)
		l.MoveToFront(two)
		l.MoveToFront(three)
		assertEqual(t, l.Len(), 3)

		expectHasExactElements(t, &l, "three", "two", "one")
	})
}

func TestMoveAfter(t *testing.T) {
	t.Run("after itself", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.PushBack("three")
		expectValidRing(t, &l)

		one := l.Front()
		two := l.Front().Next()
		three := l.Front().Next().Next()

		assertEqual(t, one.Value, "one")
		assertEqual(t, two.Value, "two")
		assertEqual(t, three.Value, "three")

		l.MoveToBack(three)
		l.MoveToBack(two)
		l.MoveToBack(one)
		assertEqual(t, l.Len(), 3)

		expectHasExactElements(t, &l, "three", "two", "one")
	})
}

func TestMoveForward(t *testing.T) {
	t.Run("overflow", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.Move(l.Front(), 3)

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "two")
		assertEqual(t, l.Back().Value, "one")
	})

	t.Run("not overflow", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.PushBack("three")
		l.Move(l.Front(), 1)

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "two")
		assertEqual(t, l.Front().Next().Value, "one")
		assertEqual(t, l.Back().Value, "three")
	})
}

func TestMoveBackwards(t *testing.T) {
	t.Run("overflow", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.Move(l.Back(), -3)

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "two")
		assertEqual(t, l.Back().Value, "one")
	})

	t.Run("not overflow", func(t *testing.T) {
		var l ringlist.List[string]

		l.PushBack("one")
		l.PushBack("two")
		l.PushBack("three")
		l.Move(l.Back(), -1)

		expectValidRing(t, &l)
		assertEqual(t, l.Front().Value, "one")
		assertEqual(t, l.Front().Next().Value, "three")
		assertEqual(t, l.Back().Value, "two")
	})
}

func TestDo(t *testing.T) {
	var l ringlist.List[string]

	l.PushBack("one")
	l.PushBack("two")
	l.PushBack("three")

	assertEqual(t, l.Len(), 3)
	expectValidRing(t, &l)

	var elems []string
	l.Do(func(e *ringlist.Element[string]) bool {
		elems = append(elems, e.Value)
		return true
	})

	assertEqual(t, elems, []string{"one", "two", "three"})
}

func expectHasExactElements[T any](t testing.TB, l *ringlist.List[T], elements ...T) {
	var elems []T

	l.Do(func(e *ringlist.Element[T]) bool {
		elems = append(elems, e.Value)

		return true
	})

	assertEqual(t, elems, elements)
}

func expectValidRing[T any](t testing.TB, l *ringlist.List[T]) {
	assertEqual(t, l.Len() > 0, true)
	assertEqual(t, l.Front(), l.Back().Next())
	assertEqual(t, l.Back(), l.Front().Prev())

	{
		expectedFront := l.Front()

		front := l.Front()

		for i := 0; i < l.Len(); i++ {
			front = front.Next()
		}

		assertEqual(t, front, expectedFront)
	}

	{
		expectedBack := l.Back()

		back := l.Back()

		for i := 0; i < l.Len(); i++ {
			back = back.Prev()
		}

		assertEqual(t, back, expectedBack)
	}
}

func assertEqual[T any](t testing.TB, a, b T) {
	if !reflect.DeepEqual(a, b) {
		t.Fatalf("expected '%v' to equal '%v'", a, b)
	}
}
