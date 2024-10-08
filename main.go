package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
)

type Point struct {
	x, y int
}

func validateGrid(grid [][]int, r, c int) bool {
	if r < 0 || r >= len(grid) || c >= len(grid[0]) || c < 0 || grid[r][c] != 1 {
		return false
	}
	return true
}
func orangesRotting(grid [][]int) int {
	q := queue.New()
	r := len(grid)
	c := len(grid[0])
	freshOrange := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == 2 {
				q.Enqueue(Point{x: i, y: j})
			}
			if grid[i][j] == 1 {
				q.Enqueue(Point{x: i, y: j})
				freshOrange += 1
			}
		}
	}
	cnt := 0
	dir_x := []int{-1, 0, 1, 0}
	dir_y := []int{0, -1, 0, 1}
	for q.Len() != 0 {
		length := q.Len()
		for i := 0; i < length; i++ {
			point := q.Dequeue().(Point)
			for j := 0; j < 4; j++ {
				new_r := point.x + dir_x[j]
				new_c := point.y + dir_y[j]
				if validateGrid(grid, point.x+dir_x[j], point.y+dir_y[j]) {
					grid[new_r][new_c] = 2
					q.Enqueue(Point{x: new_r, y: new_c})
					freshOrange -= 1
				}
			}
		}
		if q.Len() != 0 {
			cnt++
		}
	}
	if freshOrange > 0 {
		return -1
	}
	return cnt
}

// func main() {
// 	grid1 := [][]int{
// 		{2, 1, 1},
// 		{1, 1, 0},
// 		{0, 1, 1},
// 	}
// 	fmt.Println(orangesRotting(grid1)) // Output: 4

// 	grid2 := [][]int{
// 		{2, 1, 1},
// 		{0, 1, 1},
// 		{1, 0, 1},
// 	}
// 	fmt.Println(orangesRotting(grid2)) // Output: -1

// 	grid3 := [][]int{
// 		{0, 2},
// 	}
// 	fmt.Println(orangesRotting(grid3)) // Output: 0
// }

// package main

// import (
// 	"fmt"
// 	"sort"
// )

type ListNode struct {
	Val    int
	Next   *ListNode
	Bottom *ListNode
}

func merge(arr1 []int, n int, arr2 []int, m int) {
	i := n - 1
	j := m - 1
	k := len(arr1) - 1

	for i >= 0 && j >= 0 {
		if arr1[i] > arr2[j] {
			arr1[k] = arr1[i]
			i--
		} else {
			arr1[k] = arr2[j]
			j--
		}
		k--
	}
	for j >= 0 {
		arr1[k] = arr2[j]
		j--
		k--
	}
	arr1 = arr1[0 : n+m]
}

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[0]
	cnt := 1
	for slow != fast || cnt == 1 {
		slow = nums[slow]
		fast = nums[nums[fast]]
		cnt = 0
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

// func main() {
// 	// Test cases
// 	testCases := [][]int{
// 		{1, 3, 4, 2, 2},
// 		{3, 1, 3, 4, 2},
// 		{1, 1, 2},
// 		{1, 4, 3, 2, 4},
// 	}

// 	for _, testCase := range testCases {
// 		fmt.Printf("Array: %v, Duplicate: %d\n", testCase, findDuplicate(testCase))
// 	}
// }

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		for k := 0; k < n; k++ {
			matrix[k][i], matrix[k][j] = matrix[k][j], matrix[k][i]
		}

	}
}

// 1 4 7
// 2 5 8
// 3 6 9
// func main() {
// 	matrix := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}
// 	rotate(matrix)
// 	for _, row := range matrix {
// 		fmt.Println(row)
// 	}
// }

func generatePascalTriangle(numRows int) [][]int {
	triangle := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		tri := []int{}

		if i == 0 {
			tri = append(tri, 1)
		} else if i == 1 {
			tri = append(tri, 1, 1)
		} else {
			for j := 0; j <= i; j++ {
				if j == 0 || j == i {
					tri = append(tri, 1)
				} else {
					val := triangle[i-1][j] + triangle[i-1][j-1]
					tri = append(tri, val)
				}
			}
		}
		triangle = append(triangle, tri)
	}
	return triangle
}

// func main() {
// 	numRows := 5
// 	triangle := generatePascalTriangle(numRows)

// 	for _, row := range triangle {
// 		fmt.Println(row)
// 	}
// }

// func uniquePaths(m int, n int, dp [][]int) int {
// 	if m == 1 && n == 1 {
// 		return 1
// 	}
// 	if m <= 0 || n <= 0 {
// 		return 0
// 	}
// 	if dp[m][n] != -1 {
// 		return dp[m][n]
// 	}
// 	val := uniquePaths(m-1, n, dp) + uniquePaths(m, n-1, dp)
// 	dp[m][n] = val
// 	return val
// }

func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	for i := 1; i <= m; i++ {
		dp[i][1] = 1
	}
	for j := 1; j <= n; j++ {
		dp[1][j] = 1
	}
	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m][n]
}

// func main() {
// 	tests := []struct {
// 		m, n   int
// 		output int
// 	}{
// 		{3, 3, 6},
// 		{3, 7, 28},
// 		{1, 1, 1},
// 		{1, 10, 1},
// 		{10, 10, 48620},
// 		{19, 13, 86493225},
// 	}

// 	for _, test := range tests {
// 		dp := make([][]int, test.m+1)
// 		for i := 0; i < len(dp); i++ {
// 			dp[i] = make([]int, test.n+1)
// 			for j := range dp[i] {
// 				dp[i][j] = -1
// 			}
// 		}
// 		result := uniquePaths(test.m, test.n)
// 		if result == test.output {
// 			fmt.Printf("Test passed for grid %dx%d\n", test.m, test.n)
// 		} else {
// 			fmt.Printf("Test failed for grid %dx%d: expected %d, got %d\n", test.m, test.n, test.output, result)
// 		}
// 	}
// }

func myAtoi(s string) int {
	i, n := 0, len(s)
	for i < n && unicode.IsSpace(rune(s[i])) {
		i++
	}
	if i == n {
		return 0
	}
	sign := 1
	if s[i] == '+' || s[i] == '-' {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	val := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		val = val*10 + int(s[i]-'0')
		i++
	}
	return val * sign

}

// func main() {
// 	tests := []struct {
// 		input  string
// 		output int
// 	}{
// 		{"42", 42},
// 		{"   -42", -42},
// 		{"4193 with words", 4193},
// 		{"words and 987", 0},
// 		{"-91283472332", -91283472332},
// 		{"91283472332", 91283472332},
// 	}

// 	for _, test := range tests {
// 		result := myAtoi(test.input)
// 		if result == test.output {
// 			fmt.Printf("Test passed for input='%s'\n", test.input)
// 		} else {
// 			fmt.Printf("Test failed for input='%s': expected %d, got %d\n", test.input, test.output, result)
// 		}
// 	}
// }

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// func main() {
// 	tests := []struct {
// 		nums   []int
// 		output int
// 	}{
// 		{[]int{1, 1, 2}, 2},
// 		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
// 		{[]int{1, 1, 1, 1, 1}, 1},
// 		{[]int{1, 2, 3, 4, 5}, 5},
// 		{[]int{}, 0},
// 	}

// 	for _, test := range tests {
// 		result := removeDuplicates(test.nums)
// 		if result == test.output {
// 			fmt.Printf("Test passed for nums=%v\n", test.nums)
// 		} else {
// 			fmt.Printf("Test failed for nums=%v: expected %d, got %d\n", test.nums, test.output, result)
// 		}
// 	}
// }

func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0
	for _, num := range nums {
		if num == 1 {
			currentCount++
			maxCount = max(maxCount, currentCount)
		} else {
			currentCount = 0
		}
	}
	return maxCount

}

// func main() {
// 	tests := []struct {
// 		nums   []int
// 		output int
// 	}{
// 		{[]int{1, 1, 0, 1, 1, 1}, 3},
// 		{[]int{1, 0, 1, 1, 0, 1}, 2},
// 		{[]int{0, 0, 0, 0}, 0},
// 		{[]int{1, 1, 1, 1}, 4},
// 		{[]int{0, 1, 0, 1, 0, 1, 0}, 1},
// 	}

//		for _, test := range tests {
//			result := findMaxConsecutiveOnes(test.nums)
//			if result == test.output {
//				fmt.Printf("Test passed for nums=%v\n", test.nums)
//			} else {
//				fmt.Printf("Test failed for nums=%v: expected %d, got %d\n", test.nums, test.output, result)
//			}
//		}
//	}
func getValue(n, m int) int {
	val := 1
	for m > 0 {
		val = val * n
		m--
	}

	return val
}
func NthRoot(n, actualValue int) int {

	low := 1
	high := actualValue
	for low < high {
		mid := low + (high-low)/2
		rootValueMid := getValue(n, mid)
		if rootValueMid == actualValue {
			return mid
		} else if actualValue < rootValueMid {
			high = mid
		} else if actualValue > rootValueMid {
			low = mid + 1
		}
	}
	return -1
}

// func main() {
// 	testCases := []struct {
// 		n, m, expected int
// 	}{
// 		{3, 27, 3},
// 		{2, 16, 4},
// 		{2, 2, 1},
// 		{4, 69, -1}, // No integer root
// 		{3, 9, 2},
// 	}

// 	for _, tc := range testCases {
// 		result := NthRoot(tc.n, tc.m)
// 		if result == tc.expected {
// 			fmt.Printf("Test passed for NthRoot(%d, %d). Expected: %d, Got: %d\n", tc.n, tc.m, tc.expected, result)
// 		} else {
// 			fmt.Printf("Test failed for NthRoot(%d, %d). Expected: %d, Got: %d\n", tc.n, tc.m, tc.expected, result)
// 		}
// 	}
// }

// func findSingleElement(arr []int) int {
// 	low, high := 0, len(arr)-1
// 	for low < high {
// 		mid := (low + high) / 2

// 	}

// }
func isPossible(books []int, students int, currMin int) bool {
	currentSum := 0
	studentCount := 1

	for _, pages := range books {
		if pages > currMin {
			return false
		}
		if currentSum+pages > currMin {
			studentCount++
			currentSum = pages
			if studentCount > students {
				return false
			}
		} else {
			currentSum += pages
		}
	}
	return true
}

func findPages(books []int, students int) int {
	if len(books) < students {
		return -1
	}

	totalPages := 0
	maxPages := 0
	for _, pages := range books {
		totalPages += pages
		if pages > maxPages {
			maxPages = pages
		}
	}

	low := maxPages
	high := totalPages
	result := high

	for low <= high {
		mid := (low + high) / 2
		if isPossible(books, students, mid) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}

// func main() {
// 	testCases := []struct {
// 		books    []int
// 		students int
// 		expected int
// 	}{
// 		{[]int{12, 34, 67, 90}, 2, 113},
// 		{[]int{10, 20, 30, 40}, 2, 60},
// 		{[]int{10, 20, 30, 40, 50}, 2, 90},
// 		{[]int{5, 5, 5, 5, 5, 5, 5, 5}, 3, 15},
// 		{[]int{15, 20, 30}, 3, 30},
// 		{[]int{15, 20, 30}, 4, -1},         // more students than books
// 		{[]int{10, 10, 10, 10, 10}, 5, 10}, // equal books and students
// 	}

// 	for _, tc := range testCases {
// 		result := findPages(tc.books, tc.students)
// 		fmt.Printf("Books: %v, Students: %d, Expected: %d, Got: %d\n", tc.books, tc.students, tc.expected, result)
// 	}
// }

func canPlaceCows(stalls []int, cows int, minDist int) bool {
	count := 1
	lastPosition := stalls[0]

	for i := 1; i < len(stalls); i++ {
		if stalls[i]-lastPosition >= minDist {
			count++
			lastPosition = stalls[i]
			if count == cows {
				return true
			}
		}
	}

	return false
}

func findLargestMinDist(stalls []int, cows int) int {
	sort.Ints(stalls)

	low := 1
	high := stalls[len(stalls)-1] - stalls[0]
	result := 0

	for low < high {
		mid := (low + high) / 2
		if canPlaceCows(stalls, cows, mid) {
			result = mid
			low = mid + 1
		} else {
			high = mid
		}
	}

	return result
}

func nextSmallerElement(arr []int) []int {
	stk := stack.New()
	result := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		result[i] = -1
	}
	for i, num := range arr {

		for stk.Len() != 0 && arr[stk.Peek().(int)] > num {
			val := stk.Pop().(int)
			result[val] = num
		}
		stk.Push(i)
	}
	return result
}

