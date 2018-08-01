package main

import "testing"

func TestSingly_Front(t *testing.T) {

	l := NewSingly()
	if l.Front() != nil {
		t.Fatal("The first element must be empty")
	}

}

func TestSingly_Push(t *testing.T) {

	l := NewSingly()
	for i := 0; i < 100; i++ {
		l.Push(i)
	}

	e := l.Front()
	for i := 0; i < 100; i++ {
		if e == nil {
			t.Fatal("The element must not be empty")
		}
		if _, ok := e.Value().(int); !ok {
			t.Fatal("The value of the element must be a int")
		}
		eValue := e.Value().(int)
		if i != eValue {
			t.Fatalf("The value of the element must be a %d, but value is %d", i, eValue)
		}
		e = e.Next()
	}

}

func TestSingly_Reverse(t *testing.T) {

	l := NewSingly()
	for i := 0; i < 100; i++ {
		l.Push(i)
	}

	l.Reverse()

	e := l.Front()
	for i := 99; i >= 0; i-- {
		if e == nil {
			t.Fatal("The element must not be empty")
		}
		if _, ok := e.Value().(int); !ok {
			t.Fatal("The value of the element must be a int")
		}
		eValue := e.Value().(int)
		if i != eValue {
			t.Fatalf("The value of the element must be a %d, but value is %d", i, eValue)
		}
		e = e.Next()
	}

}