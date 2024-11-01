package main

// https://leetcode.com/problems/two-sum/?envType=problem-list-v2&envId=oizxjoit
//
//	input := []int{3, 2, 4}
//	target := 6
//
// fmt.Println(twoSum(input, target))
func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, num := range nums {
		otherNum := target - num

		if index, ok := hashMap[otherNum]; ok {
			return []int{index, i}
		}

		hashMap[num] = i
	}

	return nil
}
