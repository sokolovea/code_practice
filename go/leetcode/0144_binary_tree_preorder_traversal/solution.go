package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (treeNode *TreeNode) add(val int) *TreeNode {
	if treeNode == nil {
		treeNode = new(TreeNode)
		treeNode.Val = val
		return treeNode
	}
	treeNodeCopy := treeNode
	for {
		if val < treeNode.Val {
			if treeNode.Left == nil {
				treeNode.Left = new(TreeNode)
				treeNode.Left.Val = val
				break
			} else {
				treeNode = treeNode.Left
			}
		} else if val > treeNode.Val {
			if treeNode.Right == nil {
				treeNode.Right = new(TreeNode)
				treeNode.Right.Val = val
				break
			} else {
				treeNode = treeNode.Right
			}
		} else {
			break
		}
	}
	return treeNodeCopy
}

func preorderTraversalInternal(rootPtr *TreeNode, resultSlice *[]int) {
	if rootPtr == nil {
		return
	}
	preorderTraversalInternal(rootPtr.Left, resultSlice)
	*resultSlice = append(*resultSlice, (*rootPtr).Val)
	preorderTraversalInternal(rootPtr.Right, resultSlice)
}

func preorderTraversal(root *TreeNode) []int {
	var resultSlice []int = make([]int, 0, 5)
	preorderTraversalInternal(root, &resultSlice)
	return resultSlice
}
