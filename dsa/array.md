# Golang Solutions for Array Problems

This repository contains solutions to various array problems implemented in Golang. Each problem is broken into multiple files with a README file containing the problem description, examples, and the solution.

## Table of Contents

1. [Maximum Subarray](#maximum-subarray)
2. [Maximum Sum Circular Subarray](#maximum-sum-circular-subarray)
3. [Merge Intervals](#merge-intervals)
4. [Insert Interval](#insert-interval)
5. [Minimum Number of Arrows to Burst Balloons](#minimum-number-of-arrows-to-burst-balloons)
6. [Merge Sorted Array](#merge-sorted-array)
7. [Remove Element](#remove-element)
8. [Remove Duplicates from Sorted Array](#remove-duplicates-from-sorted-array)
9. [Remove Duplicates from Sorted Array II](#remove-duplicates-from-sorted-array-ii)
10. [Majority Element](#majority-element)
11. [Best Time to Buy and Sell Stock](#best-time-to-buy-and-sell-stock)
12. [Best Time to Buy and Sell Stock II](#best-time-to-buy-and-sell-stock-ii)
13. [Insert Delete GetRandom O(1)](#insert-delete-getrandom-o1)
14. [Product of Array Except Self](#product-of-array-except-self)
15. [Jump Game](#jump-game)
16. [Jump Game II](#jump-game-ii)
17. [Rotate Array](#rotate-array)
18. [H-Index](#h-index)
19. [Roman to Integer](#roman-to-integer)
20. [Integer to Roman](#integer-to-roman)
21. [Reverse Words in a String](#reverse-words-in-a-string)
22. [Find the Index of the First Occurrence in a String](#find-the-index-of-the-first-occurrence-in-a-string)
23. [Longest Common Prefix](#longest-common-prefix)
24. [Gas Station](#gas-station)
25. [Candy](#candy)
26. [Find the Duplicate Number](#find-the-duplicate-number)
27. [Sort Colors](#sort-colors)
---

## Maximum Subarray

**Problem Description**
Given an integer array `nums`, find the contiguous subarray (containing at least one number) that has the largest sum and return its sum. A subarray is a contiguous part of an array.

**Examples**
**Example 1**
Input: `nums = [-2,1,-3,4,-1,2,1,-5,4]`
Output: `6`
Explanation: `[4,-1,2,1]` has the largest sum = `6`.

**Example 2**
Input: `nums = [1]`
Output: `1`

**Example 3**
Input: `nums = [5,4,-1,7,8]`
Output: `23`


```go
package main

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := math.MinInt32
	for _, num := range nums {
		sum += num
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxSum
}
```


---



## Maximum Sum Circular Subarray

**Problem Description**
Given a circular integer array `nums` of length `n`, return the maximum possible sum of a non-empty subarray of `nums`. A circular array means the end of the array connects to the beginning of the array. A subarray may only include each element of the fixed buffer `nums` at most once.

**Examples**
**Example 1**
Input: `nums = [1,-2,3,-2]`
Output: `3`
Explanation: Subarray `[3]` has maximum sum `3`.

**Example 2**
Input: `nums = [5,-3,5]`
Output: `10`
Explanation: Subarray `[5,5]` has maximum sum `5 + 5 = 10`.

**Example 3**
Input: `nums = [-3,-2,-3]`
Output: `-2`
Explanation: Subarray `[-2]` has maximum sum `-2`.


```go
package main

import "math"

func maxSubarraySumCircular(nums []int) int {
	total, maxSum, curMax, minSum, curMin := 0, nums[0], 0, nums[0], 0
	for _, a := range nums {
		curMax = int(math.Max(float64(curMax + a), float64(a)))
		maxSum = int(math.Max(float64(maxSum), float64(curMax)))
		curMin = int(math.Min(float64(curMin + a), float64(a)))
		minSum = int(math.Min(float64(minSum), float64(curMin)))
		total += a
	}
	if maxSum > 0 {
		return int(math.Max(float64(maxSum), float64(total - minSum)))
	}
	return maxSum
}
```

---

## Merge Intervals

**Problem Description**
Given an array of intervals where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

**Examples**
**Example 1**
Input: `intervals = [[1,3],[2,6],[8,10],[15,18]]`
Output: `[[1,6],[8,10],[15,18]]`
Explanation: Since intervals `[1,3]` and `[2,6]` overlap, merge them into `[1,6]`.

**Example 2**
Input: `intervals = [[1,4],[4,5]]`
Output: `[[1,5]]`
Explanation: Intervals `[1,4]` and `[4,5]` are considered overlapping.


```go
package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	output := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := output[len(output)-1]
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				last[1] = interval[1]
			}
		} else {
			output = append(output, interval)
		}
	}
	return output
}
```

---

## Insert Interval

**Problem Description**
Given an array of non-overlapping intervals `intervals` where `intervals[i] = [starti, endi]` represent the start and the end of the ith interval and intervals are sorted in ascending order by `starti`. You are also given an interval `newInterval = [start, end]` that represents the start and end of another interval. Insert `newInterval` into intervals such that intervals is still sorted in ascending order by `starti` and intervals still does not have any overlapping intervals (merge overlapping intervals if necessary). Return intervals after the insertion.

**Examples**
**Example 1**
Input: `intervals = [[1,3],[6,9]], newInterval = [2,5]`
Output: `[[1,5],[6,9]]`

**Example 2**
Input: `intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]`
Output: `[[1,2],[3,10],[12,16]]`
Explanation: Because the new interval `[4,8]` overlaps with `[3,5],[6,7],[8,10]`.


```go
package main

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	i := 0
	n := len(intervals)
	for i < n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < n && newInterval[1] >= intervals[i][0] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

---

## Minimum Number of Arrows to Burst Balloons

**Problem Description**
There are some spherical balloons taped onto a flat wall that represents the XY-plane. The balloons are represented as a 2D integer array `points` where `points[i] = [xstart, xend]` denotes a balloon whose horizontal diameter stretches between `xstart` and `xend`. You do not know the exact y-coordinates of the balloons. Arrows can be shot up directly vertically (in the positive y-direction) from different points along the x-axis. A balloon with `xstart` and `xend` is burst by an arrow shot at `x` if `xstart <= x <= xend`. There is no limit to the number of arrows that can be shot. A shot arrow keeps traveling up infinitely, bursting any balloons in its path. Given the array `points`, return the minimum number of arrows that must be shot to burst all balloons.

**Examples**
**Example 1**
Input: `points = [[10,16],[2,8],[1,6],[7,12]]`
Output: `2`
Explanation: The balloons can be burst by 2 arrows:
- Shoot an arrow at `x = 6`, bursting the balloons `[2,8]` and `[1,6]`.
- Shoot an arrow at `x = 11`, bursting the balloons `[10,16]` and `[7,12]`.

**Example 2**
Input: `points = [[1,2],[3,4],[5,6],[7,8]]`
Output: `4`
Explanation: One arrow needs to be shot for each balloon for a total of 4 arrows.

**Example 3**
Input: `points = [[1,2],[2,3],[3,4],[4,5]]`
Output: `2`
Explanation: The balloons can be burst by 2 arrows:
- Shoot an arrow at `x = 2`, bursting the balloons `[1,2]` and `[2,3]`.
- Shoot an arrow at `x = 4`, bursting the balloons `[3,4]` and `[4,5]`.


```go
package main

import "sort"

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(a, b int) bool {
		return points[a][1] < points[b][1]
	})
	arrows := 1
	end := points[0][1]
	for _, p := range points {
		if p[0] > end {
			arrows++
			end = p[1]
		}
	}
	return arrows
}
```

---

## Merge Sorted Array

**Problem Description**
You are given two integer arrays `nums1` and `nums2`, sorted in non-decreasing order, and two integers `m` and `n`, representing the number of elements in `nums1` and `nums2` respectively. Merge `nums1` and `nums2` into a single array sorted in non-decreasing order. The final sorted array should not be returned by the function, but instead be stored inside the array `nums1`. To accommodate this, `nums1` has a length of `m + n`, where the first `m` elements denote the elements that should be merged, and the last `n` elements are set to `0` and should be ignored. `nums2` has a length of `n`.

**Examples**
**Example 1**
Input: `nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3`
Output: `[1,2,2,3,5,6]`
Explanation: The arrays we are merging are `[

1,2,3]` and `[2,5,6]`. The result of the merge is `[1,2,2,3,5,6]` with the underlined elements coming from `nums1`.

**Example 2**
Input: `nums1 = [1], m = 1, nums2 = [], n = 0`
Output: `[1]`

**Example 3**
Input: `nums1 = [0], m = 0, nums2 = [1], n = 1`
Output: `[1]`


```go
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
```

## Remove Element

**Problem Description**
Given an integer array `nums` and an integer `val`, remove all occurrences of `val` in `nums` in-place. The order of the elements may be changed. Then return the number of elements in `nums` which are not equal to `val`.

**Examples**
**Example 1**
Input: `nums = [3,2,2,3], val = 3`
Output: `2`
Explanation: Your function should return `k = 2`, with the first two elements of `nums` being `2`. It does not matter what you leave beyond the returned `k` (hence they are underscores).

**Example 2**
Input: `nums = [0,1,2,2,3,0,4,2], val = 2`
Output: `5`
Explanation: Your function should return `k = 5`, with the first five elements of `nums` containing `0`, `1`, `3`, `0`, and `4`. Note that the five elements can be returned in any order. It does not matter what you leave beyond the returned `k`.


```go
package main

func removeElement(nums []int, val int) int {
	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}
```

---

## Remove Duplicates from Sorted Array

**Problem Description**
Given an integer array `nums` sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. 

**Examples**
**Example 1**
Input: `nums = [1,1,2]`
Output: `2, nums = [1,2,_]`

**Example 2**
Input: `nums = [0,0,1,1,1,2,2,3,3,4]`
Output: `5, nums = [0,1,2,3,4,_,_,_,_,_]`


```go
package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	for right := 1; right < len(nums); right++ {
		if nums[left] != nums[right] {
			left++
			nums[left] = nums[right]
		}
	}
	return left + 1
}
```

---

## Remove Duplicates from Sorted Array II

**Problem Description**
Given an integer array `nums` sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice. The relative order of the elements should be kept the same. 

**Examples**
**Example 1**
Input: `nums = [1,1,1,2,2,3]`
Output: `5, nums = [1,1,2,2,3,_]`

**Example 2**
Input: `nums = [0,0,1,1,1,1,2,3,3]`
Output: `7, nums = [0,0,1,1,2,3,3,_,_]`


```go
package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	i := 2
	for j := 2; j < n; j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
```

---

## Majority Element

**Problem Description**
Given an array `nums` of size `n`, return the majority element. The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

**Examples**
**Example 1**
Input: `nums = [3,2,3]`
Output: `3`

**Example 2**
Input: `nums = [2,2,1,1,1,2,2]`
Output: `2`


```go
package main

func majorityElement(nums []int) int {
	count := 0
	var candidate int

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}
```

---

## Best Time to Buy and Sell Stock

**Problem Description**
You are given an array `prices` where `prices[i]` is the price of a given stock on the ith day. You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return `0`.

**Examples**
**Example 1**
Input: `prices = [7,1,5,3,6,4]`
Output: `5`
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.

**Example 2**
Input: `prices = [7,6,4,3,1]`
Output: `0`
Explanation: In this case, no transactions are done and the max profit = 0.


```go
package main

func maxProfit(prices []int) int {
	minPrice := int(^uint(0) >> 1) // Max int
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		} else if price - minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}
	return maxProfit
}
```

---

## Best Time to Buy and Sell Stock II

**Problem Description**
You are given an integer array `prices` where `prices[i]` is the price of a given stock on the ith day. On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can buy it then immediately sell it on the same day. Find and return the maximum profit you can achieve.

**Examples**
**Example 1**
Input: `prices = [7,1,5,3,6,4]`
Output: `7`
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4. Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3. Total profit is 4 + 3 = 7.

**Example 2**
Input: `prices = [1,2,3,4,5]`
Output: `4`
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4. Total profit is 4.

**Example 3**
Input: `prices = [7,6,4,3,1]`
Output: `0`
Explanation: There is no way to make a positive profit, so we never buy the stock to achieve the maximum profit of 0.


```go
package main

func maxProfit(prices []int) int {
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}
	return maxProfit
}
```

---

## Insert Delete GetRandom O(1)

**Problem Description**
Implement the `RandomizedSet` class:
- `RandomizedSet()`: Initializes the `RandomizedSet` object.
- `bool insert(int val)`: Inserts an item `val` into the set if not present. Returns `true` if the item was not present, `false` otherwise.
- `bool remove(int val)`: Removes an item `val` from the set if present. Returns `true` if the item was present, `false` otherwise.
- `int getRandom()`: Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average `O(1)` time complexity.

**Examples**
**Example 1**
Input
```
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
```
Output
```
[null, true, false, true, 2, true, false, 2]
```
Explanation
```
RandomizedSet randomizedSet =

 new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
```


```go
package main

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	valToIdx map[int]int
	values   []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().UnixNano())
	return RandomizedSet{
		valToIdx: make(map[int]int),
		values:   []int{},
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.valToIdx[val]; exists {
		return false
	}
	rs.valToIdx[val] = len(rs.values)
	rs.values = append(rs.values, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	idx, exists := rs.valToIdx[val]
	if !exists {
		return false
	}
	lastVal := rs.values[len(rs.values)-1]
	rs.values[idx] = lastVal
	rs.valToIdx[lastVal] = idx
	rs.values = rs.values[:len(rs.values)-1]
	delete(rs.valToIdx, val)
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.values[rand.Intn(len(rs.values))]
}
```

---

## Product of Array Except Self

**Problem Description**
Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`. The product of any prefix or suffix of `nums` is guaranteed to fit in a 32-bit integer. You must write an algorithm that runs in `O(n)` time and without using the division operation.

**Examples**
**Example 1**
Input: `nums = [1,2,3,4]`
Output: `[24,12,8,6]`

**Example 2**
Input: `nums = [-1,1,0,-3,3]`
Output: `[0,0,9,0,0]`


```go
package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}
	return answer
}
```

---

## Jump Game

**Problem Description**
You are given an integer array `nums`. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position. Return `true` if you can reach the last index, or `false` otherwise.

**Examples**
**Example 1**
Input: `nums = [2,3,1,1,4]`
Output: `true`
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

**Example 2**
Input: `nums = [3,2,1,0,4]`
Output: `false`
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.


```go
package main

func canJump(nums []int) bool {
	reachable := 0
	for i := 0; i < len(nums); i++ {
		if reachable < i {
			return false
		}
		reachable = max(reachable, i + nums[i])
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

---

## Jump Game II

**Problem Description**
You are given a 0-indexed array of integers `nums` of length `n`. You are initially positioned at `nums[0]`. Each element `nums[i]` represents the maximum length of a forward jump from index `i`. In other words, if you are at `nums[i]`, you can jump to any `nums[i + j]` where `0 <= j <= nums[i]` and `i + j < n`. Return the minimum number of jumps to reach `nums[n - 1]`. The test cases are generated such that you can reach `nums[n - 1]`.

**Examples**
**Example 1**
Input: `nums = [2,3,1,1,4]`
Output: `2`
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.

**Example 2**
Input: `nums = [2,3,0,1,4]`
Output: `2`


```go
package main

func jump(nums []int) int {
	n := len(nums)
	maxReachable := 0
	lastJumpedPos := 0
	jumps := 0
	for lastJumpedPos < n - 1 {
		maxReachable = max(maxReachable, i + nums[i])
		if i == lastJumpedPos {
			lastJumpedPos = maxReachable
			jumps++
		}
		i++
	}
	return jumps
}
```

---

## Rotate Array

**Problem Description**
Given an integer array `nums`, rotate the array to the right by `k` steps, where `k` is non-negative.

**Examples**
**Example 1**
Input: `nums = [1,2,3,4,5,6,7], k = 3`
Output: `[5,6,7,1,2,3,4]`
Explanation: rotate 1 step to the right: `[7,1,2,3,4,5,6]`, rotate 2 steps to the right: `[6,7,1,2,3,4,5]`, rotate 3 steps to the right: `[5,6,7,1,2,3,4]`

**Example 2**
Input: `nums = [-1,-100,3,99], k = 2`
Output: `[3,99,-1,-100]`
Explanation: rotate 1 step to the right: `[99,-1,-100,3]`, rotate 2 steps to the right: `[3,99,-1,-100]`


```go
package main

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums) - 1)
	reverse(nums, 0, k - 1)
	reverse(nums, k, len(nums) - 1)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
```

---

## H-Index

**Problem Description**
Given an array of integers `citations` where `citations[i]` is the number of citations a researcher received for their `i-th` paper, return the researcher's `h-index`.

**Examples**
**Example 1**
Input: `citations = [3,0,6,1,5]`
Output: `3`
Explanation: `3` means the researcher has `3` papers with at least `3` citations each.


```go
package main

func hIndex(citations []int) int {
	n := len(citations)
	buckets := make([]int, n + 1)
	for _, c := range citations {
		if c >= n {
			buckets[n]++
		} else {
			buckets[c]++
		}
	}
	count := 0
	for i := n; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}
	return 0
}
```

---

## Roman to Integer

**Problem Description**
Given a roman numeral, convert it to an integer. 

**Examples**
**Example 1**
Input: `s = "III"`
Output: `3`

**Example 2**
Input: `s = "LVIII"`
Output: `58`

**Example 3**
Input: `s = "MCMXCIV"`
Output: `1994`


```go
package main

func romanToInt(s string) int {
	romanMap := map[byte]int{
	

	'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := romanMap[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if romanMap[s[i]] < romanMap[s[i+1]] {
			sum -= romanMap[s[i]]
		} else {
			sum += romanMap[s[i]]
		}
	}
	return sum
}
```

---

## Integer to Roman

**Problem Description**
Given an integer, convert it to a roman numeral.

**Examples**
**Example 1**
Input: `num = 3`
Output: `"III"`

**Example 2**
Input: `num = 58`
Output: `"LVIII"`

**Example 3**
Input: `num = 1994`
Output: `"MCMXCIV"`


```go
package main

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			res += sym[i]
		}
	}
	return res
}
```

---

## Reverse Words in a String

**Problem Description**
Given an input string `s`, reverse the order of the words. A word is defined as a sequence of non-space characters. The words in `s` will be separated by at least one space. Return a string of the words in reverse order concatenated by a single space.

**Examples**
**Example 1**
Input: `s = "the sky is blue"`
Output: `"blue is sky the"`

**Example 2**
Input: `s = "  hello world  "`
Output: `"world hello"`

**Example 3**
Input: `s = "a good   example"`
Output: `"example good a"`


```go
package main

import (
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}
```

---

## Find the Index of the First Occurrence in a String

**Problem Description**
Given two strings `needle` and `haystack`, return the index of the first occurrence of `needle` in `haystack`, or `-1` if `needle` is not part of `haystack`.

**Examples**
**Example 1**
Input: `haystack = "sadbutsad", needle = "sad"`
Output: `0`

**Example 2**
Input: `haystack = "leetcode", needle = "leeto"`
Output: `-1`


```go
package main

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	for i := 0; i <= m-n; i++ {
		j := 0
		for ; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == n {
			return i
		}
	}
	return -1
}
```

---

## Longest Common Prefix

**Problem Description**
Write a function to find the longest common prefix string amongst an array of strings. If there is no common prefix, return an empty string `""`.

**Examples**
**Example 1**
Input: `strs = ["flower","flow","flight"]`
Output: `"fl"`

**Example 2**
Input: `strs = ["dog","racecar","car"]`
Output: `""`


```go
package main

import "sort"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	first, last := strs[0], strs[len(strs)-1]
	i := 0
	for i < len(first) && first[i] == last[i] {
		i++
	}
	return first[:i]
}
```

---

## Gas Station

**Problem Description**
There are `n` gas stations along a circular route, where the amount of gas at the `i-th` station is `gas[i]`. You have a car with an unlimited gas tank and it costs `cost[i]` of gas to travel from the `i-th` station to its next `(i + 1)-th` station. You begin the journey with an empty tank at one of the gas stations. Given two integer arrays `gas` and `cost`, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return `-1`. If there exists a solution, it is guaranteed to be unique.

**Examples**
**Example 1**
Input: `gas = [1,2,3,4,5], cost = [3,4,5,1,2]`
Output: `3`
Explanation:
Start at station 3 (index 3) and fill up with 4 units of gas. Travel to station 4. Your tank = 0 + 4 - 1 + 5 = 8 units of gas. Travel to station 0. Your tank = 8 - 2 + 1 = 7 units of gas. Travel to station 1. Your tank = 7 - 3 + 2 = 6 units of gas. Travel to station 2. Your tank = 6 - 4 + 3 = 5 units of gas. Travel to station 3. Your tank = 5 - 5 = 0 units of gas.

**Example 2**
Input: `gas = [2,3,4], cost = [3,4,3]`
Output: `-1`
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.


```go
package main

func canCompleteCircuit(gas []int, cost []int) int {
	totalSurplus, surplus, start := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		totalSurplus += gas[i] - cost[i]
		surplus += gas[i] - cost[i]
		if surplus < 0 {
			surplus = 0
			start = i + 1
		}
	}
	if totalSurplus < 0 {
		return -1
	}
	return start
}
```

---

## Candy

**Problem Description**
There are `n` children standing in a line. Each child is assigned a rating value given in the integer array `ratings`. You are giving candies to these children subjected to the following requirements:
1. Each child must have at least one candy.
2. Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.

**Examples**
**Example 1**
Input: `ratings = [1,0,2]`
Output: `5`
Explanation: You can allocate to the first, second, and third child with 2, 1, 2 candies respectively.

**Example 2**
Input: `ratings = [1,2,2]`
Output: `4`
Explanation: You can allocate to the first, second, and third child with 1, 2, 1 candies respectively.


```go
package main

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1] + 1)
		}
	}
	total := 0
	for _, candy := range candies {
		total += candy
	}
	return total
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

---

## Find the Duplicate Number

**Problem Description**
Given an array of integers `nums` containing `n + 1` integers where each integer is in the range `[1, n]` inclusive. There is only one repeated number in `nums`, return this repeated

 number. You must solve the problem without modifying the array `nums` and uses only constant extra space.

**Examples**
**Example 1**
Input: `nums = [1,3,4,2,2]`
Output: `2`

**Example 2**
Input: `nums = [3,1,3,4,2]`
Output: `3`


```go
package main

func findDuplicate(nums []int) int {
	tortoise := nums[0]
	hare := nums[0]

	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	tortoise = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	return hare
}
```

---

## Sort Colors

**Problem Description**
Given an array `nums` with `n` objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue. We will use the integers `0`, `1`, and `2` to represent the color red, white, and blue, respectively. You must solve this problem without using the library's sort function.

**Examples**
**Example 1**
Input: `nums = [2,0,2,1,1,0]`
Output: `[0,0,1,1,2,2]`

**Example 2**
Input: `nums = [2,0,1]`
Output: `[0,1,2]`


```go
package main

func sortColors(nums []int) {
	low, mid, high := 0, 0, len(nums) - 1
	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
		}
	}
}
```

For more problems and their solutions, navigate through the respective folders. Each folder contains a `README.md` file with the problem description, examples, and the Golang solution.

---
