package linkList2

import "fmt"

type LinkNode struct {
	Data int
	Next *LinkNode
}

// LinkedList 实质上就是 LinkNode 所有用别名，
// 不要重新命名
type LinkedList = LinkNode

func New() *LinkedList {
	// new 申请空间, 这里返回的不是nil 了
	// Next 是nil，Next.value 必然报错
	return new(LinkedList)
}

func (l *LinkedList) IsEmpty() bool {
	return l.Next == nil
}

func (l *LinkedList) Clear() {
	l = New()
}

func (l *LinkedList) Travel() {
	if l.Next == nil {
		return
	}

	n := l.Next
	for n.Next != nil {
		fmt.Printf("%d\t", n.Data)
		n = n.Next
	}
	fmt.Println(n.Data)
}

func (l *LinkedList) Length() int {
	var length int

	//if l.Next == nil{
	//	return length
	//}

	p := l.Next
	for p != nil {
		length++
		p = p.Next
	}

	return length
}

func (l *LinkedList) Append(i int) {
	if l.Next == nil {
		l.Next = &LinkNode{
			Data: i,
		}
		return
	}

	p := l.Next
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &LinkNode{Data: i}
	return
}

func (l *LinkedList) Append2(i int) {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &LinkNode{Data: i}
	return
}
func Array2LL2(arr []int) *LinkedList {
	l := New()

	p := l
	for _, i := range arr {
		p.Next = &LinkNode{Data: i}
		p = p.Next
	}

	return l
}
func Array2LL(arr []int) *LinkedList {
	l := New()

	for _, i := range arr {
		l.Append(i)
	}
	return l
}

func LL2Array(l *LinkedList) []int {
	var a []int
	if l.Next == nil {
		return a
	}

	p := l.Next
	for p != nil {
		a = append(a, p.Data)
		p = p.Next
	}
	return a
}
