package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func BuildTree2(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 || len(inorder) == 0 {
        return nil
    }
    
    tn := &TreeNode{}
    currR := postorder[len(postorder)- 1]
    tn.Val = currR
    Rindex := 0

    for i, r := range inorder {
        if r == currR {
            Rindex = i
        }
    }

    lst, rst := inorder[: Rindex], inorder[Rindex+1 :]
    lpr, rpr := postorder[: len(lst)], postorder[len(lst) : len(postorder)-1]

    tn.Left = BuildTree2(lst, lpr)
    tn.Right = BuildTree2(rst, rpr)

    return tn
}

// 4ms
// 6.02mb