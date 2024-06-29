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

These solutions provide implementations for different types of binary tree traversals: inorder, preorder, postorder, and level order, each with their respective example input and output.