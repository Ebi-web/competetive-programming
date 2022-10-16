package main

import (
	"fmt"
	"math"
)

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1 := toSlice(l1)
	s2 := toSlice(l2)
	//	s1 and s2 to int
	s1_int := 0
	for i := 0; i < len(*s1); i++ {
		s1_int += (*s1)[i] * int(math.Pow10(i))
		fmt.Print((*s1)[i] * int(math.Pow10(i)))
	}
	s2_int := 0
	for i := 0; i < len(*s2); i++ {
		s2_int += (*s2)[i] * int(math.Pow10(i))
	}
	sum := s1_int + s2_int
	//	sum to ListNode
	return newListNode(sum)
}

func newListNode(s int) *ListNode {
	slice := make([]int, 0)
	for s != 0 {
		slice = append(slice, s%10)
		s /= 10
	}
	//	slice to ListNode
	l := &ListNode{}
	currentNode := l
	for i := 0; i < len(slice); i++ {
		currentNode.Val = slice[i]
		if i == len(slice)-1 {
			break
		}
		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}
	return l
}

func toSlice(l *ListNode) *[]int {
	s := make([]int, 0)
	currentNode := l
	for {
		s = append(s, currentNode.Val)
		if currentNode.Next == nil {
			break
		}
		currentNode = currentNode.Next
	}
	return &s
}
