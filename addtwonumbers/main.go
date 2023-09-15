package addtwonumbers

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Run() {
	fmt.Println("add two numbers")

	l1 := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	res := addTwoNumbers(&l1, &l2)
	fmt.Println(res)
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	ln := &ListNode{}
	for n, sum := ln, 0; l1 != nil || l2 != nil || sum > 0; n = n.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		n.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		sum /= 10
	}

	return ln.Next
}
