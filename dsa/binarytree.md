###  Binary Tree Inorder Traversal

**Question:**

Given the root of a binary tree, return the inorder traversal of its nodes' values.

**Example 1:**
```
Input: root = [1, null, 2, 3]
Output: [1, 3, 2]
```

**Solution:**
```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	s := stack.New()
	current := root
	for current != nil || s.Len() > 0 {
		for current != nil {
			s.Push(current)
			current = current.Left
		}
		current = s.Pop().(*TreeNode)
		result = append(result, current.Val)
		current = current.Right
	}

	return result
}
```

###  Binary Tree Preorder Traversal

**Question:**

Given the root of a binary tree, return the preorder traversal of its nodes' values.

**Example 1:**
```
Input: root = [1, null, 2, 3]
Output: [1, 2, 3]
```

**Solution:**
```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	s := stack.New()
	s.Push(root)

	for s.Len() > 0 {
		node := s.Pop().(*TreeNode)
		result = append(result, node.Val)

		if node.Right != nil {
			s.Push(node.Right)
		}
		if node.Left != nil {
			s.Push(node.Left)
		}
	}

	return result
}
```

###  Binary Tree Postorder Traversal

**Question:**

Given the root of a binary tree, return the postorder traversal of its nodes' values.

**Example 1:**
```
Input: root = [1, null, 2, 3]
Output: [3, 2, 1]
```

**Example 2:**
```
Input: root = []
Output: []
```

**Example 3:**
```
Input: root = [1]
Output: [1]
```

**Solution:**
```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	stack1 := stack.New()
	stack2 := stack.New()
	stack1.Push(root)

	for stack1.Len() > 0 {
		node := stack1.Pop().(*TreeNode)
		stack2.Push(node)

		if node.Left != nil {
			stack1.Push(node.Left)
		}
		if node.Right != nil {
			stack1.Push(node.Right)
		}
	}

	for stack2.Len() > 0 {
		node := stack2.Pop().(*TreeNode)
		result = append(result, node.Val)
	}

	return result
}
```

###  Binary Tree Level Order Traversal

**Question:**

Given the root of a binary tree, return the level order traversal of its nodes' values (i.e., from left to right, level by level).

**Example 1:**
```
Input: root = [3, 9, 20, null, null, 15, 7]
Output: [[3], [9, 20], [15, 7]]
```

**Solution:**
```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderTraversal(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	q := queue.New()
	q.Enqueue(root)

	for q.Len() > 0 {
		levelSize := q.Len()
		var level []int
		for i := 0; i < levelSize; i++ {
			node := q.Dequeue().(*TreeNode)
			level = append(level, node.Val)
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		result = append(result, level)
	}

	return result
}
```


Certainly! Let's go through the problem of finding the **left view** and **right view** of a binary tree, along with example explanations.

### Left view and right view of binary tree

**Left View**: The left view of a binary tree contains the nodes that are visible when the tree is viewed from the left side. For each level of the tree, the first node encountered from the left is part of the left view.

**Right View**: The right view of a binary tree contains the nodes that are visible when the tree is viewed from the right side. For each level of the tree, the last node encountered from the right is part of the right view.

#### Examples

**Example 1**:

- **Input**:
  ```
        1
       / \
      2   3
     /   / \
    4   5   6
  ```
- **Left View Output**: [1, 2, 4]
  - **Explanation**: The nodes visible from the left side are 1 (root), 2 (first node on the second level), and 4 (first node on the third level).

- **Right View Output**: [1, 3, 6]
  - **Explanation**: The nodes visible from the right side are 1 (root), 3 (last node on the second level), and 6 (last node on the third level).

**Example 2**:

- **Input**:
  ```
      1
     /
    2
   /
  3
  ```
- **Left View Output**: [1, 2, 3]
  - **Explanation**: The nodes visible from the left side are 1 (root), 2 (first node on the second level), and 3 (first node on the third level).

- **Right View Output**: [1, 2, 3]
  - **Explanation**: The nodes visible from the right side are the same as the left view since there are no right children.

### Implementation Using Queue (BFS)

Here’s how to implement the left and right views using the `github.com/golang-collections/collections/queue` library for queue management:

```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leftView returns the left view of the binary tree.
func leftView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	q := queue.New()
	q.Enqueue(root)

	for q.Len() > 0 {
		levelSize := q.Len()
		for i := 0; i < levelSize; i++ {
			node := q.Dequeue().(*TreeNode)

			// If it's the first node of this level
			if i == 0 {
				result = append(result, node.Val)
			}

			// Add left and right children to the queue
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
	}
	return result
}

// rightView returns the right view of the binary tree.
func rightView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	q := queue.New()
	q.Enqueue(root)

	for q.Len() > 0 {
		levelSize := q.Len()
		for i := 0; i < levelSize; i++ {
			node := q.Dequeue().(*TreeNode)

			// If it's the last node of this level
			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			// Add left and right children to the queue
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
	}
	return result
}
```

To also include the **Bottom View** of a binary tree, we can use a similar approach as the Top View. The Bottom View consists of the nodes visible when the tree is viewed from the bottom. For each horizontal distance (HD), the last node encountered during a level-order traversal is part of the bottom view.



### topview bottom view of binary tree

