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


** Here the Factory Method Pattern is a suitable design pattern to use**


### ParkingSpot Abstract Class
- **Purpose:** Generic template for parking spots.
- **Derived Classes:**
  1. **HandicappedSpot:** For vehicles with a handicapped permit.
  2. **CompactSpot:** For small vehicles.
  3. **LargeSpot:** For large vehicles like SUVs and trucks.
  4. **MotorcycleSpot:** Specifically for motorcycles.
- **Common Properties:**
  - `spotID`, `isOccupied`, `vehicleType`
- **Common Methods:**
  - `assignVehicle(vehicle)`, `removeVehicle()`, `isAvailable()`

  ![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/parkingSpot.png)

### Account Abstract Class
- **Purpose:** Represents a generic user account in the parking system.
- **Derived Classes:**
  1. **Admin:**
     - **Role:** Manages the system, including user and spot management.
     - **Permissions:** Full access to system features and settings.
  2. **ParkingAgent:**
     - **Role:** Manages day-to-day operations, such as assigning spots and handling issues.
     - **Permissions:** Limited to operational tasks, less access than Admin.
- **Common Properties:**
  - `accountID`, `username`, `password`
- **Common Methods:**
  - `login()`, `logout()`, `resetPassword()`

    ![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/account.png)


### DisplayBoard Class
- **Purpose:** Shows the availability of different types of parking spots and the number of empty slots.
    ![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/displayBoard.png)

### Entrance and exit 
The Entrance class is responsible for generating the parking ticket whenever a vehicle arrives. It contains the ID attribute, since there are multiple entrances to the parking lot. It also has the getTicket() method.

The Exit class is responsible for validating the parking ticket’s payment status before allowing the vehicle to exit the parking lot. It contains the ID attribute, since there are multiple exits to the parking lot. It also has the validateTicket() method.
    ![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/entrance.png)

### ParkingTicket Class

- **Attributes:**
  - **entranceTime**: Records the time when the vehicle enters the parking lot.
  - **exitTime**: Records the time when the vehicle exits the parking lot.
  - **amount**: The total amount to be paid for the parking duration.
  - **paymentStatus**: Indicates whether the payment has been made (paid/unpaid).

These attributes are central to managing the details of a vehicle's stay in the parking system.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/parkingTicket.png)


- **ParkingLot Class**:
  - Manages entrance/exits, parking spots, and parking rates.

- **Enums**:
  - **PaymentStatus**: Paid, Unpaid, Canceled, Refunded.
  - **AccountStatus**: Active, Canceled, Closed.

- **Address Data Type**:
  - Stores street, city, state, country, and postal code.

- **Person Class**:
  - Contains name and address details.



### **Relationships Between Classes in the Parking Lot System**

#### Association
- **ParkingSpot ↔ Vehicle**: 
  - One-way association where a ParkingSpot is assigned to a Vehicle.
- **Vehicle ↔ ParkingTicket**: 
  - One-way association where a Vehicle is linked to a ParkingTicket.
- **Payment ↔ ParkingTicket**: 
  - Two-way association where a Payment is linked to a ParkingTicket and vice versa.

#### Composition
- **ParkingLot**:
  - Composes Entrance, Exit, ParkingRate, DisplayBoard, ParkingTicket, and ParkingSpot objects, meaning these elements are integral to the ParkingLot.
- **ParkingTicket**:
  - Composes a Payment object, indicating that the payment details are a part of the ticket.

#### Inheritance
- **Vehicle Class**:
  - Inherits Car, Truck, Van, and Motorcycle subclasses, representing different types of vehicles.
- **ParkingSpot Class**:
  - Inherits Handicapped, Compact, Large, and Motorcycle subclasses, representing different types of parking spots.
- **Payment Class**:
  - Inherits Cash and CreditCard subclasses, representing different payment methods.


![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/association.png)
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/designpattern/photos/composition.png)

