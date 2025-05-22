package main

// brutal
// func longestConsecutive(nums []int) int {
// 	maxLen := 0
// 	for _, num := range nums {
// 		length := 1
// 		currentNum := num
// 		for contains(nums, currentNum+1) {
// 			length++
// 			currentNum++
// 		}

// 		if length > maxLen {
// 			maxLen = length
// 		}
// 	}

// 	return maxLen
// }

// func contains(nums []int, targetNum int) bool {
// 	for _, num := range nums {
// 		if num == targetNum {
// 			return true
// 		}
// 	}

// 	return false
// }

// sorting
// func longestConsecutive(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	sort.Ints(nums)
// 	maxLen, currentLen := 1, 1
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] == nums[i-1] {
// 			continue
// 		} else if nums[i] == nums[i-1]+1 {
// 			currentLen++
// 		} else {
// 			currentLen = 1
// 		}

// 		if currentLen > maxLen {
// 			maxLen = currentLen
// 		}
// 	}

// 	return maxLen
// }

// HashSet
func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLen := 0

	for _, num := range nums {
		if !numSet[num-1] {
			currenNum := num
			currentLen := 1
			for numSet[currenNum+1] {
				currenNum++
				currentLen++
			}

			if currentLen > maxLen {
				maxLen = currentLen
			}
		}
	}
	return maxLen
}