```go

package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// HDNode stores a tree node along with its horizontal distance (HD).
type HDNode struct {
	Node *TreeNode
	HD   int
}

// topView prints the top view of the binary tree.
func topView(root *TreeNode) {
	if root == nil {
		return
	}

	// Queue to perform level order traversal
	q := queue.New()
	q.Enqueue(HDNode{Node: root, HD: 0})

	// Map to store the first node at each horizontal distance for the top view
	topViewMap := make(map[int]int)

	// Track the range of HDs for proper ordering
	minHD, maxHD := 0, 0

	for q.Len() > 0 {
		hdNode := q.Dequeue().(HDNode)
		node, hd := hdNode.Node, hdNode.HD

		// If HD is seen for the first time, store the node's value for the top view
		if _, found := topViewMap[hd]; !found {
			topViewMap[hd] = node.Val
		}

		// Update the range of HDs
		if hd < minHD {
			minHD = hd
		}
		if hd > maxHD {
			maxHD = hd
		}

		// Enqueue left and right children with updated HDs
		if node.Left != nil {
			q.Enqueue(HDNode{Node: node.Left, HD: hd - 1})
		}
		if node.Right != nil {
			q.Enqueue(HDNode{Node: node.Right, HD: hd + 1})
		}
	}

	// Print the top view from leftmost to rightmost HD
	fmt.Print("Top View: ")
	for hd := minHD; hd <= maxHD; hd++ {
		if val, found := topViewMap[hd]; found {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
}

// bottomView prints the bottom view of the binary tree.
func bottomView(root *TreeNode) {
	if root == nil {
		return
	}

	// Queue to perform level order traversal
	q := queue.New()
	q.Enqueue(HDNode{Node: root, HD: 0})

	// Map to store the last node at each horizontal distance for the bottom view
	bottomViewMap := make(map[int]int)

	// Track the range of HDs for proper ordering
	minHD, maxHD := 0, 0

	for q.Len() > 0 {
		hdNode := q.Dequeue().(HDNode)
		node, hd := hdNode.Node, hdNode.HD

		// Always update the map for bottom view
		bottomViewMap[hd] = node.Val

		// Update the range of HDs
		if hd < minHD {
			minHD = hd
		}
		if hd > maxHD {
			maxHD = hd
		}

		// Enqueue left and right children with updated HDs
		if node.Left != nil {
			q.Enqueue(HDNode{Node: node.Left, HD: hd - 1})
		}
		if node.Right != nil {
			q.Enqueue(HDNode{Node: node.Right, HD: hd + 1})
		}
	}

	// Print the bottom view from leftmost to rightmost HD
	fmt.Print("Bottom View: ")
	for hd := minHD; hd <= maxHD; hd++ {
		if val, found := bottomViewMap[hd]; found {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
}

func main() {
	// Example tree:
	//       1
	//      / \
	//     2   3
	//    / \ / \
	//   4  5 6  7
	//        \
	//         8

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Left.Right = &TreeNode{Val: 8}

	// Print the top view of the binary tree
	topView(root) // Output: Top View: 4 2 1 3 7

	// Print the bottom view of the binary tree
	bottomView(root) // Output: Bottom View: 4 2 5 8 7
}

```

### Vertical Order Travesal

Vertical Order Traversal of a Binary Tree

Given the `root` of a binary tree, calculate the vertical order traversal of the binary tree. For each node at position (row, col), its left and right children will be at positions (row + 1, col - 1) and (row + 1, col + 1) respectively. The solution should return a list of lists representing the node values grouped by columns and ordered from top to bottom and left to right.

**Example 1:**
- Input: `root = [3,9,20,null,null,15,7]`
- Output: `[[9], [3,15], [20], [7]]`
- Explanation:
  - Column -1: Only node 9 is in this column.
  - Column 0: Nodes 3 and 15 are in this column in that order from top to bottom.
  - Column 1: Only node 20 is in this column.
  - Column 2: Only node 7 is in this column.


```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"sort"
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
```




### Root to Leaf Paths
### Problem Description: Find All Root-to-Leaf Paths in a Binary Tree

Given a binary tree, the task is to find all possible paths from the root node to every leaf node. A leaf node is defined as a node that does not have any children. The paths should be printed in such a way that each path is represented as a sequence of node values separated by a space.

#### Example 1:

**Input:**
```
       1
    /     \
   2       3
```

**Output:**
```
1 2
1 3
```

**Explanation:** The binary tree has two paths from the root to the leaf nodes:
- Path 1: 1 -> 2
- Path 2: 1 -> 3

#### Example 2:

**Input:**
```
         10
       /    \
      20    30
     /  \
    40   60
```

**Output:**
```
10 20 40
10 20 60
10 30
```

```go
package main

import (
	"fmt"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findPaths is a helper function that finds all paths from root to leaf nodes.
func findPaths(node *TreeNode, path []int, paths *[][]int) {
	if node == nil {
		return
	}

	// Append the current node's value to the path
	path = append(path, node.Val)

	// If it's a leaf node, add the path to the paths list
	if node.Left == nil && node.Right == nil {
		// Make a copy of the path to avoid modification in subsequent operations
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*paths = append(*paths, pathCopy)
		return
	}

	// Recur for left and right children
	findPaths(node.Left, path, paths)
	findPaths(node.Right, path, paths)
}

// rootToLeafPaths returns all root-to-leaf paths in the binary tree as [][]int.
func rootToLeafPaths(root *TreeNode) [][]int {
	var paths [][]int
	findPaths(root, []int{}, &paths)
	return paths
}

func main() {
	// Creating a sample binary tree
	/*
	         1
	        / \
	       2   3
	      / \
	     4   5
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// Get all root-to-leaf paths
	paths := rootToLeafPaths(root)

	// Print all paths
	fmt.Println("Root to Leaf Paths:")
	for _, path := range paths {
		fmt.Println(path)
	}
}
```


### Problem Description

Given a binary tree, the task is to find the **maximum width** of the binary tree. The maximum width of a binary tree is defined as the maximum number of nodes present at any level in the tree. The width of a level is determined by the number of nodes between the leftmost and rightmost nodes at that level, including any null nodes in between. 

