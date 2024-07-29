package main

import (
	"fmt"
	"sort"

	"github.com/golang-collections/collections/queue"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// HDNode stores a tree node along with its horizontal distance (HD) and level (Lvl).
type HDNode struct {
	Node *TreeNode
	HD   int
	Lvl  int
}

// Solution provides methods to perform operations on a binary tree.
type Solution struct{}

// findVertical performs a vertical order traversal of a binary tree.
func (s *Solution) findVertical(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	// Map to store nodes based on vertical and level information
	nodes := make(map[int]map[int][]int)

	// Queue for BFS traversal, storing node along with its vertical and level
	q := queue.New()
	q.Enqueue(HDNode{Node: root, HD: 0, Lvl: 0})

	// BFS traversal
	for q.Len() > 0 {
		hdNode := q.Dequeue().(HDNode)
		node, hd, lvl := hdNode.Node, hdNode.HD, hdNode.Lvl

		// Initialize the map if not already present
		if _, exists := nodes[hd]; !exists {
			nodes[hd] = make(map[int][]int)
		}
		nodes[hd][lvl] = append(nodes[hd][lvl], node.Val)

		// Add left child with updated HD and level
		if node.Left != nil {
			q.Enqueue(HDNode{Node: node.Left, HD: hd - 1, Lvl: lvl + 1})
		}

		// Add right child with updated HD and level
		if node.Right != nil {
			q.Enqueue(HDNode{Node: node.Right, HD: hd + 1, Lvl: lvl + 1})
		}
	}

	// Prepare the final result vector by combining values from the map
	var result [][]int
	var sortedKeys []int
	fmt.Printf("%v", nodes)
	for k := range nodes {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Ints(sortedKeys)
	for _, hd := range sortedKeys {
		var col []int
		var levels []int
		for lvl := range nodes[hd] {
			levels = append(levels, lvl)
		}
		sort.Ints(levels)
		for _, lvl := range levels {
			col = append(col, nodes[hd][lvl]...)
		}
		result = append(result, col)
	}
	return result
}

// printResult prints the result of vertical order traversal.
func printResult(result [][]int) {
	for _, level := range result {
		for _, node := range level {
			fmt.Print(node, " ")
		}
		fmt.Println()
	}
}

func main() {
	// Creating a sample binary tree
	/*
	       1
	      / \
	     2   3
	    / \ / \
	   4  9 10 10
	    \
	     5
	      \
	       6
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 9}
	root.Left.Left.Right = &TreeNode{Val: 5}
	root.Left.Left.Right.Right = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 10}
	root.Right.Left = &TreeNode{Val: 10}

	solution := &Solution{}

	// Get the vertical traversal
	verticalTraversal := solution.findVertical(root)

	// Print the result
	fmt.Println("Vertical Traversal:")
	printResult(verticalTraversal)
}
