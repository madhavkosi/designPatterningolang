### Architecture of Redis

Redis (REmote DIctionary Server) is designed for high performance, flexibility, and scalability. Its architecture allows it to serve as a database, cache, and message broker. Let's delve into the architecture of Redis, exploring its components, features, and the mechanisms it employs to achieve its goals.

#### 1. **In-Memory Data Store**
- **Primary Storage in RAM**: Redis stores all its data in RAM, providing extremely fast read and write operations.
- **Persistence Options**: To ensure durability, Redis offers:
  - **RDB (Redis Database Backup)**: Periodic snapshots of the dataset saved to disk.
  - **AOF (Append-Only File)**: **Logs every write** operation received by the server, which can be replayed to reconstruct the dataset.

#### 2. **Data Structures**
Redis supports various data structures that are stored in memory and manipulated through commands:
- **Strings**: Binary-safe strings.
- **Lists**: Collections of strings sorted by insertion order.
- **Sets**: Unordered collections of unique strings.
- **Sorted Sets**: Sets ordered by a score.
- **Hashes**: Maps between string fields and string values.
- **Bitmaps**: Bit-level operations on strings.
- **HyperLogLogs**: Probabilistic data structures for cardinality estimation.
- **Geospatial Indexes**: Data structures for geospatial data and querying.

#### 3. **Client-Server Model**
- **Clients**: Applications connect to the Redis server using client libraries available for various programming languages.
- **Server**: The Redis server processes client requests, performing operations on in-memory data structures.

#### 4. **Replication**
- **Master-Slave Replication**: Redis supports asynchronous replication where a master instance replicates data to one or more slave instances. This enhances read scalability and provides redundancy.
- **Replica of a Replica**: Slaves can replicate data from other slaves, allowing complex replication topologies.

#### 5. **Sharding**
- **Partitioning**: Redis supports horizontal scaling by partitioning data across multiple Redis instances. This can be achieved using techniques like consistent hashing.

#### 6. **Persistence**
- **Snapshotting (RDB)**: Periodically saves the dataset to disk. Snapshots are compact and fast to create, but some data loss can occur between snapshots.
- **Append-Only File (AOF)**: Logs each write operation. AOF can be rewritten in the background to keep it compact. This method provides more durability at the cost of increased disk I/O.

#### 7. **High Availability**
- **Redis Sentinel**: Provides high availability and monitoring:
  - **Monitoring**: Continuously checks if your master and slave instances are working as expected.
  - **Notification**: Alerts the system administrators in case of issues.
  - **Automatic Failover**: Promotes a slave to master if the master goes down, ensuring high availability.
  - **Configuration Provider**: Allows clients to connect to the currently available master instance.

#### 8. **Cluster Mode**
- **Redis Cluster**: Enables horizontal scaling by automatically sharding data across multiple nodes. Features include:
  - **Automatic Sharding**: Data is automatically divided among multiple nodes.
  - **High Availability**: Each shard has replicas, and the cluster can continue operating if a node fails, as long as a majority of the masters are up.
  - **Fault Tolerance**: Nodes can be added or removed on the fly, and the cluster will reconfigure itself.

#### 9. **Security**
- **Authentication**: Redis supports password-based authentication using the `AUTH` command.
- **Network Security**: Redis can be configured to bind to specific IP addresses and use SSL/TLS for encrypted communication.

#### 10. **Pub/Sub Messaging**
- **Publish/Subscribe**: Redis supports a publish/subscribe messaging paradigm where clients can subscribe to channels and receive messages in real-time.

### 1. Redis as a Single Point of Failure:
To handle Redis failures, implementing a **Redis cluster** is essential. In this setup:
- Redis can be horizontally scaled by adding more nodes, with **read replicas** distributed across nodes to balance the load.
- **Redis Sentinel** monitors the cluster for node failures. If the master node fails, Sentinel promotes one of the replicas to become the new master.
- The cluster ensures **high availability** and **partition tolerance**, as Redis can continue to serve requests from other nodes even if one node fails.
- **Write-ahead logs (WAL)** can be enabled to ensure data is persisted even if the system crashes.
- Synchronization happens when the primary node recovers, ensuring data consistency across the cluster.

