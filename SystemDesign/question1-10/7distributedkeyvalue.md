### Dynamo: Introduction(distributed key value store)

#### Goal
The goal of Dynamo is to design a distributed key-value store that is highly available, scalable, and decentralized. 

#### What is Dynamo?
Dynamo is a highly available key-value store developed by Amazon for internal use. It caters to services like shopping carts and bestseller lists, which require only primary-key access to data. It offers a flexible design that lets applications choose their desired levels of availability and consistency, avoiding the limitations of traditional relational databases in terms of scalability and availability.

#### Background
Dynamo, not to be confused with DynamoDB, is a distributed key-value storage system designed for high availability and partition tolerance at the expense of strong consistency. It falls within the AP category of the CAP theorem. The primary motivation was to ensure high availability, as it correlates directly with customer satisfaction. Dynamo's design has inspired many NoSQL databases, including Cassandra, Riak, and Voldemort.

#### Design Goals
1. **Highly Available**: Ensures the system is always on, even if imperfect.
2. **Scalable**: Adding machines should proportionally improve system performance.
3. **Decentralized**: Avoids single points of failure and performance bottlenecks.
4. **Eventually Consistent**: Data is replicated optimistically, with inconsistencies resolved later to maintain high availability.

#### Dynamo's Use Cases
Dynamo is ideal for applications where strong consistency is not critical. It supports strong consistency but at a performance cost. Amazon uses Dynamo for services requiring high reliability and flexible trade-offs between availability, consistency, cost-effectiveness, and performance. It provides a simple primary-key interface, making it suitable for services that would otherwise be inefficient with relational databases.


### Dynamo: High-Level Architecture

#### Introduction
Dynamo is a Distributed Hash Table (DHT) replicated across a cluster for high availability and fault tolerance.

#### 1. Data Distribution
- **Consistent Hashing**: Distributes data among nodes, facilitating easy addition or removal of nodes.

#### 2. Data Replication and Consistency
- **Eventual Consistency**: Data is replicated optimistically, ensuring high availability.

#### 3. Handling Temporary Failures
- **Sloppy Quorum**: Replicates data to a subset of nodes to handle temporary failures, rather than requiring a strict majority.

#### 4. Inter-node Communication and Failure Detection
- **Gossip Protocol**: Nodes communicate and maintain cluster state using the gossip protocol.

#### 5. High Availability
- **Hinted Handoff**: Ensures the system remains writable by temporarily handing off data to other nodes when the primary node is unavailable.

#### 6. Conflict Resolution and Handling Permanent Failures
- **Vector Clocks**: Tracks value history to reconcile conflicts during reads.
- **Merkle Trees**: Used as an anti-entropy mechanism to handle permanent failures and ensure data consistency in the background.
#### System APIs
Dynamo clients use `put()` and `get()` operations to write and read data:

- **get(key)**: Retrieves the object associated with the given key, potentially returning conflicting versions and metadata context.
- **put(key, context, object)**: Writes the object associated with the given key to storage nodes, using context to verify object validity.

Both objects and keys are treated as byte arrays, with keys hashed using MD5 to generate a 128-bit identifier for storage node allocation.

## Design Scope for Key-Value Store

### Problem Understanding
Designing a key-value store involves balancing tradeoffs between read, write, and memory usage. Additionally, there is a need to balance consistency and availability.

### Desired Characteristics
1. **Small Key-Value Pair Size**: Each key-value pair should be less than 10 KB.
2. **Big Data Storage**: Ability to store and manage large volumes of data.
3. **High Availability**: The system should respond quickly, even in the face of failures.
4. **High Scalability**: The system should support large datasets and be easily scalable.
5. **Automatic Scaling**: The system should automatically add or remove servers based on traffic.
6. **Tunable Consistency**: The system should allow adjustments to the consistency level.
7. **Low Latency**: The system should have minimal delay in data access and operations.


## System Components

- Data Partition
- Data Replication
- Consistency
- Inconsistency Resolution
- Handling Failures
- System Architecture Diagram
- Write Path
- Read Path


## Data Partition

### Purpose
- Split large data sets into smaller partitions to store across multiple servers.

### Challenges
1. Distributing data evenly across multiple servers.
2. Minimizing data movement when nodes are added or removed.

### Technique: Consistent Hashing
- **Servers on Hash Ring**: Servers are placed on a hash ring (e.g., s0, s1, ..., s7).
- **Key Placement**: Keys are hashed onto the same ring and stored on the first server encountered moving clockwise.

### Advantages
1. **Automatic Scaling**: Servers can be added or removed automatically based on load.
2. **Heterogeneity**: Servers with higher capacity are assigned more virtual nodes.


## Data Replication

### Purpose
- Achieve high availability and reliability by replicating data across multiple servers.

### Replication Logic
- **Replication Factor (N)**: Data is replicated asynchronously over N servers.
- **Server Selection**: After a key is mapped to the hash ring, choose the first N unique servers clockwise to store data copies.

