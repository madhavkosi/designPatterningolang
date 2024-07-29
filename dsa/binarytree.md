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

### Problem Description

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


To also include the **Bottom View** of a binary tree, we can use a similar approach as the Top View. The Bottom View consists of the nodes visible when the tree is viewed from the bottom. For each horizontal distance (HD), the last node encountered during a level-order traversal is part of the bottom view.

### Implementation Overview

1. **Data Structures**:
   - **TreeNode**: Represents a node in the binary tree.
   - **HDNode**: Stores a node along with its HD.
   - **Map**: Stores the node value for each HD. In the case of the bottom view, we update the value whenever we encounter a new node at the same HD during traversal.

2. **Steps**:
   - Use a queue for BFS traversal.
   - For the top view, store the first node for each HD.
   - For the bottom view, store the last node encountered for each HD.
   - After traversal, extract values from the map in the order of HDs to get the views.


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

### Explanation

1. **Top View**:
   - For each horizontal distance (HD), the first node encountered during BFS traversal is stored in `topViewMap`.
   - The output is printed in the order of HDs from `minHD` to `maxHD`.

2. **Bottom View**:
   - Similar to the top view, but for each HD, we update the map entry with the latest node encountered. This ensures that the map holds the bottom-most node for each HD.
   - The output is printed in the order of HDs from `minHD` to `maxHD`.

3. **Main Function**:
   - The tree is constructed, and the top and bottom views are computed and printed.

This code effectively captures and prints the top and bottom views of a binary tree using level-order traversal and a map to manage the nodes at each horizontal distance.