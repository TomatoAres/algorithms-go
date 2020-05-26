package linklist

func revert(Link *LinkNode) *LinkNode {

	var prev *LinkNode

	for Link != nil {

		tmp := Link.next

		Link.next = prev

		prev = Link
		Link = tmp
	}
	return prev
}
