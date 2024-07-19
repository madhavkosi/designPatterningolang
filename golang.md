Sure! Let's dive into the basics and then the more advanced aspects of the `for` loop in Go. 

### Basic `for` Loop

The `for` loop in Go is the only loop construct and can be used in several forms.

#### Simple Loop
This is the most common form, similar to other languages.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
}
```

- `i := 0`: Initialization statement, sets up the loop variable.
- `i < 5`: Condition, the loop runs as long as this condition is true.
- `i++`: Post statement, executed at the end of every iteration.

#### Loop with Only Condition
If you omit the initialization and post statements, you get a while-like loop.

```go
package main

import "fmt"

func main() {
    i := 0
    for i < 5 {
        fmt.Println(i)
        i++
    }
}
```

#### Infinite Loop
A loop without any condition will run indefinitely.

```go
package main

func main() {
    for {
        // infinite loop
    }
}
```

### Loop with `range`

The `range` keyword is used to iterate over elements in various data structures.

#### Slice
```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    for i, num := range nums {
        fmt.Printf("Index: %d, Value: %d\n", i, num)
    }
}
```

- `i`: Index of the element.
- `num`: Value of the element.

#### Map
```go
package main

import "fmt"

func main() {
    m := map[string]int{"a": 1, "b": 2, "c": 3}
    for key, value := range m {
        fmt.Printf("Key: %s, Value: %d\n", key, value)
    }
}
```

### Advanced `for` Loop

#### Breaking and Continuing
You can control the loop using `break` and `continue`.

```go
package main

import "fmt"

func main() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break // exits the loop
        }
        if i%2 == 0 {
            continue // skips the current iteration
        }
        fmt.Println(i)
    }
}
```

#### Nested Loops
Loops can be nested to handle more complex iterations.

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }
}
```

#### Labels with Break and Continue
Labels can be used with `break` and `continue` to control outer loops.

```go
package main

import "fmt"

func main() {
outer:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if i*j > 4 {
                break outer // breaks out of the outer loop
            }
            fmt.Printf("i: %d, j: %d\n", i, j)
        }
    }
}
```

### Summary

- Basic form: `for init; condition; post { }`
- Condition-only: `for condition { }`
- Infinite loop: `for { }`
- Iterating with `range` over slices, maps, and other collections.
- Controlling loops with `break`, `continue`, and labels for more advanced flow control.

These constructs allow you to implement a wide variety of loop-based logic in Go.


Sure! Here are very short notes on `if` and `else` statements in Go:

### Basic `if` Statement
- Executes a block of code if the condition is true.
  ```go
  if condition {
      // code to execute if condition is true
  }
  ```

### `if-else` Statement
- Provides an alternative block of code if the condition is false.
  ```go
  if condition {
      // code to execute if condition is true
  } else {
      // code to execute if condition is false
  }
  ```

### `if-else if-else` Statement
- Allows multiple conditions to be checked sequentially.
  ```go
  if condition1 {
      // code to execute if condition1 is true
  } else if condition2 {
      // code to execute if condition2 is true
  } else {
      // code to execute if both conditions are false
  }
  ```

### `if` with Short Statement
- Includes an initialization statement that can be used within the condition.
  ```go
  if err := someFunction(); err != nil {
      // handle error
  }
  ```

### Summary Table

| Type                    | Syntax                                                                 | Description                              |
|-------------------------|------------------------------------------------------------------------|------------------------------------------|
| **Basic `if`**          | ```if condition { /*...*/ }```                                         | Executes code block if condition is true.|
| **`if-else`**           | ```if condition { /*...*/ } else { /*...*/ }```                        | Executes alternative code block if condition is false. |
| **`if-else if-else`**   | ```if condition1 { /*...*/ } else if condition2 { /*...*/ } else { /*...*/ }``` | Checks multiple conditions sequentially. |
| **`if` with Short Statement** | ```if initStmt; condition { /*...*/ }```                          | Includes an init statement within the condition. |

These notes cover the basic and slightly advanced usage of `if` and `else` statements in Go.

Sure! Here are concise notes on maps in Go:

### Basic Concepts

- **Definition**: A map is a collection of key-value pairs, where each key is unique.
- **Syntax**:
  ```go
  var m map[keyType]valueType
  ```

### Creating and Initializing Maps

- **Using `make`**:
  ```go
  m := make(map[string]int)
  ```
  
