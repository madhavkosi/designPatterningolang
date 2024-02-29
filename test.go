package main

import (
	"fmt"
	"math"
	"unicode"

	Stack "github.com/golang-collections/collections/stack"
)

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

//   1
// 2  3

func findCount(n *Node) int {
	if n == nil {
		return 0
	}
	return findCount(n.Left) + findCount(n.Right) + 1
}
func InorderTravesal(n *Node) {
	if n == nil {
		return
	}
	InorderTravesal(n.Left)
	fmt.Printf("val -> %d \n", n.Value)
	InorderTravesal(n.Right)
}

type Median struct {
	CurrentCount int
	Val          []int
	TotalCount   int
}

func (m *Median) findValue(n *Node) {
	if n == nil {
		return
	}
	m.findValue(n.Left)
	m.CurrentCount += 1
	if m.TotalCount%2 == 0 {
		if m.CurrentCount == m.TotalCount/2 {
			m.Val = append(m.Val, n.Value)
		} else if m.CurrentCount == m.TotalCount/2+1 {
			m.Val = append(m.Val, n.Value)
		}
	} else {
		if m.CurrentCount == m.TotalCount/2+1 {
			m.Val = append(m.Val, n.Value)
		}
	}
	m.findValue(n.Right)
}
func getMedian(node *Node) {
	value := findCount(node)
	fmt.Printf("number of node %d \n", findCount(node))
	m := &Median{TotalCount: value, CurrentCount: 0}
	m.findValue(node)
	fmt.Printf("%v", m)
	if len(m.Val) == 2 {
		average := float64(m.Val[0]+m.Val[1]) / 2.0
		fmt.Printf("Median value is: %.2f\n", average)
	}
	fmt.Printf("value is %d\n", (m.Val[0]))
}
func printTree(node *Node, level int) {
	if node == nil {
		return
	}

	printTree(node.Right, level+1)
	for i := 0; i < level; i++ {
		fmt.Print("\t")
	}
	fmt.Println(node.Value)
	printTree(node.Left, level+1)
}
func findCountValue(node *Node, curr_count *int, val *int, total_count int) {
	if node == nil {
		return
	}
	findCountValue(node.Left, curr_count, val, total_count)
	*curr_count++
	if *curr_count == total_count {
		*val = node.Value
	}
	findCountValue(node.Right, curr_count, val, total_count)
}

type something struct {
	empId  int
	salary int
}

type somethings []something

func (a somethings) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a somethings) Len() int      { return len(a) }
func (a somethings) Less(i, j int) bool {
	if a[i].empId > a[j].empId {
		return true
	} else if a[i].empId == a[j].empId {
		return a[i].empId < a[j].empId
	}
	return false
}

// 	coins := []int{186, 419, 83, 408}
// 	coinChange(coins, 62)
// }

// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = -1
	}
	val := coinCount(coins, amount, amount, dp)

	if val == math.MaxInt64 {
		return -1
	}
	return val
}

func coinCount(coins []int, amount int, curr_sum int, dp []int) int {
	if curr_sum == 0 {
		return 0
	}
	if curr_sum < 0 {
		return -1
	}
	if dp[curr_sum] != -1 {
		return dp[curr_sum]
	}
	min_value := math.MaxInt64
	for _, val := range coins {
		valCount := coinCount(coins, amount, curr_sum-val, dp)
		if valCount == math.MaxInt64 || valCount < 0 {
			continue
		}
		min_value = min(min_value, 1+valCount)
	}
	dp[curr_sum] = min_value
	return dp[curr_sum]
}

// Input: s = "12"
// Output: 2
// Explanation: "12" could be decoded as "AB" (1 2) or "L" (12).

// Input: s = "226"
// Output: 3
// Explanation: "226" could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

// Input: s = "06"
// Output: 0
// Explanation: "06" cannot be mapped to "F" because of the leading zero ("6" is different from "06").

// A message containing letters from A-Z can be encoded into numbers using the following mapping:
// 'A' -> "1"
// 'B' -> "2"
// ...
// 'Z' -> "26"
// To decode an encoded message, all the digits must be grouped then mapped back into
// letters using the reverse of the mapping above (there may be multiple ways). For example, "11106" can be mapped into:

// "AAJF" with the grouping (1 1 10 6)
// "KJF" with the grouping (11 10 6)
// Note that the grouping (1 11 06) is invalid because "06" cannot be mapped into 'F' since "6" is different from "06".

func numDecodings(s string) int {
	dp := make([]int, len(s))
	for i := range dp {
		dp[i] = -1
	}
	val := decodeString(s, 0, dp)

	fmt.Printf("%d\n", val)
	return val
}
func decodeString(s string, index int, dp []int) int {
	if index == len(s) {
		return 1
	}
	if dp[index] != -1 {
		return dp[index]
	}
	if s[index] == '0' {
		return 0
	}
	sum := 0
	sum += decodeString(s, index+1, dp)
	if index < len(s)-1 && (s[index] == '1' || (s[index] == '2' && s[index+1] <= '6')) {
		sum = sum + decodeString(s, index+2, dp)
	}
	dp[index] = sum
	return sum
}

// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.

// You have the following three operations permitted on a word:

// Insert a character
// Delete a character
// Replace a character

// Input: word1 = "intention", word2 = "execution"
// Output: 5
// Explanation:
// intention -> inention (remove 't')
// inention -> enention (replace 'i' with 'e')
// enention -> exention (replace 'n' with 'x')
// exention -> exection (replace 'n' with 'c')
// exection -> execution (insert 'u')

func main() {
	numDecodings("05")
	//minDistance("horse", "ros")
}
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1))
	for i := range dp {
		dp[i] = make([]int, len(word2))
	}
	for i, val := range dp {
		for j := range val {
			dp[i][j] = -1
		}
	}
	value := minDistanceCal(word1, len(word1)-1, word2, len(word2)-1, dp)
	fmt.Printf("%d\n", value)
	return value
}

func minDistanceCal(word1 string, index1 int, word2 string, index2 int, dp [][]int) int {
	if index1 == -1 || index2 == -1 {
		return max(index1, index2) + 1
	}
	if dp[index1][index2] != -1 {
		return dp[index1][index2]
	}
	if word1[index1] == word2[index2] {
		dp[index1][index2] = minDistanceCal(word1, index1-1, word2, index2-1, dp)
		return dp[index1][index2]
	}
	dp[index1][index2] = 1 + min(
		minDistanceCal(word1, index1-1, word2, index2-1, dp),
		min(
			minDistanceCal(word1, index1, word2, index2-1, dp),
			minDistanceCal(word1, index1-1, word2, index2, dp),
		),
	)
	return dp[index1][index2]
}

// Input: s = "3[a]2[bc]"
// Output: "aaabcbc"
func decodeStringas(s string) string {
	k := Stack.New()
	k.Push("")
	for _, char := range s {
		if unicode.IsDigit(char) {

		}
	}
	return ""
}
