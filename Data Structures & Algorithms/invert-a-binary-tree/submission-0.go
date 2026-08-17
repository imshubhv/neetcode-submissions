/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
    if root == nil || (root.Left == nil && root.Right == nil) {
        return root
    }
    var tmpPtr *TreeNode
    tmpPtr = root.Left
    root.Left = root.Right
    root.Right = tmpPtr
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}
