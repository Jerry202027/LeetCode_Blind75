package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	stairs := 5

	fmt.Println(climbStairs(stairs))

}

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