type TreeNode struct {
	Val  int
	Next *TreeNode

	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalRecursive(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	inorderTraversalRecursive(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversalRecursive(root.Right, result)
}

func inorderTraversalIterative(root *TreeNode, result *[]int) {
	stk := []*TreeNode{}
	curr := root
	for len(stk) != 0 || curr != nil {
		for curr != nil {
			stk = append(stk, curr)
			curr = curr.Left
		}
		node := stk[len(stk)-1]
		*result = append(*result, node.Val)
		stk = stk[0 : len(stk)-1]
		curr = node.Right
	}
}

// func main() {
// 	tests := []struct {
// 		name     string
// 		tree     *TreeNode
// 		expected []int
// 	}{
// 		{
// 			name:     "Empty Tree",
// 			tree:     nil,
// 			expected: []int{},
// 		},
// 		{
// 			name:     "Single Node Tree",
// 			tree:     &TreeNode{Val: 1},
// 			expected: []int{1},
// 		},
// 		{
// 			name: "Left Skewed Tree",
// 			tree: &TreeNode{
// 				Val: 3,
// 				Left: &TreeNode{
// 					Val:  2,
// 					Left: &TreeNode{Val: 1},
// 				},
// 			},
// 			expected: []int{1, 2, 3},
// 		},
// 		{
// 			name: "Right Skewed Tree",
// 			tree: &TreeNode{
// 				Val: 1,
// 				Right: &TreeNode{
// 					Val:   2,
// 					Right: &TreeNode{Val: 3},
// 				},
// 			},
// 			expected: []int{1, 2, 3},
// 		},
// 		{
// 			name: "Balanced Tree",
// 			tree: &TreeNode{
// 				Val: 2,
// 				Left: &TreeNode{
// 					Val: 1,
// 				},
// 				Right: &TreeNode{
// 					Val: 3,
// 				},
// 			},
// 			expected: []int{1, 2, 3},
// 		},
// 	}

//		for _, tt := range tests {
//			t := tt // capture range variable
//			result1 := make([]int, 0)
//			inorderTraversalRecursive(t.tree, &result1)
//			result2 := make([]int, 0)
//			inorderTraversalIterative(t.tree, &result2)
//			fmt.Printf("Test %s - Recursive: %v\n", t.name, reflect.DeepEqual(result1, t.expected))
//		}
//	}

func preorderTraversalRecursive(root *TreeNode) []int {
	var result []int
	preorder(root, &result)
	return result
}

func preorder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preorder(node.Left, result)
	preorder(node.Right, result)
}

func preorderTraversalIterative(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	stk := stack.New()
	stk.Push(root)
	for stk.Len() != 0 {
		node := stk.Pop().(*TreeNode)
		result = append(result, node.Val)
		if node.Right != nil {
			stk.Push(node.Right)
		}
		if node.Left != nil {
			stk.Push(node.Left)
		}
	}
	return result
}

// func main() {
// 	tests := []struct {
// 		name     string
// 		tree     *TreeNode
// 		expected []int
// 	}{
// 		{
// 			name:     "Empty Tree",
// 			tree:     nil,
// 			expected: []int{},
// 		},
// 		{
// 			name:     "Single Node Tree",
// 			tree:     &TreeNode{Val: 1},
// 			expected: []int{1},
// 		},
// 		{
// 			name: "Left Skewed Tree",
// 			tree: &TreeNode{
// 				Val: 3,
// 				Left: &TreeNode{
// 					Val:  2,
// 					Left: &TreeNode{Val: 1},
// 				},
// 			},
// 			expected: []int{3, 2, 1},
// 		},
// 		{
// 			name: "Right Skewed Tree",
// 			tree: &TreeNode{
// 				Val: 1,
// 				Right: &TreeNode{
// 					Val:   2,
// 					Right: &TreeNode{Val: 3},
// 				},
// 			},
// 			expected: []int{1, 2, 3},
// 		},
// 		{
// 			name: "Balanced Tree",
// 			tree: &TreeNode{
// 				Val: 2,
// 				Left: &TreeNode{
// 					Val: 1,
// 				},
// 				Right: &TreeNode{
// 					Val: 3,
// 				},
// 			},
// 			expected: []int{2, 1, 3},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t := tt // capture range variable
// 		resultRecursive := preorderTraversalRecursive(t.tree)
// 		resultIterative := preorderTraversalIterative(t.tree)

// 		fmt.Printf("Test %s - Recursive: %v\n", t.name, reflect.DeepEqual(resultRecursive, t.expected))
// 		fmt.Printf("Test %s - Iterative: %v\n", t.name, reflect.DeepEqual(resultIterative, t.expected))
// 	}
// }

// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

func leftView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := queue.New()
	q.Enqueue(root)
	var result []int

	for q.Len() > 0 {
		n := q.Len()
		for i := 0; i < n; i++ {
			node := q.Dequeue().(*TreeNode)
			if i == 0 {
				result = append(result, node.Val)
			}
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
func leftViewDfs(root *TreeNode, data *[]int, level int) {
	if root == nil {
		return
	}
	if len(*data) == level {
		*data = append(*data, root.Val)
	}
	leftViewDfs(root.Left, data, level+1)
	leftViewDfs(root.Right, data, level+1)

}

func LeftViewTree(r *TreeNode) []int {
	data := make([]int, 0)
	leftViewDfs(r, &data, 0)
	return data
}
func rightView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	q := queue.New()
	q.Enqueue(root)
	var result []int

	for q.Len() > 0 {
		n := q.Len()
		for i := 0; i < n; i++ {
			node := q.Dequeue().(*TreeNode)
			if i == n-1 {
				result = append(result, node.Val)
			}
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

// func main() {
// 	// Create a binary tree
// 	root := &TreeNode{
// 		Val: 1,
// 		Left: &TreeNode{
// 			Val:   2,
// 			Left:  &TreeNode{Val: 4},
// 			Right: &TreeNode{Val: 5},
// 		},
// 		Right: &TreeNode{
// 			Val:   3,
// 			Left:  &TreeNode{Val: 6},
// 			Right: &TreeNode{Val: 7},
// 		},
// 	}

// 	// Test the right view function
// 	rightViewResult := rightView(root)
// 	fmt.Println("Right View:", rightViewResult) // Output: [1 3 7]

// 	// Test the left view function
// 	leftViewResult := LeftViewTree(root)
// 	fmt.Println("Left View:", leftViewResult) // Output: [1 2 4]
// }

func rightViewHelper(node *TreeNode, level int, maxLevel *int, result *[]int) {
	if node == nil {
		return
	}
	if level > *maxLevel {
		*result = append(*result, node.Val)
		*maxLevel = level
	}
	rightViewHelper(node.Right, level+1, maxLevel, result)
	rightViewHelper(node.Left, level+1, maxLevel, result)
}

// Function to get the right view
func rightViewa(root *TreeNode) []int {
	var result []int
	maxLevel := -1
	rightViewHelper(root, 0, &maxLevel, &result)
	return result
}

// Helper function for the left view
func leftViewHelper(node *TreeNode, level int, maxLevel *int, result *[]int) {
	if node == nil {
		return
	}
	if level > *maxLevel {
		*result = append(*result, node.Val)
		*maxLevel = level
	}
	leftViewHelper(node.Left, level+1, maxLevel, result)
	leftViewHelper(node.Right, level+1, maxLevel, result)
}

// Function to get the left view
func leftViews(root *TreeNode) []int {
	var result []int
	maxLevel := -1
	leftViewHelper(root, 0, &maxLevel, &result)
	return result
}

func findPath(root *TreeNode, target int) []int {
	result := []int{}
	if root == nil {
		return result
	}
	dfs(root, target, []int{}, &result)
	return result
}
func dfs(root *TreeNode, target int, data []int, result *[]int) {
	if root == nil {
		return
	}
	data = append(data, root.Val)
	if root.Val == target {
		//cool thing to note
		//Variadic Arguments in Go
		*result = append(*result, data...)
		return
	}
	dfs(root.Left, target, data, result)
	dfs(root.Right, target, data, result)
}

// func main() {
// 	// Test Case 1
// 	root1 := &TreeNode{Val: 1}
// 	root1.Left = &TreeNode{Val: 2}
// 	root1.Right = &TreeNode{Val: 3}
// 	root1.Left.Left = &TreeNode{Val: 4}
// 	root1.Left.Right = &TreeNode{Val: 5}
// 	root1.Right.Left = &TreeNode{Val: 6}
// 	root1.Right.Right = &TreeNode{Val: 7}

// 	fmt.Println(findPath(root1, 5)) // Output should be [1, 2, 5]

// 	// Test Case 2
// 	root2 := &TreeNode{Val: 1}
// 	root2.Left = &TreeNode{Val: 2}
// 	root2.Right = &TreeNode{Val: 3}
// 	root2.Left.Left = &TreeNode{Val: 4}
// 	root2.Left.Right = &TreeNode{Val: 5}

// 	fmt.Println(findPath(root2, 4)) // Output should be [1, 2, 4]

// 	// Test Case 3
// 	root3 := &TreeNode{Val: 1}
// 	root3.Left = &TreeNode{Val: 2}
// 	root3.Right = &TreeNode{Val: 3}
// 	root3.Left.Left = &TreeNode{Val: 4}
// 	root3.Left.Right = &TreeNode{Val: 5}

// 	fmt.Println(findPath(root3, 3)) // Output should be [1, 3]

// 	// Test Case 4
// 	root4 := &TreeNode{Val: 1}
// 	root4.Left = &TreeNode{Val: 2}
// 	root4.Right = &TreeNode{Val: 3}
// 	root4.Left.Left = &TreeNode{Val: 4}
// 	root4.Left.Right = &TreeNode{Val: 5}
// 	root4.Left.Left.Left = &TreeNode{Val: 8}

// 	fmt.Println(findPath(root4, 8)) // Output should be [1, 2, 4, 8]

// 	// Test Case 5
// 	root5 := &TreeNode{Val: 1}
// 	root5.Left = &TreeNode{Val: 2}
// 	root5.Right = &TreeNode{Val: 3}
// 	root5.Left.Right = &TreeNode{Val: 5}

// 	fmt.Println(findPath(root5, 10)) // Output should be []

// 	// Test Case 6: Single Node Tree
// 	root6 := &TreeNode{Val: 1}

// 	fmt.Println(findPath(root6, 1)) // Output should be [1]

// 	// Test Case 7: Left Skewed Tree
// 	root7 := &TreeNode{Val: 1}
// 	root7.Left = &TreeNode{Val: 2}
// 	root7.Left.Left = &TreeNode{Val: 3}
// 	root7.Left.Left.Left = &TreeNode{Val: 4}

// 	fmt.Println(findPath(root7, 4)) // Output should be [1, 2, 3, 4]

// 	// Test Case 8: Right Skewed Tree
// 	root8 := &TreeNode{Val: 1}
// 	root8.Right = &TreeNode{Val: 2}
// 	root8.Right.Right = &TreeNode{Val: 3}
// 	root8.Right.Right.Right = &TreeNode{Val: 4}

// 	fmt.Println(findPath(root8, 4)) // Output should be [1, 2, 3, 4]
// }

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxWidth := 0
	q := queue.New()
	q.Enqueue([]interface{}{root, 0})

	for q.Len() > 0 {
		levelSize := q.Len()
		var minIndex, maxIndex int
		for i := 0; i < levelSize; i++ {
			nodeIndexPair := q.Dequeue().([]interface{})
			node := nodeIndexPair[0].(*TreeNode)
			index := nodeIndexPair[1].(int)
			if i == 0 {
				minIndex = index
			}
			if i == levelSize-1 {
				maxIndex = index
			}
			if node.Left != nil {
				q.Enqueue([]interface{}{node.Left, 2 * index})
			}
			if node.Right != nil {
				q.Enqueue([]interface{}{node.Right, 2*index + 1})
			}
		}
		maxWidth = max(maxWidth, maxIndex-minIndex+1)
	}

	return maxWidth
}

// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 3}
// 	root.Right = &TreeNode{Val: 2}
// 	root.Left.Left = &TreeNode{Val: 5}
// 	root.Left.Right = &TreeNode{Val: 3}
// 	root.Right.Right = &TreeNode{Val: 9}

//		fmt.Println(widthOfBinaryTree(root)) // Output: 4
//	}
type NodeInfo struct {
	Node *TreeNode
	dx   int
}

func topView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	q := queue.New()
	q.Enqueue(NodeInfo{Node: root, dx: 0})
	hash := make(map[int]int)

	minHD, maxHD := 0, 0

	for q.Len() > 0 {
		n := q.Len()
		for i := 0; i < n; i++ {
			node := q.Dequeue().(NodeInfo)
			// Update min and max horizontal distances
			if node.dx < minHD {
				minHD = node.dx
			}
			if node.dx > maxHD {
				maxHD = node.dx
			}
			hash[node.dx] = node.Node.Val
			if node.Node.Left != nil {
				q.Enqueue(NodeInfo{Node: node.Node.Left, dx: node.dx - 1})
			}
			if node.Node.Right != nil {
				q.Enqueue(NodeInfo{Node: node.Node.Right, dx: node.dx + 1})
			}
		}
	}
	bottomView := []int{}
	for hd := minHD; hd <= maxHD; hd++ {
		bottomView = append(bottomView, hash[hd])
	}
	return bottomView
}

// func main() {
// 	root := &TreeNode{Val: 20}
// 	root.Left = &TreeNode{Val: 8}
// 	root.Right = &TreeNode{Val: 22}
// 	root.Left.Left = &TreeNode{Val: 5}
// 	root.Left.Right = &TreeNode{Val: 3}
// 	root.Right.Left = &TreeNode{Val: 4}
// 	root.Right.Right = &TreeNode{Val: 25}
// 	root.Left.Right.Left = &TreeNode{Val: 10}
// 	root.Left.Right.Right = &TreeNode{Val: 14}

// 	fmt.Println(topView(root)) // Output: [5 10 4 14 25]
// }

// Node represents a node in the binary search tree

func (n *Node) Insert(value int) {
	for n != nil {
		if value < n.value {
			if n.left == nil {
				n.left = &Node{value: value}
				return
			} else {
				n = n.left
			}
		} else {
			if n.right == nil {
				n.right = &Node{value: value}
				return
			} else {
				n = n.right
			}
		}
	}
}
func Search(n *Node, value int) int {
	if n == nil {
		return -1

	}
	if n.value == value {
		return 1
	}
	if value < n.value {
		return Search(n.left, value)
	}
	return Search(n.right, value)

}
func (n *Node) InOrderPrint() {
	if n == nil {
		return
	}
	n.left.InOrderPrint()
	fmt.Print(n.value, " ")
	n.right.InOrderPrint()
}

// func main() {
// 	root := &Node{value: 10}
// 	root.Insert(5)
// 	root.Insert(15)
// 	root.Insert(2)
// 	root.Insert(7)
// 	root.Insert(12)
// 	root.Insert(20)

// 	root.InOrderPrint()
// 	fmt.Println()
// 	fmt.Println(Search(root, 25))
// }

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func diameter(root *TreeNode, val *int) int {
	if root == nil {
		return 0
	}
	left := diameter(root.Left, val)
	right := diameter(root.Right, val)
	max_value := max(left, right) + 1
	*val = max(*val, left+right+1)
	return max_value
}

// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 3}
// 	root.Left.Left = &TreeNode{Val: 4}
// 	root.Left.Right = &TreeNode{Val: 5}
// 	root.Right.Left = &TreeNode{Val: 6}
// 	root.Right.Right = &TreeNode{Val: 7}
// 	fmt.Println("Height of the tree:", height(root))
// 	val := 0
// 	diameter(root, &val)
// 	fmt.Println("Diameter of binary tree:", val)
// }

// heightAndBalance returns the height of the tree and whether it is balanced
func heightAndBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftHeight, isLeftBalanced := heightAndBalance(root.Left)
	if !isLeftBalanced {
		return -1, false
	}
	rightHeight, isRightBalanced := heightAndBalance(root.Right)
	if !isRightBalanced {
		return -1, false
	}
	if abs(leftHeight-rightHeight) > 1 {
		return -1, false
	}
	return max(leftHeight, rightHeight) + 1, true
}

// abs returns the absolute value of an integer
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isBalanced(root *TreeNode) bool {
	_, balanced := heightAndBalance(root)
	return balanced
}

// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 3}
// 	root.Left.Left = &TreeNode{Val: 4}
// 	root.Left.Right = &TreeNode{Val: 5}
// 	root.Right.Left = &TreeNode{Val: 6}
// 	root.Right.Right = &TreeNode{Val: 7}

// 	fmt.Println("Is the tree height-balanced?", isBalanced(root))
// }

func isIdentical(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isIdentical(root1.Left, root2.Left) && isIdentical(root1.Right, root2.Right)

}

// func main() {
// 	// Creating first binary tree
// 	root1 := &TreeNode{Val: 1}
// 	root1.Left = &TreeNode{Val: 2}
// 	root1.Right = &TreeNode{Val: 3}
// 	root1.Left.Left = &TreeNode{Val: 4}
// 	root1.Left.Right = &TreeNode{Val: 5}
// 	root1.Right.Left = &TreeNode{Val: 6}
// 	root1.Right.Right = &TreeNode{Val: 7}

// 	// Creating second binary tree
// 	root2 := &TreeNode{Val: 1}
// 	root2.Left = &TreeNode{Val: 2}
// 	root2.Right = &TreeNode{Val: 3}
// 	root2.Left.Left = &TreeNode{Val: 4}
// 	root2.Left.Right = &TreeNode{Val: 5}
// 	root2.Right.Left = &TreeNode{Val: 5}
// 	root2.Right.Right = &TreeNode{Val: 7}

// 	fmt.Println("Are the two trees identical?", isIdentical(root1, root2))
// }

func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int

	if root == nil {
		return result
	}
	q := queue.New()
	forward := true
	q.Enqueue(root)
	for q.Len() > 0 {
		n := q.Len()
		data := make([]int, n)
		j := 0
		if !forward {
			j = n - 1
		}
		for i := 0; i < n; i++ {
			node := q.Dequeue().(*TreeNode)
			data[j] = node.Val
			if forward {
				j++
			} else {
				j--
			}
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
		result = append(result, data)
		forward = !forward
	}
	return result
}

// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 3}
// 	root.Left.Left = &TreeNode{Val: 4}
// 	root.Left.Right = &TreeNode{Val: 5}
// 	root.Right.Left = &TreeNode{Val: 6}
// 	root.Right.Right = &TreeNode{Val: 7}

// 	result := zigzagLevelOrder(root)
// 	for _, level := range result {
// 		fmt.Println(level)
// 	}
// }

func maxSumPath(root *TreeNode) int {
	val := math.MinInt32
	maxSum(root, &val)
	return val
}

func maxSum(root *TreeNode, val *int) int {
	if root == nil {
		return 0
	}
	left := max(0, maxSum(root.Left, val))
	right := max(0, maxSum(root.Right, val))
	*val = max(*val, left+right+root.Val)
	return max(left, right) + root.Val
}

//     1

//  2     3

// 4  5  6  7
// func main() {
// 	root := &TreeNode{Val: -10}
// 	root.Left = &TreeNode{Val: 9}
// 	root.Right = &TreeNode{Val: 20}
// 	root.Right.Left = &TreeNode{Val: 15}
// 	root.Right.Right = &TreeNode{Val: 7}

// 	fmt.Println(maxSumPath(root)) // Output: 42
// }

func findIndexWhereElemtisgrr(arr []int, val int) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, something := range arr {
		if something > val {
			right = append(right, something)
		} else {
			left = append(left, something)
		}
	}
	return left, right
}
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	left, right := findIndexWhereElemtisgrr(preorder[1:], preorder[0])
	root.Left = bstFromPreorder(left)
	root.Right = bstFromPreorder(right)
	return root
}

