package main

import "testing"

func TestSingly_Front(t *testing.T) {

	l := NewSingly()
	if l.Front() != nil {
		t.Fatal("The first element must be empty")
	}

}

func TestSingly_Push(t *testing.T) {

	testCase := []int{1, 2, 3, 4, 5}

	l := NewSingly()
	for _, v := range testCase {
		l.Push(v)
	}

	e := l.Front()
	for _, testValue := range testCase {
		if e == nil {
			t.Fatal("The element must not be empty")
		}
		if _, ok := e.Value().(int); !ok {
			t.Fatal("The value of the element must be a int")
		}
		eValue := e.Value().(int)
		if testValue != eValue {
			t.Fatalf("The value of the element must be a %d, but value is %d", testValue, eValue)
		}
		e = e.Next()
	}

}

func TestSingly_Reverse(t *testing.T) {

	testCase := []int{1, 2, 3, 4, 5}
	testReverseCase := []int{5, 4, 3, 2, 1}

	l := NewSingly()
	for _, v := range testCase {
		l.Push(v)
	}

	l.Reverse()

	e := l.Front()
	for _, testValue := range testReverseCase {
		if e == nil {
			t.Fatal("The element must not be empty")
		}
		if _, ok := e.Value().(int); !ok {
			t.Fatal("The value of the element must be a int")
		}
		eValue := e.Value().(int)
		if testValue != eValue {
			t.Fatalf("The value of the element must be a %d, but value is %d", testValue, eValue)
		}
		e = e.Next()
	}

}