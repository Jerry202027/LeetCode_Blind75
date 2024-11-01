package main

// https://leetcode.com/problems/valid-parentheses/?envType=problem-list-v2&envId=oizxjoit

// time O(n^2), space O(n)
// func isValid(s string) bool {
// 	if len(s)%2 != 0 {
// 		return false
// 	}

// 	for {
// 		oldLen := len(s)

// 		s = strings.ReplaceAll(s, "()", "")
// 		s = strings.ReplaceAll(s, "[]", "")
// 		s = strings.ReplaceAll(s, "{}", "")

// 		if len(s) == oldLen {
// 			break
// 		}
// 	}

// 	return len(s) == 0
// }

// time O(n), space O(n)
func isValid(s string) bool {
	stack := []rune{}

	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range s {
		if pair, ok := pairs[c]; ok {
			if len(stack) > 0 && stack[len(stack)-1] == pair {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}
