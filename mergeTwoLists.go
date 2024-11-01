package main

// https://leetcode.com/problems/merge-two-sorted-lists/?envType=problem-list-v2&envId=oizxjoit

// aux func for testing
//
// list1 := createList([]int{1, 2, 4})
// list2 := createList([]int{1, 3, 4})
// result := mergeTwoLists(list1, list2)
// fmt.Println("Test 1:", listToArray(result))
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	return head
}

func listToArray(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// my concept
// time O(n), space O(1)
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}

// 	if list2 == nil {
// 		return list1
// 	}

// 	head := list1
// 	tempNode := list2

// 	if head.Val > tempNode.Val {
// 		head = list2
// 		tempNode = list1
// 	}

// 	current := head

// 	for current != nil && tempNode != nil {
// 		if current.Next == nil {
// 			current.Next = tempNode
// 			break
// 		}

// 		oldNext := current.Next
// 		if current.Next.Val > tempNode.Val {
// 			current.Next = tempNode
// 			tempNode = oldNext
// 		}

// 		current = current.Next
// 	}

// 	return head
// }

// brute-force solution
// time O(nlogn), space O(n)
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	var values []int
// 	for list1 != nil {
// 		values = append(values, list1.Val)
// 		list1 = list1.Next
// 	}

// 	for list2 != nil {
// 		values = append(values, list2.Val)
// 		list2 = list2.Next
// 	}

// 	sort.Ints(values)

// 	if len(values) == 0 {
// 		return nil
// 	}

// 	head := &ListNode{Val: values[0]}
// 	current := head
// 	for i := 1; i < len(values); i++ {
// 		current.Next = &ListNode{Val: values[i]}
// 		current = current.Next
// 	}

// 	return head
// }

// recursion
// time O(n), space O(n)
// func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
// 	if list1 == nil {
// 		return list2
// 	}

// 	if list2 == nil {
// 		return list1
// 	}

// 	if list1.Val <= list2.Val {
// 		list1.Next = mergeTwoLists(list1.Next, list2)
// 		return list1
// 	} else {
// 		list2.Next = mergeTwoLists(list1, list2.Next)
// 		return list2
// 	}
// }

// optimal
// time O(n), space O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	}

	if list2 != nil {
		current.Next = list2
	}

	return dummy.Next
}