// func main() {
// 	preorder := []int{8, 5, 1, 7, 10, 12}
// 	root := bstFromPreorder(preorder)
// 	fmt.Println(root.Val)
// }

func numberOfSubarrays(arr []int) {
	start := 0
	answer := 0
	N := len(arr)

	for start < N {
		mp := make(map[int]bool)
		mp[0] = true
		end := start
		sum := 0

		for end < N {
			sum += arr[end]
			if mp[sum] {
				break
			}
			mp[sum] = true
			end++
		}

		fmt.Println(end, start)
		answer += end - start
		start++
	}

	fmt.Println(answer)
}

// func main() {
// 	arr := []int{10, 10, 10}
// 	numberOfSubarrays(arr)
// }

// Result structure to store information returned by the helper function.
type Result struct {
	IsBST bool
	Size  int
	Min   int
	Max   int
}

// Helper function to find the largest BST.
func largestBSTSubtreeHelper(node *TreeNode, maxTreeSize *int) Result {
	if node == nil {
		return Result{IsBST: true, Size: 0, Min: math.MaxInt, Max: math.MinInt}
	}
	leftResult := largestBSTSubtreeHelper(node.Left, maxTreeSize)
	rightResult := largestBSTSubtreeHelper(node.Right, maxTreeSize)
	if leftResult.IsBST && rightResult.IsBST && node.Val > leftResult.Max && node.Val < rightResult.Min {

		*maxTreeSize = max(*maxTreeSize, leftResult.Size+rightResult.Size+1)
		fmt.Println(*maxTreeSize, leftResult.Size+rightResult.Size+1)
		return Result{IsBST: true, Size: leftResult.Size + rightResult.Size + 1, Min: min(leftResult.Min, node.Val), Max: max(rightResult.Max, node.Val)}
	}
	return Result{IsBST: false}
}

// Function to find the size of the largest BST in a binary tree.
func largestBSTSubtree(root *TreeNode) int {
	maxTreeSize := 0
	largestBSTSubtreeHelper(root, &maxTreeSize)
	return maxTreeSize
}

// Example usage:
// func main() {
// 	// Test Case 1: Empty tree
// 	// var root1 *TreeNode
// 	// fmt.Println("Test Case 1: Size of the largest BST is:", largestBSTSubtree(root1)) // Output: 0

// 	// // Test Case 2: All nodes form a valid BST
// 	// /*
// 	//    Constructing the following binary tree:
// 	//        2
// 	//       / \
// 	//      1   3
// 	// */
// 	// root2 := &TreeNode{Val: 2}
// 	// root2.Left = &TreeNode{Val: 1}
// 	// root2.Right = &TreeNode{Val: 3}
// 	// fmt.Println("Test Case 2: Size of the largest BST is:", largestBSTSubtree(root2)) // Output: 3

// 	// // Test Case 3: No subtree forms a valid BST
// 	// /*
// 	//    Constructing the following binary tree:
// 	//        10
// 	//       /  \
// 	//      15   5
// 	// */
// 	// root3 := &TreeNode{Val: 10}
// 	// root3.Left = &TreeNode{Val: 15}
// 	// root3.Right = &TreeNode{Val: 5}
// 	// fmt.Println("Test Case 3: Size of the largest BST is:", largestBSTSubtree(root3)) // Output: 1

