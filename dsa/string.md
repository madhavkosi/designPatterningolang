### Longest Substring Without Repeating Characters

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

### Solution

```go
package main

import "fmt"

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "abcabcbb"
	fmt.Printf("The length of the longest substring without repeating characters is: %d\n", lengthOfLongestSubstring(s))

	s = "bbbbb"
	fmt.Printf("The length of the longest substring without repeating characters is: %d\n", lengthOfLongestSubstring(s))

	s = "pwwkew"
	fmt.Printf("The length of the longest substring without repeating characters is: %d\n", lengthOfLongestSubstring(s))
}
```

---

### Length of the Largest Subarray with Sum K

Given an array of integers and a target sum \( K \), find the length of the largest subarray with a sum equal to \( K \).

### Example

**Example 1**:
```
Input: nums = [1, -1, 5, -2, 3], K = 3
Output: 4
Explanation: The subarray [1, -1, 5, -2] sums to 3.
```

**Example 2**:
```
Input: nums = [-2, -1, 2, 1], K = 1
Output: 2
Explanation: The subarray [-1, 2] sums to 1.
```

### Solution

```go
package main

import "fmt"

func maxSubArrayLen(nums []int, K int) int {
	prefixSumMap := make(map[int]int)
	prefixSum := 0
	maxLength := 0

	for i, num := range nums {
		prefixSum += num

		if prefixSum == K {
			maxLength = i + 1
		}

		if val, found := prefixSumMap[prefixSum-K]; found {
			maxLength = max(maxLength, i-val)
		}

		if _, found := prefixSumMap[prefixSum]; !found {
			prefixSumMap[prefixSum] = i
		}
	}

	return maxLength
}

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
}
```

---

### 4-Sum Problem

Given an array of integers `nums` and an integer `target`, return all unique quadruplets \([a, b, c, d]\) such that \(a + b + c + d = \text{target}\).

### Example

**Example 1**:
```
Input: nums = [1, 0, -1, 0, -2, 2], target = 0
Output: [[-2, -1, 1, 2], [-2, 0, 0, 2], [-1, 0, 0, 1]]
```

**Example 2**:
```
Input: nums = [2, 2, 2, 2, 2], target = 8
Output: [[2, 2, 2, 2]]
```

### Solution

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
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0)) // Output: [[-2 -1 1 2] [-2 0 0 2] [-1 0 0 1]]
	fmt.Println(fourSum([]int{2, 2, 2, 2, 2}, 8))      // Output: [[2 2 2 2]]
}
```

---

### Maximum Number of Consecutive 1's

Given a binary array, find the maximum number of consecutive `1`s in the array.

### Example

**Example 1**:
```
Input: nums = [1, 1, 0, 1, 1, 1]
Output: 3
Explanation: The maximum number of consecutive 1's is 3.
```

**Example 2**:
```
Input: nums = [1, 0, 1, 1, 0, 1]
Output: 2
Explanation: The maximum number of consecutive 1's is 2.
```

### Solution

```go
package main

import "fmt"

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
	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) // Output: 3
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1})) // Output: 2
}
```

---

### Group Anagrams

Given an array of strings, group the anagrams together.

### Example

**Example 1**:
```
Input: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

### Solution

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
	strs1 := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs1)) // Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
}
```

---

### Reverse Words in a String

Given an input string `s`, reverse the order of the words.



### Example

**Example 1**:
```
Input: "the sky is blue"
Output: "blue is sky the"
```

**Example 2**:
```
Input: "  hello world  "
Output: "world hello"
```

**Example 3**:
```
Input: "a good   example"
Output: "example good a"
```

### Solution

```go
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	s1 := "the sky is blue"
	fmt.Println(reverseWords(s1)) // Output: "blue is sky the"

	s2 := "  hello world  "
	fmt.Println(reverseWords(s2)) // Output: "world hello"

	s3 := "a good   example"
	fmt.Println(reverseWords(s3)) // Output: "example good a"
}
```

---

### Compare Version Numbers

Given two version numbers `version1` and `version2`, compare them.

### Example

**Example 1**:
```
Input: version1 = "1.01", version2 = "1.001"
Output: 0
```

**Example 2**:
```
Input: version1 = "1.0", version2 = "1.0.0"
Output: 0
```

**Example 3**:
```
Input: version1 = "0.1", version2 = "1.1"
Output: -1
```

### Solution

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
			num1, _ = strconv.Atoi(v1Part1[i])
		}
		if i < len(v2Part2) {
			num2, _ = strconv.Atoi(v2Part2[i])
		}
		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

---

### Length of the Longest Sequence of Consecutive Elements

Given an unsorted array of integers, find the length of the longest sequence of consecutive elements.

### Example

**Example 1**:
```
Input: nums = [100, 4, 200, 1, 3, 2]
Output: 4
```

**Example 2**:
```
Input: nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
Output: 9
```

### Solution

```go
package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	for _, num := range nums {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	nums1 := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums1)) // Output: 4

	nums2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums2)) // Output: 9
}
```