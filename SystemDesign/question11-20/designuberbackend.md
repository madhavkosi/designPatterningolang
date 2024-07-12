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


  ### Efficient Broadcasting of Driver's Location to Customers

#### Push Model with Notification Service
- **Publisher/Subscriber Model**:
  - When customers open the Uber app and query for nearby drivers, they are subscribed to updates from those drivers.
  - Maintain a list of subscribers (customers) for each driver.
  - Broadcast updates from DriverLocationHT to subscribed customers whenever a driver's location is updated.

#### Memory Requirement for Subscriptions
- **Assumptions**:
  - 1 million daily active customers.
  - 500,000 daily active drivers.
  - On average, each driver has 5 subscribers.
- **Memory Calculation**:
  - Store DriverID (3 bytes) and CustomerID (8 bytes) for each subscription.
  - Memory needed:
    \[
    500,000 \text{ drivers} \times 3 \text{ bytes} + 500,000 \text{ drivers} \times 5 \text{ customers/driver} \times 8 \text{ bytes/customer} = 1,500,000 \text{ bytes} + 20,000,000 \text{ bytes} = 21,500,000 \text{ bytes} \approx 21 \text{ MB}
    \]

#### Bandwidth Requirement for Broadcasting
- **Total Subscribers**:
  - 5 subscribers per driver for 500,000 drivers.
  - Total subscribers: \(5 \times 500,000 = 2,500,000\)
- **Bandwidth Calculation**:
  - Send DriverID (3 bytes) and location (16 bytes) every second.
  - Bandwidth needed:
    \[
    2,500,000 \text{ subscribers} \times 19 \text{ bytes/subscriber} = 47,500,000 \text{ bytes/second} \approx 47.5 \text{ MB/s}
    \]

### Efficient Implementation of the Notification Service

- **Implementation Methods**
  - **HTTP Long Polling**: 
    - Clients maintain a connection to the server, waiting for updates.
  - **Push Notifications**:
    - Server actively sends updates to clients.

### Adding New Publishers/Drivers

- **Dynamic Subscription**
  - **Tracking Customer's Area**
    - Customers subscribe to nearby drivers when opening the app.
    - Challenge: How to handle new drivers entering the area the customer is viewing.
    - **Solution Consideration**: Instead of pushing updates, allow clients to pull information from the server.

### Client Pull Model

- **Mechanism**
  - Clients send their current location to the server.
  - The server retrieves nearby drivers using the QuadTree structure.
  - The server returns the driver information to the client.
  - Clients update their screens with the received driver positions.
  
- **Polling Frequency**
  - Clients query the server every five seconds.
  - Reduces the number of round trips compared to a continuous push model.

### Grid Repartitioning Strategy

- **Grid Capacity Management**
  - **Repartitioning Trigger**
    - Do not repartition immediately upon reaching the maximum limit.
    - Allow grids to grow/shrink by an extra 10% before partitioning/merging.
    - Benefits: Reduces the frequency and load of partitioning/merging in high-traffic grids.


![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/uberBackend.gif)


### "Request Ride" Use Case

1. **Customer Ride Request**
   - Customer submits a ride request through the Uber app.

2. **Aggregator Server Handling**
   - An Aggregator server receives the ride request.
   - Aggregator server requests nearby drivers from QuadTree servers.

3. **Driver Selection Process**
   - Aggregator server collects the results from QuadTree servers.
   - Aggregator server sorts the drivers by their ratings.

4. **Notification to Drivers**
   - Aggregator server sends ride request notifications to the top three drivers simultaneously.
   - The first driver to accept the request is assigned the ride.
   - Other drivers receive a cancellation request.

5. **Handling Non-Responsive Drivers**
   - If none of the top three drivers respond, the Aggregator server moves to the next three drivers in the list.
   - The process repeats until a driver accepts the ride request.

6. **Customer Notification**
   - Once a driver accepts the request, the customer is notified of the assigned driver.

### Summary
- **Initial Steps**: Customer requests a ride, Aggregator server queries QuadTree servers.
- **Driver Sorting**: Drivers are sorted by ratings.
- **Notifications**: Top three drivers are notified, with cancellations sent to non-accepting drivers.
- **Fallback Mechanism**: If no response, the next set of drivers is notified.
- **Final Step**: Customer is informed of the assigned driver once a request is accepted.


### 5. Fault Tolerance and Replication

- **Replication Strategy**
  - **Primary-Secondary Model**:
    - Maintain replicas for both Driver Location servers and Notification servers.
    - Secondary servers can take over if the primary server fails.
  - **Persistent Storage**:
    - Use SSDs for fast IO operations.
    - Store data persistently to recover in case both primary and secondary servers fail.

### 6. Ranking

- **Multi-Factor Ranking**:
  - **Criteria**: Proximity, popularity, and relevance.
  - **Driver Ratings**:
    - Track overall ratings in the database and QuadTree.
    - Use an aggregated rating system (e.g., stars out of ten).

- **Ranking Process**:
  - **Query Process**:
    - Request top-rated drivers within a specified radius.
    - Each QuadTree partition returns the top 10 drivers with the highest ratings.
  - **Aggregation**:
    - Aggregator server compiles results from all partitions.
    - Determines the top 10 drivers from the combined list.
