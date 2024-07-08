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

#### System APIs
Dynamo clients use `put()` and `get()` operations to write and read data:

- **get(key)**: Retrieves the object associated with the given key, potentially returning conflicting versions and metadata context.
- **put(key, context, object)**: Writes the object associated with the given key to storage nodes, using context to verify object validity.

Both objects and keys are treated as byte arrays, with keys hashed using MD5 to generate a 128-bit identifier for storage node allocation.

---
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

---

### Data Partitioning and Consistent Hashing in Distributed Systems

**Data Partitioning** is the method of distributing data across multiple nodes in a distributed system. It addresses two main challenges:

1. **Determining Data Location**: Identifying which node stores a specific piece of data.
2. **Handling Node Changes**: Efficiently managing data movement when nodes are added or removed to minimize disruption.

A naive approach involves using a hash function to map data keys to nodes using modulo operation. However, this approach remaps all keys when nodes change, causing significant data movement.

### Consistent Hashing

**Consistent Hashing** solves this by mapping data to a ring structure where each node is assigned a range of data. This allows only a small set of keys to move when nodes are added or removed. In this system:

- Each node in the ring is assigned a token that defines its range.
- The hash of a data key determines its position in the ring and hence its storage node.

For example, with nodes having tokens 1, 26, 51, and 76, data is distributed accordingly. When nodes change, only the next node in the ring is affected.

### Virtual Nodes (Vnodes)

**Virtual Nodes** further optimize data distribution. Instead of assigning a single range to each node, the range is divided into smaller subranges (Vnodes). Each physical node manages multiple Vnodes, which:

- **Balance Load**: Evenly distribute data and load across nodes, making the system more resilient to node changes.
- **Simplify Maintenance**: Facilitate easier handling of heterogeneous clusters with nodes of varying capacities.
- **Reduce Hotspots**: Minimize the chance of data hotspots by distributing smaller ranges.

Vnodes enhance the consistent hashing scheme by ensuring smoother rebalancing and reducing the impact on replica nodes during node rebuilds. This approach maintains efficient data management in dynamic and large-scale distributed systems.

--- 

### Dynamo Replication and Handling Failures: A Comprehensive Overview

#### Optimistic Replication
Dynamo uses a method called optimistic replication to ensure high availability and durability. Here's how it works:
- **Replication Factor**: Each data item is replicated on multiple nodes, where the number of replicas is defined by the replication factor.
- **Coordinator Node**: Each key is assigned to a coordinator node, which is the first node in the hash range.
- **Replication Process**: The coordinator node stores the data locally and then replicates it to its `N-1` clockwise successor nodes on the ring.
- **Asynchronous Replication**: This replication happens asynchronously in the background, supporting an eventually consistent model. This means that replicas are not guaranteed to be identical at all times.

#### Consistent Hashing
Dynamo employs consistent hashing to distribute data across nodes:
- **Data Ownership**: Each node is responsible for a specific range of data.
- **Replication**: Each data item is replicated on `N` nodes. If one node is down, other replicas can handle the queries.
- **Preference List**: This list contains the nodes responsible for storing a particular key, including extra nodes to account for failures and ensuring only distinct physical nodes are included.

#### Sloppy Quorum
Dynamo does not enforce strict quorum requirements to enhance availability:
- **Quorum Requirements**: Traditional quorum systems can become unavailable during failures. Dynamo uses a sloppy quorum instead.
- **Operation on Healthy Nodes**: Read/write operations are performed on the first `N` healthy nodes from the preference list, which might not be the first `N` nodes encountered on the hash ring.

#### Example Scenario
In a Dynamo setup with replication factor `N = 3`:
- If Server 1 is down during a write operation, the data will be stored on Server 4 instead.
- This transfer ensures that the system remains available even during temporary failures.

#### Hinted Handoff
Hinted handoff is a mechanism to handle node unavailability:
- **Temporary Storage**: When a node is unreachable, another node temporarily stores the writes.
- **Metadata Hint**: The replica contains metadata indicating the intended recipient.
- **Periodic Scans**: Nodes periodically scan their local database to check if the intended recipient has recovered.
- **Data Transfer**: Once the original node is back online, the temporarily stored data is transferred to it, and the holding node can delete the local copy.

