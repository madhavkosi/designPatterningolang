## Structural Pattern
Structural patterns are concerned with how classes and objects are composed to form larger structures. 

### Facade Pattern

The Facade Pattern is a structural design pattern that provides a simplified interface to a complex subsystem. It defines a higher-level interface that makes the subsystem easier to use by hiding the complexities behind it. This pattern is particularly useful when working with complex libraries, APIs, or frameworks.

**Key Characteristics**

1. **Simplified Interface**: Provides a simple, high-level interface to a complex subsystem.
2. **Decoupling**: Decouples the client from the subsystem, reducing dependencies and making the code easier to maintain.
3. **Unified Interface**: Aggregates multiple interfaces of the subsystem into a single interface.

**Implementation in Go**

In Go, the Facade Pattern can be implemented by creating a struct that provides a simplified interface to the underlying subsystem.

**Example: Home Theater System**

**Step 1: Define the Subsystem Components**

```go
package main

import "fmt"

// Amplifier is a subsystem component
type Amplifier struct{}

func (a *Amplifier) On() {
	fmt.Println("Amplifier is on")
}

func (a *Amplifier) Off() {
	fmt.Println("Amplifier is off")
}

func (a *Amplifier) SetVolume(volume int) {
	fmt.Println("Amplifier volume set to", volume)
}

// DVDPlayer is a subsystem component
type DVDPlayer struct{}

func (d *DVDPlayer) On() {
	fmt.Println("DVD Player is on")
}

func (d *DVDPlayer) Off() {
	fmt.Println("DVD Player is off")
}

func (d *DVDPlayer) Play(movie string) {
	fmt.Println("Playing movie:", movie)
}

func (d *DVDPlayer) Stop() {
	fmt.Println("Stopping movie")
}

// Projector is a subsystem component
type Projector struct{}

func (p *Projector) On() {
	fmt.Println("Projector is on")
}

func (p *Projector) Off() {
	fmt.Println("Projector is off")
}

func (p *Projector) SetInput(input string) {
	fmt.Println("Projector input set to", input)
}
```

**Step 2: Create the Facade**

```go
package main

// HomeTheaterFacade is the facade that provides a simplified interface to the home theater subsystem
type HomeTheaterFacade struct {
	amp      *Amplifier
	dvd      *DVDPlayer
	projector *Projector
}

func NewHomeTheaterFacade(amp *Amplifier, dvd *DVDPlayer, projector *Projector) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:       amp,
		dvd:       dvd,
		projector: projector,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("Get ready to watch a movie...")
	h.amp.On()
	h.amp.SetVolume(10)
	h.dvd.On()
	h.dvd.Play(movie)
	h.projector.On()
	h.projector.SetInput("DVD")
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting movie theater down...")
	h.amp.Off()
	h.dvd.Stop()
	h.dvd.Off()
	h.projector.Off()
}
```

**Step 3: Use the Facade**

```go
package main

func main() {
	amp := &Amplifier{}
	dvd := &DVDPlayer{}
	projector := &Projector{}

	homeTheater := NewHomeTheaterFacade(amp, dvd, projector)
	homeTheater.WatchMovie("Inception")
	homeTheater.EndMovie()
}
```

**Explanation**

1. **Subsystem Components**: `Amplifier`, `DVDPlayer`, and `Projector` represent the complex subsystem components with their own interfaces.
2. **Facade**: `HomeTheaterFacade` provides a simplified interface to control the subsystem components. It aggregates the interfaces of the subsystem components and provides high-level methods like `WatchMovie` and `EndMovie`.
3. **Usage**: In the `main` function, the facade is used to control the home theater system. The client code interacts only with the facade, which simplifies the process of watching and ending a movie.

**When to Use the Facade Pattern**

1. **Complex Subsystems**: When you have a complex subsystem and you want to provide a simpler interface for client code.
2. **Decoupling**: When you want to decouple the client code from the subsystem to reduce dependencies and make the code easier to maintain.
3. **Unified Interface**: When you want to aggregate multiple interfaces of the subsystem into a single interface.

**Benefits**

- **Simplification**: Provides a simplified and easy-to-use interface to a complex subsystem.
- **Decoupling**: Decouples the client code from the subsystem, reducing dependencies.
- **Improved Maintainability**: Makes the code easier to maintain by hiding the complexities of the subsystem.

**Drawbacks**

- **Limited Flexibility**: May limit the flexibility of the subsystem by providing a simplified interface.
- **Overhead**: Introduces an additional layer that may add some overhead.

The Facade Pattern is useful for providing a simplified interface to a complex subsystem, decoupling the client code from the subsystem, and improving the maintainability of the code. It is widely used in various applications, such as APIs, libraries, and frameworks, to hide the complexities and provide a unified interface to the clients.






### Decorator Pattern

The Decorator Pattern is a structural design pattern that allows behavior to be added to individual objects, either statically or dynamically, without affecting the behavior of other objects from the same class. It provides a flexible alternative to subclassing for extending functionality.

**Key Characteristics**

1. **Composition**: Uses composition instead of inheritance to extend functionality.
2. **Single Responsibility**: Adheres to the Single Responsibility Principle by allowing functionality to be divided between classes with unique areas of concern.
3. **Transparency**: The decorator and the component it decorates share the same interface.