### Handling Virtual Nodes
- Ensure the first N nodes are unique physical servers to avoid fewer than N physical servers owning the replicas.

### Reliability Enhancement
- Place replicas in distinct data centers to avoid simultaneous failures due to power outages, network issues, or natural disasters.
- Data centers are connected through high-speed networks for efficient replication.


## Consistency

### Purpose
- Ensure data synchronization across multiple replicas.

### Quorum Consensus
- **Definitions**:
  - **N**: Number of replicas.
  - **W**: Write quorum size. A write is successful if acknowledged by W replicas.
  - **R**: Read quorum size. A read is successful if responses are received from at least R replicas.

### Example (N = 3)
- **W = 1**: Write acknowledged by at least one replica.
- **R = 1**: Read acknowledged by at least one replica.
- **Coordinator**: Acts as a proxy between client and nodes.

### Tradeoff
- **Latency vs. Consistency**: 
  - **W = 1 or R = 1**: Fast operations, lower consistency.
  - **W or R > 1**: Better consistency, slower operations.
  - **W + R > N**: Strong consistency (e.g., N = 3, W = 2, R = 2).
  - **W + R <= N**: Strong consistency not guaranteed.

### Configuration Examples
- **R = 1, W = N**: Optimized for fast read.
- **W = 1, R = N**: Optimized for fast write.
- **W + R > N**: Ensures strong consistency (e.g., N = 3, W = 2, R = 2).



## Consistency

### Inconsistency and Resolution
- **Problem**: Replication causes inconsistencies among replicas.
- **Solution**: Versioning and vector clocks are used to solve inconsistency problems.

### Inconsistency Example
- Original value at nodes n1 and n2 is the same.
- Server 1 changes name to "johnSanFrancisco".
- Server 2 changes name to "johnNewYork".
- Result: Conflicting values (versions v1 and v2).

### Versioning
- Treat each data modification as a new immutable version.
- Detect and reconcile conflicts with a versioning system.

### Vector Clocks
- **Definition**: A [server, version] pair associated with a data item.
- **Usage**: Determine if one version precedes, succeeds, or conflicts with another.

### Vector Clock Example
1. Client writes data item D1 to server Sx: D1([Sx, 1]).
2. Client updates D1 to D2: D2([Sx, 2]).
3. Client updates D2 to D3, handled by server Sy: D3([Sx, 2], [Sy, 1]).
4. Client updates D2 to D4, handled by server Sz: D4([Sx, 2], [Sz, 1]).
5. Client reads D3 and D4, discovers conflict, resolves it, and writes D5([Sx, 3], [Sy, 1], [Sz, 1]).

### Conflict Detection
- **Ancestor**: No conflict if Y's vector clock counters ≥ X's vector clock counters.
  - Example: D([s0, 1], [s1, 1]) is an ancestor of D([s0, 1], [s1, 2]).
- **Sibling**: Conflict if any participant in Y's vector clock has a counter < its corresponding counter in X.
  - Example: D([s0, 1], [s1, 2]) and D([s0, 2], [s1, 1]) indicate conflict.

### Downsides of Vector Clocks

1. **Client Complexity**:
   - Clients must implement logic to resolve conflicts detected through vector clocks.

2. **Growth of Vector Clocks**:
   - Vector clocks can grow large if many servers participate.
   - **Mitigation**: Set a threshold for the length of the vector clock. When the length exceeds this limit, remove the oldest pairs.
   - **Tradeoff**: This can lead to inefficiencies in reconciliation because the complete history may not be preserved. However, in practice (e.g., Amazon Dynamo), this hasn't been a significant issue.

### Practical Use:
- Systems like Amazon Dynamo, Cassandra, and BigTable use these techniques to manage consistency and resolve conflicts effectively, balancing performance and reliability.


## Handling Failures in Key-Value Stores

### Failure Detection

#### Techniques
Detecting failures in a distributed system requires careful consideration to avoid false positives. Here are the primary techniques used:

1. **Multiple Sources of Information**:
   - Simply marking a server as down because another server says so is insufficient.
   - At least two independent sources must confirm a server's failure to avoid incorrect assumptions.

2. **All-to-All Multicasting**:
   - Every server sends status updates to every other server.
   - **Advantage**: Simple and straightforward.
   - **Disadvantage**: Inefficient in systems with many servers due to high network traffic.

3. **Gossip Protocol**:
   - A decentralized method for failure detection that scales well.
   - **How It Works**:
     - **Membership List**: Each node maintains a list containing member IDs and their heartbeat counters.
     - **Heartbeat Increment**: Each node periodically increments its heartbeat counter.
     - **Random Heartbeat Propagation**: Nodes send heartbeats to random nodes, which further propagate them.
     - **Offline Detection**: If a node’s heartbeat counter has not increased for a certain period, it is marked as offline.
   - **Example** (Figure 11):
     - Node s0 maintains a list and notices that node s2’s heartbeat counter has not increased.
     - s0 sends s2’s info to random nodes.
     - Once other nodes confirm s2's heartbeat hasn’t updated, s2 is marked down and this info is propagated.

