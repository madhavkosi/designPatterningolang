## Parking Lot
### 1. **Vehicle Component**
   - **Abstract Class: `Vehicle`**
     - Attributes: `licensePlate`, `entryTime`, `exitTime`
     - Methods: `getSize()`, `getEntryTime()`, `getExitTime()`, `calculateParkingFee()`
   
   - **Concrete Classes:**
     - `Car` (inherits from `Vehicle`)
     - `Truck` (inherits from `Vehicle`)
     - `Van` (inherits from `Vehicle`)
     - `Motorcycle` (inherits from `Vehicle`)


**Enumeration vs. Abstract Class for Vehicle Representation:**

1. **Enumeration**:
   - Defines a set of named values (e.g., CAR, TRUCK, VAN, MOTORCYCLE).
   - Not ideal for object-oriented design as it requires code modification to add new types, violating the Open Closed Principle.
   - Not scalable for evolving systems.

2. **Abstract Class**:
   - Serves as a base class that can't be instantiated, but can define common methods and properties for all vehicle types.
   - Allows for easy extension by creating new subclasses for each vehicle type.
   - Adheres to the Open Closed Principle, making it a more flexible and maintainable solution.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/vehicle.png)
