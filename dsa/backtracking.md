

## Letter Combinations of a Phone Number

Given a string `digits` containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent. Return the answer in any order.

A mapping of digits to letters (similar to telephone buttons) is provided:
- '2' maps to "abc"
- '3' maps to "def"
- '4' maps to "ghi"
- '5' maps to "jkl"
- '6' maps to "mno"
- '7' maps to "pqrs"
- '8' maps to "tuv"
- '9' maps to "wxyz"

Note that the digit '1' does not map to any letters.

**Examples:**

Example 1:
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]

Example 2:
Input: digits = ""
Output: []

Example 3:
Input: digits = "2"
Output: ["a","b","c"]

**Solution in Go:**

```go
package main

import (
	"fmt"
)

var digitToLetters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	var combinations []string
	backtrack(digits, 0, "", &combinations)
	return combinations
}

func backtrack(digits string, index int, current string, combinations *[]string) {
	if index == len(digits) {
		*combinations = append(*combinations, current)
		return
	}
	letters := digitToLetters[digits[index]]
	for i := 0; i < len(letters); i++ {
		backtrack(digits, index+1, current+string(letters[i]), combinations)
	}
}
```
Certainly! Here's the problem statement followed by the provided solution in Go:

---


## Combination Sum

Given an array `candidates` of distinct integers and a `target` integer, return a list of all unique combinations of `candidates` where the chosen numbers sum to `target`. You may return the combinations in any order.

The same number may be chosen from `candidates` an unlimited number of times. Two combinations are unique if the frequency of at least one of the chosen numbers is different.

It is guaranteed that the number of unique combinations that sum up to `target` is less than 150 combinations for the given input.

**Examples:**

Example 1:
Input: `candidates = [2,3,6,7], target = 7`
Output: `[[2,2,3],[7]]`
Explanation:
- 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
- 7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:
Input: `candidates = [2,2,3,5], target = 8`
Output: `[[2,2,2,2],[2,3,3],[3,5]]`

**Solution in Go:**

```go
package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	sort.Ints(candidates)
	backtrack(candidates, target, 0, []int{}, &result)
	return result
}

func backtrack(candidates []int, target, start int, current []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, combination)
		return
	}

	for i := start; i < len(candidates); i++ {
        if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			break
		}
		current = append(current, candidates[i])
		backtrack(candidates, target-candidates[i], i, current, result)
		current = current[:len(current)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum2(candidates, target)
	fmt.Println(result) // Output: [[2 2 3] [7]]

	candidates2 := []int{2, 3, 5}
	target2 := 8
	result2 := combinationSum2(candidates2, target2)
	fmt.Println(result2) // Output: [[2 2 2 2] [2 3 3] [3 5]]
}
```

## Subsets

**Problem Description**
Given an integer array `nums` of unique elements, return all possible subsets (the power set). The solution set must not contain duplicate subsets. Return the solution in any order.

**Examples**
**Example 1**
Input: `nums = [1, 2, 3]`  
Output: `[[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]`

**Example 2**
Input: `nums = [0]`  
Output: `[[], [0]]`

**Golang Solution**
```go
package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	var result [][]int
	backtrack(nums, 0, []int{}, &result)
	return result
}

func backtrack(nums []int, start int, current []int, result *[][]int) {
	combination := make([]int, len(current))
	copy(combination, current)
	*result = append(*result, combination)

	for i := start; i < len(nums); i++ {
		current = append(current, nums[i])
		backtrack(nums, i+1, current, result)
		current = current[:len(current)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result) // Output: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]

	nums = []int{0}
	result = subsets(nums)
	fmt.Println(result) // Output: [[], [0]]
}
```

##  Permutations

**Problem Description**
Given an array `nums` of distinct integers, return all the possible permutations. You can return the answer in any order.

**Examples**
**Example 1**
Input: `nums = [1, 2, 3]`  
Output: `[[1, 2, 3], [1, 3, 2], [2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]]`

**Example 2**
Input: `nums = [0, 1]`  
Output: `[[0, 1], [1, 0]]`

**Example 3**
Input: `nums = [1]`  
Output: `[[1]]`

