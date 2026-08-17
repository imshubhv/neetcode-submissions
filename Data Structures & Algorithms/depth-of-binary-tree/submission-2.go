/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil{
        return 0
    }
    if root.Left == nil && root.Right == nil{
        return 1
    }
    var leftDepth, rightDepth int
    if root.Left != nil{
        leftDepth = 1+ maxDepth(root.Left)
    }
    if root.Right != nil{
        rightDepth = 1+ maxDepth(root.Right)
    }
    if leftDepth > rightDepth {
        return leftDepth
    } else {
        return rightDepth
    }
}