- **Using a map literal**:
  ```go
  m := map[string]int{"one": 1, "two": 2}
  ```

### Adding and Updating Elements

- **Add or update**:
  ```go
  m["key"] = value
  ```

### Accessing Elements

- **Get value**:
  ```go
  value := m["key"]
  ```

- **Check if key exists**:
  ```go
  value, ok := m["key"]
  if ok {
      // key exists
  } else {
      // key does not exist
  }
  ```

### Deleting Elements

- **Delete a key**:
  ```go
  delete(m, "key")
  ```

### Iterating Over Maps

- **Using `range`**:
  ```go
  for key, value := range m {
      fmt.Println(key, value)
  }
  ```

### Summary Table

| Operation                | Syntax                                                        | Description                          |
|--------------------------|---------------------------------------------------------------|--------------------------------------|
| **Create with `make`**   | ```m := make(map[keyType]valueType)```                        | Creates an empty map.                |
| **Create with literal**  | ```m := map[keyType]valueType{"key1": value1, "key2": value2}``` | Creates and initializes a map.       |
| **Add/Update element**   | ```m["key"] = value```                                        | Adds or updates a key-value pair.    |
| **Access element**       | ```value := m["key"]```                                       | Retrieves the value for a key.       |
| **Check existence**      | ```value, ok := m["key"]```                                   | Checks if a key exists.              |
| **Delete element**       | ```delete(m, "key")```                                        | Removes a key-value pair.            |
| **Iterate over map**     | ```for key, value := range m { /*...*/ }```                   | Iterates over all key-value pairs.   |

These notes cover the basic and common operations you can perform with maps in Go.


Sure! Here are concise notes on defining variables in Go:

### Basic Concepts

- **Variable Definition**: In Go, variables can be defined using the `var` keyword or the shorthand `:=` operator.

### Using `var`

- **Single Variable**:
  ```go
  var x int
  ```

- **With Initialization**:
  ```go
  var x int = 10
  ```

- **Type Inference**:
  ```go
  var x = 10  // Type is inferred as int
  ```

- **Multiple Variables**:
  ```go
  var a, b, c int
  ```

- **Multiple Variables with Initialization**:
  ```go
  var a, b, c = 1, 2, 3
  ```

### Using Short Variable Declaration

- **Single Variable**:
  ```go
  x := 10
  ```

- **Multiple Variables**:
  ```go
  a, b, c := 1, 2, 3
  ```

### Constants

- **Defining Constants**:
  ```go
  const pi = 3.14
  ```

- **Typed Constants**:
  ```go
  const hello string = "Hello, World!"
  ```

### Summary Table

| Type                            | Syntax                                               | Description                                   |
|---------------------------------|------------------------------------------------------|-----------------------------------------------|
| **Single variable with `var`**  | ```var x int```                                      | Declares a variable with type int.            |
| **With initialization**         | ```var x int = 10```                                 | Declares and initializes a variable.          |
| **Type inference**              | ```var x = 10```                                     | Declares with inferred type (int).            |
| **Multiple variables**          | ```var a, b, c int```                                | Declares multiple variables of type int.      |
| **Multiple with initialization**| ```var a, b, c = 1, 2, 3```                          | Declares and initializes multiple variables.  |
| **Short declaration**           | ```x := 10```                                        | Declares and initializes using shorthand.     |
| **Multiple with shorthand**     | ```a, b, c := 1, 2, 3```                             | Declares multiple variables with shorthand.   |
| **Constant**                    | ```const pi = 3.14```                                | Declares a constant.                          |
| **Typed constant**              | ```const hello string = "Hello, World!"```           | Declares a typed constant.                    |

These notes cover the basic ways to define variables and constants in Go.


Sure! Here are concise notes on structs in Go:

### Basic Concepts

- **Definition**: A struct is a composite data type that groups together variables under a single name.
- **Syntax**:
  ```go
  type StructName struct {
      Field1 FieldType1
      Field2 FieldType2
      // more fields...
  }
  ```

### Defining a Struct

- **Example**:
  ```go
  type Person struct {
      Name string
      Age  int
  }
  ```

### Creating and Initializing Structs

- **Zero Value Initialization**:
  ```go
  var p Person
  ```

- **Literal Initialization**:
  ```go
  p := Person{Name: "Alice", Age: 30}
  ```