**Golang Solution**
```go
package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var result [][]int
	visited := make([]bool, len(nums))
	backtrack(nums, []int{}, &visited, &result)
	return result
}

func backtrack(nums []int, current []int, visited *[]bool, result *[][]int) {
	if len(current) == len(nums) {
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, current)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*visited)[i] {
			continue
		}
		current = append(current, nums[i])
		(*visited)[i] = true
		backtrack(nums, current, visited, result)
		current = (current)[:len(current)-1]
		(*visited)[i] = false
	}
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println(result) // Output: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]

	nums = []int{0, 1}
	result = permute(nums)
	fmt.Println(result) // Output: [[0 1] [1 0]]

	nums = []int{1}
	result = permute(nums)
	fmt.Println(result) // Output: [[1]]
}
```
---

##  Generate Parentheses

**Problem Description**
Given `n` pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

**Examples**
**Example 1**
Input: `n = 3`  
Output: `["((()))", "(()())", "(())()", "()(())", "()()()"]`

**Example 2**
Input: `n = 1`  
Output: `["()"]`

**Golang Solution**
```go
package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var result []string
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, current string, open, close, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}
}

func main() {
	n := 3
	result := generateParenthesis(n)
	fmt.Println(result) // Output: ["((()))", "(()())", "(())()", "()(())", "()()()"]

	n = 1
	result = generateParenthesis(n)
	fmt.Println(result) // Output: ["()"]
}
```

---

## N-Queens

**Problem Description**
The n-queens puzzle is the problem of placing `n` queens on an `n x n` chessboard such that no two queens attack each other. Given an integer `n`, return all distinct solutions to the n-queens puzzle. Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space, respectively.

**Examples**
**Example 1**
Input: `n = 4`  
Output: `[[".Q..", "...Q", "Q...", "..Q."], ["..Q.", "Q...", "...Q", ".Q.."]]`

**Example 2**
Input: `n = 1`  
Output: `[["Q"]]`

**Golang Solution**
```go
package main

import (
	"fmt"
)

func solveNQueens(n int) [][]string {
	var result [][]string
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}
	backtrack(board, 0, &result)
	return result
}

func backtrack(board [][]rune, row int, result *[][]string) {
	if row == len(board) {
		addSolution(board, result)
		return
	}

	for col := 0; col < len(board); col++ {
		if isValid(board, row, col) {
			board[row][col] = 'Q'
			backtrack(board, row+1, result)
			board[row][col] = '.'
		}
	}
}

func isValid(board [][]rune, row, col int) bool {
	// Check same column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}
	// Check upper left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	// Check upper right diagonal
	for i, j := row-1, col+1; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func addSolution(board [][]rune, result *[][]string) {
	var solution []string
	for _, row := range board {
		solution = append(solution, string(row))
	}
	*result = append(*result, solution)
}

func main() {
	n := 4
	result := solveNQueens(n)
	for _, solution := range result {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println()
	}

	n = 1
	result = solveNQueens(n)
	for _, solution := range result {
		for _, row := range solution {
			fmt.Println(row)
		}
		fmt.Println()
	}
}
```

---

## Word Search

**Problem Description**
Given an `m x n` grid of characters `board` and a string `word`, return true if `word` exists in the grid. The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically neighboring. The same letter cell may not be used more than once.

**Examples**
**Example 1**
Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`, `word = "ABCCED"`  
Output: `true`

**Example 2**
Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`, `word = "SEE"`  
Output: `true`

**Example 3**
Input: `board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]`, `word = "ABCB"`  
Output: `false`

**Golang Solution**
```go
package main

import "fmt"



func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if search(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

func search(board [][]byte, i, j int, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[index] {
		return false
	}

	temp := board[i][j]
	board[i][j] = '#' // Mark as visited

	// Check in all four directions
	found := search(board, i+1, j, word, index+1) ||
		search(board, i-1, j, word, index+1) ||
		search(board, i, j+1, word, index+1) ||
		search(board, i, j-1, word, index+1)

	board[i][j] = temp // Reset the cell
	return found
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	fmt.Println(exist(board, word)) // Output: true

	word = "SEE"
	fmt.Println(exist(board, word)) // Output: true

	word = "ABCB"
	fmt.Println(exist(board, word)) // Output: false
}
```

---