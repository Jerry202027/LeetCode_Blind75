package main

// https://leetcode.com/problems/group-anagrams/?envType=problem-list-v2&envId=oizxjoit
// test case
// strs := []string{"eat","tea","tan","ate","nat","bat"}
// fmt.Println(groupAnagrams(strs))

func strSort(str string) string {
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if runes[i] > runes[j] {
				runes[i], runes[j] = runes[j], runes[i]
			}
		}
	}
	return string(runes)
}

// brute-force solution
// time O(n * k²), space O(n * k)
// func groupAnagrams(strs []string) [][]string {
// 	var anagramsGroup [][]string
// 	var sortedStrs []string
// 	visited := make(map[int]bool)
// 	for i, str := range strs {
// 		sortedStrs = append(sortedStrs, strSort(str))
// 		visited[i] = false
// 	}

// 	for i := 0; i < len(sortedStrs); i++ {
// 		if visited[i] {
// 			continue
// 		}

// 		anagrams := []string{strs[i]}
// 		visited[i] = true
// 		for j := i + 1; j < len(sortedStrs); j++ {
// 			if sortedStrs[i] == sortedStrs[j] && !visited[j] {
// 				anagrams = append(anagrams, strs[j])
// 				visited[j] = true
// 			}
// 		}

// 		anagramsGroup = append(anagramsGroup, anagrams)
// 	}

// 	return anagramsGroup
// }

// Hash Map with Sorting Keys
// time O(n * k log k), O(n * k)
// func groupAnagrams(strs []string) [][]string {
// 	hashMap := make(map[string][]string)

// 	for _, str := range strs {
// 		sorted := strSort(str)
// 		hashMap[sorted] = append(hashMap[sorted], str)
// 	}

// 	result := [][]string{}
// 	for _, group := range hashMap {
// 		result = append(result, group)
// 	}

// 	return result
// }

// Hash Map with Frequency Keys
// time O(n * k), O(n * k)
func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int][]string)

	for _, str := range strs {
		count := [26]int{}
		for _, char := range str {
			count[char-'a']++
		}
		hashMap[count] = append(hashMap[count], str)
	}

	result := [][]string{}
	for _, group := range hashMap {
		result = append(result, group)
	}

	return result
}
