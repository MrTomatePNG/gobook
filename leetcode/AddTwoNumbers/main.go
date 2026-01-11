package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l1.Next
		}
		sum := a + b + carry
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
	}
	return dummy.Next
}

// helper

func build(nums []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Print(n.Val)
		if n.Next != nil {
			fmt.Print("->")
		}
		n = n.Next
	}

	fmt.Println()
}

func main() {
	l1 := build([]int{2, 4, 3})
	l2 := build([]int{5, 6, 4})
	res := addTwoNumbers(l1, l2)
	printList(res) // 7 -> 0 -> 8
}
