package lc

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	extra := 0
	return newList(l1, l2, extra)
}

func newList(l1 *ListNode, l2 *ListNode, extra int) *ListNode {
	if l1 == nil && l2 == nil {
		if extra > 0 {
			return &ListNode{
				Val:  extra,
				Next: nil,
			}
		}
		return nil
	}

	v := 0
	next1 := new(ListNode)
	next2 := new(ListNode)
	if l1 == nil {
		v = l2.Val + extra
		next1 = nil
		next2 = l2.Next
	} else if l2 == nil {
		v = l1.Val + extra
		next1 = l1.Next
		next2 = nil
	} else {
		v = l1.Val + l2.Val + extra
		next1 = l1.Next
		next2 = l2.Next
	}

	if v > 9 {
		v = v - 10
		extra = 1
	} else {
		extra = 0
	}
	return &ListNode{Val: v, Next: newList(next1, next2, extra)}
}
