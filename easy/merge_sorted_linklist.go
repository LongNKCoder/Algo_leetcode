package main

import "fmt"

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	initNode := &ListNode{}
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		initNode.Val = list2.Val
		initNode.Next = list2.Next
		return initNode
	}
	if list2 == nil {
		initNode.Val = list1.Val
		initNode.Next = list1.Next
		return initNode
	}
	if list1.Val < list2.Val {
		initNode.Val = list1.Val
		list1 = list1.Next
	} else {
		initNode.Val = list2.Val
		list2 = list2.Next
	}
	nextNode := &ListNode{}
	initNode.Next = nextNode
	for {
		if list1 == nil {
			nextNode.Val = list2.Val
			nextNode.Next = list2.Next
			break
		}
		if list2 == nil {
			nextNode.Val = list1.Val
			nextNode.Next = list1.Next
			break
		}
		if list1.Val < list2.Val {
			nextNode.Val = list1.Val
			list1 = list1.Next
		} else {
			nextNode.Val = list2.Val
			list2 = list2.Next
		}
		nextNode.Next = &ListNode{}
		nextNode = nextNode.Next
	}
	return initNode
}

func main() {
	//list1 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val: 8,
	//			Next: &ListNode{
	//				Val: 9,
	//			},
	//		},
	//	},
	//}
	//list2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 2,
	//		Next: &ListNode{
	//			Val: 7,
	//			Next: &ListNode{
	//				Val: 17,
	//			},
	//		},
	//	},
	//}
	mergeList := mergeTwoLists(nil, &ListNode{Val: 0})
	fmt.Println(mergeList.Val)
	for mergeList.Next != nil {
		fmt.Println(mergeList.Next.Val)
		mergeList = mergeList.Next
	}
}
