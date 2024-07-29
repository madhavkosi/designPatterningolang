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

This approach uses a bottom-up recursive method to check if a binary tree is balanced. It has a time complexity of O(n) where n is the number of nodes in the tree, because each node is visited once. The space complexity is O(h) where h is the height of the tree, due to the recursive call stack.