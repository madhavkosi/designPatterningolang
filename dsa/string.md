### Longest substring without repeating characters.

Given a string, find the length of the longest substring without repeating characters.

### Example

**Example 1**:
```
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

**Example 2**:
```
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

**Example 3**:
```
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```



```go
package main

import "fmt"

// lengthOfLongestSubstring finds the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	i := 0
	j := 0
	charSet := make(map[byte]bool)
	maxLen := 0
	for i <= j && j < len(s) {
		if !charSet[s[j]] {
			charSet[s[j]] = true
			j++
			maxLen = max(maxLen, j-i)
		} else {
			delete(charSet, s[i])
			i++
		}
	}
	return maxLen
}

func main() {
	s := "abcabcbb"
	fmt.Printf("The length of the longest substring without repeating characters is: %d\n", lengthOfLongestSubstring(s))

	s = "bbbbb"
	fmt.Printf("The length of the longest substring without repeating characters is: %d\n", lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Printf("The length of the longest substring without repeating characters is: %d\n", lengthOfLongestSubstring(s))

	s = ""
	fmt.Printf("The length of the longest substring without repeating characters is: %d\n", lengthOfLongestSubstring(s))
}
```

### Length of the longest sequence of consecutive elements

To solve the problem of finding the length of the longest sequence of consecutive elements in an array, we can use an efficient approach involving a set for O(1) average-time complexity lookups.

### Approach:

1. **Use a Set**: 
   - Store all the elements of the array in a set. This helps in quickly checking the presence of a number in O(1) average time complexity.

2. **Find the Start of a Sequence**:
   - Iterate through each number in the array.
   - For each number, check if it is the start of a sequence by checking if `num - 1` is not in the set. If it is not, then `num` is the starting number of a potential sequence.

3. **Count the Consecutive Numbers**:
   - Starting from the identified starting number, count all the consecutive numbers (`num + 1`, `num + 2`, ...) that are present in the set.

4. **Update the Maximum Length**:
   - Keep track of the maximum length of consecutive numbers found.

This method ensures that each number is processed at most twice (once during set insertion and once during the sequence finding), leading to an O(n) time complexity solution.

### Implementation in Go

Here is the implementation in Go:

```go
package main

import "fmt"

// longestConsecutive finds the length of the longest sequence of consecutive integers in the array.
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	maxLength := 0
	for num := range numSet {
		// Check if 'num' is the start of a sequence
		if !numSet[num-1] {
			currentNum := num
			currentLength := 1

			// Count the length of the consecutive sequence starting from 'num'
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}

			// Update maxLength if current sequence is longer
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}

func main() {
	// Test cases
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2})) // Output: 4 (sequence: 1, 2, 3, 4)
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})) // Output: 9 (sequence: 0, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(longestConsecutive([]int{10, 5, 9, 1, 11, 8, 6, 7, 2, 3, 12})) // Output: 6 (sequence: 7, 8, 9, 10, 11, 12)
}
```

### Explanation:

1. **numSet**:
   - A set (implemented as a map with boolean values) that contains all the numbers in the input array. This helps quickly check if a number is part of the array.

2. **Finding the Start of a Sequence**:
   - For each number `num` in the set, if `num - 1` is not in the set, then `num` is the start of a sequence. This is because all previous elements are not part of the current sequence.

3. **Counting Consecutive Numbers**:
   - Start from `num` and count consecutive numbers by incrementing `currentNum` until the next number is not found in the set.

4. **Updating maxLength**:
   - Keep track of the maximum length of any consecutive sequence found during the iteration.

### Complexity:

- **Time Complexity**: O(n), where n is the number of elements in the array. Each element is processed at most twice.
- **Space Complexity**: O(n), due to the space needed to store elements in the set.


To find the length of the largest subarray with a sum equal to a given value \( K \), we can use the **prefix sum** technique along with a hash map. This approach allows us to find the required subarray in linear time, \( O(N) \), where \( N \) is the number of elements in the array.

### Approach:

1. **Prefix Sum**:
   - The prefix sum is the sum of elements from the beginning of the array up to a certain index. For an element at index \( i \), the prefix sum can be represented as `prefixSum[i]`.

2. **Using Hash Map**:
   - Use a hash map to store the first occurrence of each prefix sum. The key will be the prefix sum, and the value will be the index at which this sum first occurs.

