package main

// https://leetcode.com/problems/clone-graph/?envType=problem-list-v2&envId=oizxjoit
// test case
// original := buildTestGraph()
// fmt.Println("== Original Graph ==")
// printGraph(original)

// cloned := cloneGraph(original)
// fmt.Println("== Cloned Graph ==")
// printGraph(cloned)

// DFS
// func cloneGraph(node *Node) *Node {
// 	if node == nil {
// 		return nil
// 	}

// 	visitedMap := make(map[*Node]*Node)
// 	clone := dfsClone(node, visitedMap)
// 	return clone
// }

// func dfsClone(node *Node, visitedMap map[*Node]*Node) *Node {
// 	if n, exist := visitedMap[node]; exist {
// 		return n
// 	}

// 	clone := &Node{Val: node.Val}
// 	visitedMap[node] = clone
// 	for _, neighborNode := range node.Neighbors {
// 		clone.Neighbors = append(clone.Neighbors, dfsClone(neighborNode, visitedMap))
// 	}

// 	return clone
// }

// BFS with Queue Clone
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	cloneMap := map[*Node]*Node{}
	queue := []*Node{node}
	cloneMap[node] = &Node{Val: node.Val}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]

		for _, neighborNode := range currentNode.Neighbors {
			if _, exist := cloneMap[neighborNode]; !exist {
				cloneMap[neighborNode] = &Node{Val: neighborNode.Val}
				queue = append(queue, neighborNode)
			}

			cloneMap[currentNode].Neighbors = append(cloneMap[currentNode].Neighbors, cloneMap[neighborNode])
		}

	}

	return cloneMap[node]
}
