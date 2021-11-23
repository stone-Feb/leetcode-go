package main

import "fmt"

//Definition singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	showNode(addTwoNum(l1, l2))
}

func showNode(p *ListNode)  {
	for p != nil {
		fmt.Println(p.Val)
		p = p.Next
	}
}

//给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
func addTwoNum(l1, l2 *ListNode) (head *ListNode) {
	//新建链表，尾插入方法
	var tail *ListNode
	//进位
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		//指针后移
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		//计算进位
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		//首节点
		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
		//最后一个进位
		if carry > 0 {
			tail.Next = &ListNode{Val: carry}
		}
	}
	return
}
