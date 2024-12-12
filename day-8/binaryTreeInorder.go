package main

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) []int {
	// approach
	// root is given
	// root.Left.Left... until Val == nil is our first element
	// then append root then right unitl the parent root is found
	// then move to right half
	// this can be implemented using stack
	// lets visit stack again

	stack := []*TreeNode{}
	curr := root
	result := []int{}

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, curr.Val)

		curr = curr.Right
	}

	return result
}

// 0ms
// 4.10mb
// this took hell lotta time :(