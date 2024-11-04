package main

// https://leetcode.com/problems/reverse-bits/?envType=problem-list-v2&envId=oizxjoit

// test case
// stairs := 5
// fmt.Println(climbStairs(stairs))

// my concept
// dynamic programing
// time O(n), space O(n)
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}

// 	lookupTable := make([]int, n+1)
// 	lookupTable[1] = 1
// 	lookupTable[2] = 2

// 	for i := 3; i <= n; i++ {
// 		lookupTable[i] = lookupTable[i-1] + lookupTable[i-2]
// 	}

// 	return lookupTable[n]
// }

// ///////////////////////////////////////////////////////////////////////////////////////////////
// brute-force solution
// time O(n), space O(n) (storage for func, because using recursion)
// func climbStairs(n int) int {
// 	if n == 1 {
// 		return 1
// 	}

// 	if n == 2 {
// 		return 2
// 	}

// 	return climbStairs(n-1) + climbStairs(n-2)
// }

// /////////////////////////////////////////////////////////////////////////////////////////////////
// dynamic programing with rolling array -> save storage
// time O(n), space O(1)
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
 
	return b
}