3. **Finding the Largest Subarray**:
   - Iterate through the array and maintain a running sum (prefix sum).
   - For each element, calculate the prefix sum and check if `prefixSum - K` exists in the hash map:
     - If it exists, it means there is a subarray that sums to \( K \) between the previous occurrence of `prefixSum - K` and the current index.
     - Update the maximum length of such subarray if the current one is longer.
   - Store the current prefix sum in the hash map if it is not already present, to ensure the longest subarray is considered.

### Implementation in Go

Here is the implementation of the above approach in Go:

```go
package main

import "fmt"

// maxSubArrayLen finds the length of the largest subarray with a sum equal to K.
func maxSubArrayLen(nums []int, K int) int {
	prefixSumMap := make(map[int]int)
	prefixSum := 0
	maxLength := 0

	for i, num := range nums {
		prefixSum += num

		// If the prefix sum equals K, the subarray from the start to current index has sum K
		if prefixSum == K {
			maxLength = i + 1
		}

		// If (prefixSum - K) is found in the map, it means there is a subarray that sums to K
		if val, found := prefixSumMap[prefixSum-K]; found {
			maxLength = max(maxLength, i-val)
		}

		// Store the prefix sum in the map if not already present
		if _, found := prefixSumMap[prefixSum]; !found {
			prefixSumMap[prefixSum] = i
		}
	}

	return maxLength
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, -1, 5, -2, 3}
	K := 3
	fmt.Printf("The length of the largest subarray with sum %d is: %d\n", K, maxSubArrayLen(nums, K))

	nums = []int{-2, -1, 2, 1}
	K = 1
	fmt.Printf("The length of the largest subarray with sum %d is: %d\n", K, maxSubArrayLen(nums, K))

	nums = []int{1, 2, 3, 4, 5}
	K = 11
	fmt.Printf("The length of the largest subarray with sum %d is: %d\n", K, maxSubArrayLen(nums, K))
}
```

### Explanation:

1. **prefixSumMap**:
   - A map that stores the prefix sum as the key and the earliest index at which this sum occurs as the value.

2. **prefixSum**:
   - A running total of the sum of elements from the start up to the current index.

3. **Checking for Subarrays**:
   - For each element, we compute the current prefix sum.
   - If `prefixSum - K` exists in `prefixSumMap`, it means there is a subarray that sums to \( K \) ending at the current index.
   - We update the `maxLength` if the current subarray length (i.e., `i - val`) is greater than the previously recorded `maxLength`.

4. **Storing Prefix Sums**:
   - The prefix sum and its corresponding index are stored in `prefixSumMap` only if the prefix sum is not already present. This ensures that the first occurrence of a prefix sum is stored, which is crucial for finding the longest subarray.

### Complexity:

- **Time Complexity**: O(N), where \( N \) is the number of elements in the array. The array is traversed once.
- **Space Complexity**: O(N), as the hash map `prefixSumMap` may store up to \( N \) different prefix sums.


 involves finding all unique quadruplets \((a, b, c, d)\) in an array such that \(a + b + c + d = \text{target}\). This is an extension of the more commonly known 2-Sum and 3-Sum problems.

### The 4-Sum problem
Given an array `nums` of `n` integers and an integer `target`, return all unique quadruplets \([a, b, c, d]\) such that:
1. \(a + b + c + d = \text{target}\)
2. The quadruplets should not contain duplicate sets of numbers.

### Approach

The 4-Sum problem can be solved efficiently using a sorted array and two pointers technique, along with a nested loop. Here’s a step-by-step breakdown:

1. **Sort the Array**: Sorting helps in avoiding duplicates and simplifying the two-pointer approach.

2. **Iterate with Nested Loops**: 
   - The first two loops will fix the first two numbers of the quadruplet.
   - For the remaining two numbers, use a two-pointer technique to find pairs that sum up to the required value.

3. **Two Pointers Technique**:
   - After fixing the first two numbers (`nums[i]` and `nums[j]`), use two pointers (`left` and `right`) to find the other two numbers (`nums[left]` and `nums[right]`) such that the sum equals the target.
   - Move the pointers `left` and `right` towards each other based on the current sum.

4. **Avoid Duplicates**: 
   - Skip duplicate numbers for all four indices to ensure the quadruplets are unique.

### Implementation in Go

Here's the implementation of the 4-Sum problem in Go:

```go
package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 4 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j + 1, n - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // Output: [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // Output: [[2 2 2 2]]
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0)) // Output: [[-3 -2 2 3] [-3 -1 1 3] [-3 0 0 3] [-3 0 1 2] [-2 -1 0 3] [-2 -1 1 2] [-1 0 0 1]]
}
```

