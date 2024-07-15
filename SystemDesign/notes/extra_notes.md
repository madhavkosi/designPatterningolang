### MapReduce: Short Notes

**Definition**: MapReduce is a programming model for processing large datasets across a distributed cluster. It involves two main functions: `Map` and `Reduce`.

**Components**:
1. **Map Function**:
   - **Input**: Key-value pairs.
   - **Processing**: Processes each pair to produce intermediate key-value pairs.
   - **Output**: Intermediate key-value pairs.

2. **Shuffle and Sort**:
   - Groups and sorts intermediate key-value pairs by key, preparing them for the Reduce function.

3. **Reduce Function**:
   - **Input**: Intermediate key-value pairs.
   - **Processing**: Aggregates values for each key.
   - **Output**: Final key-value pairs.

**Example**: Word Count
- **Map**: Emits (word, 1) for each word.
- **Shuffle and Sort**: Groups counts by word.
- **Reduce**: Sums counts for each word, producing (word, total count).

### MapReduce Example: Large File Processing

**Scenario**: Counting the occurrences of each word in a large text file, such as a collection of books.

#### 1. **Map Function**
- **Input**: Split the large text file into chunks. Each chunk is processed independently.
- **Processing**: For each word in a chunk, emit a key-value pair (word, 1).
- **Output**: Intermediate key-value pairs.
  ```text
  Chunk 1: "Hello world hello"
  Map Output:
  (Hello, 1)
  (world, 1)
  (hello, 1)

  Chunk 2: "world of MapReduce"
  Map Output:
  (world, 1)
  (of, 1)
  (MapReduce, 1)
  ```

#### 2. **Shuffle and Sort**
- **Processing**: Group all intermediate key-value pairs by key, and sort them.
- **Output**: Grouped intermediate pairs.
  ```text
  (Hello, [1, 1])
  (world, [1, 1])
  (hello, [1])
  (of, [1])
  (MapReduce, [1])
  ```

#### 3. **Reduce Function**
- **Input**: Grouped intermediate key-value pairs.
- **Processing**: Sum the values for each key.
- **Output**: Final key-value pairs representing the word counts.
  ```text
  Reduce Input:
  (Hello, [1, 1])
  (world, [1, 1])
  (hello, [1])
  (of, [1])
  (MapReduce, [1])

  Reduce Output:
  (Hello, 2)
  (world, 2)
  (hello, 1)
  (of, 1)
  (MapReduce, 1)
  ```

### Summary
- **Map**: Processes each chunk of the large file to generate intermediate key-value pairs.
- **Shuffle and Sort**: Organizes and groups the intermediate pairs by key.
- **Reduce**: Aggregates the grouped pairs to produce the final word counts.

### Applications of Large File Processing with MapReduce
- **Web Indexing**: Processing vast amounts of web pages to create search indexes.
- **Log Analysis**: Analyzing logs from large-scale systems for insights.
- **Data Mining**: Extracting patterns from large datasets.
- **Machine Learning**: Preparing and processing large datasets for training models.