**Note**: The input binary tree is represented by `TreeNode`, a struct with `Val`, `Left`, and `Right` fields.

### Example

**Example 1**:
```
Input: [1, 2, 3, 4, 5, null, 7]
        1
       / \
      2   3
     / \   \
    4   5   7
Output: 3
Explanation:
The maximum width is 3, achieved at level 2 with nodes 4, 5, and 7.
```

**Example 2**:
```
Input: [1, 2, 3, null, null, null, 7]
        1
       / \
      2   3
           \
            7
Output: 2
Explanation:
The maximum width is 2, achieved at level 1 with nodes 2 and 3.
```

### Constraints

- The number of nodes in the tree is in the range `[0, 3000]`.
- `-100 <= Node.val <= 100`

### Solution in Go

To find the maximum width of a binary tree, we will perform a level-order traversal (BFS) while keeping track of the position of each node. This will help us to compute the width of each level by calculating the difference between the positions of the leftmost and rightmost nodes at that level.

Here's the Go implementation:

```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// WidthNode contains a tree node and its index in the level order traversal.
type WidthNode struct {
	Node  *TreeNode
	Index int
}

// maxWidth calculates the maximum width of the binary tree.
func maxWidth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxWidth := 0
	q := queue.New()
	q.Enqueue(WidthNode{Node: root, Index: 0})

	for q.Len() > 0 {
		levelSize := q.Len()
		first := q.Peek().(WidthNode).Index // Index of the first node at this level
		last := first                        // Initialize last with the first index

		for i := 0; i < levelSize; i++ {
			wnode := q.Dequeue().(WidthNode)
			node, index := wnode.Node, wnode.Index
			last = index // Update last to the current node's index

			// Enqueue children with their respective indices
			if node.Left != nil {
				q.Enqueue(WidthNode{Node: node.Left, Index: 2*index + 1})
			}
			if node.Right != nil {
				q.Enqueue(WidthNode{Node: node.Right, Index: 2*index + 2})
			}
		}

		// Width of the current level is last - first + 1
		levelWidth := last - first + 1
		if levelWidth > maxWidth {
			maxWidth = levelWidth
		}
	}

	return maxWidth
}

func main() {
	// Creating a sample binary tree
	/*
	        1
	       / \
	      2   3
	     / \   \
	    4   5   7
	           / \
	          8   9
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Left = &TreeNode{Val: 8}
	root.Right.Right.Right = &TreeNode{Val: 9}

	// Calculate the maximum width of the binary tree
	result := maxWidth(root)

	// Print the result
	fmt.Printf("Maximum Width of the Binary Tree: %d\n", result) // Output: 4
}
```


### Maximum Depth of Binary tree

The **maximum depth** of a binary tree is defined as the number of nodes along the longest path from the root node down to the farthest leaf node. In other words, it's the height of the tree.

### Example

**Example 1**:
```
Input: [3, 9, 20, null, null, 15, 7]
        3
       / \
      9  20
         / \
        15  7
Output: 3
```

**Example 2**:
```
Input: [1, null, 2]
        1
         \
          2
Output: 2
```

### Constraints

- The number of nodes in the tree is in the range `[0, 10^4]`.
- `-100 <= Node.val <= 100`

### Solution in Go

To find the maximum depth of a binary tree, we can use either a depth-first search (DFS) or breadth-first search (BFS) approach. Here, we'll demonstrate both approaches.

#### Depth-First Search (DFS)

The DFS approach involves recursively calculating the maximum depth of the left and right subtrees and then taking the maximum of the two.

Here's the implementation:

```go
package main

import (
	"fmt"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepthDFS calculates the maximum depth of a binary tree using DFS.
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Recursively find the depth of left and right subtrees
	leftDepth := maxDepthDFS(root.Left)
	rightDepth := maxDepthDFS(root.Right)

	// The depth of the tree is the maximum of the depths of the subtrees + 1 for the root
	return 1 + max(leftDepth, rightDepth)
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Creating a sample binary tree
	/*
	         3
	        / \
	       9  20
	         /  \
	        15   7
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// Calculate the maximum depth of the binary tree using DFS
	result := maxDepthDFS(root)

	// Print the result
	fmt.Printf("Maximum Depth of the Binary Tree (DFS): %d\n", result) // Output: 3
}
```

#### Breadth-First Search (BFS)

The BFS approach involves traversing the tree level by level and counting the number of levels.

Here's the implementation using a queue:

```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxDepthBFS calculates the maximum depth of a binary tree using BFS.
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	q := queue.New()
	q.Enqueue(root)

	for q.Len() > 0 {
		levelSize := q.Len()
		for i := 0; i < levelSize; i++ {
			node := q.Dequeue().(*TreeNode)

			// Enqueue children if they exist
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		maxDepth++
	}

	return maxDepth
}

func main() {
	// Creating a sample binary tree
	/*
	         3
	        / \
	       9  20
	         /  \
	        15   7
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// Calculate the maximum depth of the binary tree using BFS
	result := maxDepthBFS(root)

	// Print the result
	fmt.Printf("Maximum Depth of the Binary Tree (BFS): %d\n", result) // Output: 3
}
```



### Maximum Diameter of Binary tree

The **diameter** of a binary tree is defined as the length of the longest path between any two nodes in the tree. This path may or may not pass through the root. The length of a path is measured by the number of edges between the nodes.

### Example

**Example 1**:
```
Input: [1, 2, 3, 4, 5]
       1
      / \
     2   3
    / \     
   4   5    

Output: 3
Explanation: The diameter of the tree is the length of path [4,2,1,3] or [5,2,1,3], both paths have 3 edges.
```

**Example 2**:
```
Input: [1, 2]
       1
      / 
     2   
Output: 1
```

