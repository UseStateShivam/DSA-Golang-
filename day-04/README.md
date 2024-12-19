# Merge Two Sorted Lists

## Key learnings

- Node represents a single node in a LinkedList:
``` Go
type Node struct {
    value int
    newNode *Node
}
```
- A linked list is represented bye:
``` Go
type LinkedList struct {
    head *Node
}
```
- A linked list is like a bunch of nodes pointing to the following nodes.
``` plaintext
LinkedList:
    1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9
```
- A singly linked list is where the nodes are only linked in one direction. But in a doubly linked list, nodes can have bidirectional linkings.
``` plaintext
singly linked list:
    1 -> 2 -> 3 -> 4 -> 5 
doubly linked list:
    1 <-> 2 <-> 3 <-> 4 (This means doubly linked lists can be expressed as forward singly or backward singly linked as desired)
```
- These kinds of situations can be tackled useing the greedy approach. At each step, the algorithm can confidently choose the smaller of the two current elements, knowing that this choice will preserve the overall order in the merged list.
- Shared Pointer: outputNode and current initially point to the same memory location, i.e., the same ListNode. This means they share the same underlying linked list structure at the start.
``` lua
+----+    +----+    +----+    +----+
| 10 | -> | 20 | -> | 30 | -> | 40 |
+----+    +----+    +----+    +----+
  ^                              |
  |                              v
(head)                       (last node)
```
- The outputNode and current, although are both initialized with the same pointer to the memory location, the current pointer travles from the head to the last node while populating the result list, however the outputNode stays at the head itself. Thus returning the head node makes sense cause this returns the access to the complete linked list. And returning current does not make sense.

# Remove Duplicates from Sorted Array

## Key learnings

- To concatenate slices in Go:
``` Go
slice = append(slice[:index], slice[index + 1:]...) // omitted the index from the slice
```

# Remove Element

## Key learnings

- Time complexity of append func is: 
``` lua
O(1) ---> Best case, no resizing needed
O(n) ---> Worst case, resizing needed
```