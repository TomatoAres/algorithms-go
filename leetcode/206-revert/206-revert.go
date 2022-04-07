package _06_revert

func revert(Link *LinkNode) *LinkNode {
	if Link == nil || Link.next == nil {
		return Link
	}

	// 即使前指针，也是新链表
	var prev *LinkNode

	for Link != nil {
		// 也是后指针， 暂存后边的链表
		tmp := Link.next

		// 倒置
		Link.next = prev

		// 复制到新链表
		prev = Link

		// 前进：使用之前的存储的后边的链表
		Link = tmp
	}
	return prev
}

// 暂时没看明白
func revertRecursion(link *LinkNode) *LinkNode {
	if link == nil || link.next == nil {
		return link
	}
	newLink := revertRecursion(link.next)
	link.next.next = link
	link.next = nil

	return newLink
}

//func revertRecursion(link *LinkNode) *LinkNode   {
//	first := link.next
//
//	newLink := revertRecursionPart(first)
//
//	link.next = newLink
//
//	return link
//}