## Redis vs MemCached
| **Aspect**            | **Redis**                                                                                                                                   | **Memcached**                                                                                                                            |
|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------|
| **Definition**        | An in-memory data structure store, used as a database, cache, and message broker.                                                            | An in-memory key-value store for small chunks of arbitrary data (strings, objects) from results of database calls, API calls, or page rendering. |
| **Data Structures**   | Supports various data structures such as strings, hashes, lists, sets, sorted sets, bitmaps, hyperloglogs, and geospatial indexes.           | Primarily supports simple key-value pairs (strings only).                                                                                  |
| **Persistence**       | Offers persistence options, including snapshots and append-only file (AOF) logging, allowing data to be saved to disk.                      | Does not support persistence. Data is stored only in memory and is lost on restart.                                                       |
| **Replication**       | Supports master-slave replication, allowing data to be replicated to multiple slave nodes for redundancy and high availability.             | Supports a multi-threaded architecture but does not provide built-in replication.                                                         |
| **Scalability**       | Scales vertically (more memory, CPU) and horizontally with Redis Cluster, which allows partitioning across multiple Redis nodes.            | Scales horizontally by adding more instances, but lacks built-in clustering or sharding capabilities.                                     |
| **Use Cases**         | Suitable for use cases needing complex data structures, persistence, high availability, pub/sub messaging, and sophisticated data operations. | Suitable for simple caching scenarios where fast access to key-value pairs is needed, without the need for persistence or complex data structures. |
| **Performance**       | Excellent performance with support for more complex queries and data types. Offers pub/sub functionality, making it suitable for real-time analytics and messaging. | Very high performance for simple get and set operations due to its simplicity and optimized handling of key-value pairs.                    |
| **Memory Management** | Uses a sophisticated memory management model that includes eviction policies, data compression, and advanced memory handling features.      | Employs a straightforward slab allocator for memory management, which is simple but less flexible than Redis's approach.                   |
| **Advanced Features** | Includes Lua scripting, transactions, pub/sub messaging, geospatial support, and support for atomic operations on data structures.          | Focuses on simplicity and high performance for basic caching operations, without advanced features like scripting or complex data manipulation. |
| **Setup and Maintenance** | Requires more setup and maintenance due to its rich feature set and potential for more complex configurations.                                   | Easier to set up and maintain due to its simplicity and limited feature set.                                                              |


### Service Discovery with Apache Zookeeper

**Service discovery** is a mechanism used to find and connect to network services. It’s essential in dynamic and distributed environments, like microservices architectures, where services can scale up and down, and their locations can change frequently. Apache Zookeeper is a popular tool used for service discovery.

### Apache Zookeeper's Role in Service Discovery

Apache Zookeeper is a distributed coordination service that provides various primitives for building distributed applications. It is widely used for service discovery due to its ability to maintain configuration information, naming, synchronization, and providing group services.

#### How Zookeeper Works in Service Discovery

1. **Service Registration:**
   - Each service instance registers itself with Zookeeper when it starts. This involves creating a znode (a data node in Zookeeper’s hierarchical namespace) that contains information about the service instance, such as its IP address, port, and other metadata.
   - Zookeeper nodes (znodes) can store metadata and state information that services use to find each other.

2. **Service Health Monitoring:**
   - Zookeeper periodically checks the health of the registered service instances. If a service instance fails or becomes unresponsive, Zookeeper will automatically remove its znode, ensuring that only healthy instances are discoverable.

3. **Service Lookup:**
   - When a client (like a chat application user) needs to connect to a service, it queries Zookeeper to find the best available service instance based on specific criteria (e.g., geographical location, server capacity, latency).
   - Zookeeper responds with the details of the optimal service instance, enabling the client to establish a direct connection to it.


### Benefits of Using Zookeeper for Service Discovery

1. **Centralized Configuration Management:**
   - Zookeeper acts as a centralized repository for service information, making it easier to manage and update configurations.

2. **Automatic Failover:**
   - By constantly monitoring the health of registered services, Zookeeper can automatically handle failover, ensuring clients always connect to healthy service instances.

3. **Scalability:**
   - Zookeeper is designed to handle a large number of service registrations and queries, making it suitable for large-scale distributed systems.

4. **Consistency:**
   - Zookeeper guarantees sequential consistency, ensuring that clients always have a consistent view of the registered services.

## Detailed Steps Using Zookeeper in Service Discovery

#### Step-by-Step Process (Refer to Figure 11)

1. **User A Logs into the App:**
   - User A opens the chat application and logs in.

2. **Load Balancer Routes the Login Request to API Servers:**
   - The initial login request from User A is handled by a load balancer, which routes the request to one of the available API servers responsible for handling authentication.

3. **Authentication:**
   - The API server authenticates User A’s credentials against the user database.

4. **Service Discovery Finds the Best Chat Server:**
   - After successful authentication, the API server queries Zookeeper to find the best chat server for User A.
   - Zookeeper evaluates the available chat servers based on predefined criteria (e.g., geographical proximity to User A, current server load, server health status).

