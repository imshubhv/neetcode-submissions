/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	left, right:= head, head.Next
	tmp:= right
	left.Next = nil
	for tmp != nil {
		tmp = right.Next
		right.Next = left
		left= right
		if tmp != nil {
			right = tmp
		}
	}
	return right 
}