### Approach

To find the diameter of a binary tree, the main idea is to:
1. For each node, calculate the depth of the left and right subtrees.
2. The diameter at that node is the sum of the left and right depths.
3. The maximum diameter found during this process will be the answer.

We can use a recursive depth-first search (DFS) to find the depth of each subtree and update the maximum diameter found so far.

### Solution in Go

Here’s the implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Helper function to calculate the depth of the tree and update the diameter.
func diameterAndDepth(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}

	// Recursively find the depth of the left and right subtrees
	leftDepth := diameterAndDepth(node.Left, maxDiameter)
	rightDepth := diameterAndDepth(node.Right, maxDiameter)

	// Update the diameter: maximum number of edges between two leaf nodes
	*maxDiameter = max(*maxDiameter, leftDepth+rightDepth)

	// Return the depth of the tree rooted at the current node
	return 1 + max(leftDepth, rightDepth)
}

// diameterOfBinaryTree returns the diameter of the binary tree.
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	diameterAndDepth(root, &maxDiameter)
	return maxDiameter
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Creating a sample binary tree
	/*
	         1
	        / \
	       2   3
	      / \
	     4   5
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// Calculate the diameter of the binary tree
	result := diameterOfBinaryTree(root)

	// Print the result
	fmt.Printf("Diameter of the Binary Tree: %d\n", result) // Output: 3
}
```


### Check if the Binary Tree is Balanced Binary Tree

A **balanced binary tree** (also known as a height-balanced binary tree) is defined as a binary tree in which the depth of the two subtrees of every node never differs by more than 1. The task is to determine whether a given binary tree is balanced.

### Example

**Example 1**:
```
Input: [3, 9, 20, null, null, 15, 7]
        3
       / \
      9  20
         / \
        15  7
Output: true
```

**Example 2**:
```
Input: [1, 2, 2, 3, 3, null, null, 4, 4]
       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Output: false
```

### Approach

To determine if a binary tree is balanced, we can use a recursive approach:
1. For each node, calculate the height of its left and right subtrees.
2. If the difference in heights is more than 1 for any node, the tree is not balanced.
3. Additionally, a tree is not balanced if any of its subtrees is not balanced.

### Solution in Go

Here's the implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isBalancedHelper checks if the tree is balanced and returns the height of the tree.
func isBalancedHelper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	// Check if left subtree is balanced and its height
	leftBalanced, leftHeight := isBalancedHelper(root.Left)
	if !leftBalanced {
		return false, 0
	}

	// Check if right subtree is balanced and its height
	rightBalanced, rightHeight := isBalancedHelper(root.Right)
	if !rightBalanced {
		return false, 0
	}

	// Current node is balanced if the height difference is at most 1
	balanced := abs(leftHeight-rightHeight) <= 1

	// Height of the current node is max of left and right heights plus 1
	height := max(leftHeight, rightHeight) + 1

	return balanced, height
}

// isBalanced checks if the binary tree is balanced.
func isBalanced(root *TreeNode) bool {
	balanced, _ := isBalancedHelper(root)
	return balanced
}

// abs returns the absolute value of an integer.
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Creating a sample balanced binary tree
	/*
	         3
	        / \
	       9  20
	         /  \
	        15   7
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// Check if the binary tree is balanced
	fmt.Printf("Is the tree balanced? %v\n", isBalanced(root)) // Output: true

	// Creating an unbalanced binary tree
	/*
	        1
	       / \
	      2   2
	     / \
	    3   3
	   / \
	  4   4
	*/
	unbalancedRoot := &TreeNode{Val: 1}
	unbalancedRoot.Left = &TreeNode{Val: 2}
	unbalancedRoot.Right = &TreeNode{Val: 2}
	unbalancedRoot.Left.Left = &TreeNode{Val: 3}
	unbalancedRoot.Left.Right = &TreeNode{Val: 3}
	unbalancedRoot.Left.Left.Left = &TreeNode{Val: 4}
	unbalancedRoot.Left.Left.Right = &TreeNode{Val: 4}

	// Check if the binary tree is balanced
	fmt.Printf("Is the tree balanced? %v\n", isBalanced(unbalancedRoot)) // Output: false
}
```

### Lowest Common Ancestor for two given Nodes

### Example 1

**Input**: `root = [3, 5, 1, 6, 2, 0, 8, null, null, 7, 4]`, `p = 5`, `q = 1`

```
        3
       / \
      5   1
     / \ / \
    6  2 0  8
      / \
     7   4
```

**Output**: `3`

**Explanation**: The lowest common ancestor (LCA) of nodes 5 and 1 is 3 because both nodes share 3 as the deepest common ancestor.

### Example 2

**Input**: `root = [3, 5, 1, 6, 2, 0, 8, null, null, 7, 4]`, `p = 5`, `q = 4`

```
        3
       / \
      5   1
     / \ / \
    6  2 0  8
      / \
     7   4
```

**Output**: `5`

**Explanation**: The LCA of nodes 5 and 4 is 5 because node 5 is an ancestor of node 4. Therefore, the LCA is node 5 itself.

## Go Implementation

Here's the Go implementation that finds the LCA of two given nodes in a binary tree:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor returns the lowest common ancestor of nodes p and q.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	// Search for LCA in the left and right subtrees
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// If both left and right are non-null, current node is the LCA
	if left != nil && right != nil {
		return root
	}

	// Otherwise, return the non-null child
	if left != nil {
		return left
	}
	return right
}

func main() {
	// Creating a sample binary tree
	/*
	        3
	       / \
	      5   1
	     / \ / \
	    6  2 0  8
	      / \
	     7   4
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 0}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	p := root.Left       // Node with value 5
	q := root.Right      // Node with value 1
	lca := lowestCommonAncestor(root, p, q)
	fmt.Printf("LCA of nodes %d and %d is node with value %d\n", p.Val, q.Val, lca.Val)

	p = root.Left              // Node with value 5
	q = root.Left.Right.Right  // Node with value 4
	lca = lowestCommonAncestor(root, p, q)
	fmt.Printf("LCA of nodes %d and %d is node with value %d\n", p.Val, q.Val, lca.Val)
}
```


