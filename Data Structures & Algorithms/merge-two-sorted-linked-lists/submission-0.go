/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
	if list1 == nil{
		return list2
	} else if list2 == nil {
		return list1
	}
	var head *ListNode = nil
	var current *ListNode = nil
	for list1!=nil && list2!= nil {
		if list1.Val >= list2.Val {
			//list2 node will be taken
			if head == nil{
				head = list2
				current = list2
				list2 = list2.Next
			} else {
				current.Next = list2
				current = current.Next
				list2 = list2.Next
			}
		} else {
			//list1 node will be taken
			if head == nil{
				head = list1
				current = list1
				list1 = list1.Next
			}else {
				current.Next = list1
				current = current.Next
				list1 = list1.Next
			}
		}
	}
	if list1 !=nil || list2!= nil {
		if list1 != nil {
			current.Next = list1
		} else {
			current.Next = list2
		}
	}
	return head
}
