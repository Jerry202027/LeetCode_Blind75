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
