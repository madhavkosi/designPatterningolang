
**Types of Estimations:**
- **Load Estimation**: Requests per second, data volume, user traffic.
- **Storage Estimation**: Amount of storage needed.
- **Bandwidth Estimation**: Network bandwidth required.
- **Latency Estimation**: Response time and latency.
- **Resource Estimation**: Servers, CPUs, memory required.


### 1. Feature Expectations [5 min]

#### 1.1 Use Cases
- List specific use cases your service will address.

#### 1.2 Scenarios That Will Not Be Covered
- Define scenarios that are out of scope.

#### 1.3 User Demographics
- Identify who will use the service.

#### 1.4 User Estimates
- Estimate how many users will use the service.

#### 1.5 Usage Patterns
- Describe expected usage patterns (e.g., peak times, average usage).

### 2. Estimations [5 min]

#### 2.1 Throughput
- Estimate QPS (Queries Per Second) for read and write queries.

#### 2.2 Latency Expectations
- Define expected latency for read and write queries.

#### 2.3 Read/Write Ratio
- Determine the expected read/write ratio.

#### 2.4 Traffic Estimates
- Write:
  - QPS
  - Volume of data
- Read:
  - QPS
  - Volume of data

#### 2.5 Storage Estimates
- Calculate required storage capacity.

#### 2.6 Memory Estimates
- Cache usage:
  - Type of data to store in cache
  - RAM requirements
  - Number of machines needed
- Disk/SSD storage requirements

### 3. Design Goals [5 min]

#### 3.1 Latency and Throughput Requirements
- Define the required latency and throughput.

#### 3.2 Consistency vs Availability
- Decide on the balance between consistency (weak, strong, eventual) and availability (failover, replication).

### 4. High-Level Design [5-10 min]

#### 4.1 APIs for Read/Write Scenarios
- Define APIs for crucial components.

#### 4.2 Database Schema
- Outline the database schema.

#### 4.3 Basic Algorithm
- Describe the basic algorithm for data handling.

#### 4.4 High-Level Design for Read/Write Scenario
- Present the high-level design for read/write operations.

### 5. Deep Dive [15-20 min]

#### 5.1 Scaling the Algorithm
- Discuss how to scale the algorithm.

#### 5.2 Scaling Individual Components
- Availability, consistency, and scalability for each component.
- Patterns for consistency and availability.

#### 5.3 Component Integration
- DNS
- CDN (Push vs Pull)
- Load Balancers (Active-Passive, Active-Active, Layer 4, Layer 7)
- Reverse Proxy
- Application Layer Scaling (Microservices, Service Discovery)
- Database (RDBMS, NoSQL)
  - RDBMS:
    - Master-slave, Master-master, Federation, Sharding, Denormalization, SQL Tuning
  - NoSQL:
    - Key-Value, Wide-Column, Graph, Document
    - Fast-lookups:
      - RAM [Bounded size] => Redis, Memcached
      - AP [Unbounded size] => Cassandra, RIAK, Voldemort
      - CP [Unbounded size] => HBase, MongoDB, Couchbase, DynamoDB
- Caches:
  - Client caching, CDN caching, Webserver caching, Database caching, Application caching, Cache @Query level, Cache @Object level
  - Eviction policies:
    - Cache aside
    - Write through
    - Write behind
    - Refresh ahead
- Asynchronism:
  - Message queues
  - Task queues
  - Back pressure
- Communication:
  - TCP
  - UDP
  - REST
  - RPC

### 6. Justification [5 min]

#### 6.1 Throughput of Each Layer
- Justify throughput for each layer.

#### 6.2 Latency Between Layers
- Explain latency caused between each layer.

#### 6.3 Overall Latency Justification
- Provide an overall latency justification.