// 	// // Test Case 4: The largest BST is not the whole tree
// 	// /*
// 	//    Constructing the following binary tree:
// 	//        10
// 	//       /  \
// 	//      5    15
// 	//     / \     \
// 	//    1   8     7
// 	// */
// 	// root4 := &TreeNode{Val: 10}
// 	// root4.Left = &TreeNode{Val: 5}
// 	// root4.Right = &TreeNode{Val: 15}
// 	// root4.Left.Left = &TreeNode{Val: 1}
// 	// root4.Left.Right = &TreeNode{Val: 8}
// 	// root4.Right.Right = &TreeNode{Val: 7}
// 	// fmt.Println("Test Case 4: Size of the largest BST is:", largestBSTSubtree(root4)) // Output: 3

// 	// Test Case 5: Mixed tree with different BST subtrees
// 	/*
// 	   Constructing the following binary tree:
// 	        10
// 	      /    \
// 	     5      20
// 	    / \    /  \
// 	   1   8  15  25
// 	              /
// 	             22
// 	*/
// 	root5 := &TreeNode{Val: 10}
// 	root5.Left = &TreeNode{Val: 5}
// 	root5.Right = &TreeNode{Val: 20}
// 	root5.Left.Left = &TreeNode{Val: 1}
// 	root5.Left.Right = &TreeNode{Val: 8}
// 	root5.Right.Left = &TreeNode{Val: 15}
// 	root5.Right.Right = &TreeNode{Val: 25}
// 	root5.Right.Right.Left = &TreeNode{Val: 22}
// 	fmt.Println("Test Case 5: Size of the largest BST is:", largestBSTSubtree(root5)) // Output: 8 (subtree rooted at node 20)
// }

// lowestCommonAncestor finds the lowest common ancestor of two nodes in a BST.

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

// func main() {
// 	// Creating a sample binary tree:
// 	//         3
// 	//       /   \
// 	//      5     1
// 	//     / \   / \
// 	//    6   2 0   8
// 	//       / \
// 	//      7   4
// 	root := &TreeNode{Val: 3}
// 	root.Left = &TreeNode{Val: 5}
// 	root.Right = &TreeNode{Val: 1}
// 	root.Left.Left = &TreeNode{Val: 6}
// 	root.Left.Right = &TreeNode{Val: 2}
// 	root.Right.Left = &TreeNode{Val: 0}
// 	root.Right.Right = &TreeNode{Val: 8}
// 	root.Left.Right.Left = &TreeNode{Val: 7}
// 	root.Left.Right.Right = &TreeNode{Val: 4}

// 	p := root.Left             // Node with value 5
// 	q := root.Left.Right.Right // Node with value 4

// 	lca := lowestCommonAncestor(root, p, q)
// 	if lca != nil {
// 		fmt.Printf("The LCA of node %d and node %d is node %d\n", p.Val, q.Val, lca.Val)
// 	} else {
// 		fmt.Println("LCA not found")
// 	}
// }

type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

var head, prev *DoublyListNode

func treeToDoublyList(root *TreeNode) *DoublyListNode {
	if root == nil {
		return nil
	}

	inOrder(root)
	return head
}

func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	inOrder(node.Left)
	newNode := &DoublyListNode{Val: node.Val}
	if prev == nil {
		head = newNode
		head.Prev = newNode
		head.Next = newNode
	} else {
		temp := prev.Next
		newNode.Prev = prev
		newNode.Prev.Next = newNode
		newNode.Next = temp

	}
	prev = newNode
	inOrder(node.Right)
}

func printDoublyList(head *DoublyListNode) {
	if head == nil {
		return
	}

	fmt.Print("Doubly Linked List in Forward Order: ")
	curr := head
	for {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
		if curr == head {
			break
		}
	}
	fmt.Println()

}

// 1 -> <- 2

// func main() {
// 	// Creating a sample binary tree:
// 	//         4
// 	//       /   \
// 	//      2     5
// 	//     / \
// 	//    1   3
// 	root := &TreeNode{Val: 4}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 5}
// 	root.Left.Left = &TreeNode{Val: 1}
// 	root.Left.Right = &TreeNode{Val: 3}
// 	head := treeToDoublyList(root)
// 	printDoublyList(head)
// }

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	var s string = "1"

	for i := 2; i <= n; i++ {
		sb := strings.Builder{}
		j, k := 0, 0
		for j <= k && k < len(s) {
			for k < len(s) && s[j] == s[k] {
				k++
			}
			count := k - j
			value := s[j]
			sb.WriteString(fmt.Sprintf("%d%c", count, value))
			fmt.Println(sb)
			j = k
		}
		s = sb.String()
	}
	return s
}

// func main() {
// 	n := 6

// 	fmt.Println(countAndSay(n))
// }

func dfsG(node int, adj map[int][]int, vis []bool, ls *[]int) {
	vis[node] = true
	*ls = append(*ls, node)
	for _, value := range adj[node] {
		if !vis[value] {
			dfsG(value, adj, vis, ls)
		}
	}
}
func dfsOfGraph(node int, adj map[int][]int) []int {
	visited := make([]bool, node)
	fmt.Println(visited)
	ls := make([]int, 0)
	for i := 0; i < node; i++ {
		if !visited[i] {
			dfsG(i, adj, visited, &ls)
		}
	}
	return ls
}
func bfsGraph(node int, adj map[int][]int) []int {
	visited := make([]bool, node)
	ls := make([]int, 0)
	q := queue.New()
	for i := 0; i < node; i++ {
		if !visited[i] {
			q.Enqueue(i)
			for q.Len() != 0 {
				data := q.Dequeue().(int)
				visited[data] = true
				ls = append(ls, data)
				for _, val := range adj[data] {
					if !visited[val] {
						q.Enqueue(val)
					}
				}
			}
		}
	}
	return ls
}

func addEdge(adj map[int][]int, u, v int) {
	if _, exist := adj[u]; !exist {
		adj[u] = make([]int, 0)
	}
	if _, exist := adj[v]; !exist {
		adj[v] = make([]int, 0)
	}
	adj[u] = append(adj[u], v)
	adj[v] = append(adj[v], u)
}

func printAns(ans []int) {
	for _, val := range ans {
		fmt.Print(val, " ")
	}
	fmt.Println()
}

// func main() {
// 	adj := make(map[int][]int)

// 	addEdge(adj, 0, 2)
// 	addEdge(adj, 2, 4)
// 	addEdge(adj, 0, 1)
// 	addEdge(adj, 0, 3)
// 	fmt.Println(adj)

// 	ans := bfsGraph(5, adj)
// 	printAns(ans)
// }

type GraphNode struct {
	Val    int
	Parent int
}

func bfsGraphCycleDetect(node int, adj map[int][]int) bool {
	visited := make([]bool, node)
	q := queue.New()
	for i := 0; i < node; i++ {
		if !visited[i] {
			q.Enqueue(GraphNode{Val: i, Parent: -1})
			for q.Len() != 0 {
				node := q.Dequeue().(GraphNode)
				visited[node.Val] = true
				for _, val := range adj[node.Val] {
					if !visited[val] {
						visited[val] = true
						q.Enqueue(GraphNode{Val: val, Parent: node.Val})
					} else if val != node.Parent {
						return true
					}
				}
			}
		}
	}
	return false
}
func dfsCycleGraph(node int, parent int, adj map[int][]int, visited []bool) bool {
	visited[node] = true
	for _, adjacentNode := range adj[node] {
		if !visited[adjacentNode] {
			if dfsCycleGraph(adjacentNode, node, adj, visited) {
				return true
			}
		} else if parent != adjacentNode {
			return true
		}
	}
	return false
}
func dfsGraphCycleDetect(node int, adj map[int][]int) bool {
	visited := make([]bool, node)

	for i := 0; i < node; i++ {
		if !visited[i] {
			if dfsCycleGraph(i, -1, adj, visited) {
				return true
			}
		}
	}
	return false
}

// func main() {
// 	// Test case 1: Graph with no cycles
// 	adj1 := map[int][]int{
// 		0: {1},
// 		1: {0, 2},
// 		2: {1, 3},
// 		3: {2},
// 	}
// 	fmt.Println("Test Case 1 (No cycle):", dfsGraphCycleDetect(4, adj1)) // Expected: false

// 	// Test case 2: Graph with a simple cycle
// 	adj2 := map[int][]int{
// 		0: {1, 2},
// 		1: {0, 2},
// 		2: {0, 1},
// 	}
// 	fmt.Println("Test Case 2 (Simple cycle):", dfsGraphCycleDetect(3, adj2)) // Expected: true

// 	// Test case 3: Disconnected graph with one component containing a cycle
// 	adj3 := map[int][]int{
// 		0: {1},
// 		1: {0, 2},
// 		2: {1, 3},
// 		3: {2},
// 		4: {5, 6},
// 		5: {4, 6},
// 		6: {4, 5},
// 	}
// 	fmt.Println("Test Case 3 (Disconnected with one cycle):", dfsGraphCycleDetect(7, adj3)) // Expected: true

// 	// Test case 4: Graph with multiple cycles
// 	adj4 := map[int][]int{
// 		0: {1, 2},
// 		1: {0, 2},
// 		2: {0, 1, 3},
// 		3: {2, 4},
// 		4: {3, 5},
// 		5: {4, 6},
// 		6: {5},
// 		7: {8, 9},
// 		8: {7, 9},
// 		9: {7, 8},
// 	}
// 	fmt.Println("Test Case 4 (Multiple cycles):", dfsGraphCycleDetect(10, adj4)) // Expected: true

// 	// Test case 5: Empty graph
// 	adj5 := map[int][]int{}
// 	fmt.Println("Test Case 5 (Empty graph):", dfsGraphCycleDetect(0, adj5)) // Expected: false
// }

func detectCycleInDirectedGraph(numNodes int, adj map[int][]int) bool {
	inDegreeNodes := make([]int, numNodes)
	for _, adjacentNodes := range adj {
		for _, adjacentNode := range adjacentNodes {
			inDegreeNodes[adjacentNode]++
		}
	}
	q := queue.New()
	for _, node := range inDegreeNodes {
		if inDegreeNodes[node] == 0 {
			q.Enqueue(node)
		}
	}
	cnt := 0
	for q.Len() != 0 {
		cnt++
		node := q.Dequeue().(int)
		for _, adjacentNode := range adj[node] {
			inDegreeNodes[adjacentNode]--
			if inDegreeNodes[adjacentNode] == 0 {
				q.Enqueue(adjacentNode)
			}
		}
	}
	return cnt != numNodes
}
func dfsCycleDirected(node int, adj map[int][]int, visited, recStack []bool) bool {
	visited[node] = true
	recStack[node] = true

	for _, adjacentNode := range adj[node] {
		if !visited[adjacentNode] {
			if dfsCycleDirected(adjacentNode, adj, visited, recStack) {
				return true
			}
		} else if recStack[adjacentNode] {
			return true
		}
	}

	recStack[node] = false
	return false
}

func detectCycleInDGGraph(numNodes int, adj map[int][]int) bool {
	visited := make([]bool, numNodes)
	recStack := make([]bool, numNodes)

	for i := 0; i < numNodes; i++ {
		if !visited[i] {
			if dfsCycleDirected(i, adj, visited, recStack) {
				return true
			}
		}
	}
	return false
}

// func main() {
// 	adj := map[int][]int{
// 		0: {1},
// 		1: {2},
// 		2: {0, 3},
// 		3: {4},
// 		4: {5},
// 		5: {3},
// 	}
// 	numNodes := 6
// 	fmt.Println(detectCycleInDirectedGraph(numNodes, adj)) // Output: true
// }

func topologicalSort(numNodes int, adj map[int][]int) ([]int, bool) {
	inDegree := make([]int, numNodes)
	for _, neighbors := range adj {
		for _, node := range neighbors {
			inDegree[node]++
		}
	}

	q := queue.New()
	for i := 0; i < numNodes; i++ {
		if inDegree[i] == 0 {
			q.Enqueue(i)
		}
	}

	var topoOrder []int
	count := 0
	for q.Len() > 0 {
		node := q.Dequeue().(int)
		topoOrder = append(topoOrder, node)
		count++

		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				q.Enqueue(neighbor)
			}
		}
	}

	if count != numNodes {
		return nil, false // graph has a cycle
	}
	return topoOrder, true // successful topological sort
}

