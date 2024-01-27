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

type ListItem struct {
	Key   Key
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.head
}

func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := ListItem{Value: v}
	interim := l.head
	newItem.Next = interim
	l.head = &newItem
	if l.len == 0 {
		l.tail = l.head
	} else {
		interim.Prev = &newItem
	}
	l.len++
	return &newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := ListItem{Value: v}
	interim := l.tail
	newItem.Prev = interim
	l.tail = &newItem
	if l.len == 0 {
		l.tail = l.head
	} else {
		interim.Next = &newItem
	}
	l.len++
	return &newItem
}

func (l *list) Remove(i *ListItem) {
	prev := i.Prev
	next := i.Next
	if prev != nil {
		prev.Next = next
	} else {
		l.head = next
	}
	if next != nil {
		next.Prev = prev
	} else {
		l.tail = prev
	}
	l.len--
	i.Prev = nil
	i.Next = nil
}

func (l *list) MoveToFront(i *ListItem) {
	switch {
	case i.Next == nil:
		i.Prev.Next = nil
		interim := l.head
		l.tail = interim
	case i.Prev == nil:
		return
	}
	interim := l.head
	l.head = i
	l.head.Next = interim
	l.head.Prev = nil
	interim.Prev = l.head
}

func NewList() List {
	return new(list)
}
