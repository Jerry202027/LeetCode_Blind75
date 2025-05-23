package main

// https://leetcode.com/problems/longest-palindromic-substring/description/?envType=problem-list-v2&envId=oizxjoit
// test case
// testCases := []string{
// 	"babad",
// 	"cbbd",
// 	"a",
// 	"ac",
// 	"abcba",
// 	"aaaa",
// }

// for _, s := range testCases {
// 	fmt.Println("Input:", s)
// 	fmt.Println("Output:", longestPalindrome(s))
// 	fmt.Println("-----")
// }

// Brute Force
// func longestPalindrome(s string) string {
// 	maxLen := 0
// 	result := ""
// 	for i := 0; i < len(s); i++ {
// 		for j := i; j < len(s); j++ {
// 			subStr := s[i : j+1]
// 			if isPalindrome(subStr) {

// 				if len(subStr) > maxLen {
// 					maxLen = len(subStr)
// 					result = subStr
// 				}

// 			}
// 		}
// 	}
// 	return result
// }

// func isPalindrome(s string) bool {
// 	left, right := 0, len(s)-1
// 	for left < right {
// 		if s[left] != s[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		len1 := expandFromCenter(s, i, i)
		len2 := expandFromCenter(s, i, i+1)
		maxLen := max(len1, len2)

		if maxLen > end-start+1 {
			start = i - (maxLen-1)/2
			end = i + (maxLen)/2
		}
	}

	return s[start : end+1]
}

func expandFromCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
