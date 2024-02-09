package main

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func New() *LinkedList {
	return &LinkedList{Head: nil, Length: 0}
}