//	func main() {
//		adj := map[int][]int{
//			0: {1},
//			1: {2},
//			2: {3},
//			3: {},
//		}
//		numNodes := 4
//		order, isDAG := topologicalSort(numNodes, adj)
//		if isDAG {
//			fmt.Println("Topological Order:", order)
//		} else {
//			fmt.Println("Graph has a cycle")
//		}
//	}
var dx []int = []int{-1, 0, 1, 0}
var dy []int = []int{0, 1, 0, -1}

func dfsNumberOfIsland(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	for i := 0; i < 4; i++ {
		new_x := x + dx[i]
		new_y := y + dy[i]
		if Isvalid(new_x, new_y, grid) {
			dfsNumberOfIsland(grid, new_x, new_y)
		}
	}
}

func Isvalid(new_x int, new_y int, grid [][]byte) bool {
	n := len(grid)
	m := len(grid[0])
	if new_x < 0 || new_x >= n || new_y < 0 || new_y >= m || grid[new_x][new_y] == '0' {
		return false
	}
	return true
}
func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	cnt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				cnt++
				dfsNumberOfIsland(grid, i, j)
			}
		}
	}
	return cnt
}

func distinctNumberOfIsland(grid [][]int, x, y int, origin_x, origin_y int, sb *strings.Builder) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return
	}
	grid[x][y] = 0
	sb.WriteString(strconv.Itoa(x - origin_x))
	sb.WriteString(strconv.Itoa(y - origin_y))
	for i := 0; i < 4; i++ {
		new_x := x + dx[i]
		new_y := y + dy[i]

		distinctNumberOfIsland(grid, new_x, new_y, origin_x, origin_y, sb)

	}
}

func numDistinctIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	uniqueIsland := make(map[string]bool)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sb := &strings.Builder{}
				distinctNumberOfIsland(grid, i, j, i, j, sb)
				// fmt.Println(sb.String())
				uniqueIsland[sb.String()] = true
			}
		}
	}
	return len(uniqueIsland)
}
func IsvalidCheck(new_x int, new_y int, grid [][]int) bool {
	n := len(grid)
	m := len(grid[0])
	if new_x < 0 || new_x >= n || new_y < 0 || new_y >= m || grid[new_x][new_y] == '0' {
		return false
	}
	return true
}

// func main() {
// 	tests := []struct {
// 		grid     [][]int
// 		expected int
// 	}{
// 		{
// 			grid: [][]int{
// 				{1, 1, 0, 0, 0},
// 				{1, 1, 0, 0, 0},
// 				{0, 0, 0, 1, 1},
// 				{0, 0, 0, 1, 1},
// 			},
// 			expected: 1,
// 		},
// 		{
// 			grid: [][]int{
// 				{1, 1, 0, 1, 1},
// 				{1, 0, 0, 0, 1},
// 				{0, 0, 0, 1, 0},
// 				{1, 1, 0, 1, 1},
// 			},
// 			expected: 4,
// 		},
// 		{
// 			grid: [][]int{
// 				{1, 0, 0, 1},
// 				{0, 0, 0, 0},
// 				{0, 0, 1, 1},
// 				{0, 1, 1, 0},
// 			},
// 			expected: 3,
// 		},
// 		{
// 			grid: [][]int{
// 				{1, 1, 1},
// 				{1, 0, 1},
// 				{1, 1, 1},
// 			},
// 			expected: 1,
// 		},
// 		{
// 			grid: [][]int{
// 				{1, 1, 0, 0, 0},
// 				{0, 1, 0, 0, 1},
// 				{1, 0, 0, 1, 1},
// 				{0, 0, 0, 0, 0},
// 				{1, 1, 0, 1, 1},
// 			},
// 			expected: 2,
// 		},
// 		{
// 			grid: [][]int{
// 				{0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0},
// 				{0, 0, 0, 0, 0},
// 			},
// 			expected: 0,
// 		},
// 	}

// 	for _, test := range tests {
// 		result := numDistinctIslands(test.grid)
// 		if result != test.expected {
// 			fmt.Printf("For grid %v, expected %d but got %d", test.grid, test.expected, result)
// 		}
// 	}
// }

func changeRecursive(amount int, coins []int, startIndex int, dp [][]int) int {
	if amount == 0 {
		return 1
	}
	if amount < 0 {
		return 0
	}
	if dp[startIndex][amount] != -1 {
		return dp[startIndex][amount]
	}
	sum := 0
	for i := startIndex; i < len(coins); i++ {
		sum += changeRecursive(amount-coins[i], coins, i, dp)
	}
	dp[startIndex][amount] = sum
	return sum
}

func change(amount int, coins []int) int {
	dp := make([][]int, len(coins))
	for i := range dp {
		dp[i] = make([]int, amount+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return changeRecursive(amount, coins, 0, dp)
}

// func main() {
// 	amount := 5
// 	coins := []int{1, 2, 5}
// 	fmt.Println("Number of combinations:", change(amount, coins))
// }

// func recursiveLIS(nums []int, prev, startIndex, count int) int {
// 	if startIndex == len(nums) {
// 		return count
// 	}
// 	valueToreturn := count
// 	for i := startIndex; i < len(nums); i++ {
// 		if nums[i] > prev {
// 			valueToreturn = max(valueToreturn, recursiveLIS(nums, nums[i], i+1, count+1))
// 		}
// 	}
// 	return valueToreturn
// }
// func lengthOfLIS(nums []int) int {
// 	return recursiveLIS(nums, math.MinInt, 0, 0)

// }

func recursiveLIS(nums []int, prevIndex, startIndex int, dp [][]int) int {
	if startIndex == len(nums) {
		return 0
	}
	if dp[prevIndex+1][startIndex] != -1 {
		return dp[prevIndex+1][startIndex]
	}
	valueToreturn := 0
	for i := startIndex; i < len(nums); i++ {
		if prevIndex == -1 || nums[i] > nums[prevIndex] {
			valueToreturn = max(valueToreturn, 1+recursiveLIS(nums, i, i+1, dp))
		}
	}
	dp[prevIndex+1][startIndex] = valueToreturn
	return valueToreturn
}
func lengthOfLIS(nums []int) int {
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, len(nums)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return recursiveLIS(nums, -1, 0, dp)
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
// 	fmt.Println("Length of Longest Increasing Subsequence:", lengthOfLIS(nums))
// }

// Function to determine if there's a subset of nums that sums up to target

func subsetSumRecursive(nums []int, target, startIndex int, dp [][]int) bool {
	if startIndex >= len(nums) {
		return false
	}
	if target < 0 {
		return false
	}
	if target == 0 {
		return true
	}
	return subsetSumRecursive(nums, target, startIndex+1, dp) || subsetSumRecursive(nums, target-nums[startIndex], startIndex+1, dp)
}
func subsetSum(nums []int, target int) bool {
	dp := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]int, target+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return subsetSumRecursive(nums, target, 0, dp)
}

// func main() {
// 	nums := []int{1, 2, 3, 7}
// 	target := 6
// 	fmt.Println("Is there a subset with the given sum?", subsetSum(nums, target)) // Output: true

// 	nums = []int{1, 2, 7, 1, 5}
// 	target = 10
// 	fmt.Println("Is there a subset with the given sum?", subsetSum(nums, target)) // Output: true

// 	nums = []int{1, 3, 4, 8}
// 	target = 6
// 	fmt.Println("Is there a subset with the given sum?", subsetSum(nums, target)) // Output: false
// }

// Function to determine the maximum obtainable value by cutting up the rod

func RecursiveRodCutting(prices []int, n, startIndex int, dp [][]int) int {
	if n == 0 {
		return 0
	}
	if n < 0 {
		return math.MinInt
	}
	if dp[n][startIndex] != -1 {
		return dp[n][startIndex]
	}
	sum := 0
	for i := startIndex; i < n; i++ {
		sum = max(sum, prices[i]+RecursiveRodCutting(prices, n-(i+1), i, dp))
	}
	dp[n][startIndex] = sum
	return sum
}
func rodCutting(prices []int, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, len(prices)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return RecursiveRodCutting(prices, n, 0, dp)
}

// func main() {
// 	// Test case 1
// 	prices1 := []int{1, 5, 8, 9, 10, 17, 17, 20}
// 	n1 := 8
// 	fmt.Printf("Maximum obtainable value for prices1: %d\n", rodCutting(prices1, n1)) // Output: 22

// 	// Test case 2: Single Element Price List
// 	prices2 := []int{5}
// 	n2 := 1
// 	fmt.Printf("Maximum obtainable value for prices2: %d\n", rodCutting(prices2, n2)) // Output: 5

// 	// Test case 3: All Equal Prices
// 	prices3 := []int{2, 2, 2, 2}
// 	n3 := 4
// 	fmt.Printf("Maximum obtainable value for prices3: %d\n", rodCutting(prices3, n3)) // Output: 8

// 	// Test case 4: Increasing Prices
// 	prices4 := []int{1, 3, 7, 8, 9}
// 	n4 := 5
// 	fmt.Printf("Maximum obtainable value for prices4: %d\n", rodCutting(prices4, n4)) // Output: 10

// 	// Test case 5: Decreasing Prices
// 	prices5 := []int{8, 7, 6, 5}
// 	n5 := 4
// 	fmt.Printf("Maximum obtainable value for prices5: %d\n", rodCutting(prices5, n5)) // Output: 32

// 	// Test case 6: Large Length with Random Prices
// 	prices6 := []int{3, 5, 8, 9, 10, 17, 17, 20, 24, 30}
// 	n6 := 10
// 	fmt.Printf("Maximum obtainable value for prices6: %d\n", rodCutting(prices6, n6)) // Output: 30
// }

func RecursiveEditDistance(word1, word2 string, i, j int, dp [][]int) int {
	if i == len(word1) {
		return len(word2) - j
	}
	if j == len(word2) {
		return len(word1) - i
	}
	if dp[i][j] != -1 {
		return dp[i][j]
	}
	if word1[i] == word2[j] {
		dp[i][j] = RecursiveEditDistance(word1, word2, i+1, j+1, dp)
		return dp[i][j]
	}
	dp[i][j] = 1 + minThree(RecursiveEditDistance(word1, word2, i+1, j+1, dp), RecursiveEditDistance(word1, word2, i, j+1, dp), RecursiveEditDistance(word1, word2, i+1, j, dp))
	return dp[i][j]
}

func minThree(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < c {
		return b
	}
	return c
}

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return RecursiveEditDistance(word1, word2, 0, 0, dp)
}

// Helper function to find the minimum of three integers

// func main() {
// 	// Test cases
// 	word1 := "horse"
// 	word2 := "ros"
// 	fmt.Printf("Edit Distance between '%s' and '%s': %d\n", word1, word2, minDistance(word1, word2)) // Output: 3

// 	word1 = "intention"
// 	word2 = "execution"
// 	fmt.Printf("Edit Distance between '%s' and '%s': %d\n", word1, word2, minDistance(word1, word2)) // Output: 5

// 	word1 = "abc"
// 	word2 = "yabd"
// 	fmt.Printf("Edit Distance between '%s' and '%s': %d\n", word1, word2, minDistance(word1, word2)) // Output: 2

// 	word1 = "sunday"
// 	word2 = "saturday"
// 	fmt.Printf("Edit Distance between '%s' and '%s': %d\n", word1, word2, minDistance(word1, word2)) // Output: 3
// }

// Serialize converts a tree to a string
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
// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 3}
// 	root.Right.Left = &TreeNode{Val: 4}
// 	root.Right.Right = &TreeNode{Val: 5}

// 	serialized := Serialize(root)
// 	fmt.Println("Serialized:", serialized) //"1,2,#,#,3,4,#,#,5,#,#"

// }

func DistinctNumbersInWindow(nums []int, k int) []int {
	if len(nums) < k {
		return []int{}
	}
	frequencyMap := make(map[int]int)
	for i := 0; i < k; i++ {
		frequencyMap[nums[i]]++
	}
	i := 0
	j := k - 1
	result := make([]int, 0)
	for j < len(nums) {
		result = append(result, len(frequencyMap))
		if j == len(nums)-1 {
			break
		}
		if frequencyMap[nums[i]] == 1 {
			delete(frequencyMap, nums[i])
		} else {
			frequencyMap[nums[i]] -= 1
		}
		i++
		j++
		frequencyMap[nums[j]] += 1
	}
	return result
}

// Example usage
// func main() {
// 	nums := []int{1, 2, 1, 3, 4, 2, 3}
// 	k := 4
// 	result := DistinctNumbersInWindow(nums, k)
// 	fmt.Printf("Distinct numbers in each window: %v\n", result)
// }

func InOrderTraversal(root *TreeNode, elems *[]int) {
	if root == nil {
		return
	}
	InOrderTraversal(root.Left, elems)
	*elems = append(*elems, root.Val)
	InOrderTraversal(root.Right, elems)
}

// FindPairWithSum finds a pair with a given sum in the BST
func FindPairWithSum(root *TreeNode, target int) (int, int, bool) {
	elems := []int{}
	InOrderTraversal(root, &elems)

	left, right := 0, len(elems)-1

	for left < right {
		sum := elems[left] + elems[right]
		if sum == target {
			return elems[left], elems[right], true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return 0, 0, false
}

// // Example usage
// func main() {
// 	root := &TreeNode{Val: 5}
// 	root.Left = &TreeNode{Val: 3}
// 	root.Right = &TreeNode{Val: 7}
// 	root.Left.Left = &TreeNode{Val: 2}
// 	root.Left.Right = &TreeNode{Val: 4}
// 	root.Right.Left = &TreeNode{Val: 6}
// 	root.Right.Right = &TreeNode{Val: 8}

// 	target := 9
// 	a, b, found := FindPairWithSum(root, target)
// 	if found {
// 		fmt.Printf("Pair found: (%d, %d)\n", a, b)
// 	} else {
// 		fmt.Println("No pair found")
// 	}
// }

func checkChildrenSumTree(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	leftVal := 0
	if root.Left != nil {
		leftVal = root.Left.Val
	}
	rightVal := 0
	if root.Right != nil {
		rightVal = root.Right.Val
	}
	return (root.Val == leftVal+rightVal) && checkChildrenSumTree(root.Left) && checkChildrenSumTree(root.Right)
}

func flatten(root *TreeNode) {
	var prev *TreeNode = nil
	Flattern(root, &prev)
}
func Flattern(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}
	Flattern(root.Right, prev)
	Flattern(root.Left, prev)
	root.Right = *prev
	root.Left = nil
	*prev = root
}

// Helper function to print the flattened tree
func printFlattenedTree(root *TreeNode) {
	for root != nil {
		fmt.Print(root.Val, " -> ")
		root = root.Right
	}
	fmt.Println("nil")
}

// // Example usage
// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 5}
// 	root.Left.Left = &TreeNode{Val: 3}
// 	root.Left.Right = &TreeNode{Val: 4}
// 	root.Right.Right = &TreeNode{Val: 6}

// 	flatten(root)

//		fmt.Print("Flattened tree: ")
//		printFlattenedTree(root)
//	}
func setZeroes(matrix [][]int) {
	n, m := len(matrix), len(matrix[0])
	rowMat := 1
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			rowMat = 0
		}
	}
	calMat := 1
	for j := 0; j < m; j++ {
		if matrix[0][j] == 0 {
			calMat = 0
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowMat == 0 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
	if calMat == 0 {
		for j := 0; j < m; j++ {
			matrix[0][j] = 0
		}
	}

}

// func main() {
// 	matrix := [][]int{
// 		{1, 1, 1},
// 		{1, 0, 1},
// 		{1, 1, 1},
// 	}

// 	setZeroes(matrix)
// 	fmt.Println(matrix) // Output: [[1 0 1] [0 0 0] [1 0 1]]
// }

func findRepeatAndMissing(arr []int) (int, int) {
	n := len(arr)
	sumN := n * (n + 1) / 2
	sumSqN := n * (n + 1) * (2*n + 1) / 6

	sum := 0
	sumSq := 0

	for _, num := range arr {
		sum += num
		sumSq += num * num
	}

	// Differences between expected and actual sums
	diffSum := sumN - sum       // x - y
	diffSumSq := sumSqN - sumSq // x^2 - y^2

	// x + y = diffSumSq / diffSum
	sumXY := diffSumSq / diffSum

	// Solving for x and y
	missing := (diffSum + sumXY) / 2
	repeat := sumXY - missing

	return repeat, missing
}

// func main() {
// 	arr := []int{4, 3, 6, 2, 1, 1}
// 	repeat, missing := findRepeatAndMissing(arr)
// 	fmt.Printf("Repeated Number: %d, Missing Number: %d\n", repeat, missing)
// }

func majorityElement(nums []int) int {
	maxNum := nums[0]
	count := 1
	for _, num := range nums[1:] {
		if num == maxNum {
			count++
		} else if count == 0 {
			maxNum = num
			count = 1
		} else {
			count--
		}
	}
	return maxNum
}

// func main() {
// 	arr := []int{2, 2, 1, 1, 1, 2, 2}
// 	majority := majorityElement(arr)
// 	if majority != -1 {
// 		fmt.Printf("Majority Element: %d\n", majority)
// 	} else {
// 		fmt.Println("No Majority Element found.")
// 	}
// }

// func majorityElementN3(nums []int) []int {
// 	count1, count2, number1, number2 := 0, 0, -1, -1
// 	for _, num := range nums {
// 		if num == number1 {
// 			count1++
// 		} else if num == number2 {
// 			count2++
// 		} else if count1 == 0 {
// 			number1 = num
// 			count1++
// 		} else if count2 == 0 {
// 			number2 = num
// 			count2++
// 		} else {
// 			count1--
// 			count2--
// 		}
// 	}
// }

// // func main() {
// // 	nums := []int{3, 2, 3, 4, 4, 4, 2, 4, 4}
// // 	fmt.Println(majorityElementN3(nums)) // Output: [4]
// // }

func myPow(x float64, n int) float64 {
	dp := make(map[int]float64, 0)
	if n < 0 {
		return 1 / myPowDp(x, -n, dp)
	}
	return myPowDp(x, n, dp)
}
func myPowDp(x float64, n int, dp map[int]float64) float64 {
	if dp[n] != 0 {
		return dp[n]
	}
	if n == 0 {
		return float64(1)
	}
	if n%2 != 0 {
		value := float64(x) * myPow(x, n-1)
		dp[n] = value
		return value
	}
	value := myPow(x, n/2) * myPow(x, n/2)
	dp[n] = value
	return value
}

// func main() {
// 	fmt.Println(myPow(2.0, 10)) // Output: 1024.0
// 	fmt.Println(myPow(2.0, -2)) // Output: 0.25
// 	fmt.Println(myPow(2.1, 3))  // Output: 9.261
// }

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix)
	m := len(matrix[0])

	left := 0
	right := n*m - 1
	for left < right {
		mid := (left + right) / 2
		if matrix[mid/n][mid%m] == target {
			return true
		} else if matrix[mid/m][mid%m] < target {
			left = mid + 1
		} else {
			right = mid
		}

	}
	return matrix[left/m][left%m] == target
}