### Check if two trees are identical

To determine if two binary trees are identical, we need to check if both trees have the same structure and the same node values at each corresponding position.

### Example

**Example 1**:
```
Tree 1:          Tree 2:
    1                1
   / \              / \
  2   3            2   3
```
Output: `true`
Explanation: The two trees are identical because they have the same structure and node values.

**Example 2**:
```
Tree 1:          Tree 2:
    1                1
   / \              / \
  2   3            3   2
```
Output: `false`
Explanation: The two trees are not identical because the node values at the second level are different.

### Approach

To check if two trees are identical, we can use a recursive approach:
1. If both nodes are `nil`, they are considered identical at that level.
2. If one of the nodes is `nil` and the other is not, the trees are not identical.
3. If the values of the nodes are different, the trees are not identical.
4. Recursively check the left and right subtrees for both trees.

### Solution in Go

Here's the Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isIdentical checks if two binary trees are identical.
func isIdentical(root1, root2 *TreeNode) bool {
	// If both nodes are nil, they are identical
	if root1 == nil && root2 == nil {
		return true
	}

	// If one of the nodes is nil, they are not identical
	if root1 == nil || root2 == nil {
		return false
	}

	// Check if the values are the same and recursively check left and right subtrees
	return (root1.Val == root2.Val) &&
		isIdentical(root1.Left, root2.Left) &&
		isIdentical(root1.Right, root2.Right)
}

func main() {
	// Creating two identical sample binary trees
	/*
	        1                1
	       / \              / \
	      2   3            2   3
	*/
	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}

	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}

	// Check if the two trees are identical
	fmt.Printf("Are the two trees identical? %v\n", isIdentical(root1, root2)) // Output: true

	// Creating two different sample binary trees
	/*
	        1                1
	       / \              / \
	      2   3            3   2
	*/
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	root3.Right = &TreeNode{Val: 3}

	root4 := &TreeNode{Val: 1}
	root4.Left = &TreeNode{Val: 3}
	root4.Right = &TreeNode{Val: 2}

	// Check if the two trees are identical
	fmt.Printf("Are the two trees identical? %v\n", isIdentical(root3, root4)) // Output: false
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents each node in the binary tree with an integer value (`Val`) and pointers to the left and right children.

2. **isIdentical Function**:
   - This function checks if two trees rooted at `root1` and `root2` are identical.
   - It first checks if both nodes are `nil`. If they are, the trees are considered identical at that point.
   - If one node is `nil` and the other is not, the trees are not identical.
   - If the values of `root1` and `root2` are different, the trees are not identical.
   - The function then recursively checks the left and right subtrees.

3. **main Function**:
   - Two identical binary trees are created and compared using the `isIdentical` function, which returns `true`.
   - Two different binary trees are created and compared using the `isIdentical` function, which returns `false`.

### Complexity

- **Time Complexity**: O(n), where n is the number of nodes in the smaller of the two trees. This is because each node is compared once.
- **Space Complexity**: O(h), where h is the height of the tree, due to the recursive call stack. In the worst case, h can be equal to n (skewed tree), making the space complexity O(n).


### Problem Description

**Zigzag (or Spiral) Traversal** of a binary tree involves traversing the nodes level by level, but alternating the direction of traversal for each level. Specifically, the nodes at the first level are traversed from left to right, the nodes at the second level are traversed from right to left, the nodes at the third level are traversed from left to right, and so on.

### Example

**Example 1**:
```
Input: [3, 9, 20, null, null, 15, 7]
        3
       / \
      9  20
         / \
        15  7
Output: [[3], [20, 9], [15, 7]]
```

**Example 2**:
```
Input: [1, 2, 3, 4, 5, 6, 7]
        1
       / \
      2   3
     / \ / \
    4  5 6  7
Output: [[1], [3, 2], [4, 5, 6, 7]]
```

### Solution Approach

To achieve zigzag traversal:
1. Use a double-ended queue (deque) to facilitate easy addition and removal of nodes from both ends.
2. Use a boolean variable to track the current traversal direction (`leftToRight`). Start with `leftToRight = true`.
3. For each level:
   - If `leftToRight` is `true`, traverse from left to right and add children to the end of the deque.
   - If `leftToRight` is `false`, traverse from right to left and add children to the front of the deque.
4. Toggle the `leftToRight` flag after processing each level.


