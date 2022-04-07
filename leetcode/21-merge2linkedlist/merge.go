package _1_merge2linkedlist

import (
	"github.com/TomatoAres/algorithms-go/leetcode/linklist2"
)

func merge(l1, l2 *linkList2.LinkedList) *linkList2.LinkedList {

	if l1.Next == nil {
		return l2
	}
	if l2.Next == nil {
		return l1
	}
	p1 := l1.Next
	p2 := l2.Next

	// result head
	result := linkList2.New()
	p := &linkList2.LinkNode{}

	for p1 != nil && p2 != nil {
		if p1.Data > p2.Data {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	result.Next = p
	return result

}
