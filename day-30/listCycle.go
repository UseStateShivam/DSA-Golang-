package main

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

 func HasCycle(head *ListNode) bool {
    slow := head
    fast := head
    for slow != nil && slow.Next != nil && fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    return false
}

// 0ms
// 6.15mb