### Explanation:

1. **Sorting**:
   - The array is first sorted to facilitate the two-pointer approach and to help in avoiding duplicates.

2. **Nested Loops**:
   - The outermost loop (`i`) fixes the first number.
   - The second loop (`j`) fixes the second number.
   - The two-pointer technique is used to find the remaining two numbers such that the sum of the four numbers equals the target.

3. **Two Pointers**:
   - `left` starts just after `j` and `right` starts from the end of the array.
   - If the sum of the four numbers equals the target, the quadruplet is added to the result.
   - If the sum is less than the target, `left` is incremented to increase the sum.
   - If the sum is more than the target, `right` is decremented to decrease the sum.

4. **Avoiding Duplicates**:
   - After finding a valid quadruplet, `left` and `right` are moved to skip over duplicate values.

### Complexity:

- **Time Complexity**: O(N^3), where N is the number of elements in the array. This is because we have two nested loops and a two-pointer scan within the innermost loop.
- **Space Complexity**: O(1) for extra space, not including the space used for storing the result.

This implementation ensures that all unique quadruplets are found and that duplicates are avoided.


To find the maximum number of consecutive `1`s in a binary array, we can use a simple linear scan. This problem involves iterating through the array and counting consecutive `1`s, while keeping track of the maximum count encountered.

### Problem Statement
Given a binary array (an array containing only `0`s and `1`s), find the maximum number of consecutive `1`s in the array.

### Approach

1. **Initialize Counters**:
   - Use a counter `currentCount` to count the current streak of consecutive `1`s.
   - Use `maxCount` to store the maximum streak found so far.

2. **Iterate Through the Array**:
   - For each element in the array:
     - If the element is `1`, increment `currentCount`.
     - If the element is `0`, update `maxCount` if `currentCount` is greater, and reset `currentCount` to 0.
   - After the loop, update `maxCount` one last time in case the array ends with a streak of `1`s.

3. **Return Result**:
   - The value of `maxCount` will be the maximum number of consecutive `1`s in the array.

### Implementation in Go

Here is the implementation:

```go
package main

import "fmt"

// findMaxConsecutiveOnes finds the maximum number of consecutive 1's in the binary array.
func findMaxConsecutiveOnes(nums []int) int {
	maxCount := 0
	currentCount := 0

	for _, num := range nums {
		if num == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}

	return maxCount
}

func main() {
	// Test cases
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) // Output: 3
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1})) // Output: 2
	fmt.Println(findMaxConsecutiveOnes([]int{0, 0, 0, 0, 0, 0})) // Output: 0
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 1, 1, 1, 1})) // Output: 6
	fmt.Println(findMaxConsecutiveOnes([]int{}))                 // Output: 0
}
```

### Explanation:

1. **Initialization**:
   - `maxCount` keeps track of the maximum number of consecutive `1`s found so far.
   - `currentCount` counts the current streak of `1`s.

2. **Iterating Through the Array**:
   - For each element:
     - If it is `1`, `currentCount` is incremented. If `currentCount` exceeds `maxCount`, update `maxCount`.
     - If it is `0`, reset `currentCount` to 0 as the streak of consecutive `1`s is broken.

3. **Final Update**:
   - After the loop, `maxCount` is returned, containing the maximum number of consecutive `1`s found.

### Complexity:

- **Time Complexity**: O(N), where N is the number of elements in the array. The array is traversed once.
- **Space Complexity**: O(1). The solution uses a constant amount of extra space regardless of the size of the input.

This solution efficiently finds the maximum number of consecutive `1`s in a binary array with a straightforward linear pass through the array.


### Problem Statement: Group Anagrams

Given an array of strings, group the anagrams together. You can return the result in any order.

### Input

- An array of strings `strs`.

### Output

- A list of lists, where each list contains anagrams grouped together.

### Example 1:

**Input:**  
`strs = ["eat", "tea", "tan", "ate", "nat", "bat"]`  
**Output:**  
`[["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]`

### Example 2:

**Input:**  
`strs = [""]`  
**Output:**  
`[[""]]`

### Example 3:

**Input:**  
`strs = ["a"]`  
**Output:**  
`[["a"]]`

---

