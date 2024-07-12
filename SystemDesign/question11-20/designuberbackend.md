# Designing Uber Backend

## 1. Introduction

**Uber** enables customers to book drivers for taxi rides using their personal cars. Both customers and drivers communicate through the Uber app on their smartphones.

## 2. Requirements and Goals

### User Types

1. **Drivers**
2. **Customers**

### Functional Requirements

- **Drivers**:
  - Notify the service of their current location and availability regularly.
  
- **Customers**:
  - View nearby available drivers.
  - Request a ride and notify nearby drivers of a customer ready for pickup.
  - Track the driver’s location in real-time once a ride is accepted.

- **Ride Management**:
  - Upon ride acceptance, both the customer and the driver can see each other's location.
  - The driver marks the journey complete upon reaching the destination and becomes available for the next ride.

### Non-Functional Requirements

- **Scalability**: Handle millions of users and drivers efficiently.
- **Real-Time**: Update locations and handle ride requests in real-time.
- **Reliability**: Ensure the service is reliable and highly available.
- **Low Latency**: Minimize delays in communication and location updates.

## 3. Capacity Estimation and Constraints

- **Users**:
  - 300M customers
  - 1M drivers
  - 1M daily active customers
  - 500K daily active drivers
  - 1M daily rides

- **Location Updates**:
  - Each active driver sends a location update every 3 seconds.


### 4. Basic System Design and Algorithm

We will modify the solution discussed in 'Designing Yelp' to adapt it for the "Uber" use cases. The key difference is that our QuadTree was not initially designed for frequent updates.

#### Issues with Dynamic Grid Solution

1. **Frequent Location Updates**:
   - Active drivers report their locations every three seconds.
   - Updating the QuadTree for each change in the driver's position requires significant time and resources.
   - Process to update a driver’s location:
     - Determine the right grid based on the driver’s previous location.
     - If the new position is outside the current grid, remove the driver from the current grid and reinsert into the correct grid.
     - If the new grid reaches its driver limit, repartition the grid.

2. **Real-time Location Propagation**:
   - Quick mechanism needed to propagate the current location of nearby drivers to active customers.
   - During a ride, the system needs to notify both driver and passenger of the car’s current location.

#### Challenges with QuadTree
- **Efficiency in Updates**:
  - Although the QuadTree helps find nearby drivers quickly, it does not guarantee fast updates.

### Key Points:
- Modifying the existing QuadTree solution for frequent updates.
- Addressing time and resource constraints for driver location updates.
- Ensuring real-time location updates for drivers and passengers.
- Improving efficiency in the QuadTree updates.


### Updating the QuadTree and DriverLocationHT

#### Problem with Frequent QuadTree Updates
- Updating the QuadTree with every driver location update is resource-intensive and time-consuming.
- Since drivers report their locations every three seconds, updates outnumber queries for nearby drivers.

#### Proposed Solution: Using a Hash Table for Latest Positions
- **Hash Table (DriverLocationHT)**: Store the latest reported positions of all drivers.
- **QuadTree Updates**: Update the QuadTree less frequently, ensuring it reflects the driver's location within 15 seconds.

#### Memory Requirement for DriverLocationHT
- Each record in the hash table needs to store:
  - DriverID: 3 bytes (for 1 million drivers)
  - Old latitude: 8 bytes
  - Old longitude: 8 bytes
  - New latitude: 8 bytes
  - New longitude: 8 bytes
  - **Total per record**: 35 bytes
- For 1 million drivers:
  - Memory needed = 1,000,000 drivers * 35 bytes/driver = 35,000,000 bytes (approximately 35 MB)

### Summary of Key Points:
- **Frequent Updates**: Instead of updating the QuadTree with every driver location report, maintain the latest position in a hash table (DriverLocationHT).
- **QuadTree Update Frequency**: Ensure the QuadTree is updated within 15 seconds to balance between update frequency and resource efficiency.
- **Memory Calculation**: For 1 million drivers, DriverLocationHT requires approximately 35 MB of memory.



### Bandwidth Consumption for Location Updates

#### Calculation of Bandwidth Usage
- Each update from a driver includes:
  - DriverID and location: \(3 + 16 = 19\) bytes
- Updates are received every three seconds from 500,000 daily active drivers.
- Bandwidth usage per three seconds:
  \[
  500,000 \text{ drivers} \times 19 \text{ bytes/driver} = 9,500,000 \text{ bytes} \approx 9.5 \text{ MB}
  \]

### Distributing DriverLocationHT for Scalability and Fault Tolerance

#### Reasons for Distribution
- **Scalability**: Even though one server can handle current memory and bandwidth requirements, distributing the load improves performance.
- **Fault Tolerance**: Distributing data ensures the system remains functional if one server fails.

#### Distribution Strategy
- **DriverID-based Distribution**: 
  - Distribute DriverLocationHT across multiple servers using DriverID, ensuring random and even distribution.

#### Functions of Driver Location Servers
1. **Broadcast Driver Updates**:
   - Upon receiving a location update, broadcast the information to all interested customers.

2. **Notify QuadTree Server**:
   - Notify the respective QuadTree server to refresh the driver's location every 15 seconds.

### Summary of Key Points:
- **Bandwidth Consumption**: Receiving location updates from 500,000 drivers every three seconds uses approximately 9.5 MB of bandwidth.
- **DriverLocationHT Distribution**:
  - Distribute across multiple servers for scalability and fault tolerance.
  - Use DriverID for random distribution.
- **Server Functions**:
  - Broadcast driver location updates to customers.
  - Notify QuadTree server for periodic updates.