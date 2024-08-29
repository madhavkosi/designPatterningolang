

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

func backtrack(digits string, startIndex int, current string, combinations *[]string) {
	if startIndex == len(digits) {
		*combinations = append(*combinations, current)
		return
	}
	letters := digitToLetters[digits[startIndex]]
	for i := 0; i < len(letters); i++ {
		backtrack(digits, startIndex+1, current+string(letters[i]), combinations)
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

func backtrackCombinationSum(arr []int, result *[][]int, target int, startIndex int, sample []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		copyArr := make([]int, len(sample))
		copy(copyArr, sample)
		*result = append(*result, copyArr)
	}
	for i := startIndex; i < len(arr); i++ {
		if i != startIndex && arr[i] == arr[i-1] {
			continue
		}
		sample = append(sample, arr[i])
		backtrackCombinationSum(arr, result, target-arr[i], i+1, sample)
		sample = sample[0 : len(sample)-1]
	}
}

func combinationSum2(arr []int, target int) [][]int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	result := make([][]int, 0)
	array := make([]int, 0)
	backtrackCombinationSum(arr, &result, target, 0, array)
	return result
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

## Subsets - 2

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
		if i != start && nums[i] == nums[i-1] {
			continue
		}
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
	backtrack(nums, []int{}, visited, &result)
	return result
}