```go
package main

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// zigzagLevelOrder returns the zigzag level order traversal of a binary tree.
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var result [][]int
	q := queue.New()
	q.Enqueue(root)
	leftToRight := true

	for q.Len() > 0 {
		levelSize := q.Len()
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := q.Dequeue().(*TreeNode)

			// Place the node value based on the current traversal direction
			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}

			// Enqueue children nodes for the next level
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		result = append(result, level)
		leftToRight = !leftToRight
	}

	return result
}

func main() {
	// Creating a sample binary tree
	/*
	        3
	       / \
	      9  20
	         / \
	        15  7
	*/
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// Get the zigzag level order traversal
	result := zigzagLevelOrder(root)

	// Print the result
	fmt.Println("Zigzag Level Order Traversal:")
	for _, level := range result {
		fmt.Println(level)
	}
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary tree with integer value `Val` and pointers to the left and right children.

2. **zigzagLevelOrder Function**:
   - This function performs a zigzag (or spiral) level order traversal of the binary tree.
   - A queue (`q`) from `github.com/golang-collections/collections/queue` is used to manage the nodes level by level.
   - A boolean variable `leftToRight` indicates the current direction of traversal:
     - `true` for left-to-right traversal.
     - `false` for right-to-left traversal.
   - For each level:
     - Nodes are dequeued, and their values are added to the `level` slice in the appropriate order based on `leftToRight`.
     - Child nodes are enqueued for processing in the next level.
   - After processing each level, the direction (`leftToRight`) is toggled.

3. **main Function**:
   - Constructs a sample binary tree.
   - Calls `zigzagLevelOrder` to get the zigzag level order traversal result.
   - Prints the result, displaying the nodes at each level in the required order.

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is visited once during the traversal.
- **Space Complexity**: O(n), where n is the number of nodes in the tree. The maximum number of nodes stored in the queue at any time could be half the number of nodes, plus the space needed for the result list.


### Problem Description

The **Boundary Traversal** of a binary tree involves visiting all the boundary nodes of the tree in an anti-clockwise direction, starting from the root. The boundary traversal can be divided into three parts:

1. **Left Boundary**: Nodes on the left boundary of the tree excluding the leaf nodes.
2. **Leaf Nodes**: All the leaf nodes from left to right.
3. **Right Boundary**: Nodes on the right boundary of the tree excluding the leaf nodes, visited in bottom-up order.

### Example

**Example 1**:
```
Input: [1, 2, 3, 4, 5, 6, 7, null, null, 8, 9]
        1
       / \
      2   3
     / \ / \
    4  5 6  7
      / \
     8   9
Output: [1, 2, 4, 8, 9, 6, 7, 3]
```

**Explanation**:
- Left Boundary: 1, 2 (excluding 4 as it’s a leaf)
- Leaf Nodes: 4, 8, 9, 6, 7
- Right Boundary: 3 (excluding 7 as it’s a leaf), added in reverse order.

### Solution Approach

To implement boundary traversal:
1. **Left Boundary**: Traverse from the root to the leftmost node, excluding leaf nodes. Stop if a leaf node is encountered.
2. **Leaf Nodes**: Use any traversal method (such as DFS) to find and collect all leaf nodes.
3. **Right Boundary**: Traverse from the root to the rightmost node, excluding leaf nodes, and collect these nodes. Finally, reverse the collected right boundary nodes.

### Solution in Go

Here's the Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// addLeftBoundary adds the left boundary nodes to the result, excluding leaf nodes.
func addLeftBoundary(root *TreeNode, result *[]int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	*result = append(*result, root.Val)
	if root.Left != nil {
		addLeftBoundary(root.Left, result)
	} else {
		addLeftBoundary(root.Right, result)
	}
}

// addLeafNodes adds all leaf nodes to the result.
func addLeafNodes(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*result = append(*result, root.Val)
		return
	}

	addLeafNodes(root.Left, result)
	addLeafNodes(root.Right, result)
}

// addRightBoundary adds the right boundary nodes to the result, excluding leaf nodes.
func addRightBoundary(root *TreeNode, result *[]int) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}

	if root.Right != nil {
		addRightBoundary(root.Right, result)
	} else {
		addRightBoundary(root.Left, result)
	}

	*result = append(*result, root.Val)
}

// boundaryOfBinaryTree returns the boundary traversal of the binary tree.
func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	if !(root.Left == nil && root.Right == nil) {
		result = append(result, root.Val)
	}

	addLeftBoundary(root.Left, &result)
	addLeafNodes(root, &result)
	addRightBoundary(root.Right, &result)

	return result
}

func main() {
	// Creating a sample binary tree
	/*
	        1
	       / \
	      2   3
	     / \ / \
	    4  5 6  7
	      / \
	     8   9
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	root.Left.Right.Left = &TreeNode{Val: 8}
	root.Left.Right.Right = &TreeNode{Val: 9}

	// Get the boundary traversal
	result := boundaryOfBinaryTree(root)

	// Print the result
	fmt.Println("Boundary Traversal:")
	fmt.Println(result)
}
```

### Problem Description

The **Maximum Sum Path** in a binary tree is the path with the highest sum of node values. This path can start and end at any nodes in the tree and doesn't necessarily pass through the root. It can involve only one node if that node has the highest value in the tree.

### Example

**Example 1**:
```
Input: [1, 2, 3]
        1
       / \
      2   3
Output: 6
Explanation: The maximum path sum is 2 -> 1 -> 3.
```

**Example 2**:
```
Input: [-10, 9, 20, null, null, 15, 7]
       -10
       /  \
      9   20
         /  \
        15   7
Output: 42
Explanation: The maximum path sum is 15 -> 20 -> 7.
```

### Solution Approach

To find the maximum sum path in a binary tree:
1. **Define a recursive function** that computes the maximum path sum that "ends" at the current node and also updates the global maximum path sum found so far.
2. The maximum path sum for a given node can be obtained by:
   - Taking the maximum of the sum from the left subtree and the right subtree plus the current node's value.
   - The path can be split into a left and right subtree, meaning the path can include both children.
3. During the recursive traversal, keep track of the global maximum path sum.

### Implementation in Go

Here's the Go implementation:

