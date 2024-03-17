package main

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("popping the stack returns the correct values", func(t *testing.T) {
		s := Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Push(3)

		v, _ := s.Pop()
		if v != 3 {
			t.Fatalf("wanted %d got %d", 3, v)
		}
		v, _ = s.Pop()
		if v != 2 {
			t.Fatalf("wanted %d got %d", 2, v)
		}

		v, _ = s.Pop()
		if v != 1 {
			t.Fatalf("wanted %d got %d", 1, v)
		}

		_, ok := s.Pop()
		if ok {
			t.Fatal("Expected ok to be false")
		}
	})
	t.Run("popping the stack returns the correct values", func(t *testing.T) {
		s := Stack[int]{}
		s.Push(1)
		s.Push(2)
		s.Push(3)

		if !reflect.DeepEqual(s.vals, []int{1, 2, 3}) {
			t.Fatalf("wanted %v got %v", "1,2,3", s.vals)
		}
	})

	t.Run("returns true when value is in stack", func(t *testing.T) {
		s := Stack[string]{}

		s.Push("Hello")
		s.Push("World")

		f := s.Contains("Hello")
		if f != true {
			t.Fatalf("Could not find %s in stack", s)
		}
	})

	t.Run("returns false when value is not in stack", func(t *testing.T) {
		s := Stack[string]{}

		s.Push("Hello")
		s.Push("World")

		v := s.Contains("Not-Here")
		if v == true {
			t.Fatalf("got %v, expeted %v when calling contains", true, v)
		}
	})

}