Fallback in case of cache failure:
- If Redis is unavailable in certain scenarios, the system can query the **NoSQL database** directly, but this could introduce latency.
- To mitigate the performance bottleneck, **graceful degradation** can be implemented, where non-critical data can skip cache lookups and continue functioning without impacting the user experience.

### 2. Handling Hot Keys in Redis:
To prevent **hot key** overload on a single Redis node:
- **Sharding**: Distribute hot keys across multiple Redis nodes by using consistent hashing or sharding keys at the application level.
- **Replication**: Use Redis replicas to offload read traffic for hot keys. Implement intelligent routing to send read requests to replicas.
- **Caching strategies**: Use **local cache** (e.g., in-memory cache on application nodes) for frequently accessed hot keys to minimize Redis load.
- **Key splitting**: For large data objects that become hot, split the key into smaller subkeys, distributing the load more evenly across nodes.


### Redis Consistency and Network Partitioning:
In Redis, network partitioning can cause a **split-brain** scenario, where both the original master and a newly promoted replica accept writes. To prevent this:
- **Quorum-based voting**: Use Sentinel with quorum voting, where a majority of nodes must agree on promoting a new master. If not enough nodes are reachable to form a quorum, no promotion happens, preventing split-brain.
- **Automatic failover timeout**: Configure a longer timeout for Sentinel to avoid premature promotion during transient partitions.
- **Strong consistency**: In critical situations, Redis can operate in **append-only mode** (AOF) with fsync to disk after every write. This ensures data durability across nodes during recovery.
- **Recovery strategy**: When the original master re-joins, it gets demoted to a replica, and data is synchronized from the current master, ensuring consistency.
### Redis Internals

#### 1. **Event Loop**
Redis uses a single-threaded event loop for handling requests. This event loop is based on the Reactor pattern and uses multiplexing to handle multiple clients concurrently. This approach provides simplicity and high performance under typical workloads.

#### 2. **Data Eviction Policies**
When Redis reaches the maximum memory limit, it uses eviction policies to free up space. Policies include:
- **volatile-lru**: Evict the least recently used keys with an expiration set.
- **allkeys-lru**: Evict the least recently used keys, regardless of expiration.
- **volatile-random**: Evict random keys with an expiration set.
- **allkeys-random**: Evict random keys, regardless of expiration.
- **volatile-ttl**: Evict keys with the shortest time-to-live.

#### 3. **Snapshotting and Forking**
When Redis performs a snapshot, it uses a copy-on-write mechanism to fork the process. This allows the parent process to continue serving clients while the child process saves the snapshot to disk. This approach minimizes the impact on performance during snapshotting.

### Conclusion

Redis's architecture is designed for high performance, flexibility, and scalability. Its in-memory data store, combined with support for various data structures, replication, sharding, persistence options, and high availability features, makes it a versatile tool for numerous applications. Whether used as a cache, a database, or a message broker, Redis can be tailored to meet the specific needs of your application.


### Architecture of Memcached

Memcached is a high-performance, distributed memory caching system designed to speed up dynamic web applications by alleviating database load. It is simple, yet highly effective, and its architecture is built to ensure high throughput and low latency. Here's a detailed look at the architecture of Memcached:

#### 1. **Client-Server Model**

**Client**:
- **Library Integration**: Applications use Memcached client libraries to interact with the Memcached server. These libraries are available for various programming languages (e.g., Python, Java, PHP, C, Ruby).
- **Hashing**: The client typically uses a consistent hashing mechanism to distribute keys across multiple Memcached servers in a cluster.

