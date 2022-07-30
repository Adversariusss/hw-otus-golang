package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type doublelist struct {
	first *ListItem
	last  *ListItem
	size  int
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

func (list *doublelist) Len() int {
	if list == nil {
		return 0
	}
	return list.size
}

func (list *doublelist) Front() interface{} {
	return list.first
}

func (list *doublelist) Back() interface{} {
	return list.last
}

func (list *doublelist) Clear() {
	list.first = nil
	list.last = nil
	list.size = 0
}

func (list *doublelist) PushFront(values interface{}) *ListItem {
	//for _, value := range values {
	newListItem := &ListItem{Value: values, Next: list.first, Prev: nil}
	if list.size == 0 {
		list.first = newListItem
		list.last = newListItem
	} else {
		list.first.Prev = newListItem
		list.first = newListItem
	}
	list.size++
	return list.first
}

func (list *doublelist) PushBack(values interface{}) *ListItem {

	newElement := &ListItem{Value: values, Prev: list.last, Next: nil}
	if list.size == 0 {
		list.first = newElement
		list.last = newElement
	} else {
		list.last.Next = newElement
		list.last = newElement
	}
	list.size++
	return list.last
}

func (list *doublelist) Remove(i *ListItem) {

	if list.size == 1 {
		list.Clear()
		return
	}

	if i.Value != nil {
		switch {
		case i.Prev == nil && i.Next == nil:
			list.first = nil
			list.first = nil
		case i.Prev == nil:
			i.Next.Prev = nil
			list.first = i.Next
		case i.Next == nil:
			i.Prev.Next = nil
			list.last = i.Prev
		default:
			i.Next.Prev = i.Prev
			i.Prev.Next = i.Next
		}

		list.size--
	}
}

func (list *doublelist) MoveToFron(i *ListItem) {
	list.Remove((i))
	list.PushFront(i.Value)
}

func NewList() List {
	return new(list)
}
