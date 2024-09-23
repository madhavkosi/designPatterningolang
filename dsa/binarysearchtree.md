### Populate Next Right pointers of Tree

In a binary tree, each node contains an additional pointer called `nextRight` (or simply `next`). This pointer should be set to point to the next right node in the same level. If there is no next right node, the `nextRight` pointer should be set to `nil`.

For example, given the following tree:
```
    1
   / \
  2   3
 / \   \
4   5   7
```

The tree should be transformed to:
```
    1 -> nil
   / \
  2 -> 3 -> nil
 / \   \
4-> 5 -> 7 -> nil
```
### Implementation in Go

Here's the Go implementation:

```go

// connect populates the nextRight pointers to point to the next right node in the same level.
func connect(root *TreeNode) {
	if root == nil {
		return
	}

	// Initialize a queue for level order traversal
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		var prev *TreeNode

		// Process all nodes at the current level
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			// If there is a previous node, link it to the current node
			if prev != nil {
				prev.NextRight = node
			}
			prev = node

			// Enqueue the left and right children
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// The last node in the level should point to nil
		if prev != nil {
			prev.NextRight = nil
		}
	}
}
```




### Key in BST

In a Binary Search Tree (BST), each node contains a key, and the keys in the left subtree are less than the node's key, while the keys in the right subtree are greater than the node's key. Given a key, the task is to find the node with the given key in the BST. If the key exists, return the node; otherwise, return `nil`.

### Example

**Example 1**:
```
BST:
      4
     / \
    2   6
   / \ / \
  1  3 5  7

Key: 5
Output: Node with value 5
```

**Example 2**:
```
BST:
      4
     / \
    2   6
   / \ / \
  1  3 5  7

Key: 10
Output: nil (Key not found)
```

### Solution Approach


### Implementation in Go

Here's the Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// searchBST searches for a given key in the BST and returns the node if found, otherwise returns nil.
func searchBST(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// If the key is found, return the current node
	if root.Val == key {
		return root
	}

	// If the key is less than the current node's value, search in the left subtree
	if key < root.Val {
		return searchBST(root.Left, key)
	}

	// If the key is greater than the current node's value, search in the right subtree
	return searchBST(root.Right, key)
}

```



### Complexity:

- **Time Complexity**: O(h), where h is the height of the tree. In a balanced BST, this is O(log n), where n is the number of nodes.
- **Space Complexity**: O(h), due to the recursion stack. In the worst case (skewed tree), the space complexity is O(n).


### Check BST

A Binary Search Tree (BST) is a binary tree where each node has the following properties:
1. The left subtree of a node contains only nodes with keys less than the node's key.
2. The right subtree of a node contains only nodes with keys greater than the node's key.
3. Both the left and right subtrees must also be binary search trees.

Given a binary tree, the task is to determine whether it is a BST or not.

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

// isBSTUtil is a utility function to check if the tree is a BST.
func isBSTUtil(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	// Check the current node value against min and max range
	if root.Val <= min || root.Val >= max {
		return false
	}

	// Check the left and right subtrees with updated ranges
	return isBSTUtil(root.Left, min, root.Val) && isBSTUtil(root.Right, root.Val, max)
}

// isBST checks if the binary tree is a BST.
func isBST(root *TreeNode) bool {
	return isBSTUtil(root, math.MinInt64, math.MaxInt64)
}

```

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is visited once.
- **Space Complexity**: O(h), where h is the height of the tree, due to the recursion stack. In the worst case (a skewed tree), the space complexity is O(n).

### Lowest Common Ancestor (LCA) BST

### Implementation in Go

Here’s the Go implementation using a DFS approach for a BST:

```go
package main

import "fmt"

// lowestCommonAncestor finds the LCA of two given nodes in a BST.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// Base case: if root is nil, return nil
	if root == nil {
		return nil
	}

	// If both p and q are less than root, then LCA lies in the left subtree
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	// If both p and q are greater than root, then LCA lies in the right subtree
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	// If one node is on the left and the other is on the right, root is the LCA
	// This also covers the case where one of the nodes is the root
	return root
}

```
### Construct BST from given keys


Given an array of integers representing the keys, the task is to construct a Binary Search Tree (BST) from these keys. The properties of a BST are:
1. The left subtree of a node contains only nodes with keys less than the node's key.
2. The right subtree of a node contains only nodes with keys greater than the node's key.
3. Both the left and right subtrees must also be binary search trees.

### Solution Approach

To construct a BST from an array of keys:
1. **Initialize the Tree**: Start with an empty tree.
2. **Insert Each Key**:
   - For each key, start from the root.
   - If the key is less than the current node's key, move to the left child.
   - If the key is greater than the current node's key, move to the right child.
   - Insert the key in the correct position (as a left or right child of a leaf node).

### Implementation in Go

Here's the Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// insertNode inserts a new node with the given value into the BST.
func insertNode(root *TreeNode, key int) *TreeNode {
	// If the tree is empty, return a new node
	if root == nil {
		return &TreeNode{Val: key}
	}

	// Otherwise, recur down the tree
	if key < root.Val {
		root.Left = insertNode(root.Left, key)
	} else if key > root.Val {
		root.Right = insertNode(root.Right, key)
	}

	// Return the (unchanged) root pointer
	return root
}

