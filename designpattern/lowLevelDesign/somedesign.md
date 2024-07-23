### Data Hiding in Object-Oriented Programming

**Definition**:
- Data hiding involves masking a class's internal operations, providing only an interface for interaction.
- Prevents unauthorized access or modification of a class's contents.

**Goals**:
- Ensure the internal workings of a class are not exposed.
- Enable communication between classes without revealing underlying algorithms.

### Components of Data Hiding

1. **Encapsulation**:
   - Binding data and methods within a single unit (class).
   - Hide object state and representation from the outside.
   - Declare class variables as private to restrict access.
   - Use public methods (getters and setters) to interact with the encapsulated data.

2. **Abstraction**:
   - Simplifies complexity by exposing only essential features.
   - Details on abstraction will be discussed in the next lesson.

### Implementing Encapsulation

- Create a class with private data members.
- Provide public methods for accessing and modifying these data members.
- Example of a `Movie` class with private attributes `title`, `year`, and `genre`.
  - Use public methods like `getTitle()` to access private variables.

### Advantages of Encapsulation

- Easier to modify and maintain classes.
- Control over which data members are hidden or accessible.
- Flexibility in setting variables as read-only or write-only.

### Summary

- Data hiding is crucial for secure and maintainable code in OOP.
- Encapsulation and abstraction are key techniques for achieving data hiding.



### Abstraction in Object-Oriented Programming

**Definition**:
- Abstraction simplifies a program's structure by exposing only necessary details and hiding irrelevant information.
- Focuses on what an object does, not how it does it.

**Example**:
- **TV Volume Button**: Increases volume without needing to understand the TV's internal circuitry.
- **Vehicle Accelerator**: Increases car speed without needing to understand the mechanical process.

**Implementation of Abstraction in Programming Languages**:
- Define a class with necessary attributes and methods.
- Hide implementation details within the class, exposing only the necessary interface.

Example: `Circle` class
```java
class Circle {
  private double radius;
  private double pi;
  
  // Constructor
  public Circle() {
    radius = 0;
    pi = 3.142;
  }

  // Method to calculate area
  public double area() {
    return pi * radius * radius;
  }

  // Method to calculate perimeter
  public double perimeter() {
    return 2 * pi * radius;
  }
}
```
- Only the `radius` needs to be provided; `pi` is hidden and constant.
- `area()` and `perimeter()` methods are available for calculations without exposing implementation details.

**Advantages of Abstraction**:
- Reduces system complexity for the user.
- Makes the code extendable and reusable.
- Enhances modularity of the application.
- Improves code maintainability.

**Abstraction vs. Encapsulation**:

| **Abstraction** | **Encapsulation** |
|-----------------|--------------------|
| Focuses on the design level of the system. | Focuses on the implementation level of the system. |
| Hides unnecessary data to simplify the structure. | Restricts access to data to prevent misuse. |
| Highlights what the object performs. | Deals with the internal working of the object. |
| Achieved using interfaces and abstract classes. | Achieved using getter and setter functions. |

Next, we will explore another important principle of object-oriented programming—inheritance.


### Inheritance in Object-Oriented Programming

**Definition**:
- Inheritance allows the creation of a new class from an existing class.
- The new class (derived class) inherits all public attributes and methods from the existing class (base class).

**The IS-A Relationship**:
- Used when there is an IS-A relationship between objects.
- Example:
  - `Square` IS-A `Shape`
  - `Dog` IS-A `Animal`
  - `Car` IS-A `Vehicle`

### Modes of Inheritance
- **Access Modifiers**: Define the scope of data members and methods for other classes and the main program.

### Types of Inheritance

1. **Single Inheritance**:
   - One class extends from a single parent class.
   - Example: `FuelCar` IS-A `Vehicle`

2. **Multiple Inheritance**:
   - A class is derived from more than one base class.
   - Example: `HybridCar` IS-A `FuelCar` and IS-A `ElectricCar`
   - Note: Not supported in Java, C#, and JavaScript through classes.

3. **Multi-Level Inheritance**:
   - A class is derived from another class that is itself derived from another class.
   - Example:
     - `FuelCar` IS-A `Vehicle`
     - `GasolineCar` IS-A `FuelCar`