- **Named Fields**:
  ```go
  p := Person{
      Name: "Bob",
      Age:  25,
  }
  ```

- **Partial Initialization**:
  ```go
  p := Person{Name: "Charlie"}
  ```

### Accessing and Modifying Fields

- **Access Fields**:
  ```go
  fmt.Println(p.Name, p.Age)
  ```

- **Modify Fields**:
  ```go
  p.Age = 31
  ```

### Anonymous Structs

- **Definition and Initialization**:
  ```go
  p := struct {
      Name string
      Age  int
  }{Name: "Eve", Age: 28}
  ```

### Embedded Structs

- **Definition**:
  ```go
  type Address struct {
      City, State string
  }

  type Employee struct {
      Name    string
      Address // embedded struct
      Age     int
  }
  ```

- **Usage**:
  ```go
  e := Employee{
      Name: "John",
      Address: Address{
          City:  "New York",
          State: "NY",
      },
      Age: 30,
  }
  fmt.Println(e.City, e.State) // Access embedded fields directly
  ```

### Methods on Structs

- **Definition**:
  ```go
  func (p Person) Greet() {
      fmt.Println("Hello, my name is", p.Name)
  }
  ```

- **Usage**:
  ```go
  p.Greet()
  ```

### Summary Table

| Type                         | Syntax                                                                 | Description                                      |
|------------------------------|------------------------------------------------------------------------|--------------------------------------------------|
| **Define Struct**            | ```type Person struct { Name string; Age int }```                      | Defines a struct named `Person` with fields.     |
| **Zero Value Initialization**| ```var p Person```                                                     | Creates a zero-initialized struct.               |
| **Literal Initialization**   | ```p := Person{Name: "Alice", Age: 30}```                              | Initializes struct with values.                  |
| **Access Fields**            | ```p.Name```                                                          | Accesses the `Name` field of the struct.         |
| **Modify Fields**            | ```p.Age = 31```                                                      | Modifies the `Age` field of the struct.          |
| **Anonymous Struct**         | ```p := struct { Name string; Age int }{Name: "Eve", Age: 28}```       | Defines and initializes an anonymous struct.     |
| **Embedded Struct**          | ```type Employee struct { Address; Name string; Age int }```           | Embeds one struct within another.                |
| **Method on Struct**         | ```func (p Person) Greet() { fmt.Println(p.Name) }```                  | Defines a method on a struct type.               |

These notes cover the basic and some advanced usage of structs in Go.


Sure! Here are concise notes on interfaces in Go:

### Basic Concepts

- **Definition**: An interface is a type that specifies a set of method signatures but does not implement them. Types that implement all methods of an interface are said to satisfy that interface.
- **Syntax**:
  ```go
  type InterfaceName interface {
      Method1(param1 paramType1) returnType1
      Method2(param2 paramType2) returnType2
      // more methods...
  }
  ```

### Defining an Interface

- **Example**:
  ```go
  type Shape interface {
      Area() float64
      Perimeter() float64
  }
  ```

### Implementing an Interface

- A type implements an interface by providing implementations for all of the interface's methods.
- **Example**:
  ```go
  type Rectangle struct {
      Width, Height float64
  }

  func (r Rectangle) Area() float64 {
      return r.Width * r.Height
  }

  func (r Rectangle) Perimeter() float64 {
      return 2 * (r.Width + r.Height)
  }
  ```

### Using an Interface

- **Example**:
  ```go
  func printShapeInfo(s Shape) {
      fmt.Println("Area:", s.Area())
      fmt.Println("Perimeter:", s.Perimeter())
  }
  ```

- **Usage**:
  ```go
  r := Rectangle{Width: 3, Height: 4}
  printShapeInfo(r)
  ```

### Empty Interface

- **Definition**: The empty interface `interface{}` can hold values of any type.
- **Example**:
  ```go
  var i interface{}
  i = 42
  fmt.Println(i)
  i = "hello"
  fmt.Println(i)
  ```

### Type Assertions

- **Syntax**: Used to retrieve the underlying value of an interface.
- **Example**:
  ```go
  var i interface{} = "hello"
  str, ok := i.(string)
  if ok {
      fmt.Println(str)
  }
  ```

### Type Switch

- **Syntax**: Used to determine the type of an interface value.
- **Example**:
  ```go
  func describe(i interface{}) {
      switch v := i.(type) {
      case int:
          fmt.Println("int:", v)
      case string:
          fmt.Println("string:", v)
      default:
          fmt.Println("unknown type")
      }
  }
  ```