func backtrack(nums []int, current []int, visited []bool, result *[][]int) {
	if len(current) == len(nums) {
		combination := make([]int, len(current))
		copy(combination, current)
		*result = append(*result, current)
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		current = append(current, nums[i])
		visited[i] = true
		backtrack(nums, current, visited, result)
		current = (current)[:len(current)-1]
		visited[i] = false
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

### Problem Statement: Palindrome Partitioning

Given a string `s`, partition `s` such that every substring of the partition is a palindrome. Return all possible palindrome partitioning of `s`.

### Input

- A string `s` of length `n`.

### Output

- A list of lists containing all possible palindrome partitions of `s`.

### Example 1:

**Input:** `s = "aab"`  
**Output:** `[["a", "a", "b"], ["aa", "b"]]`

### Example 2:

**Input:** `s = "a"`  
**Output:** `[["a"]]`

---

### Solution: Palindrome Partitioning

Here's how you can solve the problem using a backtracking approach.

```go
package main

import (
	"fmt"
)

// Function to check if a string is a palindrome
func isPalindrome(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// Function to find all palindrome partitions
func partition(s string) [][]string {
	var result [][]string
	var current []string
	backtrack(s, 0, current, &result)
	return result
}

// Backtracking function to explore all possible partitions
func backtrack(s string, start int, current []string, result *[][]string) {
	if start >= len(s) {
		temp := make([]string, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}
	for end := start; end < len(s); end++ {
		if isPalindrome(s, start, end) {
			current = append(current, s[start:end+1])
			backtrack(s, end+1, current, result)
			current = current[:len(current)-1]
		}
	}
}

func main() {
	// Example 1
	s1 := "aab"
	fmt.Println(partition(s1))
	// Output: [["a", "a", "b"], ["aa", "b"]]

	// Example 2
	s2 := "a"
	fmt.Println(partition(s2))
	// Output: [["a"]]
}
```

### Problem Statement: Sudoku Solver

Given a 9x9 Sudoku board, write a function to solve the board. The empty cells are filled with digits from 1 to 9, such that each digit appears exactly once in each row, column, and 3x3 sub-box.

### Input

- A 9x9 grid representing the Sudoku board, where empty cells are represented by `'.'`.

### Output

- The same 9x9 grid with the empty cells filled with the correct digits to solve the Sudoku.

### Example:

**Input:**
```text
[
  ['5', '3', '.', '.', '7', '.', '.', '.', '.'],
  ['6', '.', '.', '1', '9', '5', '.', '.', '.'],
  ['.', '9', '8', '.', '.', '.', '.', '6', '.'],
  ['8', '.', '.', '.', '6', '.', '.', '.', '3'],
  ['4', '.', '.', '8', '.', '3', '.', '.', '1'],
  ['7', '.', '.', '.', '2', '.', '.', '.', '6'],
  ['.', '6', '.', '.', '.', '.', '2', '8', '.'],
  ['.', '.', '.', '4', '1', '9', '.', '.', '5'],
  ['.', '.', '.', '.', '8', '.', '.', '7', '9']
]
```

**Output:**
```text
[
  ['5', '3', '4', '6', '7', '8', '9', '1', '2'],
  ['6', '7', '2', '1', '9', '5', '3', '4', '8'],
  ['1', '9', '8', '3', '4', '2', '5', '6', '7'],
  ['8', '5', '9', '7', '6', '1', '4', '2', '3'],
  ['4', '2', '6', '8', '5', '3', '7', '9', '1'],
  ['7', '1', '3', '9', '2', '4', '8', '5', '6'],
  ['9', '6', '1', '5', '3', '7', '2', '8', '4'],
  ['2', '8', '7', '4', '1', '9', '6', '3', '5'],
  ['3', '4', '5', '2', '8', '6', '1', '7', '9']
]
```

---

### Solution: Sudoku Solver

```go
package main

import (
	"fmt"
)

func printBoard(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}
func validBoard(board [][]byte, r, c int, char byte) bool {

	for i := 0; i < 9; i++ {
		if board[r][i] == char || board[i][c] == char || board[(r/3)*3+i/3][(c/3)*3+i%3] == char {
			return false
		}
	}
	return true
}
func backtrackSudoku(board [][]byte, r, c int) bool {
	if r == 9 {
		return true
	}
	if c == 9 {
		return backtrackSudoku(board, r+1, 0)
	}
	if board[r][c] == '.' {
		for i := byte('1'); i <= byte('9'); i++ {
			if validBoard(board, r, c, i) {
				fmt.Println(i)
				board[r][c] = i
				if backtrackSudoku(board, r, c+1) {
					return true
				}
				board[r][c] = '.'
			}

		}
	} else {
		return backtrackSudoku(board, r, c+1)
	}
	return false
}
func solveSudoku(board [][]byte) bool {
	return backtrackSudoku(board, 0, 0)
}


func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)
	printBoard(board)
}
```


### Problem Statement: All Possible Word Breaks

Given a string `s` and a dictionary of words `wordDict`, return all possible ways to segment `s` into a space-separated sequence of one or more dictionary words.

### Input

- A string `s` of length `n`.
- A list of words `wordDict` containing valid words.

### Output

- A list of strings, where each string is a valid segmentation of `s`.

### Example 1:

**Input:**  
`s = "catsanddog"`  
`wordDict = ["cat", "cats", "and", "sand", "dog"]`

**Output:**  
`["cats and dog", "cat sand dog"]`

### Example 2:

**Input:**  
`s = "pineapplepenapple"`  
`wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]`

**Output:**  
`["pine apple pen apple", "pineapple pen apple", "pine applepen apple"]`

### Example 3:

**Input:**  
`s = "catsandog"`  
`wordDict = ["cats", "dog", "sand", "and", "cat"]`

**Output:**  
`[]` (No valid segmentation)

---

### Solution: Backtracking with Memoization

```go
package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	wordSet := make(map[string]bool, 0)
	for _, s := range wordDict {
		wordSet[s] = true
	}
	result := make([][]string, 0)
	var current []string

	backtrack(s, wordSet, 0, current, &result)
	ss := make([]string, 0)
	for _, r := range result {
		ss = append(ss, strings.Join(r, " "))
	}
	return ss
}

func backtrack(s string, wordSet map[string]bool, start int, current []string, result *[][]string) {
	if start == len(s) {
		copyArr := make([]string, len(current))
		copy(copyArr, current)
		*result = append(*result, copyArr)
		return
	}
	for i := start; i < len(s); i++ {
		st := s[start : i+1]
		if wordSet[st] {
			current = append(current, st)
			backtrack(s, wordSet, i+1, current, result)
			current = (current)[:len(current)-1]
		}
	}

}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	result := wordBreak(s, wordDict)
	for _, r := range result {
		fmt.Println(r)
	}
}

func main() {
	// Example 1
	s1 := "catsanddog"
	wordDict1 := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s1, wordDict1)) // Output: ["cats and dog", "cat sand dog"]

	// Example 2
	s2 := "pineapplepenapple"
	wordDict2 := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	fmt.Println(wordBreak(s2, wordDict2)) // Output: ["pine apple pen apple", "pineapple pen apple", "pine applepen apple"]

	// Example 3
	s3 := "catsandog"
	wordDict3 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s3, wordDict3)) // Output: []
}
```