**Server**:
- **Daemon Process**: Memcached runs as a daemon process on one or more servers. Each Memcached instance stores data in memory and handles requests from clients.
- **TCP/UDP**: Memcached supports both TCP and UDP protocols for communication.

#### 2. **Data Storage**

**Key-Value Store**:
- **In-Memory Storage**: Data is stored in memory, which allows for extremely fast access times. Each piece of data is stored as a key-value pair.
- **Ephemeral Storage**: Memcached does not provide persistence; data is stored only in RAM and is lost if the server is restarted.

**Slab Allocator**:
- **Memory Management**: Memcached uses a slab allocator to manage memory efficiently. Memory is pre-allocated into chunks of fixed sizes (slabs) to reduce fragmentation.
- **Chunks and Slabs**: Data is stored in chunks, which are grouped into slabs based on size classes. Each slab contains chunks of a specific size.

#### 3. **Distributed Architecture**

**Horizontal Scalability**:
- **Scaling Out**: Memcached can scale horizontally by adding more servers to the cluster. Each server operates independently, storing a portion of the data.
- **Client-Side Distribution**: The client library is responsible for distributing keys across the servers. This is typically done using consistent hashing to ensure even distribution and minimal rehashing when nodes are added or removed.

**No Replication**:
- **Simplicity**: Memcached does not natively support replication. Each server is responsible for a subset of the data, and there is no built-in redundancy.
- **External Management**: High availability and data redundancy must be managed externally, if needed (e.g., using client-side replication or a higher-level caching strategy).

#### 4. **Cache Management**

**Expiration**:
- **TTL (Time to Live)**: Each key-value pair can have an expiration time (TTL) set. After this time, the data is automatically evicted from the cache.
- **Lazy Eviction**: Expired items are only removed when accessed or when memory is needed for new data.

**Eviction Policy**:
- **Least Recently Used (LRU)**: Memcached uses an LRU eviction policy to manage memory. When the cache is full, the least recently used items are evicted to make space for new data.

#### 5. **Concurrency**

**Multi-Threading**:
- **Concurrent Requests**: Memcached supports handling multiple concurrent connections using a multi-threaded architecture.
- **Event Handling**: Uses the `libevent` library for asynchronous event notification, which helps manage multiple connections efficiently.

#### 6. **Security**

**Authentication**:
- **SASL Authentication**: Memcached supports Simple Authentication and Security Layer (SASL) for authentication, which can be enabled if needed.

**Network Security**:
- **Access Control**: Typically, Memcached should be deployed within a trusted network or use firewall rules to restrict access.

### Memcached Workflow

1. **Client Request**:
   - The client application uses a Memcached client library to send a request to the Memcached server.
   - The client library hashes the key and determines which server in the cluster should handle the request.

2. **Server Processing**:
   - The Memcached server receives the request and processes it. If it's a read request (e.g., `GET`), it looks up the key in memory and returns the value.
   - If it's a write request (e.g., `SET`), it stores the key-value pair in memory, managing the memory allocation using the slab allocator.

3. **Response**:
   - The server sends the response back to the client.

4. **Cache Management**:
   - If the server's memory is full, the LRU eviction policy is applied to remove the least recently used items.
   - Expired items are lazily evicted when accessed or when memory is required for new data.

### Example Use Case

1. **Web Page Caching**:
   - **Scenario**: A high-traffic website needs to serve dynamic content quickly.
   - **Implementation**: The application caches rendered HTML fragments in Memcached. On subsequent requests, the application retrieves these fragments from Memcached, reducing the need to re-render the page and query the database.
   - **Benefit**: This reduces the load on the database and application servers, leading to faster page load times and improved user experience.

### Conclusion

Memcached's architecture is designed for simplicity, high performance, and scalability. Its in-memory data storage, slab allocator for memory management, and client-side distribution for horizontal scalability make it an effective solution for caching in high-performance web applications. However, it lacks built-in persistence and replication, which need to be managed externally if required.

