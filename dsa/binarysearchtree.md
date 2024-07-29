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

### Solution Approach

To populate the `nextRight` pointers:
1. **Level Order Traversal**: Use a level order traversal approach where you process each level of the tree before moving to the next. This can be done using a queue.
2. **Linking Nodes**: While processing each level, link each node to the next node in the queue.
3. **Use a Sentinel Node**: To mark the end of each level, a sentinel node (`nil`) can be used to identify when to move to the next level.

### Implementation in Go

Here's the Go implementation:

```go
package main

import "fmt"

// TreeNode represents a node in the binary tree with an additional next pointer.
type TreeNode struct {
	Val       int
	Left      *TreeNode
	Right     *TreeNode
	NextRight *TreeNode
}

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

// printLevels prints the nodes of the tree level by level using the nextRight pointers.
func printLevels(root *TreeNode) {
	start := root
	for start != nil {
		curr := start
		start = nil
		for curr != nil {
			fmt.Print(curr.Val, " ")
			if start == nil {
				if curr.Left != nil {
					start = curr.Left
				} else if curr.Right != nil {
					start = curr.Right
				}
			}
			curr = curr.NextRight
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
	     / \   \
	    4   5   7
	*/
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}

	// Connect the nextRight pointers
	connect(root)

	// Print the levels using nextRight pointers
	fmt.Println("Levels of the tree using nextRight pointers:")
	printLevels(root)
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary tree with an integer value (`Val`), pointers to the left and right children (`Left` and `Right`), and an additional pointer (`NextRight`) that points to the next node on the same level.

2. **connect Function**:
   - This function populates the `nextRight` pointers for each node in the tree.
   - A queue is used for a level-order traversal (BFS).
   - For each node at a given level, the `NextRight` pointer is set to point to the next node in the queue.
   - The last node in each level has its `NextRight` pointer set to `nil`.

3. **printLevels Function**:
   - This function prints the nodes of the tree level by level, using the `NextRight` pointers for traversal.

4. **main Function**:
   - Constructs a sample binary tree.
   - Calls `connect` to populate the `NextRight` pointers.
   - Prints the levels of the tree using `NextRight` pointers to verify the result.

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is visited once.
- **Space Complexity**: O(n) in the worst case, where the queue might hold all the nodes in the last level.



### Problem Description

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

To search for a given key in a BST:
1. **Start at the root**: If the tree is empty, return `nil`.
2. **Compare the key** with the current node's key:
   - If the key is equal to the current node's key, the search is successful, and the node is returned.
   - If the key is less than the current node's key, search in the left subtree.
   - If the key is greater than the current node's key, search in the right subtree.

This approach leverages the BST property and allows for efficient searching with a time complexity of O(h), where h is the height of the tree. In a balanced BST, this complexity is O(log n), where n is the number of nodes.

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

func main() {
	// Creating a sample binary search tree
	/*
	        4
	       / \
	      2   6
	     / \ / \
	    1  3 5  7
	*/
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}

	// Search for the key in the BST
	key := 5
	node := searchBST(root, key)
	if node != nil {
		fmt.Printf("Node with key %d found.\n", node.Val)
	} else {
		fmt.Printf("Node with key %d not found.\n", key)
	}

	// Search for a key not in the BST
	key = 10
	node = searchBST(root, key)
	if node != nil {
		fmt.Printf("Node with key %d found.\n", node.Val)
	} else {
		fmt.Printf("Node with key %d not found.\n", key)
	}
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary search tree with an integer value (`Val`) and pointers to the left and right children.

2. **searchBST Function**:
   - This function searches for the given `key` in the BST.
   - It starts from the `root` and compares the `key` with the current node's value.
   - If the `key` is found (i.e., `root.Val == key`), the node is returned.
   - If the `key` is less than the current node's value, the function recursively searches in the left subtree.
   - If the `key` is greater than the current node's value, the function recursively searches in the right subtree.
   - If the `root` is `nil`, it means the `key` is not present in the BST, and the function returns `nil`.

3. **main Function**:
   - Constructs a sample binary search tree.
   - Calls `searchBST` to search for a given key in the BST and prints whether the key is found or not.

### Complexity:

- **Time Complexity**: O(h), where h is the height of the tree. In a balanced BST, this is O(log n), where n is the number of nodes.
- **Space Complexity**: O(h), due to the recursion stack. In the worst case (skewed tree), the space complexity is O(n).


### Problem Description

A Binary Search Tree (BST) is a binary tree where each node has the following properties:
1. The left subtree of a node contains only nodes with keys less than the node's key.
2. The right subtree of a node contains only nodes with keys greater than the node's key.
3. Both the left and right subtrees must also be binary search trees.

Given a binary tree, the task is to determine whether it is a BST or not.

### Solution Approach

To check if a binary tree is a BST, we can use a recursive approach:
1. For each node, we need to ensure that all nodes in the left subtree have values less than the node's value and all nodes in the right subtree have values greater than the node's value.
2. We can maintain a range `[min, max]` for each node, where:
   - Initially, `min` is negative infinity, and `max` is positive infinity.
   - For the left child of a node, the new `max` becomes the node's value.
   - For the right child of a node, the new `min` becomes the node's value.

This approach ensures that every node adheres to the BST properties throughout the tree.

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

func main() {
	// Creating a sample binary tree that is a BST
	/*
	        4
	       / \
	      2   6
	     / \ / \
	    1  3 5  7
	*/
	root1 := &TreeNode{Val: 4}
	root1.Left = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 6}
	root1.Left.Left = &TreeNode{Val: 1}
	root1.Left.Right = &TreeNode{Val: 3}
	root1.Right.Left = &TreeNode{Val: 5}
	root1.Right.Right = &TreeNode{Val: 7}

	// Check if the tree is a BST
	fmt.Printf("Tree 1 is a BST: %v\n", isBST(root1)) // Output: true

	// Creating a sample binary tree that is not a BST
	/*
	        4
	       / \
	      2   6
	     / \ / \
	    1  3 5  8
	              /
	             7
	*/
	root2 := &TreeNode{Val: 4}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 6}
	root2.Left.Left = &TreeNode{Val: 1}
	root2.Left.Right = &TreeNode{Val: 3}
	root2.Right.Left = &TreeNode{Val: 5}
	root2.Right.Right = &TreeNode{Val: 8}
	root2.Right.Right.Left = &TreeNode{Val: 7}

	// Check if the tree is a BST
	fmt.Printf("Tree 2 is a BST: %v\n", isBST(root2)) // Output: false
}
```

