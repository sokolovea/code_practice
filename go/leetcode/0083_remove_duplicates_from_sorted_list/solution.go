package removeduplicatesfromsortedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	nodeValSet := make(map[int]bool)
	root := head
	prev_head := head
	for head != nil {
		isInSet := nodeValSet[head.Val]
		if isInSet {
			prev_head.Next = head.Next
		} else {
			nodeValSet[head.Val] = true
			prev_head = head
		}
		head = head.Next
	}
	return root
}
