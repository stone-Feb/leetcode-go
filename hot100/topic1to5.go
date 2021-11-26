package topic1to5

type ListNode struct {
	Val int
	Next *ListNode
}

//1.两数之和
//给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for k, v := range nums {
		if x, ok := hashTable[target-v]; ok {
			return []int{x, k}
		}
		hashTable[v] = k
	}
	return nil
}

//2.两树相加
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

//TODO ：3.无重复字符的最长子串
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int{

	return 0
}