// func main() {
// 	matrix := [][]int{
// 		{1, 2},
// 	}

// 	fmt.Println(searchMatrix(matrix, 1)) // Output: true
// 	fmt.Println(searchMatrix(matrix, 5)) // Output: true
// }

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	len := 0
	cur := head
	prev := cur
	for cur != nil {
		len = len + 1
		prev = cur
		cur = cur.Next
	}
	k = k % len
	prev.Next = head
	curr := head
	for i := 1; i <= len-k-1; i++ {
		curr = curr.Next
	}
	dat := curr.Next
	curr.Next = nil
	return dat
}

// func main() {
// 	// Create a linked list: 1 -> 2 -> 3 -> 4 -> 5
// 	head := &ListNode{Val: 1}
// 	head.Next = &ListNode{Val: 2}
// 	head.Next.Next = &ListNode{Val: 3}
// 	head.Next.Next.Next = &ListNode{Val: 4}
// 	head.Next.Next.Next.Next = &ListNode{Val: 5}
// 	// Rotate the list by 2 places
// 	head = rotateRight(head, 2)

// 	// Print the rotated list
// 	printList(head) // Output: 4 -> 5 -> 1 -> 2 -> 3 -> nil
// }

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		maxArea = max(maxArea, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// func main() {
// 	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
// 	fmt.Println(maxArea(height)) // Output: 49
// }

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	leftMaxHeight := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		leftMaxHeight[i] = max(leftMaxHeight[i-1], height[i-1])
	}
	rightMaxHeight := make([]int, len(height))
	for i := len(height) - 2; i >= 0; i-- {
		rightMaxHeight[i] = max(rightMaxHeight[i+1], height[i+1])
	}

	waterTrapped := 0

	for i := 0; i < len(height); i++ {
		waterTrapped += max(0, min(leftMaxHeight[i], rightMaxHeight[i])-height[i])
	}
	return waterTrapped
}

// func main() {
// 	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
// 	fmt.Println(trap(height))
// }

func findContentChildren(g []int, s []int) int {
	sort.Slice(g, func(i, j int) bool {
		return g[i] < g[j]
	})
	sort.Ints(s)

	child, cookie := 0, 0

	for child < len(g) && cookie < len(s) {
		corrChild := g[child]
		corrCookie := s[cookie]
		if corrCookie >= corrChild {
			child++
			cookie++
		} else {
			cookie++
		}

	}
	return child

}

// func main() {
// 	g := []int{4, 5, 6}
// 	s := []int{1, 1}
// 	fmt.Println(findContentChildren(g, s)) // Output: 1

// 	g = []int{1, 2}
// 	s = []int{1, 2, 3}
// 	fmt.Println(findContentChildren(g, s)) // Output: 2
// }

func minCoins(coins []int, amount int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	coinCount := 0

	for _, num := range coins {
		if amount >= num {
			coinCount = coinCount + amount/num
			amount = amount % num
		}
	}
	if amount == 0 {
		return coinCount
	}
	return -1
}

// func main() {
// 	coins := []int{25, 10, 5, 1}
// 	amount := 49
// 	fmt.Println(minCoins(coins, amount)) // Output: 7

// 	coins = []int{9, 6, 5, 1}
// 	amount = 11
// 	fmt.Println(minCoins(coins, amount)) // Output: 3
// }

