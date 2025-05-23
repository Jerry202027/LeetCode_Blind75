package main

// https://leetcode.com/problems/container-with-most-water/description/?envType=problem-list-v2&envId=oizxjoit
// test case
// testCases := [][]int{
// 	{1, 8, 6, 2, 5, 4, 8, 3, 7}, // expected: 49
// 	{1, 1},                      // expected: 1
// 	{4, 3, 2, 1, 4},             // expected: 16
// 	{1, 2, 1},                   // expected: 2
// 	{2, 3, 10, 5, 7, 8, 9},      // expected: 36
// 	{1, 2, 4, 3},                // expected: 4
// }

// for _, heights := range testCases {
// 	fmt.Println("Input:  ", heights)
// 	fmt.Println("Output: ", maxArea(heights))
// 	fmt.Println("-----")
// }

// Brute Force
// func maxArea(height []int) int {
// 	maxArea := 0
// 	for i := 0; i < len(height); i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			minHeight := min(height[i], height[j])
// 			width := j - i
// 			currentArea := minHeight * width

// 			if currentArea > maxArea {
// 				maxArea = currentArea
// 			}
// 		}
// 	}
// 	return maxArea
// }

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for right-left > 0 {
		minHeight := min(height[left], height[right])
		width := right - left
		currentArea := minHeight * width
		if currentArea > maxArea {
			maxArea = currentArea
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
