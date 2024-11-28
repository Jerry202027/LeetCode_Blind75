package main

// https://leetcode.com/problems/reverse-linked-list/?envType=problem-list-v2&envId=oizxjoit
// test case
// list := createList([]int{1, 2, 3, 4, 5})
// result := reverseList(list)
// fmt.Println(listToArray(result))

// my concepet
// time O(n), space O(n)
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}

// 	var nodeArr []*ListNode
// 	current := head
// 	for current != nil {
// 		nodeArr = append(nodeArr, current)
// 		current = current.Next
// 	}

// 	reverseHead := nodeArr[len(nodeArr)-1]
// 	for i := len(nodeArr) - 1; i >= 0; i-- {
// 		if i == 0 {
// 			nodeArr[i].Next = nil
// 			break
// 		}

// 		nodeArr[i].Next = nodeArr[i-1]
// 	}

// 	return reverseHead
// }

// //////////////////////////////////////////////////////////////////////////////////////////////

// time O(n), space O(1)
// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	current := head

// 	for current != nil {
// 		next := current.Next
// 		current.Next = prev
// 		prev = current
// 		current = next
// 	}

// 	return prev
// }

////////////////////////////////////////////////////////////////////////////////////////////////////

// brute-force solution
// time O(n), space O(1)
// func reverseList(head *ListNode) *ListNode {
// 	var reverseHead *ListNode
// 	for head != nil {
// 		tempNode := &ListNode{Val: head.Val}
// 		tempNode.Next = reverseHead
// 		reverseHead = tempNode
// 		head = head.Next
// 	}

// 	return reverseHead
// }

////////////////////////////////////////////////////////////////////////////////////////////////////

// recursion
// time O(n), space O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reverseHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return reverseHead
}
 