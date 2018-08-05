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

func (l Singly) Front() Element {
	return *l.root
}

func (l *Singly) Reverse() {

	curr := l.root
	var prev *Element

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.root = prev

}

func (l *Singly) Rotate(r int) {

	if r == 0 {
		return
	}

	curr := l.root

	cnt := 1
	for cnt < r && curr != nil {
		curr = curr.next
		cnt++
	}

	if curr == nil {
		return
	}

	rthElement := curr

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = l.root
	rthElement.next = nil

}

func (l *Singly) Middle() Element {

	fastElement := l.root
	slowElement := l.root

	if l.root != nil {

		fastElement = fastElement.next

		for fastElement != nil && fastElement.next != nil && fastElement.next.next != nil {
			fastElement = fastElement.next.next
			slowElement = slowElement.next
		}

	}

	return *slowElement

}

func NewSingly() *Singly {
	return &Singly{}
}