### Summary Table

| Concept                  | Syntax                                                                  | Description                                      |
|--------------------------|-------------------------------------------------------------------------|--------------------------------------------------|
| **Define Interface**     | ```type Shape interface { Area() float64; Perimeter() float64 }```      | Defines an interface named `Shape` with methods. |
| **Implement Interface**  | ```func (r Rectangle) Area() float64 { /*...*/ }```                     | Implements `Area` method for `Rectangle` type.   |
| **Use Interface**        | ```func printShapeInfo(s Shape) { /*...*/ }```                          | Function that takes an interface as a parameter. |
| **Empty Interface**      | ```var i interface{}; i = 42; i = "hello"```                            | Holds values of any type.                        |
| **Type Assertion**       | ```str, ok := i.(string)```                                             | Retrieves the underlying value of an interface.  |
| **Type Switch**          | ```switch v := i.(type) { case int: /*...*/ }```                        | Determines the type of an interface value.       |

These notes cover the basic and some advanced usage of interfaces in Go.

Sure! Here are concise notes on sorting in Go, focusing on basic and advanced sorting techniques.

### Basic Concepts

- **Definition**: Sorting is the process of arranging elements in a certain order (ascending or descending).

### Sorting Built-in Types

- **Slice of Integers**:
  ```go
  import "sort"
  
  ints := []int{3, 1, 4, 1, 5, 9}
  sort.Ints(ints)
  fmt.Println(ints) // Output: [1 1 3 4 5 9]
  ```

- **Slice of Strings**:
  ```go
  import "sort"
  
  strs := []string{"apple", "orange", "banana"}
  sort.Strings(strs)
  fmt.Println(strs) // Output: ["apple", "banana", "orange"]
  ```

- **Slice of Floats**:
  ```go
  import "sort"
  
  floats := []float64{3.1, 1.4, 1.5, 9.2}
  sort.Float64s(floats)
  fmt.Println(floats) // Output: [1.4 1.5 3.1 9.2]
  ```

### Custom Sorting

- **Using `sort.Slice`**:
  ```go
  import "sort"
  
  people := []struct {
      Name string
      Age  int
  }{
      {"Alice", 30},
      {"Bob", 25},
      {"Charlie", 35},
  }
  
  sort.Slice(people, func(i, j int) bool {
      return people[i].Age < people[j].Age
  })
  fmt.Println(people) // Output: [{Bob 25} {Alice 30} {Charlie 35}]
  ```

### Sorting with Interface

- **Implementing `sort.Interface`**:
  ```go
  import "sort"
  
  type ByAge []struct {
      Name string
      Age  int
  }
  
  func (a ByAge) Len() int           { return len(a) }
  func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
  func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
  
  people := []struct {
      Name string
      Age  int
  }{
      {"Alice", 30},
      {"Bob", 25},
      {"Charlie", 35},
  }
  
  sort.Sort(ByAge(people))
  fmt.Println(people) // Output: [{Bob 25} {Alice 30} {Charlie 35}]
  ```

### Summary Table

| Type                          | Syntax                                                                                       | Description                                    |
|-------------------------------|----------------------------------------------------------------------------------------------|------------------------------------------------|
| **Sort Integers**             | ```sort.Ints(ints)```                                                                        | Sorts a slice of integers in ascending order.  |
| **Sort Strings**              | ```sort.Strings(strs)```                                                                     | Sorts a slice of strings in ascending order.   |
| **Sort Floats**               | ```sort.Float64s(floats)```                                                                  | Sorts a slice of floats in ascending order.    |
| **Custom Sort with `sort.Slice`** | ```sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })```      | Custom sorting using `sort.Slice` function.    |
| **Implementing `sort.Interface`** | ```type ByAge []struct { Name string; Age int }; func (a ByAge) Less(i, j int) bool { ... }``` | Custom sorting by implementing `sort.Interface`.|

These notes cover the basic and advanced techniques for sorting in Go.


Sure! Here are concise notes on pointers in Go:

### Basic Concepts

- **Definition**: A pointer is a variable that holds the memory address of another variable.
- **Syntax**:
  ```go
  var p *int  // p is a pointer to an int
  ```

### Getting the Address

- **Using `&` Operator**: The address-of operator `&` is used to get the address of a variable.
  ```go
  x := 42
  p := &x  // p now holds the address of x
  ```