#### Conflict Resolution
Due to the nature of sloppy quorum:
- **Divergence**: Data can diverge, with concurrent writes being accepted by non-overlapping sets of nodes.
- **Conflicts**: Multiple conflicting values for the same key can exist, leading to potential stale or conflicting reads.
- **Vector Clocks**: Dynamo uses vector clocks to resolve these conflicts, allowing the system to manage and reconcile divergent data effectively.

### Summary
Dynamo's replication strategy ensures high availability and durability through:
- Optimistic replication with asynchronous updates.
- Consistent hashing to distribute data and handle node failures.
- Sloppy quorum to maintain operations during temporary failures.
- Hinted handoff to accept writes even when nodes are unreachable, ensuring eventual consistency.
- Conflict resolution using vector clocks to handle data divergence.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/replication.svg)

---

### Conflict Resolution in Dynamo: A Detailed Exploration

#### Clock Skew

**Clock skew** is the variance in the time kept by different clocks in a distributed system. Here’s how it can cause inconsistencies:

- **Single Machine**: Assumes a linear progression of time (t1 < t2), enabling straightforward versioning.
- **Distributed System**: Different machines have unsynchronized clocks, so time t on one machine doesn’t necessarily happen before time t+1 on another. This makes relying on wall clock timestamps unreliable for versioning.

### Vector Clocks

Instead of wall clock timestamps, Dynamo uses **vector clocks** to track the causality between different versions of a data item. Here’s how they work:

1. **Structure**: A vector clock is a list of (node, counter) pairs.
2. **Versioning**: Each version of a data item is associated with a vector clock.
3. **Causality**: By comparing vector clocks, the system can determine if one version is an ancestor of another or if they are concurrent and conflicting.

### How Vector Clocks Handle Conflicts

#### Example Scenario

1. **Initial Write**:
   - **Server A** writes key `k1` with value `foo`, version `[A:1]`. This is replicated to **Server B**.

2. **Subsequent Write**:
   - **Server A** writes key `k1` with value `bar`, version `[A:2]`. This is also replicated to **Server B**.

3. **Network Partition**:
   - **Server A** and **Server B** cannot communicate.

4. **Concurrent Writes**:
   - **Server A** writes key `k1` with value `baz`, version `[A:3]`.
   - **Server B** writes key `k1` with value `bax`, version `[B:1]`.

5. **Network Heals**:
   - **Server A** and **Server B** synchronize. They detect two versions: `[A:3]` and `[A:2, B:1]`.

6. **Conflict Detection**:
   - **Server A** and **Server B** recognize the versions are conflicting and return both to the client for reconciliation.
   
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/keystore.svg)

#### Conflict Resolution Process

1. **Client-Side Reconciliation**:
   - The client receives conflicting versions and must merge them. For example, it might decide which value to keep based on application-specific logic.

2. **Semantic Reconciliation**:
   - The client merges different branches of data evolution. For example, merging shopping cart items ensures no items are lost.

3. **Truncation**:
   - Dynamo truncates vector clocks when they grow too large. This is a potential issue for maintaining eventual consistency if older vector clocks necessary for reconciliation are deleted.

### Conflict-Free Replicated Data Types (CRDTs)

**CRDTs** are designed to resolve conflicts automatically, ensuring strong eventual consistency. Here’s how they work:

1. **Modeling Data**: Data is modeled such that concurrent changes can be applied in any order, yielding the same result.
2. **Example**: Amazon’s shopping cart:
   - Adding items A and B can be done in any order. Both additions result in a cart containing A and B.
   - Removing items is modeled as a negative add operation.

### Last-Write-Wins (LWW)

Dynamo and systems like Apache Cassandra often use a simpler, though less reliable, conflict resolution strategy: **last-write-wins** (LWW):

