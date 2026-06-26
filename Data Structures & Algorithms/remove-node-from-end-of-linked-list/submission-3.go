/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head.Next == nil{
		return nil
	}
	//find length
	tracker:= head
	count:=0
	for tracker!= nil {
		count++
		tracker = tracker.Next
	}
	tracker = head
	prevNode := head
	if count-n == 0 {
		head = head.Next
		return head
	}
	for i:=0 ; i<=count-n-1; i++ {
		prevNode = tracker
		tracker = tracker.Next
	}
	prevNode.Next = tracker.Next
	return head
}