### Dereferencing

- **Using `*` Operator**: The dereference operator `*` is used to access the value stored at the address a pointer is holding.
  ```go
  fmt.Println(*p)  // Prints the value of x (42)
  *p = 21          // Sets x to 21
  ```

### Working with Functions

- **Passing Pointers to Functions**: Pointers can be passed to functions to allow the function to modify the original variable.
  ```go
  func increment(x *int) {
      *x++
  }

  func main() {
      a := 10
      increment(&a)
      fmt.Println(a)  // Prints 11
  }
  ```

### Pointers to Structs

- **Accessing Struct Fields**: When using pointers to structs, you can use the `.` operator to access fields directly (Go automatically dereferences the pointer).
  ```go
  type Person struct {
      Name string
      Age  int
  }

  func main() {
      p := &Person{Name: "Alice", Age: 30}
      fmt.Println(p.Name)  // Prints "Alice"
      p.Age = 31
      fmt.Println(p.Age)   // Prints 31
  }
  ```

### Nil Pointers

- **Definition**: A pointer that does not point to any memory location is called a nil pointer.
  ```go
  var p *int  // p is nil
  if p == nil {
      fmt.Println("p is nil")
  }
  ```

### Pointer Arithmetic

- **Not Supported**: Go does not support pointer arithmetic (e.g., `p++` to move to the next memory location).

### Summary Table

| Operation                | Syntax                                          | Description                                       |
|--------------------------|-------------------------------------------------|---------------------------------------------------|
| **Declare Pointer**      | ```var p *int```                                | Declares a pointer to an int.                     |
| **Get Address**          | ```p = &x```                                    | Assigns the address of `x` to `p`.                |
| **Dereference**          | ```*p = 21```                                   | Sets the value at the address `p` holds to 21.    |
| **Function Argument**    | ```increment(&a)```                             | Passes the address of `a` to a function.          |
| **Pointer to Struct**    | ```p := &Person{Name: "Alice", Age: 30}```      | Creates a pointer to a `Person` struct.           |
| **Access Struct Field**  | ```p.Name```                                    | Accesses the `Name` field of the `Person` struct. |
| **Nil Pointer**          | ```var p *int```                                | Declares a nil pointer.                           |

### Full Example

Here’s a full example demonstrating various pointer operations:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func increment(x *int) {
    *x++
}

func main() {
    x := 42
    p := &x
    fmt.Println(*p) // Prints 42
    *p = 21
    fmt.Println(x)  // Prints 21

    a := 10
    increment(&a)
    fmt.Println(a) // Prints 11

    person := &Person{Name: "Alice", Age: 30}
    fmt.Println(person.Name) // Prints "Alice"
    person.Age = 31
    fmt.Println(person.Age)  // Prints 31

    var pNil *int
    if pNil == nil {
        fmt.Println("pNil is nil")
    }
}
```

This example covers pointer declaration, address-of operator, dereferencing, passing pointers to functions, working with pointers to structs, and checking for nil pointers.


Certainly! Let's dive into the `fmt` package in Go, which provides formatting functions for input and output. The `fmt` package is widely used for printing formatted strings to the console, reading input, and working with formatted I/O.

### Basic Concepts

- **Importing `fmt`**: To use the `fmt` package, you need to import it:
  ```go
  import "fmt"
  ```

### Printing Functions

#### Basic Printing

- **Print**: Prints the arguments as they are.
  ```go
  fmt.Print("Hello, World!")
  fmt.Print("Number:", 42)
  ```

- **Println**: Prints the arguments followed by a newline.
  ```go
  fmt.Println("Hello, World!")
  fmt.Println("Number:", 42)
  ```

- **Printf**: Prints formatted output.
  ```go
  name := "Alice"
  age := 30
  fmt.Printf("Name: %s, Age: %d\n", name, age)
  ```

### Formatting Verbs

- **String**: `%s`
- **Integer**: `%d`
- **Float**: `%f`
- **Boolean**: `%t`
- **Type-specific**: `%v` (default format), `%+v` (include field names for structs), `%#v` (Go-syntax representation)

#### Examples

