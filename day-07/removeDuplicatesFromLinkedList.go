package main

// https://leetcode.com/u/gimangi/ credits to this guy for the solution cuase I definitly gave up

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	for curr := head; curr != nil; curr = curr.Next {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		}
		// In the worst case, the inner loop skips duplicates, but each node is still only visited once across the entire list, ensuring that the total number of operations for both loops is still proportional to the number of nodes (O(n)).
	}
	return head
}

// 0ms
// 4.88mb