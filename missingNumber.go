package main

//https://leetcode.com/problems/missing-number/?envType=problem-list-v2&envId=oizxjoit

// time O(n^2), space O(1)
// func missingNumber(nums []int) int {
// 	isFounded := false
// 	for i := 0; i <= len(nums); i++ {
// 		for _, n := range nums {
// 			if i == n {
// 				isFounded = true
// 				break
// 			}
// 		}

// 		if !isFounded {
// 			return i
// 		}
// 	}

// 	return -1
// }

// time O(n), space O(n)
// func missingNumber(nums []int) int {
// 	hashMap := make(map[int]bool)

// 	for _, n := range nums {
// 		hashMap[n] = true
// 	}

// 	for i := 0; i <= len(nums); i++ {
// 		if !hashMap[i] {
// 			return i
// 		}
// 	}
// 	return -1
// }

// time O(n), space O(1)
func missingNumber(nums []int) int {
	expectedSum := len(nums) * (len(nums) + 1) / 2
	sum := 0
	for _, n := range nums {
		sum += n
	}

	return expectedSum - sum
}
