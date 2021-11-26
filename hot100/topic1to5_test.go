package topic1to5

import (
	"testing"
)

func TestAddTwoNum(t *testing.T){
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
	addTwoNum(l1, l2)
}

func TestLengthOfLongestSubstring(t *testing.T)  {
	s := "abcabcbb"
	lengthOfLongestSubstring(s)
}