### Temporary Failures

When a node or network component temporarily fails, the system needs to ensure availability despite the failure.

#### Sloppy Quorum
- **Objective**: Improve availability by relaxing strict quorum requirements.
- **Operation**:
  - **Writes**: Choose the first W healthy servers for writes.
  - **Reads**: Choose the first R healthy servers for reads.
  - **Offline Servers**: Ignored during these operations.
- **Example** (Figure 12):
  - If node s2 is down, reads and writes are temporarily handled by node s3.
  - When s2 is back online, node s3 transfers the changes back to s2 (hinted handoff).

#### Hinted Handoff
- **Process**:
  - When a server is unavailable, another server temporarily processes its requests.
  - Changes are tracked and handed back to the original server once it is back online.
- **Benefit**: Ensures that data consistency is eventually restored while maintaining high availability.

### Permanent Failures

Permanent node failures require more robust mechanisms to ensure data consistency and integrity.

#### Anti-Entropy Protocol
- **Purpose**: Synchronize replicas to maintain consistency.
- **Technique**: Merkle trees are used to detect and resolve inconsistencies efficiently.

##### Merkle Trees
- **Definition**: A hash tree where each non-leaf node is labeled with the hash of its child nodes’ values.
- **Advantages**:
  - Efficient and secure verification of large data structures.
  - Only the differing parts of the data need to be synchronized.
- **Steps to Build**:
  1. **Divide Key Space into Buckets**:
     - Key space is divided to manage depth (Figure 13).
  2. **Hash Each Key in a Bucket**:
     - Uniform hashing method used for each key (Figure 14).
  3. **Create a Single Hash Node per Bucket**:
     - Aggregate the hashes within each bucket (Figure 15).
  4. **Build Tree Upwards**:
     - Calculate hashes of child nodes recursively until the root is formed (Figure 16).

##### Inconsistency Detection
- **Comparison**:
  - Start with root hashes. If they match, data is identical.
  - If root hashes differ, traverse down the tree, comparing child hashes to pinpoint differences.
- **Efficiency**: The amount of data synchronized is proportional to the differences, not the total data size.

### Data Center Outage

#### Strategy
Data center outages, due to various reasons like power outages, network failures, or natural disasters, can severely impact availability. Here's how to handle them:

1. **Replication Across Multiple Data Centers**:
   - Ensure data is replicated in different geographic locations.
   - Users can still access data even if one data center is completely offline.
2. **High-Speed Networks**:
   - Data centers are connected via high-speed networks to facilitate quick and efficient data replication and access.

By employing these strategies, key-value stores can achieve high availability, reliability, and consistency, ensuring robust performance even in the face of failures.


## System Architecture Diagram

### Main Features

1. **Client Communication**:
   - Clients interact with the key-value store using simple APIs: `get(key)` and `put(key, value)`.

2. **Coordinator Node**:
   - Acts as a proxy between the client and the key-value store.
   - Manages client requests and directs them to the appropriate nodes.

3. **Node Distribution**:
   - Nodes are distributed on a ring using consistent hashing.
   - Ensures even distribution of data and efficient scaling.

4. **Decentralization**:
   - The system is fully decentralized.
   - Nodes can be added or moved automatically without manual intervention.

5. **Data Replication**:
   - Data is replicated across multiple nodes to ensure high availability and fault tolerance.

6. **No Single Point of Failure**:
   - Every node has the same set of responsibilities, eliminating single points of failure.
   - Ensures robustness and reliability of the system.


![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/keyvalue1.svg)
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/keyvalue2.svg)



## Write Path

### Steps (Figure 19)

1. **Commit Log**:
   - The write request is first persisted in a commit log file to ensure durability.
   
2. **Memory Cache**:
   - Data is then saved in the memory cache (often referred to as a memtable).

3. **Flushing to SSTable**:
   - When the memory cache is full or reaches a predefined threshold, the data is flushed to an SSTable on disk.
   - SSTable is a sorted list of <key, value> pairs.

### Summary
- Write operations involve persisting data to a commit log, caching in memory, and eventually flushing to disk as SSTables to maintain data durability and efficiency.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/keyvalue3.svg)


## Read Path

### Steps (Figures 20 and 21)

1. **Check Memory Cache**:
   - After a read request is directed to a specific node, the system first checks if the data is in the memory cache.
   - If the data is found in the memory cache, it is returned to the client.

2. **Check Disk**:
   - If the data is not in memory, the system proceeds to check the disk.

3. **Bloom Filter**:
   - The bloom filter helps determine which SSTables might contain the key.
   
4. **Retrieve from SSTables**:
   - SSTables identified by the bloom filter are queried to retrieve the data.
   
5. **Return Data to Client**:
   - The retrieved data is then returned to the client.

### Summary
- Read operations first check the memory cache for data. If not found, the system uses bloom filters to efficiently locate the data in SSTables on disk, ensuring quick and accurate data retrieval.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/keyvalue4.svg)
