package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=problem-list-v2&envId=oizxjoit
// test case
// testCases := []string{
// 	"abcabcbb",
// 	"bbbbb",
// 	"pwwkew",
// 	"",
// 	"a",
// 	"ab",
// 	"dvdf",
// }
// for _, s := range testCases {
// 	fmt.Println("Input:", s)
// 	fmt.Println("Output:", lengthOfLongestSubstring(s))
// 	fmt.Println("-----")
// }

// Brute Force
// func lengthOfLongestSubstring(s string) int {
// 	maxLen := 0

// 	for i := 0; i < len(s); i++ {
// 		hashTable := make(map[byte]bool)
// 		for j := i; j < len(s); j++ {
// 			if hashTable[s[j]] {
// 				break
// 			}

// 			hashTable[s[j]] = true
// 			currentLen := j - i + 1
// 			if currentLen > maxLen {
// 				maxLen = currentLen
// 			}
// 		}

// 	}

// 	return maxLen
// }

// Sliding window using hash map
// func lengthOfLongestSubstring(s string) int {
// 	maxLen := 0
// 	charIndex := make(map[byte]int)
// 	start := 0

// 	for i := 0; i < len(s); i++ {
// 		if idx, exist := charIndex[s[i]]; exist && idx >= start {
// 			start = idx + 1
// 		}

// 		charIndex[s[i]] = i
// 		currentLen := i - start + 1
// 		if currentLen > maxLen {
// 			maxLen = currentLen
// 		}

// 	}

// 	return maxLen
// }

// Sliding window using set
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	charSet := make(map[byte]bool)
	left := 0

	for right := 0; right < len(s); right++ {
		for charSet[s[right]] {
			delete(charSet, s[left])
			left++
		}

		charSet[s[right]] = true
		currentLen := right - left + 1
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}

	return maxLen
}
