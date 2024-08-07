## Behavioral Pattern
Behavior patterns are concerted with algorithms and the assignment of responsibilities between objects. Behavioral patterns describe not just the patterns of objects or classes but also the patterns of communication between them.


### Strategy Pattern
The Strategy Pattern is a behavioral design pattern that enables selecting an algorithm's behavior at runtime. It defines a family of algorithms, encapsulates each one, and makes them interchangeable. The pattern allows the algorithm to vary independently from the clients that use it.

**Key Characteristics**
1. **Encapsulation**: Encapsulates different algorithms in separate classes.
2. **Interchangeability**: Allows switching between different algorithms at runtime.
3. **Separation of Concerns**: Separates the algorithm implementation from the context that uses it.

**Implementation in Go**
In Go, the Strategy Pattern can be implemented by defining an interface for the strategy and then creating concrete implementations of this interface.

**Example: Sorting Algorithms**
**Step 1: Define the Strategy Interface**
```go
package main

import "fmt"

// SortStrategy is the strategy interface for different sorting algorithms
type SortStrategy interface {
	Sort([]int)
}
```

**Step 2: Create Concrete Strategies**
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

**Step 3: Create the Context**
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

**Step 4: Use the Strategy**
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
**Explanation**
1. **Strategy Interface**: `SortStrategy` defines the method `Sort` that all concrete strategies must implement.
2. **Concrete Strategies**: `BubbleSort` and `QuickSort` implement the `SortStrategy` interface.
3. **Context**: `SortContext` uses a `SortStrategy` and can switch strategies at runtime using `SetStrategy`.
4. **Usage**: In the `main` function, different sorting strategies are set and used to sort the same array.

**When to Use the Strategy Pattern**
1. **Multiple Algorithms**: When a class has multiple behaviors, and you want to switch between them dynamically.
2. **Avoid Conditional Statements**: When you want to avoid complex conditional statements for selecting different behaviors.
3. **Encapsulation**: When you want to encapsulate related algorithms into separate classes.

**Benefits**
- **Flexibility**: Easily switch between different algorithms at runtime.
- **Code Reusability**: Reuse individual strategies across different contexts.
- **Separation of Concerns**: Separate the algorithm implementation from the context.

**Drawbacks**
- **Complexity**: Can introduce additional classes and interfaces, making the code more complex.
- **Overhead**: May introduce a slight overhead due to the additional abstraction.

The Strategy Pattern is useful for scenarios where you need to select and change algorithms dynamically. It promotes flexibility, reusability, and separation of concerns by encapsulating algorithms into separate classes.


### Observer Pattern
The Observer Pattern is a behavioral design pattern that defines a one-to-many dependency between objects so that when one object (the subject) changes state, all its dependents (observers) are notified and updated automatically. This pattern is useful for implementing distributed event-handling systems.

**Key Characteristics**
1. **Subject Interface**: Maintains a list of observers and notifies them of any state changes.
2. **Observer Interface**: Defines an updating interface for objects that should be notified of changes in the subject.
3. **Loose Coupling**: Allows the subject and observers to interact with minimal knowledge of each other.
Sure! Let's consider another example of the Observer Pattern, this time in the context of a stock price monitoring system.

**Example: Stock Price Monitoring**
**Step 1: Define the Observer Interface**
```go
package main

// Observer is the interface for objects that should be notified of changes
type Observer interface {
	Update(stockName string, price float64)
}
```

**Step 2: Define the Subject Interface**
```go
package main

// Subject is the interface for the object that maintains a list of observers
type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}
```

**Step 3: Create Concrete Subject**
```go
package main

// StockData is the concrete subject that maintains state and notifies observers
type StockData struct {
	observers map[string][]Observer
	prices    map[string]float64
}

func NewStockData() *StockData {
	return &StockData{
		observers: make(map[string][]Observer),
		prices:    make(map[string]float64),
	}
}

func (s *StockData) RegisterObserver(stockName string, o Observer) {
	s.observers[stockName] = append(s.observers[stockName], o)
}

func (s *StockData) RemoveObserver(stockName string, o Observer) {
	observers := s.observers[stockName]
	for i, observer := range observers {
		if observer == o {
			s.observers[stockName] = append(observers[:i], observers[i+1:]...)
			break
		}
	}
}

func (s *StockData) NotifyObservers(stockName string) {
	for _, observer := range s.observers[stockName] {
		observer.Update(stockName, s.prices[stockName])
	}
}

func (s *StockData) SetPrice(stockName string, price float64) {
	s.prices[stockName] = price
	s.NotifyObservers(stockName)
}
```

**Step 4: Create Concrete Observers**
```go
package main

import "fmt"

// Investor is a concrete observer
type Investor struct {
	name string
}

func (i *Investor) Update(stockName string, price float64) {
	fmt.Printf("Investor %s notified. New price of %s: %.2f\n", i.name, stockName, price)
}
```

**Step 5: Use the Observer Pattern**
```go
package main

func main() {
	stockData := NewStockData()

	investor1 := &Investor{name: "Alice"}
	investor2 := &Investor{name: "Bob"}

	stockData.RegisterObserver("GOOG", investor1)
	stockData.RegisterObserver("GOOG", investor2)
	stockData.RegisterObserver("AAPL", investor1)

	stockData.SetPrice("GOOG", 1500.0)
	stockData.SetPrice("AAPL", 300.0)
	stockData.SetPrice("GOOG", 1550.0)
}
```