**Example: Pizza Ordering System**

**Step 1: Define the Component Interface**

```go
package main

import "fmt"

// Pizza is the component interface
type Pizza interface {
	Cost() float64
	Description() string
}
```

**Step 2: Create Concrete Components**

```go
package main

// PlainPizza is a concrete component
type PlainPizza struct{}

func (p *PlainPizza) Cost() float64 {
	return 5.0
}

func (p *PlainPizza) Description() string {
	return "Plain pizza"
}
```

**Step 3: Create Decorators**

```go
package main

// PizzaDecorator is a base decorator that embeds a Pizza component
type PizzaDecorator struct {
	pizza Pizza
}

func (d *PizzaDecorator) Cost() float64 {
	return d.pizza.Cost()
}

func (d *PizzaDecorator) Description() string {
	return d.pizza.Description()
}

// CheeseDecorator is a concrete decorator for adding cheese
type CheeseDecorator struct {
	PizzaDecorator
}

func NewCheeseDecorator(pizza Pizza) *CheeseDecorator {
	return &CheeseDecorator{
		PizzaDecorator: PizzaDecorator{pizza: pizza},
	}
}

func (d *CheeseDecorator) Cost() float64 {
	return d.pizza.Cost() + 1.5
}

func (d *CheeseDecorator) Description() string {
	return d.pizza.Description() + ", cheese"
}

// PepperoniDecorator is a concrete decorator for adding pepperoni
type PepperoniDecorator struct {
	PizzaDecorator
}

func NewPepperoniDecorator(pizza Pizza) *PepperoniDecorator {
	return &PepperoniDecorator{
		PizzaDecorator: PizzaDecorator{pizza: pizza},
	}
}

func (d *PepperoniDecorator) Cost() float64 {
	return d.pizza.Cost() + 2.0
}

func (d *PepperoniDecorator) Description() string {
	return d.pizza.Description() + ", pepperoni"
}

// MushroomDecorator is a concrete decorator for adding mushrooms
type MushroomDecorator struct {
	PizzaDecorator
}

func NewMushroomDecorator(pizza Pizza) *MushroomDecorator {
	return &MushroomDecorator{
		PizzaDecorator: PizzaDecorator{pizza: pizza},
	}
}

func (d *MushroomDecorator) Cost() float64 {
	return d.pizza.Cost() + 1.0
}

func (d *MushroomDecorator) Description() string {
	return d.pizza.Description() + ", mushrooms"
}
```

**Step 4: Use the Decorators**

```go
package main

func main() {
	pizza := &PlainPizza{}
	fmt.Printf("Cost: $%.2f, Description: %s\n", pizza.Cost(), pizza.Description())

	pizzaWithCheese := NewCheeseDecorator(pizza)
	fmt.Printf("Cost: $%.2f, Description: %s\n", pizzaWithCheese.Cost(), pizzaWithCheese.Description())

	pizzaWithCheeseAndPepperoni := NewPepperoniDecorator(pizzaWithCheese)
	fmt.Printf("Cost: $%.2f, Description: %s\n", pizzaWithCheeseAndPepperoni.Cost(), pizzaWithCheeseAndPepperoni.Description())

	pizzaWithAllToppings := NewMushroomDecorator(pizzaWithCheeseAndPepperoni)
	fmt.Printf("Cost: $%.2f, Description: %s\n", pizzaWithAllToppings.Cost(), pizzaWithAllToppings.Description())
}
```

**Explanation**

1. **Component Interface**: `Pizza` defines the methods `Cost` and `Description` that all concrete components and decorators must implement.
2. **Concrete Component**: `PlainPizza` implements the `Pizza` interface and represents a basic pizza.
3. **Base Decorator**: `PizzaDecorator` embeds a `Pizza` component and implements the `Pizza` interface, delegating calls to the wrapped component.
4. **Concrete Decorators**: `CheeseDecorator`, `PepperoniDecorator`, and `MushroomDecorator` extend `PizzaDecorator` to add functionality (cheese, pepperoni, and mushrooms) to the pizza. They override the `Cost` and `Description` methods to include their specific behavior.
5. **Usage**: In the `main` function, decorators are used to add cheese, pepperoni, and mushrooms to the pizza dynamically. Each decorator wraps the previous one, extending its functionality.

**When to Use the Decorator Pattern**

1. **Adding Responsibilities**: When you need to add responsibilities to objects dynamically and transparently.
2. **Extension**: When subclassing would create an explosion of subclasses to support every combination of behaviors.
3. **Single Responsibility**: When you want to divide functionality between classes with unique areas of concern.

**Benefits**

- **Flexibility**: Allows behaviors to be added and combined at runtime.
- **Adheres to Open/Closed Principle**: Classes can be extended without modifying existing code.
- **Single Responsibility Principle**: Each decorator class handles a specific concern.

**Drawbacks**

- **Complexity**: Can lead to a large number of small classes that are hard to understand and maintain.
- **Debugging**: Can be more difficult to debug due to the number of objects involved.

The Decorator Pattern is useful for scenarios where you need to add functionality to objects dynamically and transparently. It provides a flexible alternative to subclassing and adheres to design principles such as Single Responsibility and Open/Closed.