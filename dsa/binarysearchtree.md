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


Here is the solution for finding a pair with a given sum in a Binary Search Tree (BST):

---

**Find a Pair with Given Sum in BST**

Problem Description  
Given a binary search tree (BST) and a target sum, the task is to find two nodes in the BST whose sum is equal to the given target. If such a pair exists, return `true`; otherwise, return `false`.

**Examples**

Example 1  
Input: A BST with inorder traversal [1, 2, 3, 4, 5, 6, 7], target = 9  
Output: `true` (since 2 + 7 = 9)

Example 2  
Input: A BST with inorder traversal [1, 2, 3, 4, 5], target = 10  
Output: `false` (no such pair exists)

**Approach**

1. Perform an in-order traversal of the BST to get a sorted array of elements.
2. Use the two-pointer technique to find if there is a pair whose sum equals the target.

**Implementation in Go**

```go
package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorderTraversal is a helper function to perform an in-order traversal
// and store the elements in a slice.
func inorderTraversal(root *TreeNode, elems *[]int) {
	if root == nil {
		return
	}

	// Traverse the left subtree
	inorderTraversal(root.Left, elems)

	// Append the root's value to the slice
	*elems = append(*elems, root.Val)

	// Traverse the right subtree
	inorderTraversal(root.Right, elems)
}

// findPairWithSum uses two-pointer technique on the sorted array to find a pair
// that sums up to the target.
func findPairWithSum(root *TreeNode, target int) bool {
	// Perform in-order traversal to get elements in sorted order
	var elems []int
	inorderTraversal(root, &elems)

	// Use two-pointer technique to find a pair with the given sum
	left, right := 0, len(elems)-1
	for left < right {
		sum := elems[left] + elems[right]
		if sum == target {
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}

```


Here’s the updated solution with the addition of a function to find the kth largest element in a Binary Search Tree (BST), along with the existing functionality to find the kth smallest element:

---

**Find kth Smallest and kth Largest Element in BST**

Problem Description  
Given a binary search tree (BST), write functions to find:
- The kth smallest element.
- The kth largest element.

The kth smallest element can be found by performing an in-order traversal (left-root-right), while the kth largest element can be found by performing a reverse in-order traversal (right-root-left).

**Examples**

Example 1 (kth Smallest)  
Input: A BST with inorder traversal [1, 2, 3, 4, 5, 6, 7], k = 3  
Output: 3

Example 2 (kth Largest)  
Input: A BST with inorder traversal [1, 2, 3, 4, 5, 6, 7], k = 2  
Output: 6

**Implementation in Go**

```go
package main

import "fmt"

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func inorderTraversal(root *TreeNode, k *int, result *int) {
	if root == nil {
		return
	}

	inorderTraversal(root.Left, k, result)

	*k--
	if *k == 0 {
		*result = root.Val
		return
	}

	inorderTraversal(root.Right, k, result)
}

func reverseInorderTraversal(root *TreeNode, k *int, result *int) {
	if root == nil {
		return
	}

	reverseInorderTraversal(root.Right, k, result)

	*k--
	if *k == 0 {
		*result = root.Val
		return
	}

	reverseInorderTraversal(root.Left, k, result)
}

func kthSmallest(root *TreeNode, k int) int {
	var result int
	inorderTraversal(root, &k, &result)
	return result
}

func kthLargest(root *TreeNode, k int) int {
	var result int
	reverseInorderTraversal(root, &k, &result)
	return result
}

```

**Explanation**:
- **`kthSmallest` function**: Uses in-order traversal (left-root-right) to find the kth smallest element.
- **`kthLargest` function**: Uses reverse in-order traversal (right-root-left) to find the kth largest element.
- The traversal continues until the kth element is found by decrementing `k` with each visited node. When `k` reaches 0, the current node's value is stored as the result.

This solution efficiently finds both the kth smallest and largest elements in O(h + k) time, where `h` is the height of the tree and `k` is the number of elements processed.


Here is a solution to **serialize** and **deserialize** a Binary Search Tree (BST) in Go. Serialization converts the BST into a string representation, and deserialization reconstructs the BST from this string.

---

**Serialize and Deserialize BST**

Problem Description  
Serialization is the process of converting a data structure or object into a format that can be easily stored or transmitted. Deserialization is the reverse process, converting the serialized format back into the original data structure. Given a binary search tree (BST), write functions to serialize the tree into a string and deserialize it back into the BST.

**Examples**

Example 1  
Input: A BST with nodes [5, 3, 7, 2, 4, 6, 8]  
Output: Serialized string and the reconstructed BST

**Approach**

1. **Serialization**:  
   Perform a pre-order traversal (root-left-right) and convert the nodes into a string.
   
2. **Deserialization**:  
   Reconstruct the BST from the pre-order traversal string.

**Implementation in Go**

```go
func Serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return fmt.Sprintf("%d,%s,%s", root.Val, Serialize(root.Left), Serialize(root.Right))
}

// Deserialize converts a string back to a tree
func Deserialize(data string) *TreeNode {
	str := strings.Split(data, ",")
	index := -1
	return DeserializeDfs(&index, str)
}
func DeserializeDfs(index *int, str []string) *TreeNode {
	*index = *index + 1
	if str[*index] == "#" {
		return nil
	}
	val, _ := strconv.Atoi(str[*index])
	root := &TreeNode{Val: val}
	root.Left = DeserializeDfs(index, str)
	root.Right = DeserializeDfs(index, str)
	return root
}

// // Example usage
func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	serialized := Serialize(root)
	fmt.Println("Serialized:", serialized) //"1,2,#,#,3,4,#,#,5,#,#"

}
```

**Space Complexity**:
- O(n) for storing the serialized string and the tree nodes during deserialization.