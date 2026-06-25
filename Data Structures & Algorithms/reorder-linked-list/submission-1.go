func reorderList(head *ListNode) {
    if head == nil || head.Next == nil {
        return
    }

    length := findLength(head)

    // correct midpoint
    midNode := head
    for count := 1; count < (length+1)/2; count++ {
        midNode = midNode.Next
    }

    second := reverse(midNode.Next)
    midNode.Next = nil

    first := head

    for first != nil && second != nil {
        firstNext := first.Next
        secondNext := second.Next

        first.Next = second
        second.Next = firstNext

        first = firstNext
        second = secondNext
    }
}

func findLength(head *ListNode) int {
    length := 0
    for head != nil {
        length++
        head = head.Next
    }
    return length
}

func reverse(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    node := reverse(head.Next)
    head.Next.Next = head
    head.Next = nil
    return node
}