**Explanation**
1. **Observer Interface**: `Observer` defines the `Update` method that all concrete observers must implement.
2. **Subject Interface**: `Subject` defines methods for registering, removing, and notifying observers.
3. **Concrete Subject**: `StockData` maintains state and notifies observers of any changes in stock prices. It uses maps to manage multiple stock prices and their respective observers.
4. **Concrete Observer**: `Investor` implements the `Observer` interface to get updates from the subject. It has a `name` field to distinguish between different investors.
5. **Usage**: In the `main` function, observers (investors) are registered to the subject (stock data). They get updates whenever the stock prices change.

**When to Use the Observer Pattern**
1. **State Changes**: When an object needs to notify other objects about state changes.
2. **Event Handling**: When implementing distributed event-handling systems.
3. **Loose Coupling**: When you want to maintain loose coupling between objects.

**Benefits**
- **Loose Coupling**: Reduces the dependencies between objects.
- **Scalability**: Easily add or remove observers without changing the subject.
- **Reusability**: Reuse individual observers across different subjects.

**Drawbacks**
- **Complexity**: Can introduce additional complexity with the management of observers.
- **Performance**: May lead to performance issues with a large number of observers.

The Observer Pattern is a powerful design pattern for creating event-driven systems and ensuring that changes in one object can be communicated to dependent objects without tight coupling. It is widely used in various applications such as GUI frameworks, real-time systems, and distributed event handling.

### Chain of Responsibility Pattern

The Chain of Responsibility Pattern is a behavioral design pattern that allows an object to pass a request along a chain of potential handlers until the request is handled. Each handler in the chain can either handle the request or pass it to the next handler in the chain. This pattern decouples the sender of the request from its receivers.

**Key Characteristics**

1. **Chain of Handlers**: Handlers are linked to form a chain.
2. **Decoupling**: The sender of the request is decoupled from the receivers.
3. **Request Passing**: Each handler decides whether to handle the request or pass it to the next handler in the chain.

**When to Use the Chain of Responsibility Pattern**

1. **Multiple Handlers**: When multiple objects can handle a request, and you want to decouple the sender from the receivers.
2. **Dynamic Handler Chain**: When you want to dynamically specify the chain of handlers at runtime.
3. **Request Processing**: When you want to process a request with a sequence of handlers.

**Benefits**

- **Decoupling**: Decouples the sender of a request from its receivers.
- **Flexibility**: Allows adding or removing handlers dynamically.
- **Responsibility Sharing**: Multiple handlers can process the request without tight coupling.

**Drawbacks**

- **Potential Performance Issues**: Can introduce performance issues if the chain is long or the handlers are slow.
- **Complexity**: Can make the code more complex due to the setup of the chain.

The Chain of Responsibility Pattern is useful for scenarios where multiple handlers can process a request, and you want to decouple the sender from the receivers. It provides flexibility and promotes responsibility sharing by allowing handlers to either process the request or pass it to the next handler in the chain.


**Example: Support Ticket System**

**Step 1: Define the Handler Interface**

```go
package main

// SupportHandler is the handler interface
type SupportHandler interface {
	SetNext(handler SupportHandler)
	HandleRequest(level int, message string)
}
```

**Step 2: Create Concrete Handlers**

```go
package main

import "fmt"

// BaseHandler is the base struct that implements the common logic for setting the next handler
type BaseHandler struct {
	next SupportHandler
}

func (h *BaseHandler) SetNext(next SupportHandler) {
	h.next = next
}

func (h *BaseHandler) HandleRequest(level int, message string) {
	if h.next != nil {
		h.next.HandleRequest(level, message)
	}
}

// LevelOneSupport is a concrete handler for Level 1 support
type LevelOneSupport struct {
	BaseHandler
}

func (h *LevelOneSupport) HandleRequest(level int, message string) {
	if level == 1 {
		fmt.Println("Level 1 Support: Handling request -", message)
	} else {
		h.BaseHandler.HandleRequest(level, message)
	}
}

// LevelTwoSupport is a concrete handler for Level 2 support
type LevelTwoSupport struct {
	BaseHandler
}

func (h *LevelTwoSupport) HandleRequest(level int, message string) {
	if level == 2 {
		fmt.Println("Level 2 Support: Handling request -", message)
	} else {
		h.BaseHandler.HandleRequest(level, message)
	}
}

// LevelThreeSupport is a concrete handler for Level 3 support
type LevelThreeSupport struct {
	BaseHandler
}

func (h *LevelThreeSupport) HandleRequest(level int, message string) {
	if level == 3 {
		fmt.Println("Level 3 Support: Handling request -", message)
	} else {
		h.BaseHandler.HandleRequest(level, message)
	}
}
```

**Step 3: Use the Chain of Responsibility**

```go
package main

func main() {
	levelOne := &LevelOneSupport{}
	levelTwo := &LevelTwoSupport{}
	levelThree := &LevelThreeSupport{}

	// Set up the chain: Level 1 -> Level 2 -> Level 3
	levelOne.SetNext(levelTwo)
	levelTwo.SetNext(levelThree)

	// Test the chain with different request levels
	levelOne.HandleRequest(1, "Password reset request")
	levelOne.HandleRequest(2, "Software installation request")
	levelOne.HandleRequest(3, "System outage report")
	levelOne.HandleRequest(4, "Unrecognized support request")
}
```

**Explanation**

1. **Handler Interface**: `SupportHandler` defines the `SetNext` and `HandleRequest` methods that all concrete handlers must implement.
2. **Concrete Handlers**: `LevelOneSupport`, `LevelTwoSupport`, and `LevelThreeSupport` are concrete handlers that process support requests based on the complexity level. They extend `BaseHandler` to handle the chain logic.
3. **Base Handler**: `BaseHandler` provides common logic for setting the next handler and passing the request to the next handler in the chain.
4. **Usage**: In the `main` function, the handlers are linked to form a chain. Different support requests are processed by the appropriate handler based on the complexity level.
