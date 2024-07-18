
### Singleton Pattern

The Singleton Pattern is a creational design pattern that ensures a class has only one instance and provides a global point of access to that instance. This is useful when exactly one object is needed to coordinate actions across a system.

### Key Characteristics

1. **Single Instance**: Ensures that only one instance of the class exists.
2. **Global Access**: Provides a global point of access to the instance.
3. **Lazy Initialization**: The instance is created only when it is needed for the first time.

### Implementation in Go


```go
package singleton

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}
var singleton *single

type single struct {
	val string
}

func (s single) values() {
	fmt.Println("Abc")
}

func NewSingleObject() *single {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = &single{val: "abc"}
			fmt.Printf("new created \n")
		} else {
			fmt.Printf("already created \n")
		}
	} else {
		fmt.Printf("already created \n")
	}
	return singleton
}

```



### When to Use the Singleton Pattern

1. **Resource Management**: When you need to control access to a shared resource such as a configuration object, database connection, or a file.
2. **Global State**: When you need a single point of access to some global state or service.
3. **Control Access**: When you need to ensure that only one instance of a class is created to prevent conflicting operations or state.

### Factory Method
1. **Factory Method**:
   - Defines an interface for creating an object, but lets subclasses decide which class to instantiate.
   - Promotes loose coupling by reducing the dependency of application code on concrete classes.

### When to Use the Factory Method Pattern

1. **When the exact type of the object cannot be determined until runtime**:
   - If your application needs to decide which class to instantiate at runtime, the Factory Method pattern allows this decision to be deferred to subclasses or implementing classes.
   
   Example: A document editor that supports different types of documents (e.g., Word, PDF, Text) where the type of document to be created depends on user input.

2. **When you want to isolate the client from the concrete implementation classes**:
   - The pattern promotes loose coupling by ensuring that the client interacts with a common interface rather than directly with concrete classes.
   
   Example: A logging framework where the client uses a Logger interface and the concrete implementation (e.g., FileLogger, ConsoleLogger) is determined by a configuration.

3. **When you have a group of related classes and want to localize the knowledge of their instantiation**:
   - If you have several classes that are part of the same family and share a common interface or base class, the Factory Method pattern centralizes the creation logic.
   
   Example: GUI frameworks where different widgets (e.g., buttons, text fields) are created depending on the platform (e.g., Windows, macOS).

4. **When the creation process involves logic beyond simply instantiating a class**:
   - If the instantiation process involves more complex setup steps or validation, the Factory Method pattern encapsulates this complexity within the factory.
   
   Example: A network connection manager that requires specific setup steps for different types of connections (e.g., HTTP, FTP).

5. **When you anticipate future extensions to the creation process**:
   - The Factory Method pattern makes it easier to introduce new types of objects without modifying existing code.
   
   Example: An e-commerce platform where new types of payment methods can be added without changing the core order processing logic.

### Benefits of Using the Factory Method Pattern

- **Encapsulation of Object Creation**: It encapsulates the creation logic, making the codebase more modular and easier to manage.
- **Scalability**: New types of products can be added with minimal changes to existing code.
- **Maintainability**: Centralizes the instantiation logic, making it easier to update or change.
- **Loose Coupling**: Reduces dependencies between the client code and concrete classes.

### Example Scenarios

1. **Plugin Architecture**:
   - You have a media player that supports different types of media formats through plugins. The Factory Method pattern can help dynamically load and create the appropriate media plugin at runtime.

2. **Report Generation System**:
   - A reporting system that can generate different types of reports (e.g., PDF, Excel, HTML). The Factory Method pattern can determine which report generator to use based on user selection or configuration.

3. **Transport Logistics System**:
   - A logistics system that needs to create different types of transportation modes (e.g., Truck, Ship, Airplane) depending on the delivery requirements. The Factory Method can choose the appropriate transport class based on criteria like distance and cost.



Certainly! Here's a concise one-page implementation and explanation of the Factory Method pattern in Go, including the main usage example.

### Factory Pattern in Go

```go
package factory

import "fmt"

// Printer is the product interface
type Printer interface {
	Print()
}

// BlackAndWhitePrinter is a concrete product
type BlackAndWhitePrinter struct{}

func (b BlackAndWhitePrinter) Print() {
	fmt.Println("black printer")
}

// ColorPrinter is a concrete product
type ColorPrinter struct{}

func (c ColorPrinter) Print() {
	fmt.Println("color printer")
}

// PrinterFactory is the factory function
func PrinterFactory(PrintType string) (Printer, error) {
	switch PrintType {
	case "black printer":
		return BlackAndWhitePrinter{}, nil
	case "color printer":
		return ColorPrinter{}, nil
	default:
		return nil, fmt.Errorf("Printer type does not exist")
	}
}
```

### Usage Example

```go
package main

import (
	"fmt"
	"path/to/your/module/factory" // Replace with the correct import path
)

func main() {
	printer, err := factory.PrinterFactory("black printer")
	if err != nil {
		fmt.Println(err)
		return
	}
	printer.Print()

	printer, err = factory.PrinterFactory("color printer")
	if err != nil {
		fmt.Println(err)
		return
	}
	printer.Print()

	printer, err = factory.PrinterFactory("unknown printer")
	if err != nil {
		fmt.Println(err)
		return
	}
	printer.Print()
}
```