func connect(root *TreeNode) *TreeNode {
	q := queue.New()
	q.Enqueue(root)
	for q.Len() > 0 {
		n := q.Len()
		prev := &TreeNode{}
		for i := 0; i < n; i++ {
			node := q.Dequeue().(*TreeNode)
			prev.Next = node
			prev = prev.Next
			if node.Left != nil {
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
	}
	return root
}
func conneckt(root *TreeNode) *TreeNode {
	q := make([]*TreeNode, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		prev := &TreeNode{}
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			prev.Next = node
			prev = prev.Next
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return root
}

// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 3}
// 	root.Left.Left = &TreeNode{Val: 4}
// 	root.Left.Right = &TreeNode{Val: 5}
// 	root.Right.Right = &TreeNode{Val: 7}

//		conneckt(root)
//		fmt.Println(root.Left.Next.Val)       // Output: 3
//		fmt.Println(root.Left.Left.Next.Val)  // Output: 5
//		fmt.Println(root.Left.Right.Next.Val) // Output: 7
//		fmt.Println(root.Right.Right.Next)    // Output: <nil>
//	}
type Node struct {
	Val   int
	value int
	left  *Node
	right *Node
	Next  *Node
	Prev  *Node
	Child *Node
	next  *Node
	prev  *Node
	key   int
}

func printList(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
func flattenK(head *Node) *Node {
	curr := head
	for curr != nil {
		if curr.Child != nil {
			nextElem := curr.Child
			curr.Child = nil
			lastElem := nextElem
			for lastElem.Next != nil {
				lastElem = lastElem.Next
			}
			if curr.Next != nil {
				curr.Next.Prev = lastElem
				lastElem.Next = curr.Next
			}
			curr.Next = nextElem
			nextElem.Prev = curr
		}
		curr = curr.Next
	}
	return head
}

// 1 <-> 2 <-> 3 <-> 4
//
//		|
//		7 <-> 8 <-> 9
//	    	  |
//	     	 11 <-> 12
// func main() {
// 	// Example creation of the multilevel doubly linked list structure
// 	head := &Node{Val: 1}
// 	head.Next = &Node{Val: 2}
// 	head.Next.Prev = head
// 	head.Next.Next = &Node{Val: 3}
// 	head.Next.Next.Prev = head.Next
// 	head.Next.Next.Next = &Node{Val: 4}
// 	head.Next.Next.Next.Prev = head.Next.Next

// 	head.Next.Child = &Node{Val: 7}
// 	head.Next.Child.Next = &Node{Val: 8}
// 	head.Next.Child.Next.Prev = head.Next.Child
// 	head.Next.Child.Next.Child = &Node{Val: 11}
// 	head.Next.Child.Next.Child.Next = &Node{Val: 12}
// 	head.Next.Child.Next.Child.Next.Prev = head.Next.Child.Next.Child
// 	head.Next.Child.Next.Next = &Node{Val: 9}
// 	head.Next.Child.Next.Next.Prev = head.Next.Child.Next

// 	flattenedHead := flattenK(head)
// 	printList(flattenedHead)
// }

func calculateSpan(prices []int) []int {
	stk := stack.New()
	span := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		for stk.Len() > 0 && prices[stk.Peek().(int)] < prices[i] {
			stk.Pop()
		}
		if stk.Len() == 0 {
			span[i] = i + 1
		} else {
			span[i] = i - stk.Peek().(int)
		}
		stk.Push(i)
	}
	return span
}

// func main() {
// 	prices := []int{100, 80, 60, 70, 60, 75, 85}
// 	span := calculateSpan(prices)
// 	fmt.Println(span) // Output: [1, 1, 1, 2, 1, 4, 6]
// }

func insertSorted(s *stack.Stack, val int) {
	if s.Len() == 0 || s.Peek().(int) < val {
		s.Push(val)
		return
	}
	temp := s.Pop().(int)
	insertSorted(s, val)
	s.Push(temp)
}

func sortStack(s *stack.Stack) {
	if s.Len() == 0 {
		return
	}
	temp := s.Pop().(int)
	sortStack(s)
	insertSorted(s, temp)
}

// // 3 1 4 2
// // 1 3 4
// func main() {
// 	s := stack.New()
// 	s.Push(3)
// 	s.Push(1)
// 	s.Push(4)
// 	s.Push(2)

// 	sortStack(s)

// 	for s.Len() > 0 {
// 		fmt.Println(s.Pop())
// 	}
// }

type Stack struct {
	q *queue.Queue
}

func NewStack() *Stack {
	return &Stack{q: queue.New()}
}

func (s *Stack) Push(val int) {
	len := s.q.Len()
	s.q.Enqueue(val)
	for i := 0; i < len; i++ {
		val = s.q.Dequeue().(int)
		s.q.Enqueue(val)
	}
}

func (s *Stack) Pop() int {
	if s.q.Len() == 0 {
		return -1
	}
	val := s.q.Dequeue().(int)
	return val
}

func (s *Stack) Top() int {
	if s.q.Len() == 0 {
		return -1
	}
	return s.q.Peek().(int)
}

func (s *Stack) IsEmpty() bool {
	return s.q.Len() == 0

}

//  1
//   1 2
//   1 2 3
//  1 2  3 4

//  2 1

//  2 1 3 2 1

//  4 3 2 1 5

// func main() {
// 	stack := NewStack()
// 	stack.Push(1)
// 	stack.Push(2)
// 	stack.Push(3)

// 	fmt.Println(stack.Pop()) // Output: 3
// 	fmt.Println(stack.Top()) // Output: 2
// 	fmt.Println(stack.Pop()) // Output: 2
// 	fmt.Println(stack.Pop()) // Output: 1
// }

type Queue struct {
	s1 *stack.Stack
	s2 *stack.Stack
}

func NewQueue() *Queue {
	return &Queue{
		s1: stack.New(),
		s2: stack.New(),
	}
}

func (q *Queue) Enqueue(val int) {
	for q.s1.Len() > 0 {
		q.s2.Push(q.s1.Pop().(int))
	}
	q.s1.Push(val)
	for q.s2.Len() > 0 {
		q.s1.Push(q.s2.Pop().(int))
	}
}

func (q *Queue) Dequeue() int {
	return q.s1.Pop().(int)
}

func (q *Queue) Front() int {
	return q.s1.Peek().(int)
}

func (q *Queue) IsEmpty() bool {
	return q.s1.Len() == 0 && q.s2.Len() == 0
}

// func main() {
// 	queue := NewQueue()
// 	queue.Enqueue(1)
// 	queue.Enqueue(2)
// 	queue.Enqueue(3)

//		fmt.Println(queue.Dequeue()) // Output: 1
//		fmt.Println(queue.Front())   // Output: 2
//		fmt.Println(queue.Dequeue()) // Output: 2
//		fmt.Println(queue.Dequeue()) // Output: 3
//	}
type queueData struct {
	X    int
	Y    int
	node *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q := make([]queueData, 0)
	mp := make(map[int]map[int][]int, 0)
	q = append(q, queueData{X: 0, Y: 0, node: root})
	for len(q) > 0 {
		dat := q[0]
		q = q[1:]
		_, exist := mp[dat.X]
		if !exist {
			mp[dat.X] = make(map[int][]int)
		}

		mp[dat.X][dat.Y] = append(mp[dat.X][dat.Y], dat.node.Val)
		if dat.node.Left != nil {
			q = append(q, queueData{X: dat.X - 1, Y: dat.Y + 1, node: dat.node.Left})
		}
		if dat.node.Right != nil {
			q = append(q, queueData{X: dat.X + 1, Y: dat.Y + 1, node: dat.node.Right})

		}
	}
	keySorted := []int{}
	for key := range mp {
		keySorted = append(keySorted, key)
	}
	sort.Slice(keySorted, func(i, j int) bool {
		return keySorted[i] < keySorted[j]
	})
	verticalTree := make([][]int, 0)
	for _, key := range keySorted {
		da := mp[key]
		innerKey := []int{}
		for key := range da {
			innerKey = append(innerKey, key)
		}
		sort.Slice(innerKey, func(i, j int) bool {
			return innerKey[i] < innerKey[j]
		})
		arrData := []int{}
		for _, key := range innerKey {
			sort.Slice(da[key], func(i, j int) bool {
				return da[key][i] < da[key][j]
			})
			arrData = append(arrData, da[key]...)
		}
		verticalTree = append(verticalTree, arrData)
	}
	return verticalTree
}

// func main() {
// 	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 3}
// 	root.Left.Left = &TreeNode{Val: 4}
// 	root.Left.Right = &TreeNode{Val: 5}
// 	root.Right.Left = &TreeNode{Val: 6}
// 	root.Right.Right = &TreeNode{Val: 7}

// 	verticalOrder(root)
// }

func leftTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := queue.New()
	q.Enqueue(root)
	lef := make([]int, 0)
	for q.Len() > 0 {
		node := q.Dequeue().(*TreeNode)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Left != nil && node.Right != nil {
			lef = append(lef, node.Val)
		}
	}
	return lef
}
func LeafNode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	q := queue.New()
	q.Enqueue(root)
	bou := make([]int, 0)
	for q.Len() > 0 {
		node := q.Dequeue().(*TreeNode)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
		if node.Left == nil && node.Right == nil {
			bou = append(bou, node.Val)
		}
	}
	return bou
}
func boundaryTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	boundaryTraversal := make([]int, 0)
	boundaryTraversal = append(boundaryTraversal, leftTraversal(root)...)
	boundaryTraversal = append(boundaryTraversal, LeafNode(root)...)
	return boundaryTraversal
}

// func main() {
// 	root := &TreeNode{Val: 1}
// 	root.Left = &TreeNode{Val: 2}
// 	root.Right = &TreeNode{Val: 3}
// 	root.Left.Left = &TreeNode{Val: 4}
// 	root.Left.Right = &TreeNode{Val: 5}
// 	root.Right.Left = &TreeNode{Val: 6}
// 	root.Right.Right = &TreeNode{Val: 7}
// 	root.Left.Right.Left = &TreeNode{Val: 8}
// 	root.Left.Right.Right = &TreeNode{Val: 9}

// 	fmt.Println(boundaryTraversal(root)) // Output: [1 2 4 8 9 6 7 3]
// }

// Input:Inorder: [40, 20, 50, 10, 60, 30], Preorder: [10, 20, 40, 50, 30, 60]

func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int, 0)
	for index, val := range inorder {
		inorderMap[val] = index
	}
	pStartIndex, pEndIndex := 0, len(preorder)-1
	iStartIndex, iEndIndex := 0, len(inorder)-1
	return buildTreeRec(preorder, inorder, inorderMap, pStartIndex, pEndIndex, iStartIndex, iEndIndex)
}

// preorder := []int{3, 9, 20, 15, 7}
//
//	inorder := []int{9, 3, 15, 20, 7}
func buildTreeRec(preorder []int, inorder []int, inorderMap map[int]int, pStartIndex, pEndIndex, iStartIndex, iEndIndex int) *TreeNode {
	if pStartIndex > pEndIndex || iStartIndex > iEndIndex {
		return nil
	}
	root := preorder[pStartIndex]
	inorderPlace := inorderMap[root]

	node := &TreeNode{Val: root}
	lenLeft := inorderPlace - iStartIndex
	node.Left = buildTreeRec(preorder, inorder, inorderMap, pStartIndex+1, pStartIndex+lenLeft, iStartIndex, inorderPlace)
	node.Right = buildTreeRec(preorder, inorder, inorderMap, pStartIndex+lenLeft+1, pEndIndex, inorderPlace+1, iEndIndex)
	return node
}

// func buildTree(preorder []int, inorder []int) *TreeNode {
// 	if len(inorder) == 0 || len(preorder) == 0 {
// 		return nil
// 	}
// 	var index int
// 	for i, key := range inorder {
// 		if key == preorder[0] {
// 			index = i
// 			break
// 		}
// 	}
// 	inorderLeft := inorder[0:index]
// 	inorderRight := inorder[index+1:]
// 	preorderLeft := preorder[1 : len(inorderLeft)+1]
// 	preorderRight := preorder[len(inorderLeft)+1:]
// 	return &TreeNode{Val: preorder[0], Left: buildTree(preorderLeft, inorderLeft), Right: buildTree(preorderRight, inorderRight)}
// }

func preorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	preorderTraversal(root.Left)
	preorderTraversal(root.Right)
}

func inorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversal(root.Left)
	fmt.Printf("%d ", root.Val)
	inorderTraversal(root.Right)
}

// func main() {
// 	preorder := []int{3, 9, 20, 15, 7}
// 	inorder := []int{9, 3, 15, 20, 7}

// 	root := buildTree(preorder, inorder)

// 	fmt.Println("Preorder of constructed tree:")
// 	preorderTraversal(root)

//		fmt.Println("\nInorder of constructed tree:")
//		inorderTraversal(root)
//	}
func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}
	val := 0
	for i := len(s) - 1; i >= 0; i-- {
		if i != len(s)-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			val = val - romanMap[s[i]]
		} else {
			val = val + romanMap[s[i]]
		}
	}
	return val
}

// func main() {
// 	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
// }

func longestPalindrome(s string) string {
	ans := ""
	for i := 0; i < len(s); i++ {
		ans1 := expandAroundCenter(s, i, i)
		ans2 := expandAroundCenter(s, i, i+1)
		if len(ans) < len(ans1) {
			ans = ans1
		}
		if len(ans) < len(ans2) {
			ans = ans2
		}
	}
	return ans
}
func expandAroundCenter(s string, i, j int) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	return s[i+1 : j]
}

