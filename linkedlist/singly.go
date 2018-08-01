package main

type Element struct {
	next  *Element
	value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Value() interface{} {
	return e.value
}

type Singly struct {
	root  *Element
}

func (l *Singly) Push(v interface{}) {

	tmp := &Element{value: v}

	if l.root == nil {
		l.root = tmp
	} else {
		e := l.root
		for e.next != nil {
			e = e.next
		}
		e.next = tmp
	}

}

func (l Singly) Front() *Element {
	return l.root
}

func (l *Singly) Reverse() {

	current := l.root
	var prev *Element

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	l.root = prev

}

func NewSingly() *Singly {
	return &Singly{}
}