```go
package main

import (
	"fmt"
	"math"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// maxSumPathDFS is a helper function to find the maximum path sum that goes through the node.
func maxSumPathDFS(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	// Recursively get the maximum sum paths for the left and right subtrees
	leftSum := max(0, maxSumPathDFS(node.Left, maxSum))  // If negative, take 0
	rightSum := max(0, maxSumPathDFS(node.Right, maxSum)) // If negative, take 0

	// Update the maximum sum with the maximum path sum with root at the current node
	currentMax := leftSum + rightSum + node.Val
	*maxSum = max(*maxSum, currentMax)

	// Return the maximum path sum "ending" at this node
	return max(leftSum, rightSum) + node.Val
}

// maxPathSum returns the maximum path sum in the binary tree.
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	maxSumPathDFS(root, &maxSum)
	return maxSum
}

// max returns the maximum of two integers.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Creating a sample binary tree
	/*
	        -10
	        /  \
	       9   20
	          /  \
	         15   7
	*/
	root := &TreeNode{Val: -10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// Get the maximum path sum
	result := maxPathSum(root)

	// Print the result
	fmt.Printf("Maximum Path Sum: %d\n", result) // Output: 42
}
```

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is visited once during the traversal.
- **Space Complexity**: O(h), where h is the height of the tree, due to the recursion stack. In the worst case (a skewed tree), the space complexity is O(n).

### Problem Description

Given two arrays, `preorder` and `inorder`, that represent the preorder and inorder traversal of a binary tree, construct and return the binary tree.

- **Preorder traversal** visits nodes in the order: root -> left subtree -> right subtree.
- **Inorder traversal** visits nodes in the order: left subtree -> root -> right subtree.

### Example

**Example 1**:
```
Preorder: [3, 9, 20, 15, 7]
Inorder: [9, 3, 15, 20, 7]

Output: 
        3
       / \
      9  20
         / \
        15  7
```

**Example 2**:
```
Preorder: [-1]
Inorder: [-1]

Output: 
        -1
```


### Optimal Solution Approach

1. **Use a Hash Map**: Create a map to store the value-to-index mappings for the inorder traversal. This allows for quick lookup of the root's index in the inorder list.
2. **Recursive Construction**:
   - Start with the entire range of the preorder and inorder lists.
   - The first element in the preorder list is the root of the tree.
   - Use the hash map to find the index of this root in the inorder list.
   - The left subtree will consist of elements before this index in the inorder list, and the right subtree will consist of elements after this index.
   - Recursively apply the same process to construct the left and right subtrees.

### Implementation in Go

Here's the optimized Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTreeHelper is a helper function for the recursive construction of the binary tree.
func buildTreeHelper(preorder []int, inorder []int, inorderMap map[int]int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	// The first element in the preorder slice is the root
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	// Find the index of the root in the inorder map
	inorderRootIndex := inorderMap[rootVal]

	// Number of nodes in the left subtree
	leftTreeSize := inorderRootIndex - inStart

	// Recursively build the left and right subtrees
	root.Left = buildTreeHelper(preorder, inorder, inorderMap, preStart+1, preStart+leftTreeSize, inStart, inorderRootIndex-1)
	root.Right = buildTreeHelper(preorder, inorder, inorderMap, preStart+leftTreeSize+1, preEnd, inorderRootIndex+1, inEnd)

	return root
}

// buildTree constructs the binary tree from preorder and inorder traversal.
func buildTree(preorder []int, inorder []int) *TreeNode {
	// Create a map for quick lookup of index in inorder traversal
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	return buildTreeHelper(preorder, inorder, inorderMap, 0, len(preorder)-1, 0, len(inorder)-1)
}

// printInorder prints the tree in inorder for verification
func printInorder(node *TreeNode) {
	if node == nil {
		return
	}
	printInorder(node.Left)
	fmt.Print(node.Val, " ")
	printInorder(node.Right)
}