```go
fmt.Printf("String: %s\n", "Hello")
fmt.Printf("Integer: %d\n", 123)
fmt.Printf("Float: %.2f\n", 3.14159)
fmt.Printf("Boolean: %t\n", true)
fmt.Printf("Default format: %v\n", struct{ Name string }{"Alice"})
fmt.Printf("Struct with field names: %+v\n", struct{ Name string }{"Alice"})
fmt.Printf("Go-syntax representation: %#v\n", struct{ Name string }{"Alice"})
```

### Scanning Functions

- **Scan**: Reads space-separated values.
  ```go
  var name string
  var age int
  fmt.Print("Enter your name and age: ")
  fmt.Scan(&name, &age)
  fmt.Printf("Name: %s, Age: %d\n", name, age)
  ```

- **Scanln**: Reads space-separated values, stopping at a newline.
  ```go
  var name string
  var age int
  fmt.Print("Enter your name and age: ")
  fmt.Scanln(&name, &age)
  fmt.Printf("Name: %s, Age: %d\n", name, age)
  ```

- **Scanf**: Reads formatted input.
  ```go
  var name string
  var age int
  fmt.Print("Enter your name and age: ")
  fmt.Scanf("%s %d", &name, &age)
  fmt.Printf("Name: %s, Age: %d\n", name, age)
  ```

### Example Program

Here’s a complete example demonstrating various `fmt` functions:

```go
package main

import (
    "fmt"
)

func main() {
    // Basic printing
    fmt.Print("Hello, World!")
    fmt.Println("Hello, World!")
    fmt.Printf("Name: %s, Age: %d\n", "Alice", 30)

    // Formatting verbs
    fmt.Printf("String: %s\n", "Hello")
    fmt.Printf("Integer: %d\n", 123)
    fmt.Printf("Float: %.2f\n", 3.14159)
    fmt.Printf("Boolean: %t\n", true)
    fmt.Printf("Default format: %v\n", struct{ Name string }{"Alice"})
    fmt.Printf("Struct with field names: %+v\n", struct{ Name string }{"Alice"})
    fmt.Printf("Go-syntax representation: %#v\n", struct{ Name string }{"Alice"})

    // Scanning input
    var name string
    var age int
    fmt.Print("Enter your name and age: ")
    fmt.Scan(&name, &age)
    fmt.Printf("Name: %s, Age: %d\n", name, age)

    fmt.Print("Enter your name and age: ")
    fmt.Scanln(&name, &age)
    fmt.Printf("Name: %s, Age: %d\n", name, age)

    fmt.Print("Enter your name and age: ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("Name: %s, Age: %d\n", name, age)
}
```

### Summary Table

| Function          | Syntax                                     | Description                                 |
|-------------------|--------------------------------------------|---------------------------------------------|
| **Print**         | `fmt.Print("Hello, World!")`               | Prints the arguments as they are.           |
| **Println**       | `fmt.Println("Hello, World!")`             | Prints the arguments followed by a newline. |
| **Printf**        | `fmt.Printf("Name: %s, Age: %d\n", name, age)` | Prints formatted output.                    |
| **Scan**          | `fmt.Scan(&name, &age)`                    | Reads space-separated values.               |
| **Scanln**        | `fmt.Scanln(&name, &age)`                  | Reads space-separated values, stopping at a newline. |
| **Scanf**         | `fmt.Scanf("%s %d", &name, &age)`          | Reads formatted input.                      |

These notes cover the basic and some advanced usage of the `fmt` package in Go for formatted I/O operations.



### Detailed Notes on Slices vs Arrays in Go

### Summary Table

| Feature           | Array                            | Slice                                     |
|-------------------|----------------------------------|-------------------------------------------|
| **Size**          | Fixed                            | Dynamic                                   |
| **Declaration**   | `var arr [3]int`                 | `var slice []int`                         |
| **Initialization**| `arr := [3]int{1, 2, 3}`         | `slice := []int{1, 2, 3}`                 |
| **Access**        | `arr[0]`                         | `slice[0]`                                |
| **Modification**  | `arr[0] = 10`                    | `slice[0] = 10`                           |
| **Slicing**       | `slice := arr[1:4]`              | `subSlice := slice[1:4]`                  |
| **Appending**     | Not allowed                      | `slice = append(slice, 4, 5)`             |
| **Memory**        | Fixed allocation                 | Dynamic allocation with possible reallocations |
| **Function Pass** | Passed by value                  | Passed by reference                       |

By understanding the differences and appropriate use cases for arrays and slices, you can make better decisions when designing your Go programs.