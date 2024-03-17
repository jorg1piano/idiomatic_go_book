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

}
