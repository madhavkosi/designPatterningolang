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