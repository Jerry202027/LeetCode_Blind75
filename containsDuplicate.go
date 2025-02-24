package main

// https://leetcode.com/problems/contains-duplicate/?envType=problem-list-v2&envId=oizxjoit
// test case
// nums :=[]int{1,1,1,3,3,4,3,2,4,2}
// fmt.Println(containsDuplicate(result))

// my concept(hash map, optimal solution)
// time O(n), space O(n)
// func containsDuplicate(nums []int) bool {
// 	hashMap := make(map[int]bool)
// 	for _, num := range nums {
// 		if hashMap[num] {
// 			return true
// 		} else {
// 			hashMap[num] = true
// 		}
// 	}
// 	return false
// }

////////////////////////////////////////////////////////////////////////////////////////////////////

// brute-force solution
// time O(n²), space O(1)
// func containsDuplicate(nums []int) bool {
// 	for i, num := range nums {
// 		for j := i + 1; j < len(nums); j++ {
// 			if num == nums[j] {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// sorting solution
// time O(n log n), space O(log n)
func containsDuplicate(nums []int) bool {
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if num == nums[j] {
				return true
			}
		}
	}

	return false
}


