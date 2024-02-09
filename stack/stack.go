package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
}

type Stack struct {
	Top    *Node
	Length int
}

func (s *Stack) Len() int {
	return s.Length
}

func New() *Stack {
	return &Stack{
		Top:    nil,
		Length: 0,
	}
}

func (s *Stack) Push(val int) {
	n := &Node{Value: val, Prev: s.Top}
	s.Top = n
	s.Length++
}

func (s *Stack) Pop() int {
	if s.Length == 0 {
		fmt.Println("No element to pop")
		return 0
	}
	n := s.Top
	s.Top = n.Prev
	s.Length--
	return n.Value
}

func (s *Stack) Peek() int {
	if s.Length == 0 {
		fmt.Println("Stack is empty")
		return 0
	}
	return s.Top.Value
}

func main() {
	s := New()
	s.Push(10)
	fmt.Println(s.Peek())
	s.Push(30)
	fmt.Println(s.Peek())
	s.Pop()
	s.Pop()
	fmt.Println(s.Peek())
}
