package reorderlist

import (
	"strconv"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (listNodePtr *ListNode) String() string {
	if listNodePtr == nil {
		return ""
	}
	var output strings.Builder
	listNodeCurrentPtr := listNodePtr
	for {
		output.WriteString("[")
		output.WriteString(strconv.Itoa((*listNodeCurrentPtr).Val))
		output.WriteString("]")
		listNodeCurrentPtr = listNodeCurrentPtr.Next
		if listNodeCurrentPtr == nil {
			break
		}
		output.WriteString("->")
	}
	return output.String()
}

func reorderListBySlice(head *ListNode) {
	if head != nil {
		nodesPtrSlice := make([]*ListNode, 0, 1000)
		tempHead := head
		for tempHead != nil {
			nodesPtrSlice = append(nodesPtrSlice, tempHead)
			tempHead = tempHead.Next
		}
		for i := range len(nodesPtrSlice) / 2 {
			nodesPtrSlice[i].Next = nodesPtrSlice[len(nodesPtrSlice)-i-1]
			nodesPtrSlice[len(nodesPtrSlice)-i-1].Next = nodesPtrSlice[i+1]
		}
		nodesPtrSlice[len(nodesPtrSlice)/2].Next = nil
	}
}

func reverseList(headPtr *ListNode) *ListNode {
	var reversedHeadPtr *ListNode = nil
	tempPtr := headPtr
	for tempPtr != nil {
		oldReversedHeadPtr := reversedHeadPtr
		reversedHeadPtr = new(ListNode)
		reversedHeadPtr.Val = tempPtr.Val
		reversedHeadPtr.Next = oldReversedHeadPtr
		tempPtr = tempPtr.Next
	}
	return reversedHeadPtr
}

func mergeTwoLists(firstListPtr *ListNode, secondListPtr *ListNode) *ListNode {
	var head *ListNode = nil
	var isLastElementFromFirst bool
	firstListPtrCopy, secondListPtrCopy := firstListPtr, secondListPtr
	if firstListPtrCopy != nil {
		isLastElementFromFirst = true
		head = firstListPtrCopy
		firstListPtrCopy = firstListPtrCopy.Next
	} else if secondListPtrCopy != nil {
		isLastElementFromFirst = false
		head = secondListPtrCopy
		secondListPtrCopy = secondListPtrCopy.Next
	}
	var lastSelectedElementPtr *ListNode = head
	for firstListPtrCopy != nil || secondListPtrCopy != nil {
		if isLastElementFromFirst && secondListPtrCopy != nil {
			lastSelectedElementPtr.Next = secondListPtrCopy
			isLastElementFromFirst = false
			secondListPtrCopy = secondListPtrCopy.Next
		} else if !isLastElementFromFirst && firstListPtrCopy != nil {
			lastSelectedElementPtr.Next = firstListPtrCopy
			isLastElementFromFirst = true
			firstListPtrCopy = firstListPtrCopy.Next
		} else {
			if firstListPtrCopy != nil {
				lastSelectedElementPtr.Next = firstListPtrCopy
			} else if secondListPtrCopy != nil {
				lastSelectedElementPtr.Next = secondListPtrCopy
			}
			break
		}
		lastSelectedElementPtr = lastSelectedElementPtr.Next
	}
	return head
}

// Bugs
func reorderListByList(head *ListNode) {
	if head == nil {
		return
	}
	beforeSlow, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		beforeSlow = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondList := reverseList(slow)
	beforeSlow.Next = nil
	head = mergeTwoLists(head, secondList)
}

var reorderList func(head *ListNode) = reorderListBySlice
