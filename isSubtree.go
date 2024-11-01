package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.com/problems/subtree-of-another-tree/?envType=problem-list-v2&envId=oizxjoit

// test case
// root := &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{5, nil, nil}}
// subRoot := &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}
// fmt.Println(isSubtree(root, subRoot))

// my concept(brute-force solution)
// time O(mn), space O(h), h: the height of the tree
// func isSameTree(a, b *TreeNode) bool {
// 	if a == nil && b == nil {
// 		return true
// 	}

// 	if a == nil || b == nil {
// 		return false
// 	}

// 	return a.Val == b.Val && isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
// }

// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
// 	if root == nil {
// 		return subRoot == nil
// 	}

// 	if isSameTree(root, subRoot) {
// 		return true
// 	}

// 	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// String Matching Solution
// time O(mn), space O(m+n)
// func preorderSerialize(root *TreeNode, builder *strings.Builder) {
// 	if root == nil {
// 		builder.WriteString("#,")
// 		return
// 	}

// 	builder.WriteString(fmt.Sprintf("(%d),", root.Val))
// 	preorderSerialize(root.Left, builder)
// 	preorderSerialize(root.Right, builder)

// }

// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
// 	if root == nil {
// 		return subRoot == nil
// 	}

// 	var rootStr, subStr strings.Builder
// 	preorderSerialize(root, &rootStr)
// 	preorderSerialize(subRoot, &subStr)

// 	return strings.Contains(rootStr.String(), subStr.String())
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// String Matching Solution with KMP Search
// time O(m+n), space O(m+n)
// func preorderSerialize(root *TreeNode, builder *strings.Builder) {
// 	if root == nil {
// 		builder.WriteString("nil, ")
// 		return
// 	}

// 	builder.WriteString(fmt.Sprintf("(%d), ", root.Val))
// 	preorderSerialize(root.Left, builder)
// 	preorderSerialize(root.Right, builder)

// }

// func buildKMPTable(pattern string) []int {
// 	// 例如 pattern = "ABABC"
// 	// table[0] = 0 (一定是0，因為沒有真前綴/後綴)
// 	// table[1] = 0 ("A" 沒有相同的真前綴/後綴)
// 	// table[2] = 0 ("AB" 沒有相同的真前綴/後綴)
// 	// table[3] = 1 ("ABA" 的真前綴/後綴 "A" 長度為1)
// 	// table[4] = 2 ("ABAB" 的真前綴/後綴 "AB" 長度為2)
// 	// table[5] = 0 ("ABABC" 沒有相同的真前綴/後綴)
// 	table := make([]int, len(pattern))
// 	table[0] = 0
// 	length := 0
// 	i := 1
// 	for i < len(pattern) {
// 		if pattern[i] == pattern[length] {
// 			length++
// 			table[i] = length
// 			i++
// 		} else {
// 			if length != 0 {
// 				length = table[length-1]
// 			} else {
// 				table[i] = 0
// 				i++
// 			}
// 		}
// 	}
// 	return table
// }

// func kmpSearch(text, pattern string) bool {
// 	// text = "ABABABCABABAC"
// 	// pattern = "ABABAC"
// 	if len(pattern) == 0 {
// 		return true
// 	}

// 	if len(text) == 0 {
// 		return false
// 	}

// 	table := buildKMPTable(pattern)
// 	j := 0
// 	i := 0

// 	for i < len(text) {
// 		if pattern[j] == text[i] {
// 			i++
// 			j++
// 		}

// 		if j == len(pattern) {
// 			return true
// 		} else if i < len(text) && pattern[j] != text[i] {
// 			if j != 0 {
// 				j = table[j-1]
// 			} else {
// 				i++
// 			}
// 		}
// 	}

// 	return false
// }

// func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
// 	if root == nil {
// 		return subRoot == nil
// 	}

// 	var rootStr, subStr strings.Builder
// 	preorderSerialize(root, &rootStr)
// 	preorderSerialize(subRoot, &subStr)

// 	return kmpSearch(rootStr.String(), subStr.String())
// }

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// Hashing and Serialization
// time O(n * m), space O(m+n)
func serializeAndHash(root *TreeNode) string {
	if root == nil {
		return "#"
	}

	leftHash := serializeAndHash(root.Left)
	rightHash := serializeAndHash(root.Right)

	hash := fmt.Sprintf("%d,%s,%s", root.Val, leftHash, rightHash)

	return hash
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}

	if root == nil {
		return false
	}

	subRootHash := serializeAndHash(subRoot)

	return hasMatchingSubtree(root, subRootHash)
}

func hasMatchingSubtree(root *TreeNode, subRootHash string) bool {

	if serializeAndHash(root) == subRootHash {
		return true
	}

	return hasMatchingSubtree(root.Left, subRootHash) || hasMatchingSubtree(root.Right, subRootHash)
}
