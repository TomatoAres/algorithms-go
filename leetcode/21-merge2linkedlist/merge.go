package _1_merge2linkedlist

import (
	"github.com/TomatoAres/algorithms-go/leetcode/linklist2"
)

type LinkedList = linkList2.LinkedList
type LinkedNode = linkList2.LinkNode

func Merge(l1, l2 *LinkedList) LinkedList {
	//if l1.Next == nil {
	//	return l2
	//}
	//if l2.Next == nil {
	//	return l1
	//}
	//p1 := l1.Next
	//p2 := l2.Next

	// result head
	//result := linkList2.New()
	//p := &linkList2.LinkNode{}
	//
	//for p1 != nil && p2 != nil {
	//	if p1.Data > p2.Data {
	//		p.Next = p2
	//		p2 = p2.Next
	//	} else {
	//		p.Next = p1
	//		p1 = p1.Next
	//	}
	//	p = p.Next
	//}
	//if p1 != nil {
	//	p.Next = p1
	//}
	//if p2 != nil {
	//	p.Next = p2
	//}
	//result.Next = p
	//return result
	return LinkedNode{}
}

func Merge2(l1 *LinkedNode, l2 *LinkedNode) *LinkedNode {
	head := &LinkedNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Data < l2.Data {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}

func MergeTwoLists(l1 *LinkedList, l2 *LinkedList) *LinkedList {

	if l1.Next == nil {
		return l2
	}
	if l2.Next == nil {
		return l1
	}
	p1 := l1.Next
	p2 := l2.Next
	if p1.Data < p2.Data {
		p1.Next = MergeTwoLists(p1.Next, p2)
		return l1
	}
	p2.Next = MergeTwoLists(p1, p2.Next)
	return l2
}
