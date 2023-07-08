package list_test

import (
	"testing"

	"github.com/mgnsk/list"
	. "github.com/onsi/gomega"
)

func TestPushFront(t *testing.T) {
	var l list.List[int]

	g := NewWithT(t)

	l.PushFront(list.NewElement(0))
	g.Expect(l.Len()).To(Equal(1))

	l.PushFront(list.NewElement(1))
	g.Expect(l.Len()).To(Equal(2))

	expectValidRing(g, &l)
}

func TestPushBack(t *testing.T) {
	var l list.List[int]

	g := NewWithT(t)

	l.PushFront(list.NewElement(0))
	g.Expect(l.Len()).To(Equal(1))

	l.PushFront(list.NewElement(1))
	g.Expect(l.Len()).To(Equal(2))

	expectValidRing(g, &l)
}

func TestMoveToFront(t *testing.T) {
	t.Run("moving the back element", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.MoveToFront(l.Back())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("moving the middle element", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.PushBack(list.NewElement("three"))
		l.MoveToFront(l.Front().Next())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("three"))
	})
}

func TestMoveToBack(t *testing.T) {
	t.Run("moving the front element", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.MoveToBack(l.Front())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("moving the middle element", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.PushBack(list.NewElement("three"))
		l.MoveToBack(l.Front().Next())

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("one"))
		g.Expect(l.Back().Value).To(Equal("two"))
	})
}

func TestMoveForward(t *testing.T) {
	t.Run("overflow", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.Move(l.Front(), 3)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("not overflow", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.PushBack(list.NewElement("three"))
		l.Move(l.Front(), 1)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Front().Next().Value).To(Equal("one"))
		g.Expect(l.Back().Value).To(Equal("three"))
	})
}

func TestMoveBackwards(t *testing.T) {
	t.Run("overflow", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.Move(l.Back(), -3)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("two"))
		g.Expect(l.Back().Value).To(Equal("one"))
	})

	t.Run("not overflow", func(t *testing.T) {
		var l list.List[string]

		g := NewWithT(t)

		l.PushBack(list.NewElement("one"))
		l.PushBack(list.NewElement("two"))
		l.PushBack(list.NewElement("three"))
		l.Move(l.Back(), -1)

		expectValidRing(g, &l)
		g.Expect(l.Front().Value).To(Equal("one"))
		g.Expect(l.Front().Next().Value).To(Equal("three"))
		g.Expect(l.Back().Value).To(Equal("two"))
	})
}

func TestDo(t *testing.T) {
	var l list.List[string]

	g := NewWithT(t)

	l.PushBack(list.NewElement("one"))
	l.PushBack(list.NewElement("two"))
	l.PushBack(list.NewElement("three"))

	g.Expect(l.Len()).To(Equal(3))
	expectValidRing(g, &l)

	var elems []string
	l.Do(func(e *list.Element[string]) bool {
		elems = append(elems, e.Value)
		return true
	})

	g.Expect(elems).To(Equal([]string{"one", "two", "three"}))
}

func expectValidRing[T any](g *WithT, l *list.List[T]) {
	g.Expect(l.Len()).To(BeNumerically(">", 0))
	g.Expect(l.Front()).To(Equal(l.Back().Next()))
	g.Expect(l.Back()).To(Equal(l.Front().Prev()))

	{
		expectedFront := l.Front()

		front := l.Front()

		for i := 0; i < l.Len(); i++ {
			front = front.Next()
		}

		g.Expect(front).To(Equal(expectedFront))
	}

	{
		expectedBack := l.Back()

		back := l.Back()

		for i := 0; i < l.Len(); i++ {
			back = back.Prev()
		}

		g.Expect(back).To(Equal(expectedBack))
	}
}