### Solution: Using a Hash Map

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	anagramlist := make(map[string][]string)
	for _, str := range strs {
		k := strings.Split(str, "")
		sort.Slice(k, func(i, j int) bool {
			return k[i] < k[j]
		})
		updated := strings.Join(k, "")
		anagramlist[updated] = append(anagramlist[updated], str)
	}
	groupAnagram := make([][]string, 0)
	for _, value := range anagramlist {
		groupAnagram = append(groupAnagram, value)

	}
	return groupAnagram
}
func main() {
	// Example 1
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs1)) // Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]

	// Example 2
	strs2 := []string{""}
	fmt.Println(groupAnagrams(strs2)) // Output: [[""]]

	// Example 3
	strs3 := []string{"a"}
	fmt.Println(groupAnagrams(strs3)) // Output: [["a"]]
}
```


### Problem Statement: Reverse Words in a String

Given an input string `s`, reverse the order of the words. A word is defined as a sequence of non-space characters. The words in `s` are separated by one or more spaces. Return a string of the words in reverse order concatenated by a single space.

### Input

- A string `s` containing words separated by spaces.

### Output

- A string with the words in reverse order, with a single space separating each word.

### Example 1:

**Input:**  
`s = "the sky is blue"`  
**Output:**  
`"blue is sky the"`

### Example 2:

**Input:**  
`s = "  hello world  "`  
**Output:**  
`"world hello"`

### Example 3:

**Input:**  
`s = "a good   example"`  
**Output:**  
`"example good a"`



Here's the implementation in Go:

```go
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Step 1: Trim the input string to remove leading and trailing spaces
	s = strings.TrimSpace(s)
	
	// Step 2: Split the string into a slice of words
	words := strings.Fields(s)
	
	// Step 3: Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	
	// Step 4: Join the words with a single space
	return strings.Join(words, " ")
}

func main() {
	// Example 1
	s1 := "the sky is blue"
	fmt.Println(reverseWords(s1)) // Output: "blue is sky the"

	// Example 2
	s2 := "  hello world  "
	fmt.Println(reverseWords(s2)) // Output: "world hello"

	// Example 3
	s3 := "a good   example"
	fmt.Println(reverseWords(s3)) // Output: "example good a"
}
```

### Problem Statement: Compare Version Numbers

Given two version numbers, `version1` and `version2`, compare them.

- If `version1 > version2`, return `1`.
- If `version1 < version2`, return `-1`.
- If `version1 == version2`, return `0`.

A version number consists of one or more numbers separated by dots. Each number can have multiple digits, and leading zeros are ignored. Version strings are non-empty and contain only digits and dots.

### Input

- Two strings `version1` and `version2` representing version numbers.

### Output

- An integer representing the comparison result: `1`, `-1`, or `0`.

### Example 1:

**Input:**  
`version1 = "1.01"`, `version2 = "1.001"`  
**Output:**  
`0`  
**Explanation:** Both versions are considered equal because leading zeros are ignored.

### Example 2:

**Input:**  
`version1 = "1.0"`, `version2 = "1.0.0"`  
**Output:**  
`0`  
**Explanation:** Trailing zeros in version strings are ignored.

### Example 3:

**Input:**  
`version1 = "0.1"`, `version2 = "1.1"`  
**Output:**  
`-1`  
**Explanation:** `version1` is less than `version2`.

---

### Compare Version Numbers

The idea is to split both version numbers by the dot (`.`), compare each corresponding segment (converted to an integer), and determine which version is greater.

Here’s the implementation in Go:

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)


func compareVersion(version1 string, version2 string) int {

	v1Part1 := strings.Split(version1, ".")
	v2Part2 := strings.Split(version2, ".")
	maxIter := max(len(v1Part1), len(v2Part2))
	for i := 0; i < maxIter; i++ {
		num1 := 0
		num2 := 0
		if i < len(v1Part1) {
			num1, _ := strconv.Atoi(v1Part1[i])
		}
		if i < len(v2Part2) {
			num2, _ := strconv.Atoi(v2Part2[i])
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

func main() {
	// Example 1
	version1a, version2a := "1.01", "1.001"
	fmt.Println(compareVersion(version1a, version2a)) // Output: 0

	// Example 2
	version1b, version2b := "1.0", "1.0.0"
	fmt.Println(compareVersion(version1b, version2b)) // Output: 0

	// Example 3
	version1c, version2c := "0.1", "1.1"
	fmt.Println(compareVersion(version1c, version2c)) // Output: -1

	// Example 4
	version1d, version2d := "1.2", "1.10"
	fmt.Println(compareVersion(version1d, version2d)) // Output: -1

	// Example 5
	version1e, version2e := "1.0.1", "1"
	fmt.Println(compareVersion(version1e, version2e)) // Output: 1
}
```
