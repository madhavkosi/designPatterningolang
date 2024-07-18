### Strategy Pattern

The Strategy Pattern is a behavioral design pattern that enables selecting an algorithm's behavior at runtime. It defines a family of algorithms, encapsulates each one, and makes them interchangeable. The pattern allows the algorithm to vary independently from the clients that use it.

### Key Characteristics

1. **Encapsulation**: Encapsulates different algorithms in separate classes.
2. **Interchangeability**: Allows switching between different algorithms at runtime.
3. **Separation of Concerns**: Separates the algorithm implementation from the context that uses it.

### Implementation in Go

In Go, the Strategy Pattern can be implemented by defining an interface for the strategy and then creating concrete implementations of this interface.

### Example: Sorting Algorithms

### Step 1: Define the Strategy Interface

```go
package main

import "fmt"

// SortStrategy is the strategy interface for different sorting algorithms
type SortStrategy interface {
	Sort([]int)
}
```

### Step 2: Create Concrete Strategies

```go
package main

import "sort"

// BubbleSort is a concrete strategy for bubble sort
type BubbleSort struct{}

func (b *BubbleSort) Sort(data []int) {
	n := len(data)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Println("Sorted using Bubble Sort:", data)
}

// QuickSort is a concrete strategy for quick sort
type QuickSort struct{}

func (q *QuickSort) Sort(data []int) {
	sort.Ints(data)
	fmt.Println("Sorted using Quick Sort:", data)
}
```

### Step 3: Create the Context

```go
package main

// SortContext is the context that uses a SortStrategy
type SortContext struct {
	strategy SortStrategy
}

func (s *SortContext) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *SortContext) Sort(data []int) {
	s.strategy.Sort(data)
}
```

### Step 4: Use the Strategy

```go
package main

func main() {
	data := []int{64, 34, 25, 12, 22, 11, 90}

	context := &SortContext{}

	// Use bubble sort strategy
	bubbleSort := &BubbleSort{}
	context.SetStrategy(bubbleSort)
	context.Sort(data)

	// Use quick sort strategy
	data = []int{64, 34, 25, 12, 22, 11, 90} // Reset data
	quickSort := &QuickSort{}
	context.SetStrategy(quickSort)
	context.Sort(data)
}
```
### Explanation

1. **Strategy Interface**: `SortStrategy` defines the method `Sort` that all concrete strategies must implement.
2. **Concrete Strategies**: `BubbleSort` and `QuickSort` implement the `SortStrategy` interface.
3. **Context**: `SortContext` uses a `SortStrategy` and can switch strategies at runtime using `SetStrategy`.
4. **Usage**: In the `main` function, different sorting strategies are set and used to sort the same array.

### When to Use the Strategy Pattern

1. **Multiple Algorithms**: When a class has multiple behaviors, and you want to switch between them dynamically.
2. **Avoid Conditional Statements**: When you want to avoid complex conditional statements for selecting different behaviors.
3. **Encapsulation**: When you want to encapsulate related algorithms into separate classes.

### Benefits

- **Flexibility**: Easily switch between different algorithms at runtime.
- **Code Reusability**: Reuse individual strategies across different contexts.
- **Separation of Concerns**: Separate the algorithm implementation from the context.

### Drawbacks

- **Complexity**: Can introduce additional classes and interfaces, making the code more complex.
- **Overhead**: May introduce a slight overhead due to the additional abstraction.

The Strategy Pattern is useful for scenarios where you need to select and change algorithms dynamically. It promotes flexibility, reusability, and separation of concerns by encapsulating algorithms into separate classes.

