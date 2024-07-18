
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

### Builder Pattern

The Builder Pattern is a creational design pattern that allows you to construct complex objects step by step. Unlike other creational patterns, the Builder Pattern doesn't require products to have a common interface. This pattern is particularly useful when creating an object involves many steps or when an object can be created in multiple configurations.

### Key Characteristics

1. **Step-by-Step Construction**: Constructs the object step by step.
2. **Complex Object Creation**: Useful for creating complex objects with multiple configurations.
3. **Separation of Concerns**: Separates the construction of an object from its representation.

### Implementation in Go

In Go, the Builder Pattern can be implemented using a combination of struct and methods to build complex objects. Here’s an example:

#### Example: Building a Computer

### Step 1: Define the Product

```go
package main

import "fmt"

// Computer is the complex object we want to build.
type Computer struct {
	CPU       string
	GPU       string
	RAM       int
	Storage   int
	PowerSupply string
}

func (c Computer) String() string {
	return fmt.Sprintf("Computer: CPU=%s, GPU=%s, RAM=%dGB, Storage=%dGB, PowerSupply=%s", c.CPU, c.GPU, c.RAM, c.Storage, c.PowerSupply)
}
```

### Step 2: Create the Builder

```go
package main

// ComputerBuilder is the builder responsible for constructing a Computer.
type ComputerBuilder struct {
	computer Computer
}

func NewComputerBuilder() *ComputerBuilder {
	return &ComputerBuilder{}
}

func (b *ComputerBuilder) SetCPU(cpu string) *ComputerBuilder {
	b.computer.CPU = cpu
	return b
}

func (b *ComputerBuilder) SetGPU(gpu string) *ComputerBuilder {
	b.computer.GPU = gpu
	return b
}

func (b *ComputerBuilder) SetRAM(ram int) *ComputerBuilder {
	b.computer.RAM = ram
	return b
}

func (b *ComputerBuilder) SetStorage(storage int) *ComputerBuilder {
	b.computer.Storage = storage
	return b
}

func (b *ComputerBuilder) SetPowerSupply(powerSupply string) *ComputerBuilder {
	b.computer.PowerSupply = powerSupply
	return b
}

func (b *ComputerBuilder) Build() Computer {
	return b.computer
}
```

### Step 3: Use the Builder

```go
package main

func main() {
	builder := NewComputerBuilder()
	computer := builder.SetCPU("Intel i7").
		SetGPU("NVIDIA GTX 3080").
		SetRAM(32).
		SetStorage(1000).
		SetPowerSupply("750W").
		Build()

	fmt.Println(computer)
}
```

### Explanation

1. **Computer Struct**: The `Computer` struct represents the complex object that we want to build.
2. **ComputerBuilder Struct**: The `ComputerBuilder` struct provides methods to set the various properties of the `Computer` object. Each method returns the builder itself to allow for method chaining.
3. **Build Method**: The `Build` method returns the final `Computer` object.

### When to Use the Builder Pattern

1. **Complex Construction**: When the construction process of an object is complex and involves many steps.
2. **Multiple Representations**: When you need to create different representations of the same object.
3. **Immutability**: When you want to ensure that an object is immutable once it is constructed.

### Benefits

- **Control over the Construction Process**: Provides fine-grained control over the construction process.
- **Readable Code**: Makes the code more readable and maintainable by separating the construction logic.
- **Reusability**: Allows reusability of the construction process for different types of products.

### Drawbacks

- **Overhead**: Can introduce additional complexity and overhead if the object construction is simple.
- **Verbose**: Can make the code more verbose compared to simple constructors or factory methods.

The Builder Pattern is especially useful when dealing with complex objects that require multiple steps to construct or when you need to create different variations of an object. It helps in maintaining a clear and concise construction process, making the code more modular and easier to understand.


### Prototype Pattern

The Prototype Pattern is a creational design pattern that allows you to create new objects by copying an existing object, known as the prototype. This pattern is particularly useful when the creation of an object is a costly operation and the object already exists in a state that can be cloned.

### Key Characteristics

1. **Clone Method**: Objects are created by copying an existing prototype.
2. **Prototype Interface**: Defines the method to clone objects.
3. **Reduced Overhead**: Useful for reducing the overhead of creating objects from scratch.

### Implementation in Go

In Go, the Prototype Pattern can be implemented by defining an interface with a `Clone` method and creating concrete types that implement this interface.

#### Example: Cloning Shapes

### Step 1: Define the Prototype Interface

```go
package main

import "fmt"

// Prototype is the interface that all prototypes must implement
type Prototype interface {
	Clone() Prototype
	GetDetails() string
}
```

### Step 2: Create Concrete Prototypes

```go
package main

// Circle is a concrete prototype
type Circle struct {
	Radius int
	Color  string
}

func (c *Circle) Clone() Prototype {
	return &Circle{
		Radius: c.Radius,
		Color:  c.Color,
	}
}

func (c *Circle) GetDetails() string {
	return fmt.Sprintf("Circle: Radius=%d, Color=%s", c.Radius, c.Color)
}

// Rectangle is another concrete prototype
type Rectangle struct {
	Width  int
	Height int
	Color  string
}

func (r *Rectangle) Clone() Prototype {
	return &Rectangle{
		Width:  r.Width,
		Height: r.Height,
		Color:  r.Color,
	}
}

func (r *Rectangle) GetDetails() string {
	return fmt.Sprintf("Rectangle: Width=%d, Height=%d, Color=%s", r.Width, r.Height, r.Color)
}
```

### Step 3: Use the Prototype

```go
package main

func main() {
	// Create an initial circle
	originalCircle := &Circle{
		Radius: 5,
		Color:  "Red",
	}

	// Clone the circle
	clonedCircle := originalCircle.Clone()

	// Create an initial rectangle
	originalRectangle := &Rectangle{
		Width:  10,
		Height: 20,
		Color:  "Blue",
	}

	// Clone the rectangle
	clonedRectangle := originalRectangle.Clone()

	fmt.Println(originalCircle.GetDetails())
	fmt.Println(clonedCircle.GetDetails())

	fmt.Println(originalRectangle.GetDetails())
	fmt.Println(clonedRectangle.GetDetails())
}
```

### Explanation

1. **Prototype Interface**: The `Prototype` interface declares the `Clone` method that all concrete prototypes must implement.
2. **Concrete Prototypes**: The `Circle` and `Rectangle` structs implement the `Prototype` interface by providing their own `Clone` method, which creates a copy of the object.
3. **Usage**: In the `main` function, we create initial instances of `Circle` and `Rectangle`, then clone these instances using the `Clone` method. The `GetDetails` method is used to print the details of the original and cloned objects.

### When to Use the Prototype Pattern

1. **Costly Object Creation**: When creating a new object is a costly operation in terms of resources or time.
2. **Multiple Configurations**: When an object can have one of several possible configurations, and you need to create multiple objects in different configurations.
3. **Decoupling**: When you want to decouple the creation of objects from their specific classes.

### Benefits

- **Performance**: Reduces the overhead of creating objects from scratch.
- **Flexibility**: Allows for easy creation of complex objects.
- **Decoupling**: Decouples the client code from the classes of the objects being instantiated.

### Drawbacks

- **Cloning Complexity**: Cloning complex objects with circular references or deep hierarchies can be challenging.
- **Memory Consumption**: May lead to increased memory consumption if many copies of large objects are created.

The Prototype Pattern is a useful design pattern when you need to create objects based on a template or when the cost of creating an object is high. It provides a mechanism to clone existing objects, allowing for flexible and efficient object creation.