func main() {
	// Example 1
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println("Inorder traversal of constructed tree:")
	printInorder(root) // Output: 9 3 15 20 7
	fmt.Println()

	// Example 2
	preorder = []int{-1}
	inorder = []int{-1}
	root = buildTree(preorder, inorder)
	fmt.Println("Inorder traversal of constructed tree:")
	printInorder(root) // Output: -1
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary tree with integer value (`Val`) and pointers to the left and right children.

2. **buildTree Function**:
   - Initializes a map (`inorderMap`) that maps each value in the inorder list to its index. This allows for O(1) time complexity for index lookups.
   - Calls the `buildTreeHelper` function to construct the tree.

3. **buildTreeHelper Function**:
   - This function is responsible for recursively constructing the tree.
   - It takes the current ranges of the preorder and inorder lists, and the map `inorderMap`.
   - The first element in the current `preorder` range is the root of the tree.
   - It finds the root's index in the inorder list using the map.
   - The size of the left subtree is calculated to determine the ranges for the left and right subtrees in both the preorder and inorder lists.
   - The function then recursively constructs the left and right subtrees and returns the root node.

4. **printInorder Function**:
   - A helper function to print the inorder traversal of the constructed tree, useful for verification.

5. **main Function**:
   - Constructs a binary tree from given `preorder` and `inorder` arrays.
   - Prints the inorder traversal of the constructed tree to verify correctness.

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is processed once, and the index lookups in the inorder list are O(1) due to the map.
- **Space Complexity**: O(n) for the map and the recursion stack. The recursion stack space can go up to O(h), where h is the height of the tree, in the worst case (skewed tree), which is O(n).



### Problem Description

Given two arrays, `inorder` and `postorder`, which represent the inorder and postorder traversal of a binary tree, the task is to construct and return the binary tree.

- **Inorder traversal** visits nodes in the order: left subtree -> root -> right subtree.
- **Postorder traversal** visits nodes in the order: left subtree -> right subtree -> root.

### Example

**Example 1**:
```
Input:
Inorder: [9, 3, 15, 20, 7]
Postorder: [9, 15, 7, 20, 3]

Output:
        3
       / \
      9  20
        /  \
       15   7
```

**Example 2**:
```
Input:
Inorder: [-1]
Postorder: [-1]

Output:
        -1
```

### Solution Approach

To construct the binary tree from inorder and postorder traversals:
1. The last element of the `postorder` list is the root node of the tree.
2. In the `inorder` list, the elements to the left of the root node represent the left subtree, and the elements to the right represent the right subtree.
3. Recursively repeat the process for the left and right subtrees using the corresponding segments from the `inorder` and `postorder` lists.

To optimize the solution:
- Use a hash map to store the indices of the elements in the `inorder` list for quick lookup. This allows for O(1) lookup time to find the root index in the `inorder` list, reducing the overall time complexity to O(n).

### Implementation in Go

Here's the optimized Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// buildTreeHelper is a helper function for the recursive construction of the binary tree.
func buildTreeHelper(inorder, postorder []int, inorderMap map[int]int, postStart, postEnd, inStart, inEnd int) *TreeNode {
	if postStart > postEnd || inStart > inEnd {
		return nil
	}

	// The last element in the postorder slice is the root
	rootVal := postorder[postEnd]
	root := &TreeNode{Val: rootVal}

	// Find the index of the root in the inorder map
	inorderRootIndex := inorderMap[rootVal]

	// Number of nodes in the left subtree
	leftTreeSize := inorderRootIndex - inStart

	// Recursively build the left and right subtrees
	root.Left = buildTreeHelper(inorder, postorder, inorderMap, postStart, postStart+leftTreeSize-1, inStart, inorderRootIndex-1)
	root.Right = buildTreeHelper(inorder, postorder, inorderMap, postStart+leftTreeSize, postEnd-1, inorderRootIndex+1, inEnd)

	return root
}

// buildTree constructs the binary tree from inorder and postorder traversal.
func buildTree(inorder []int, postorder []int) *TreeNode {
	// Create a map for quick lookup of index in inorder traversal
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	return buildTreeHelper(inorder, postorder, inorderMap, 0, len(postorder)-1, 0, len(inorder)-1)
}

// printInorder prints the tree in inorder for verification
func printInorder(node *TreeNode) {
	if node == nil {
		return
	}
	printInorder(node.Left)
	fmt.Print(node.Val, " ")
	printInorder(node.Right)
}

func main() {
	// Example 1
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
	fmt.Println("Inorder traversal of constructed tree:")
	printInorder(root) // Output: 9 3 15 20 7
	fmt.Println()

	// Example 2
	inorder = []int{-1}
	postorder = []int{-1}
	root = buildTree(inorder, postorder)
	fmt.Println("Inorder traversal of constructed tree:")
	printInorder(root) // Output: -1
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary tree with an integer value (`Val`) and pointers to the left and right children.

2. **buildTree Function**:
   - Initializes a map (`inorderMap`) that maps each value in the inorder list to its index. This allows for O(1) time complexity for index lookups.
   - Calls the `buildTreeHelper` function to construct the tree.

3. **buildTreeHelper Function**:
   - This function is responsible for recursively constructing the tree.
   - It takes the current ranges of the `inorder` and `postorder` lists, and the map `inorderMap`.
   - The last element in the current `postorder` range is the root of the tree.
   - It finds the root's index in the `inorder` list using the map.
   - The size of the left subtree is calculated to determine the ranges for the left and right subtrees in both the `inorder` and `postorder` lists.
   - The function then recursively constructs the left and right subtrees and returns the root node.

4. **printInorder Function**:
   - A helper function to print the inorder traversal of the constructed tree, useful for verification.

5. **main Function**:
   - Constructs a binary tree from given `inorder` and `postorder` arrays.
   - Prints the inorder traversal of the constructed tree to verify correctness.

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is processed once, and the index lookups in the `inorder` list are O(1) due to the map.
- **Space Complexity**: O(n) for the map and the recursion stack. The recursion stack space can go up to O(h), where h is the height of the tree, in the worst case (skewed tree), which is O(n).


### Problem Description

A **symmetric binary tree** is a binary tree in which the left and right subtrees are mirror images of each other. The tree is symmetric if the left subtree is a mirror reflection of the right subtree.

For example, the following tree is symmetric:

```
        1
       / \
      2   2
     / \ / \
    3  4 4  3
```

But the following tree is not symmetric:

```
        1
       / \
      2   2
       \   \
       3    3
```

### Solution Approach

To determine if a binary tree is symmetric, we can use a recursive approach:
1. Define a helper function `isMirror` that checks if two trees are mirror images of each other.
2. The `isMirror` function compares:
   - The values of the two nodes.
   - The left subtree of the first node with the right subtree of the second node.
   - The right subtree of the first node with the left subtree of the second node.

If the tree is empty, it is symmetric by definition.

### Implementation in Go

Here's the Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// isSymmetric checks if the binary tree is symmetric.
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror checks if two trees are mirror images of each other.
func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return (t1.Val == t2.Val) && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

func main() {
	// Creating a sample symmetric binary tree
	/*
	        1
	       / \
	      2   2
	     / \ / \
	    3  4 4  3
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}

	// Check if the binary tree is symmetric
	fmt.Printf("Is the tree symmetric? %v\n", isSymmetric(root)) // Output: true

	// Creating a sample non-symmetric binary tree
	/*
	        1
	       / \
	      2   2
	       \   \
	       3    3
	*/
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 2}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Right.Right = &TreeNode{Val: 3}

	// Check if the binary tree is symmetric
	fmt.Printf("Is the tree symmetric? %v\n", isSymmetric(root2)) // Output: false
}
```

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is visited once to compare values and recursively check subtrees.
- **Space Complexity**: O(h), where h is the height of the tree, due to the recursion stack. In the worst case (a skewed tree), the space complexity is O(n).