1. **Wall Clock Timestamp**: Conflicts are resolved by choosing the version with the most recent timestamp.
2. **Drawbacks**:
   - LWW can lead to data loss if conflicting writes occur simultaneously.
   - It essentially discards one of the conflicting updates, akin to flipping a coin to decide which version to keep.

---

### The Life of Dynamo’s put() & get() Operations

Dynamo handles `get()` and `put()` requests through a well-defined process designed to ensure availability, durability, and consistency. Here’s a detailed look into how Dynamo manages these operations, including strategies for choosing the coordinator node, the consistency protocol, and the specifics of the `put()` and `get()` processes.

#### Strategies for Choosing the Coordinator Node

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/keystore2.svg)

Dynamo clients can use two strategies to choose a node for their requests:

1. **Load Balancer Strategy**:
   - **Description**: Clients route their requests through a generic load balancer.
   - **Advantages**: Scalability and loose coupling.
   - **Disadvantages**: The load balancer might forward the request to a node not in the preference list, causing an extra hop.

2. **Partition-Aware Client Library**:
   - **Description**: Clients use a library that routes requests directly to the appropriate coordinator node.
   - **Advantages**: Lower latency by directly contacting the node holding the required data.
   - **Disadvantages**: Less control over load distribution and request handling by Dynamo.

#### Consistency Protocol

Dynamo uses a quorum-like system for its consistency protocol, defined by parameters \(N\), \(R\), and \(W\):
- **\(N\)**: Number of replicas.
- **\(R\)**: Minimum number of nodes that must participate in a successful read.
- **\(W\)**: Minimum number of nodes that must participate in a successful write.

Common configurations include:
- \(N = 3\), \(R = 2\), \(W = 2\)
- \(N = 3\), \(R = 3\), \(W = 1\): Fast reads, slow writes, not very durable.
- \(N = 3\), \(R = 1\), \(W = 3\): Fast writes, slow reads, durable.

The latency of operations depends on the slowest replica involved. Lower values of \(R\) and \(W\) can improve latency but increase the risk of inconsistency and reduce durability.

#### put() Process

1. **Version and Vector Clock**: The coordinator generates a new data version and updates the vector clock.
2. **Local Storage**: The coordinator saves the new data locally.
3. **Replication**: The coordinator sends the write request to \(W\) highest-ranked healthy nodes from the preference list.
4. **Confirmation**: The `put()` operation is considered successful after receiving \(W\) confirmations.

#### get() Process

1. **Request Data**: The coordinator requests the data version from \(R\) highest-ranked healthy nodes from the preference list.
2. **Wait for Replies**: The coordinator waits until \(R\) replies are received.
3. **Causal Versions**: The coordinator uses vector clocks to handle causal data versions.
4. **Return Data**: All relevant data versions are returned to the caller.

#### Request Handling through State Machine

Each client request results in creating a state machine on the node that received the client request. The state machine handles:
- Identifying responsible nodes for a key.
- Sending requests and waiting for responses.
- Potential retries.
- Processing replies and packaging the response.

For read operations:
1. **Send Read Requests**: To the nodes.
2. **Wait for Responses**: Until the minimum required responses are received.
3. **Fail Request**: If too few replies are received within the time limit.
4. **Gather Data Versions**: And determine which ones to return.
5. **Syntactic Reconciliation**: Generate an opaque write context if versioning is enabled.
6. **Read Repair**: Update nodes with the latest version if stale versions were returned.

#### Load Distribution

To avoid uneven load distribution:
- Any of the top \(N\) nodes in the preference list can coordinate writes.
- The coordinator for a write operation is often the node that responded fastest to the preceding read operation, increasing the chances of achieving "read-your-writes" consistency.

### Summary

Dynamo's approach to `put()` and `get()` operations ensures high availability and eventual consistency through:
- Multiple strategies for choosing the coordinator node.
- A quorum-like consistency protocol with configurable \(N\), \(R\), and \(W\) parameters.
- Detailed processes for handling writes (`put()`) and reads (`get()`).
- The use of state machines for efficient request handling and consistency maintenance.
- Optimizations for load distribution and maintaining data consistency.

Would you like for me to generate a downloadable Word document of these notes?

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