// func main() {
// 	s := "babad"
// 	fmt.Println(longestPalindrome(s)) // Output: "bab" or "aba"
// }

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = commonPrefix(prefix, strs[i])
		if prefix == "" {
			break
		}
	}

	return prefix
}

func commonPrefix(str1, str2 string) string {
	minLen := min(len(str1), len(str2))

	for i := 0; i < minLen; i++ {
		if str1[i] != str2[i] {
			return str1[:i]
		}
	}

	return str1[:minLen]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	strs := []string{"flower", "flow", "flight"}
// 	fmt.Println(longestCommonPrefix(strs)) // Output: "fl"
// }

func climbStairs(n int) int {
	dp := make([]int, n+1)
	return climbStair(n, dp)
}
func climbStair(n int, dp []int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = climbStair(n-1, dp) + climbStair(n-2, dp)
	return dp[n]
}

// func main() {
// 	// Test with 5 steps

// 	fmt.Println(climbStairs(45)) // Output: 8
// }

func robRec(nums []int) int {
	dp := make([]int, len(nums)+1)
	return maxSumNonAdjacent(nums, len(nums)-1, dp)
}
func robIter(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	dp := make([]int, len(nums)+1)
	dp[0] = nums[0]
	dp[1] = nums[1]
	for i := 2; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[len(nums)-1]
}
func maxSumNonAdjacent(nums []int, n int, dp []int) int {
	if n < 0 {
		return 0
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = max(nums[n]+maxSumNonAdjacent(nums, n-2, dp), maxSumNonAdjacent(nums, n-1, dp))
	return dp[n]
}

// func main() {
// 	// Example: array with values [3, 2, 5, 10, 7]
// 	nums := []int{3, 2, 5, 10, 7}
// 	fmt.Println(rob(nums)) // Output: 15 (3 + 5 + 7)
// }

func asteroidCollision(asteroids []int) []int {
	stack := []int{}

	for _, asteroid := range asteroids {
		collide := true

		// Process the asteroid
		for collide && asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			top := stack[len(stack)-1]

			// Compare the top of the stack with the current asteroid
			if top < -asteroid {
				// Stack asteroid is smaller, pop it and continue checking
				stack = stack[:len(stack)-1]
			} else if top == -asteroid {
				// Both asteroids are of the same size, pop the stack and destroy both
				stack = stack[:len(stack)-1]
				collide = false
			} else {
				// Stack asteroid is larger, destroy the current asteroid
				collide = false
			}
		}

		// If no collision or asteroid is moving right, push it to the stack
		if collide {
			stack = append(stack, asteroid)
		}
	}

	return stack
}

// func main() {
// 	asteroids := []int{5, 10, -5}
// 	fmt.Println(asteroidCollision(asteroids)) // Output: [5, 10]

// 	asteroids2 := []int{8, -8}
// 	fmt.Println(asteroidCollision(asteroids2)) // Output: []

// 	asteroids3 := []int{10, 2, -5}
// 	fmt.Println(asteroidCollision(asteroids3)) // Output: [10]

//		asteroids4 := []int{-2, -1, 1, 2}
//		fmt.Println(asteroidCollision(asteroids4)) // Output: [-2, -1, 1, 2]
//	}'
//
// abacabac

// func PrintA(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	toggle := true
// 	for {
// 		<-chA
// 		fmt.Println("a")
// 		if toggle {
// 			chB <- true
// 		} else {
// 			chC <- true
// 		}
// 		toggle = !toggle
// 	}
// }

// func PrintB(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-chB:
// 			fmt.Println("b")
// 			chA <- true
// 		case <-exit1:
// 			fmt.Println("exitb")
// 			return
// 		}
// 	}
// }

// func PrintC(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for {
// 		select {
// 		case <-chC:
// 			fmt.Println("c")
// 			chA <- true
// 		case <-exit1:
// 			fmt.Println("exitc")
// 			return
// 		}
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	chA = make(chan bool)
// 	chB = make(chan bool)
// 	chC = make(chan bool)
// 	exit1 = make(chan bool)

// 	wg.Add(3)
// 	go PrintA(&wg)
// 	go PrintB(&wg)
// 	go PrintC(&wg)

// 	// Start the first goroutine by sending a value to chA
// 	chA <- true

//		// Wait for all goroutines to complete
//		wg.Wait()
//	}
func worker1(ch chan string) {
	for {
		time.Sleep(1 * time.Second) // Simulate work
		ch <- "Worker 1 completed task"
	}
}

// Worker 2 simulates completing a task in 2 seconds
func worker2(ch chan string) {
	for {
		time.Sleep(2 * time.Second) // Simulate work
		ch <- "Worker 2 completed task"
	}
}

type PrintStruct struct {
	chA   chan bool
	chB   chan bool
	chC   chan bool
	exit1 chan bool
}

func NewPrintStruct() *PrintStruct {
	return &PrintStruct{
		chA: make(chan bool),
		chB: make(chan bool),
		chC: make(chan bool),
	}

}

func (p *PrintStruct) PrintA(ch chan string) {
	toggle := true
	for {
		<-p.chA
		ch <- "0"
		if toggle {
			p.chB <- true
		} else {
			p.chC <- true
		}
		toggle = !toggle
	}
}

func (p *PrintStruct) PrintB(ch chan string) {
	for {
		<-p.chB
		ch <- "1"
		p.chA <- true
	}
}

func (p *PrintStruct) PrintC(ch chan string) {
	for {
		<-p.chC
		ch <- "2"
		p.chA <- true
	}
}

// func main() {
// 	// Create channels for worker 1 and worker 2
// 	ch1 := make(chan string)
// 	pr := NewPrintStruct()
// 	// Launch the workers
// 	go pr.PrintA(ch1)
// 	go pr.PrintB(ch1)
// 	go pr.PrintC(ch1)

// 	pr.chA <- true
// 	// Simulate receiving tasks and handling results
// 	for i := 0; i < 6; i++ {
// 		select {
// 		case result := <-ch1:
// 			fmt.Println(result)
// 		}
// 	}
// }

type PrintFuzz struct {
	ch1 chan bool
	ch2 chan bool
	ch3 chan bool
	ch4 chan bool
}

func NewPrintFuzz() *PrintFuzz {
	return &PrintFuzz{
		ch1: make(chan bool),
		ch2: make(chan bool),
		ch3: make(chan bool),
		ch4: make(chan bool),
	}

}

func (p *PrintFuzz) PrintFuz(ch chan string) {
	for {
		<-p.ch1 //block
		ch <- "fuzz"
		p.ch4 <- true
	}
}

func (p *PrintFuzz) PrintBuzz(ch chan string) {
	for {
		<-p.ch2
		ch <- "buzz"
		p.ch4 <- true
	}
}

func (p *PrintFuzz) Printfizzbuzz(ch chan string) {
	for {
		<-p.ch3
		ch <- "fuzzbuzz"
		p.ch4 <- true
	}
}
func (p *PrintFuzz) PrintNumber(ch chan string) {
	i := 1
	<-p.ch4
	for {
		if i%3 == 0 && i%5 == 0 {
			p.ch3 <- true
			<-p.ch4
		} else if i%5 == 0 {
			p.ch2 <- true
			<-p.ch4
		} else if i%3 == 0 {
			p.ch1 <- true
			<-p.ch4
		} else {
			ch <- fmt.Sprintf("%d", i)
		}
		i++
	}
}

// func main() {
// 	// Create channels for worker 1 and worker 2
// 	ch1 := make(chan string)
// 	pr := NewPrintFuzz()
// 	// Launch the workers
// 	go pr.PrintFuz(ch1)
// 	go pr.PrintBuzz(ch1)
// 	go pr.Printfizzbuzz(ch1)
// 	go pr.PrintNumber(ch1)

// 	pr.ch4 <- true
// 	for i := 0; i < 15; i++ {
// 		select {
// 		case result := <-ch1:
// 			fmt.Printf("%s ", result)
// 		}
// 	}
// }

type PrintH2O struct {
	ch1 chan bool
	ch2 chan bool
	mu  *sync.Mutex
}

func NewPrintH2O() *PrintH2O {
	return &PrintH2O{
		ch1: make(chan bool),
		ch2: make(chan bool),
		mu:  &sync.Mutex{},
	}
}

func (p *PrintH2O) PrintHydrogen(ch chan string) {
	cnt := 0
	for {
		if cnt == 2 {
			<-p.ch1
		}
		ch <- "H"
		cnt++
	}
}

func (p *PrintH2O) PrintOxygen(ch chan string) {
	cnt := 0
	for {
		if cnt == 1 {
			<-p.ch2
		}
		ch <- "0"
		cnt++
	}
}

// func main() {
// 	// Create channels for worker 1 and worker 2
// 	ch1 := make(chan string)
// 	pr := NewPrintH2O()
// 	go pr.PrintHydrogen(ch1)
// 	go pr.PrintOxygen(ch1)
// 	for i := 0; i < 3; i++ {
// 		select {
// 		case result := <-ch1:
// 			fmt.Printf("%s ", result)
// 		}
// 	}
// }

// type PrintH2O struct {
// 	ch1 chan bool
// 	ch2 chan bool
// 	mu  *sync.Mutex
// }

// func NewPrintH2O() *PrintH2O {
// 	return &PrintH2O{
// 		ch1: make(chan bool),
// 		ch2: make(chan bool),
// 		mu:  &sync.Mutex{},
// 	}
// }

// func (p *PrintH2O) PrintHydrogen(ch chan string) {
// 	cnt := 0
// 	for {
// 		if cnt == 2 {
// 			<-p.ch1
// 		}
// 		ch <- "H"
// 		cnt++
// 	}
// }

// func (p *PrintH2O) PrintOxygen(ch chan string) {
// 	cnt := 0
// 	for {
// 		if cnt == 1 {
// 			<-p.ch2
// 		}
// 		ch <- "0"
// 		cnt++
// 	}
// }

//	func main() {
//		// Create channels for worker 1 and worker 2
//		ch1 := make(chan string)
//		pr := NewPrintH2O()
//		go pr.PrintHydrogen(ch1)
//		go pr.PrintOxygen(ch1)
//		for i := 0; i < 3; i++ {
//			select {
//			case result := <-ch1:
//				fmt.Printf("%s ", result)
//			}
//		}
//	}
type someMinMax struct {
	Min int
	Max int
}

func findMax(arr []int, ch chan someMinMax) {
	curr_max := arr[0]
	for _, x := range arr {
		curr_max = max(curr_max, x)
	}
	ch <- someMinMax{Max: curr_max}
}

// func main() {
// 	ch := make(chan someMinMax)
// 	arr := []int{1, 9, 2, 3, 4, 5, 6}

// 	go findMax(arr[0:3], ch)
// 	go findMax(arr[3:6], ch)
// 	go findMax(arr[6:], ch)

// 	maxVal := math.MinInt32
// 	for i := 0; i < 3; i++ {
// 		select {
// 		case result := <-ch:
// 			fmt.Println(result)
// 		}
// 	}
// 	fmt.Println(maxVal)
// }

// func main() {
// 	pingChan := make(chan string)
// 	pongChan := make(chan string)
// 	lastTimeout := make(chan string)

// 	// Goroutine to send "ping" every second
// 	go func() {
// 		for {
// 			time.Sleep(5 * time.Second)
// 			pingChan <- "ping"
// 		}
// 	}()

// 	// Goroutine to send "pong" every 2 seconds
// 	go func() {
// 		for {
// 			time.Sleep(6 * time.Second)
// 			pongChan <- "pong"
// 		}
// 	}()
// 	go func() {
// 		for {
// 			time.Sleep(6 * time.Second)
// 			lastTimeout <- "Timeout! No message received for 3 seconds."
// 		}
// 	}()

// 	for {
// 		select {
// 		case msg := <-pingChan:
// 			fmt.Println("Received:", msg)
// 		case msg := <-pongChan:
// 			fmt.Println("Received:", msg)
// 		case msg := <-lastTimeout:
// 			fmt.Println("Timeout! No message received for 3 seconds.", msg)
// 			return
// 		}
// 	}
// }
