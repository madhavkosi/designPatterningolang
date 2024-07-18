### Understanding SOLID Principles in Go

The SOLID principles are a set of design principles intended to make software designs more understandable, flexible, and maintainable. Let's explore each of these principles in the context of Go (Golang).

---

### 1. Single Responsibility Principle (SRP)

**Definition:** A class should have only one reason to change, meaning it should have only one job or responsibility.

**Go Example:**
```go
package main

import (
    "fmt"
)

type User struct {
    Name  string
    Email string
}

// Notification handles sending notifications
type Notification struct{}

func (n *Notification) SendEmail(user User, message string) {
    // logic to send email
    fmt.Printf("Sending email to %s: %s\n", user.Email, message)
}

func main() {
    user := User{Name: "John Doe", Email: "john.doe@example.com"}
    notifier := Notification{}
    notifier.SendEmail(user, "Welcome to our service!")
}
```
In this example, `User` struct represents user data, and `Notification` struct handles the responsibility of sending notifications. This adheres to the SRP by separating concerns.

---

### 2. Open/Closed Principle (OCP)

**Definition:** Software entities should be open for extension but closed for modification.

**Go Example:**
```go
package main

import (
    "fmt"
)

type AreaCalculator interface {
    Area() float64
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    shapes := []AreaCalculator{
        Rectangle{Width: 5, Height: 10},
        Circle{Radius: 7},
    }
    
    for _, shape := range shapes {
        fmt.Println("Area:", shape.Area())
    }
}
```
The `AreaCalculator` interface allows for the addition of new shapes without modifying existing code, adhering to the OCP.

---

### 3. Liskov Substitution Principle (LSP)

**Definition:** Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.

**Go Example:**
```go
package main

import "fmt"

type Shape interface {
    Area() float64
}

type Square struct {
    Side float64
}

func (s Square) Area() float64 {
    return s.Side * s.Side
}

type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func PrintArea(shape Shape) {
    fmt.Println("Area:", shape.Area())
}

func main() {
    square := Square{Side: 4}
    rectangle := Rectangle{Width: 5, Height: 6}
    
    PrintArea(square)
    PrintArea(rectangle)
}
```
Both `Square` and `Rectangle` structs implement the `Shape` interface, and the `PrintArea` function can accept either without any issue, satisfying the LSP.

---

### 4. Interface Segregation Principle (ISP)

**Definition:** No client should be forced to depend on methods it does not use.

**Go Example:**
```go
package main

import "fmt"

type Printer interface {
    Print()
}

type Scanner interface {
    Scan()
}

type AllInOnePrinter struct{}

func (a AllInOnePrinter) Print() {
    fmt.Println("Printing document...")
}

func (a AllInOnePrinter) Scan() {
    fmt.Println("Scanning document...")
}

type SimplePrinter struct{}

func (s SimplePrinter) Print() {
    fmt.Println("Printing document...")
}

func main() {
    aioPrinter := AllInOnePrinter{}
    simplePrinter := SimplePrinter{}
    
    aioPrinter.Print()
    aioPrinter.Scan()
    
    simplePrinter.Print()
    // simplePrinter.Scan() // Not available, adhering to ISP
}
```
The interfaces are split into smaller, more specific interfaces, so clients only need to implement the methods they use.

---

### 5. Dependency Inversion Principle (DIP)

**Definition:** High-level modules should not depend on low-level modules. Both should depend on abstractions.

**Go Example:**
```go
package main

import "fmt"

// Abstraction
type Logger interface {
    Log(message string)
}

// Low-level module
type FileLogger struct{}

func (f FileLogger) Log(message string) {
    fmt.Println("Logging to file:", message)
}

// High-level module
type UserService struct {
    logger Logger
}

func (u UserService) CreateUser(name string) {
    // User creation logic
    u.logger.Log("User created: " + name)
}

func main() {
    fileLogger := FileLogger{}
    userService := UserService{logger: fileLogger}
    userService.CreateUser("John Doe")
}
```
In this example, `UserService` depends on the `Logger` interface, not on the `FileLogger` directly, adhering to DIP.

---
