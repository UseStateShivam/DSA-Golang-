package main

type TreeNode struct {
	Val   int
	Left *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    tn := &TreeNode{}
    currR := preorder[0]
    tn.Val = currR
    Rindex := 0
    for i, r := range inorder {
        if r == currR {
            Rindex = i  
            break
        } 
    }
    lst, rst := inorder[: Rindex], inorder[Rindex + 1 :]
    lpr, rpr := preorder[1 : len(lst)+1], preorder[1+len(lst):]
    tn.Left = BuildTree(lpr, lst)
    tn.Right = BuildTree(rpr, rst)
    return tn
}

// 0ms
// 5.92mb