// constructBST constructs a BST from the given keys.
func constructBST(keys []int) *TreeNode {
	var root *TreeNode
	for _, key := range keys {
		root = insertNode(root, key)
	}
	return root
}

// inorderTraversal prints the inorder traversal of the tree.
func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inorderTraversal(root.Right)
	}
}

```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary search tree with an integer value (`Val`) and pointers to the left and right children (`Left` and `Right`).

2. **insertNode Function**:
   - This function inserts a new node with the given key into the BST.
   - If the tree is empty, it creates a new node with the given key and returns it.
   - Otherwise, it recursively finds the correct position for the new key by comparing it with the current node's key.
   - If the key is less than the current node's key, it goes to the left subtree; if greater, it goes to the right subtree.

3. **constructBST Function**:
   - Takes an array of integers (`keys`) and constructs a BST by calling `insertNode` for each key.
   - The function initializes an empty BST and inserts each key sequentially.

4. **inorderTraversal Function**:
   - A utility function to print the inorder traversal of the BST. Inorder traversal of a BST results in a sorted sequence of the node values.

5. **main Function**:
   - Provides example keys, constructs the BST, and prints the inorder traversal of the constructed BST to verify correctness.

### Complexity:

- **Time Complexity**:
  - Insertion in the worst case takes O(h) time, where h is the height of the tree. In the worst case (unbalanced tree), this can be O(n), where n is the number of nodes. However, if the tree is balanced, the complexity is O(log n).
  - Constructing the BST for n keys takes O(n * h) time in the worst case, which simplifies to O(n^2) for an unbalanced tree. In a balanced scenario, it would be O(n log n).

- **Space Complexity**: 
  - O(h) due to the recursion stack for the insertions. In the worst case (a completely unbalanced tree), this could be O(n).


### Construct a BST from a preorder traversal

Given an array representing the preorder traversal of a binary search tree (BST), the task is to construct the BST. In a preorder traversal, the nodes are visited in the following order: root, left subtree, and then right subtree.


### Implementation in Go

Here's the Go implementation:

```go

// constructBSTFromPreorderHelper is a helper function to construct a BST from preorder traversal.
func constructBSTFromPreorderHelper(preorder []int, idx *int, min, max int) *TreeNode {
	// Base case: if all elements are processed or the current element is out of bounds
	if *idx >= len(preorder) {
		return nil
	}

	val := preorder[*idx]
	if val < min || val > max {
		return nil
	}

	// Create a new node with the current element and increment index
	*idx++
	node := &TreeNode{Val: val}

	// Elements in the left subtree must be smaller than the current node's value
	node.Left = constructBSTFromPreorderHelper(preorder, idx, min, val)

	// Elements in the right subtree must be greater than the current node's value
	node.Right = constructBSTFromPreorderHelper(preorder, idx, val, max)

	return node
}

// constructBSTFromPreorder constructs a BST from a preorder traversal array.
func constructBSTFromPreorder(preorder []int) *TreeNode {
	idx := 0
	return constructBSTFromPreorderHelper(preorder, &idx, math.MinInt64, math.MaxInt64)
}

// inorderTraversal prints the inorder traversal of the tree.
func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inorderTraversal(root.Right)
	}
}

```
### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each element of the preorder array is processed exactly once.
- **Space Complexity**: O(h), where h is the height of the tree, due to the recursion stack. In the worst case (a skewed tree), this could be O(n). However, if the BST is balanced, the space complexity is O(log n).



### Finding Floor and Ceil in a BST

To complement the floor function, we can also implement a function to find the **ceil** of a given key in a Binary Search Tree (BST). The **ceil** of a key is the smallest key in the BST that is greater than or equal to the given key. If no such key exists, the ceil is considered to be `nil`.

### Solution Approach for Ceil

To find the ceil of a key in a BST, the approach is similar to finding the floor:
1. If the current node's key is equal to the given key, then the ceil is the current node.
2. If the current node's key is less than the given key, then we search in the right subtree because all keys in the left subtree are smaller.
3. If the current node's key is greater than the given key, then the current node could be the ceil, but there might be a closer value in the left subtree.

### Implementation in Go

Here’s the Go implementation for finding both floor and ceil in a BST:

```go
package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findFloor finds the floor of a given key in the BST.
func findFloor(root *TreeNode, key int) *TreeNode {
	var floor *TreeNode

	for root != nil {
		if root.Val == key {
			return root
		} else if root.Val > key {
			root = root.Left
		} else {
			floor = root
			root = root.Right
		}
	}

	return floor
}

// findCeil finds the ceil of a given key in the BST.
func findCeil(root *TreeNode, key int) *TreeNode {
	var ceil *TreeNode

	for root != nil {
		if root.Val == key {
			return root
		} else if root.Val < key {
			root = root.Right
		} else {
			ceil = root
			root = root.Left
		}
	}

	return ceil
}
```

### Complexity:

- **Time Complexity**: O(h), where h is the height of the tree. In the best case (balanced BST), this is O(log n), where n is the number of nodes. This is because we might traverse from the root to a leaf node.
- **Space Complexity**: O(1). The functions use a constant amount of extra space for storing pointers (`floor`, `ceil`, and the traversal pointer). The iterative approach ensures that there is no additional space needed for recursion stacks.