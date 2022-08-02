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
	len   int
}

type ListItem struct {
	Value interface{}
	Prev  *ListItem
	Next  *ListItem
}

func NewList() List {
	return new(doublelist)
}

func (list *doublelist) Len() int {
	if list == nil {
		return 0
	}
	return list.len
}

func (list *doublelist) Front() *ListItem {
	return list.first
}

func (list *doublelist) Back() *ListItem {
	return list.last
}

func (list *doublelist) Clear() {
	list.first = nil
	list.last = nil
	list.len = 0
}

func (list *doublelist) PushFront(values interface{}) *ListItem {
	//for _, value := range values {
	newListItem := &ListItem{Value: values, Next: list.first, Prev: nil}
	if list.len == 0 {
		list.first = newListItem
		list.last = newListItem
	} else {
		list.first.Prev = newListItem
		list.first = newListItem
	}
	list.len++
	return list.first
}

func (list *doublelist) PushBack(values interface{}) *ListItem {

	newElement := &ListItem{Value: values, Prev: list.last, Next: nil}
	if list.len == 0 {
		list.first = newElement
		list.last = newElement
	} else {
		list.last.Next = newElement
		list.last = newElement
	}
	list.len++
	return list.last
}

func (list *doublelist) Remove(i *ListItem) {

	if list.len == 1 {
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

		list.len--
	}
}

func (list *doublelist) MoveToFront(i *ListItem) {
	list.Remove((i))
	list.PushFront(i.Value)
}