5. **Zookeeper Returns the Server Info to the API Server:**
   - Zookeeper returns the information of the most suitable chat server (e.g., Chat Server 2) to the API server.

6. **API Server Provides Chat Server Info to User A:**
   - The API server then sends the details of Chat Server 2 (e.g., IP address and port) back to User A.

7. **User A Connects to Chat Server 2 via WebSocket:**
   - Using the provided information, User A establishes a WebSocket connection directly to Chat Server 2 for real-time communication.

### Conclusion

Apache Zookeeper plays a crucial role in service discovery by providing a robust mechanism to register, monitor, and lookup services. In the context of a chat application, Zookeeper ensures that clients are connected to the most optimal chat server, enhancing performance and reliability. Its ability to manage configurations and monitor service health makes it an indispensable tool for dynamic and distributed environments.


## Various Server Type

| **Server Type**      | **CPU Cores** | **RAM**         | **Storage**      | **Bandwidth**     | **Requests per Second (RPS)**   |
|----------------------|---------------|-----------------|------------------|-------------------|---------------------------------|
| **Web Servers**      |               |                 |                  |                   |                                 |
| Small                | 2-4           | 4-8 GB          | 100 GB - 1 TB    | 1 TB/month        | 100 - 1,000                     |
| Medium               | 4-8           | 8-16 GB         | 1-5 TB           | 5 TB/month        | 1,000 - 5,000                   |
| Large                | 8-16          | 16-64 GB        | 5-10 TB          | 10 TB/month+      | 5,000 - 50,000                  |
| **Database Servers** |               |                 |                  |                   |                                 |
| Small                | 4-8           | 16-32 GB        | 1-5 TB SSD       | N/A               | 100 - 1,000                     |
| Medium               | 8-16          | 32-64 GB        | 5-10 TB SSD      | N/A               | 1,000 - 5,000                   |
| Large                | 16-32         | 64-256 GB       | 10-100 TB SSD    | N/A               | 5,000 - 50,000                  |
| **File Servers**     |               |                 |                  |                   |                                 |
| Small                | 2-4           | 4-8 GB          | 1-10 TB          | N/A               | N/A (Dependent on file size)    |
| Medium               | 4-8           | 8-16 GB         | 10-50 TB         | N/A               | N/A (Dependent on file size)    |
| Large                | 8-16          | 16-64 GB        | 50-500 TB+       | N/A               | N/A (Dependent on file size)    |
| **Application Servers** |           |                 |                  |                   |                                 |
| Small                | 4-8           | 8-16 GB         | 500 GB - 2 TB    | N/A               | 100 - 1,000                     |
| Medium               | 8-16          | 16-32 GB        | 2-5 TB           | N/A               | 1,000 - 5,000                   |
| Large                | 16-32         | 32-128 GB       | 5-10 TB+         | N/A               | 5,000 - 50,000                  |
| **Cloud Servers**    |               |                 |                  |                   |                                 |
| Small                | 2-8           | 4-16 GB         | 100 GB - 1 TB SSD| Variable          | 100 - 1,000                     |
| Medium               | 8-32          | 16-64 GB        | 1-10 TB SSD      | Variable          | 1,000 - 10,000                  |
| Large                | 32-64+        | 64-512+ GB      | 10+ TB SSD       | Variable          | 10,000 - 100,000+               |

### Notes:
- **Requests per Second (RPS):** The RPS values are approximate and can vary significantly based on the server's specific hardware, software configuration, and workload characteristics.
- **Storage:** SSDs are preferred for better performance, especially for database and application servers.
- **Bandwidth:** Web servers typically have higher bandwidth requirements due to data transfer to/from users. File and database servers may have lower bandwidth needs but high internal network speeds.
- **Scalability:** Cloud servers offer the advantage of scalability, allowing resources to be adjusted dynamically based on demand.

This table provides a general guideline. The actual performance and capacity can vary based on specific use cases, optimizations, and real-world conditions.