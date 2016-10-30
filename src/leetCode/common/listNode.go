package common

type ListNode struct {
	Val int
	Next *ListNode
}

func MakeArrayToLinkList(arr []int) *ListNode {
	l := &ListNode{}
	p := l
	for _, a := range arr {
		p.Next = &ListNode{Val: a}
		p = p.Next
	}
	return l.Next
}