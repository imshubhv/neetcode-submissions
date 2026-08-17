/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	var maxx int  = 0
	height(root, &maxx)
	return maxx
}

func height(root *TreeNode, maxx *int) int{
	if root == nil{
		return 0
	}
	leftHeight := height(root.Left, maxx)
	rightHeight := height(root.Right, maxx)
	if *maxx < leftHeight+rightHeight {
		*maxx = leftHeight+rightHeight
	}
	return 1+max(leftHeight, rightHeight)
}
