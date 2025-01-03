package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func RotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }
    n := 1
    curr := head
    for curr.Next != nil {
        curr = curr.Next
        n++
    }
    k = k % n
    if k == 0 {
        return head
    }
    newTailPosi := n-k-1
    newTail := head
    for i := 0; i < newTailPosi; i++ {
        newTail = newTail.Next
    }
    newHead := newTail.Next
    newTail.Next = nil 
    curr.Next = head
    return newHead
}

// 0ms
// 4.44mb