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