### Complexity:

- **Time Complexity**: O(n), where n is the number of nodes in the tree. Each node is visited once.
- **Space Complexity**: O(h), where h is the height of the tree, due to the recursion stack. In the worst case (a skewed tree), the space complexity is O(n).


To find the **Lowest Common Ancestor (LCA)** of two nodes in a Binary Search Tree (BST) using a depth-first search (DFS) approach, we can take advantage of the BST properties:

1. **BST Property**: In a BST, for any given node:
   - The left subtree contains only nodes with values less than the node's key.
   - The right subtree contains only nodes with values greater than the node's key.

Using these properties, we can optimize the LCA search:

- Start at the root and compare the values of the nodes `p` and `q` with the current node's value.
- If both `p` and `q` are less than the current node's value, the LCA must be in the left subtree.
- If both `p` and `q` are greater than the current node's value, the LCA must be in the right subtree.
- If `p` and `q` are on opposite sides of the current node, or one of them is the current node, then the current node is the LCA.

### Implementation in Go

Here’s the Go implementation using a DFS approach for a BST:

```go
package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func main() {
	// Creating a sample binary search tree
	/*
	        6
	       / \
	      2   8
	     / \ / \
	    0  4 7  9
	      / \
	     3   5
	*/
	root := &TreeNode{Val: 6}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 0}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Right.Left = &TreeNode{Val: 3}
	root.Left.Right.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 7}
	root.Right.Right = &TreeNode{Val: 9}

	// Finding LCA of nodes with values 2 and 8
	p := &TreeNode{Val: 2}
	q := &TreeNode{Val: 8}
	lca := lowestCommonAncestor(root, p, q)
	if lca != nil {
		fmt.Printf("LCA of nodes %d and %d is node with value %d\n", p.Val, q.Val, lca.Val)
	} else {
		fmt.Printf("LCA of nodes %d and %d not found\n", p.Val, q.Val)
	}

	// Finding LCA of nodes with values 2 and 4
	p = &TreeNode{Val: 2}
	q = &TreeNode{Val: 4}
	lca = lowestCommonAncestor(root, p, q)
	if lca != nil {
		fmt.Printf("LCA of nodes %d and %d is node with value %d\n", p.Val, q.Val, lca.Val)
	} else {
		fmt.Printf("LCA of nodes %d and %d not found\n", p.Val, q.Val)
	}
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary search tree with an integer value (`Val`) and pointers to the left and right children.

2. **lowestCommonAncestor Function**:
   - This function determines the LCA of two nodes `p` and `q` in a BST.
   - **Base Case**: If the `root` is `nil`, the function returns `nil`.
   - If both `p` and `q` are less than the `root`'s value, the LCA must be in the left subtree, so the function recursively calls itself with the left child.
   - If both `p` and `q` are greater than the `root`'s value, the LCA must be in the right subtree, so the function recursively calls itself with the right child.
   - If one node is on the left and the other on the right, or one of the nodes is the `root`, then the current `root` is the LCA.

3. **main Function**:
   - Constructs a sample BST and finds the LCA for different pairs of nodes. It prints the result to verify the correctness.

### Complexity:

- **Time Complexity**: O(h), where h is the height of the tree. This is O(log n) for a balanced BST, where n is the number of nodes.
- **Space Complexity**: O(h), due to the recursion stack. In the worst case (a skewed tree), the space complexity could be O(n).


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

func main() {
	// Example keys to construct the BST
	keys := []int{6, 2, 8, 0, 4, 7, 9, 3, 5}

	// Construct the BST
	root := constructBST(keys)

	// Print the inorder traversal of the BST
	fmt.Println("Inorder traversal of the constructed BST:")
	inorderTraversal(root) // Output: 0 2 3 4 5 6 7 8 9
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

### Solution Approach

To construct a BST from a preorder traversal, we can utilize the properties of BSTs and the sequence of preorder traversal:
1. **Root Identification**: The first element in the preorder array is the root of the BST.
2. **Subtree Construction**:
   - All elements following the root in the preorder array that are less than the root will form the left subtree.
   - All elements greater than the root will form the right subtree.
3. **Recursive Construction**:
   - Recursively apply the above logic to construct the left and right subtrees.

### Implementation in Go

Here's the Go implementation:

```go
package main