4. **Hierarchical Inheritance**:
   - More than one class extends from the same base class.
   - Example:
     - `FuelCar` IS-A `Vehicle`
     - `ElectricCar` IS-A `Vehicle`

5. **Hybrid Inheritance**:
   - A combination of more than one type of inheritance.
   - Example:
     - `FuelCar` IS-A `Vehicle`
     - `ElectricCar` IS-A `Vehicle`
     - `HybridCar` IS-A `FuelCar` and IS-A `ElectricCar`


### Advantages of Inheritance

1. **Reusability**:
   - Avoids duplication of methods in child classes that also occur in parent classes.

2. **Code Modification**:
   - Ensures changes are localized and avoids inconsistencies.

3. **Extensibility**:
   - Allows extending the base class as per the derived class requirements.
   - Facilitates upgrading or enhancing specific parts without changing core attributes.

4. **Data Hiding**:
   - The base class can keep some data private, preventing alteration by the derived class (encapsulation).

Next, we will learn about another important principle of object-oriented programming—polymorphism.



### Polymorphism in Object-Oriented Programming

**Introduction to Polymorphism**:
- Derived from Greek words "poly" (many) and "morph" (forms).
- Allows an object to take many forms and behaviors.
- Example: `Animal` class with different implementations of a `makeNoise` method for `Lion`, `Deer`, `Dog`, etc.

### Types of Polymorphism

1. **Dynamic Polymorphism (Runtime Polymorphism)**:
   - Methods have the same name, return type, and parameters in the base and derived classes.
   - Call to an overridden method is decided at runtime.
   - Achieved through method overriding.

   **Method Overriding**:
   - Subclass provides a specific implementation of a method defined in its parent class.

   Example in Java:
   ```java
   class Animal {
       public void makeNoise() {
           System.out.println("Animal noise");
       }
   }

   class Lion extends Animal {
       @Override
       public void makeNoise() {
           System.out.println("Lion roars");
       }
   }
   ```

2. **Static Polymorphism (Compile-Time Polymorphism)**:
   - Achieved through method overloading or operator overloading.
   - Polymorphism is resolved at compile-time.

   **Method Overloading**:
   - Multiple methods in the same class with the same name but different parameters.

   Example in Java:
   ```java
   class MathOperations {
       public int add(int a, int b) {
           return a + b;
       }

       public int add(int a, int b, int c) {
           return a + b + c;
       }
   }
   ```

   **Operator Overloading**:
   - Operators act differently based on the data types of operands.
   - Supported in languages like C++ and Python, but not in Java and JavaScript.

   Example in C++:
   ```cpp
   class ComplexNumber {
       float real;
       float imaginary;

   public:
       ComplexNumber(float r, float i) : real(r), imaginary(i) {}

       ComplexNumber operator + (const ComplexNumber &c) {
           return ComplexNumber(real + c.real, imaginary + c.imaginary);
       }
   };
   ```

### Dynamic Polymorphism vs. Static Polymorphism

| **Static Polymorphism**                       | **Dynamic Polymorphism**                    |
|-----------------------------------------------|---------------------------------------------|
| Resolved at compile-time.                     | Resolved at runtime.                        |
| Achieved using method overloading.            | Achieved using method overriding.           |
| Increases code readability.                   | Provides specific implementation in derived classes. |
| Arguments must differ in type or number.      | Arguments must be the same.                 |
| Return type does not matter.                  | Return type must be the same.               |
| Private and sealed methods can be overloaded. | Private and sealed methods cannot be overridden. |
| Better performance due to compile-time binding. | Worse performance due to runtime binding.   |

### Advantages of Polymorphism

- **Flexibility**: Allows objects to be treated as instances of their parent class.
- **Reusability**: Promotes code reuse by allowing a single method to work with different types.
- **Maintainability**: Easier to manage and update the code.

### Summary

Polymorphism is a key concept in OOP that enhances flexibility, reusability, and maintainability by allowing objects to take on multiple forms. Understanding and implementing both dynamic and static polymorphism enables more robust and adaptable code design.