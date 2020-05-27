package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *listItem                  // первый Item
	Back() *listItem                   // последний Item
	PushFront(v interface{}) *listItem // добавить значение в начало
	PushBack(v interface{}) *listItem  // добавить значение в конец
	Remove(i *listItem)                // удалить элемент
	MoveToFront(i *listItem)           // переместить элемент в начало

}

type listItem struct {
	Value interface{} // значение
	Next  *listItem   // следующий элемент
	Prev  *listItem   // предыдущий элемент

}

type list struct {
	first  *listItem
	last   *listItem
	length int
}

func NewList() List {
	return &list{}
}

func (l list) Len() int {
	return l.length
}

func (l list) Front() *listItem {
	return l.first
}

func (l list) Back() *listItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *listItem {
	item := &listItem{Value: v, Next: nil, Prev: l.first}

	if l.first == nil {
		l.last = item
	} else {
		l.first.Next = item
	}
	item.Prev = l.first
	l.first = item
	l.length++

	return item
}

func (l *list) PushBack(v interface{}) *listItem {
	item := &listItem{Value: v, Next: l.last, Prev: nil}

	if l.last == nil {
		l.first = item
	} else {
		l.last.Prev = item
	}

	item.Next = l.last
	l.last = item
	l.length++

	return item
}

func (l *list) Remove(i *listItem) {
	switch {
	case i.Prev == nil && i.Next == nil:
		l.first = nil
		l.last = nil
	case i.Prev == nil:
		l.last = i.Next
		i.Next.Prev = nil
	case i.Next == nil:
		l.first = i.Prev
		i.Prev.Next = nil
	default:
		i.Prev.Next = i.Next
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil

	l.length--
}

func (l *list) MoveToFront(i *listItem) {
	if i == l.first {
		return
	}

	l.Remove(i)
	l.PushFront(i.Value)
}
