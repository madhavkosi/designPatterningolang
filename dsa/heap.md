

**Find the Median from a Data Stream**

Design a data structure that supports the following operations efficiently:
1. **AddNum(int num)**: Inserts a new number into the data stream.
2. **FindMedian() float64**: Returns the median of all elements so far.

The median is:
- The middle value in a sorted list if the number of elements is odd.
- The average of the two middle values if the number of elements is even.

### Example

```go

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFind struct {
	minHeap *minHeap
	maxHeap *maxHeap
}

func Const() MedianFind {
	maxHeap := &maxHeap{}
	minHeap := &minHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFind{minHeap, maxHeap}
}

func (mf *MedianFind) AddNM(num int) {
	heap.Push(mf.maxHeap, num)
	x := heap.Pop(mf.maxHeap)
	heap.Push(mf.minHeap, x)

	if len(*mf.minHeap) > len(*mf.maxHeap) {
		x := heap.Pop(mf.minHeap)
		heap.Push(mf.maxHeap, x)
	}
}

func (mf *MedianFind) FindMedian() float64 {
	if len(*mf.minHeap) == len(*mf.maxHeap) {
		va1 := (*mf.minHeap)[0]
		va2 := (*mf.maxHeap)[0]
		return float64(va1+va2) / float64(2)
	}
	return float64((*mf.maxHeap)[0])
}

func main() {
	mf := Const()
	mf.AddNM(1)
	mf.AddNM(2)
	fmt.Println(mf.FindMedian()) // 1.5
	mf.AddNM(3)
	fmt.Println(mf.FindMedian()) // 2
}

```

## Merge k Sorted Linked Lists

### Problem Statement

Given \( k \) sorted linked lists, each with \( n_1, n_2, \ldots, n_k \) elements, merge them into a single sorted linked list.

### Example

**Input:**
- List 1: `1 -> 4 -> 5`
- List 2: `1 -> 3 -> 4`
- List 3: `2 -> 6`

**Output:**
- Merged List: `1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6`

### Explanation:
- The first elements of all the lists are `1, 1, 2`. The smallest is `1`, so we add `1` to the merged list.
- We then move to the next element in the first list (`4`) and repeat the process, comparing `4, 1, 2`.
- Continue this process until all elements from all lists are merged into a single sorted list.

```go

type Element struct {
	Val  int
	Next *ListNode
}
type MinHeap []Element

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Element)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func mergeKLists(lists []*ListNode) *ListNode {
	minheap := &MinHeap{}
	heap.Init(minheap)
	for _, node := range lists {
		if node != nil {
			heap.Push(minheap, Element{Val: node.Val, Next: node})
		}
	}
	dummyNode := ListNode{}
	curr := &dummyNode
	for len(*minheap) > 0 {
		node := heap.Pop(minheap).(Element)
		curr.Next = node.Next
		curr = curr.Next
		forwardNode := node.Next.Next
		if forwardNode != nil {
			heap.Push(minheap, Element{Val: forwardNode.Val, Next: forwardNode})
		}
	}
	return dummyNode.Next
}

```


To find the \( k \) largest and \( k \) smallest elements in an unsorted list, we can use various approaches depending on the size of the list and \( k \). Here are efficient methods for both tasks:

### 1. Finding the \( k \) Largest Elements
We can use a **min-heap** (priority queue) of size \( k \) to keep track of the \( k \) largest elements. As we traverse the list:
- If the heap contains fewer than \( k \) elements, we add the current element.
- Otherwise, if the current element is larger than the smallest element in the heap, we remove the smallest element and add the current element.

This method has a time complexity of \( O(n \log k) \).

### 2. Finding the \( k \) Smallest Elements
Similarly, we can use a **max-heap** of size \( k \) to find the \( k \) smallest elements. We invert the comparison by storing negative values in the heap:
- If the heap contains fewer than \( k \) elements, we add the current element (negated).
- Otherwise, if the current element is smaller than the largest element in the heap (smallest when negated), we remove the largest and add the current element (negated).

This method also has a time complexity of \( O(n \log k) \).

Let's look at Go code examples for both:

### Code for Finding \( k \) Largest Elements

```go
package main

import (
	"container/heap"
	"fmt"
)

func findKLargest(nums []int, k int) []int {
	minHeap := &IntHeap{}
	heap.Init(minHeap)

	for _, num := range nums {
		heap.Push(minHeap, num)
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	return *minHeap
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println("k largest elements:", findKLargest(nums, k)) // Output: [5, 6]
}
```

### Code for Finding \( k \) Smallest Elements

```go
// MinHeap is a type alias for a slice of ints

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	len_value := len(*h)
	val := (*h)[len_value-1]
	*h = (*h)[0 : len_value-1]
	return val
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}
func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return val
}
func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, num := range nums {
		heap.Push(minHeap, num)
		if len(*minHeap) > k {
			heap.Pop(minHeap)
		}
	}
	return (*minHeap)[0]
}
func findKthSmallest(nums []int, k int) int {
	maxHeap := &maxHeap{}
	heap.Init(maxHeap)
	for _, num := range nums {
		heap.Push(maxHeap, num)
		if len(*maxHeap) > k {
			heap.Pop(maxHeap)
		}
	}
	return (*maxHeap)[0]
}
```

### Explanation:
- **Finding \( k \) Largest**: The min-heap keeps the smallest \( k \) elements at the top, allowing us to discard smaller elements efficiently.
- **Finding \( k \) Smallest**: The max-heap stores the \( k \) smallest elements by inverting comparisons, allowing larger elements to be discarded.

This code will correctly find the \( k \) largest and \( k \) smallest elements in any unsorted list.



**top \( k \) most frequent elements** in an array, you can follow these steps:

### Approach

1. **Frequency Map**: 
   - Use a hash map (or dictionary) to count the frequency of each element in the array.

2. **Min-Heap**:
   - Use a min-heap of size \( k \) to keep track of the top \( k \) elements based on frequency.
   - As you iterate through the frequency map, maintain the heap by removing the element with the smallest frequency when the size exceeds \( k \).

3. **Result Extraction**:
   - The heap will contain the top \( k \) frequent elements after processing all elements.

### Time Complexity
- **O(N log k)**, where \( N \) is the number of elements in the array.

### Example

**Input:**
- Array: `[1, 1, 1, 2, 2, 3]`
- \( k \): `2`

**Output:**
- `[1, 2]`

**Explanation:**
- Element `1` appears 3 times.
- Element `2` appears 2 times.
- Element `3` appears 1 time.
- The top 2 most frequent elements are `1` and `2`.

### Go Code Implementation

```go
package main

import (
	"container/heap"
	"fmt"
)

// ElementFrequency holds the element and its frequency
type ElementFrequency struct {
	element   int
	frequency int
}

// MinHeap implements a min-heap for ElementFrequency
type MinHeap []ElementFrequency

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].frequency < h[j].frequency }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(ElementFrequency))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// Step 1: Frequency map
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	// Step 2: Min-Heap to keep top k elements
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	for element, freq := range frequencyMap {
		heap.Push(minHeap, ElementFrequency{element, freq})
		if minHeap.Len() > k {
			heap.Pop(minHeap)
		}
	}

	// Step 3: Extract the results from the heap
	result := make([]int, 0, k)
	for minHeap.Len() > 0 {
		result = append(result, heap.Pop(minHeap).(ElementFrequency).element)
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println("Top k frequent elements:", topKFrequent(nums, k)) // Output: [1, 2]
}
```