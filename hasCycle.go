package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/linked-list-cycle/?envType=problem-list-v2&envId=oizxjoit
func hasCycle(head *ListNode) bool {
	// brute-force solution
	// visitedNode := make(map[*ListNode]bool)
	// currentNode := head
	// for currentNode != nil {
	// 	if visitedNode[currentNode] {
	// 		return true
	// 	}

	// 	visitedNode[currentNode] = true
	// 	currentNode = currentNode.Next
	// }

	// return false

	// optimal
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