import (
	"fmt"
	"math"
)

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func main() {
	// Example preorder traversal to construct the BST
	preorder := []int{8, 5, 1, 7, 10, 12}

	// Construct the BST
	root := constructBSTFromPreorder(preorder)

	// Print the inorder traversal of the constructed BST
	fmt.Println("Inorder traversal of the constructed BST:")
	inorderTraversal(root) // Output: 1 5 7 8 10 12
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary search tree with an integer value (`Val`) and pointers to the left and right children (`Left` and `Right`).

2. **constructBSTFromPreorderHelper Function**:
   - A recursive helper function that constructs a BST from the preorder traversal array.
   - The function maintains a pointer (`idx`) to the current index in the preorder array.
   - It checks whether the current value falls within the permissible range (`min` and `max`) for the node. This range ensures that the BST properties are maintained.
   - If the value is within the range, a new `TreeNode` is created with this value, and the index is incremented.
   - The function then recursively constructs the left and right subtrees by updating the permissible range:
     - The left subtree can only contain values less than the current node's value.
     - The right subtree can only contain values greater than the current node's value.

3. **constructBSTFromPreorder Function**:
   - This function initializes the index (`idx`) and calls the helper function with the entire permissible range of values (`math.MinInt64` to `math.MaxInt64`).

4. **inorderTraversal Function**:
   - A utility function to print the inorder traversal of the BST. Inorder traversal of a BST should result in a sorted sequence of the node values, which can be used to verify the correctness of the constructed tree.

5. **main Function**:
   - Demonstrates the construction of a BST from a given preorder traversal and prints the inorder traversal of the resulting BST.

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

func main() {
	// Creating a sample binary search tree
	/*
	        20
	       /  \
	      8    22
	     / \
	    4   12
	       /  \
	      10   14
	*/
	root := &TreeNode{Val: 20}
	root.Left = &TreeNode{Val: 8}
	root.Right = &TreeNode{Val: 22}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 12}
	root.Left.Right.Left = &TreeNode{Val: 10}
	root.Left.Right.Right = &TreeNode{Val: 14}

	key := 13

	// Finding floor
	floorNode := findFloor(root, key)
	if floorNode != nil {
		fmt.Printf("Floor of %d is %d\n", key, floorNode.Val)
	} else {
		fmt.Printf("Floor of %d does not exist\n", key)
	}

	// Finding ceil
	ceilNode := findCeil(root, key)
	if ceilNode != nil {
		fmt.Printf("Ceil of %d is %d\n", key, ceilNode.Val)
	} else {
		fmt.Printf("Ceil of %d does not exist\n", key)
	}
}
```

### Explanation:

1. **TreeNode Structure**:
   - Represents a node in the binary search tree with an integer value (`Val`) and pointers to the left and right children (`Left` and `Right`).

2. **findFloor Function**:
   - This function finds the floor of a given `key` in the BST.
   - It iteratively traverses the tree, updating the `floor` variable whenever it encounters a node with a value less than or equal to the `key` and continues the search for a potentially closer floor value in the right subtree.

3. **findCeil Function**:
   - This function finds the ceil of a given `key` in the BST.
   - It iteratively traverses the tree, updating the `ceil` variable whenever it encounters a node with a value greater than or equal to the `key` and continues the search for a potentially closer ceil value in the left subtree.

4. **main Function**:
   - Constructs a sample BST and finds both the floor and ceil for a given `key`, printing the results.

### Complexity:

- **Time Complexity**: O(h), where h is the height of the tree. In the best case (balanced BST), this is O(log n), where n is the number of nodes. This is because we might traverse from the root to a leaf node.
- **Space Complexity**: O(1). The functions use a constant amount of extra space for storing pointers (`floor`, `ceil`, and the traversal pointer). The iterative approach ensures that there is no additional space needed for recursion stacks.