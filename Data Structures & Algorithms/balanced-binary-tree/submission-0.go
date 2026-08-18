/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	var result bool = true
	height(root, &result)
	return result
}

func height(root *TreeNode, result *bool) int {
	if root == nil {
		return 0
	}
	left := height(root.Left, result)
	right := height(root.Right, result)

	if abs(left-right) >=2 {
		*result = false
	}
	return 1+ max(left,right)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

