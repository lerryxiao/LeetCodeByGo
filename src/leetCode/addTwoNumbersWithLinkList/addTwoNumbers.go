package AddTowNumbersLinkList

import (
	. "leetCode/common"
)


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    h := &ListNode{}
	p := h
	carry := 0
	r := 0
	for l1 != nil || l2 != nil  {
		r = 0
		if l1 != nil && l2 != nil {
			r = l1.Val + l2.Val
			p.Next = &ListNode{Val: (r + carry) % 10}
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil {
			r = l1.Val
			p.Next = &ListNode{Val: (r + carry) %10}
			l1 = l1.Next
		}else if l2 != nil {
			r = l2.Val
			p.Next = &ListNode{Val: (r + carry) % 10}
			l2 = l2.Next
		}
		p = p.Next
		carry = (r + carry) / 10
	}
	if carry == 1 {
		p.Next = &ListNode{Val : carry}
	}
	return h.Next
}