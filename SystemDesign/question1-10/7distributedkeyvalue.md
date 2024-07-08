### Dynamo: Introduction(distributed key value store)

**Goal**
The goal of Dynamo is to design a distributed key-value store that is highly available, scalable, and decentralized. 

**What is Dynamo?**
Dynamo is a highly available key-value store developed by Amazon for internal use. It caters to services like shopping carts and bestseller lists, which require only primary-key access to data. It offers a flexible design that lets applications choose their desired levels of availability and consistency, avoiding the limitations of traditional relational databases in terms of scalability and availability.

**Background**
Dynamo, not to be confused with DynamoDB, is a distributed key-value storage system designed for high availability and partition tolerance at the expense of strong consistency. It falls within the AP category of the CAP theorem. The primary motivation was to ensure high availability, as it correlates directly with customer satisfaction. Dynamo's design has inspired many NoSQL databases, including Cassandra, Riak, and Voldemort.

**Design Goals**
1. **Highly Available**: Ensures the system is always on, even if imperfect.
2. **Scalable**: Adding machines should proportionally improve system performance.
3. **Decentralized**: Avoids single points of failure and performance bottlenecks.
4. **Eventually Consistent**: Data is replicated optimistically, with inconsistencies resolved later to maintain high availability.

**Dynamo's Use Cases**
Dynamo is ideal for applications where strong consistency is not critical. It supports strong consistency but at a performance cost. Amazon uses Dynamo for services requiring high reliability and flexible trade-offs between availability, consistency, cost-effectiveness, and performance. It provides a simple primary-key interface, making it suitable for services that would otherwise be inefficient with relational databases.

**System APIs**
Dynamo clients use `put()` and `get()` operations to write and read data:

- **get(key)**: Retrieves the object associated with the given key, potentially returning conflicting versions and metadata context.
- **put(key, context, object)**: Writes the object associated with the given key to storage nodes, using context to verify object validity.

Both objects and keys are treated as byte arrays, with keys hashed using MD5 to generate a 128-bit identifier for storage node allocation.

---
### Dynamo: High-Level Architecture

**Introduction**
Dynamo is a Distributed Hash Table (DHT) replicated across a cluster for high availability and fault tolerance.

**1. Data Distribution**
- **Consistent Hashing**: Distributes data among nodes, facilitating easy addition or removal of nodes.

**2. Data Replication and Consistency**
- **Eventual Consistency**: Data is replicated optimistically, ensuring high availability.

**3. Handling Temporary Failures**
- **Sloppy Quorum**: Replicates data to a subset of nodes to handle temporary failures, rather than requiring a strict majority.

**4. Inter-node Communication and Failure Detection**
- **Gossip Protocol**: Nodes communicate and maintain cluster state using the gossip protocol.

**5. High Availability**
- **Hinted Handoff**: Ensures the system remains writable by temporarily handing off data to other nodes when the primary node is unavailable.

**6. Conflict Resolution and Handling Permanent Failures**
- **Vector Clocks**: Tracks value history to reconcile conflicts during reads.
- **Merkle Trees**: Used as an anti-entropy mechanism to handle permanent failures and ensure data consistency in the background.

---

### Data Partitioning and Consistent Hashing in Distributed Systems

**Data Partitioning** is the method of distributing data across multiple nodes in a distributed system. It addresses two main challenges:

1. **Determining Data Location**: Identifying which node stores a specific piece of data.
2. **Handling Node Changes**: Efficiently managing data movement when nodes are added or removed to minimize disruption.

A naive approach involves using a hash function to map data keys to nodes using modulo operation. However, this approach remaps all keys when nodes change, causing significant data movement.

**Consistent Hashing**
**Consistent Hashing** solves this by mapping data to a ring structure where each node is assigned a range of data. This allows only a small set of keys to move when nodes are added or removed. In this system:

- Each node in the ring is assigned a token that defines its range.
- The hash of a data key determines its position in the ring and hence its storage node.

For example, with nodes having tokens 1, 26, 51, and 76, data is distributed accordingly. When nodes change, only the next node in the ring is affected.

**Virtual Nodes (Vnodes)**
**Virtual Nodes** further optimize data distribution. Instead of assigning a single range to each node, the range is divided into smaller subranges (Vnodes). Each physical node manages multiple Vnodes, which:

- **Balance Load**: Evenly distribute data and load across nodes, making the system more resilient to node changes.
- **Simplify Maintenance**: Facilitate easier handling of heterogeneous clusters with nodes of varying capacities.
- **Reduce Hotspots**: Minimize the chance of data hotspots by distributing smaller ranges.

Vnodes enhance the consistent hashing scheme by ensuring smoother rebalancing and reducing the impact on replica nodes during node rebuilds. This approach maintains efficient data management in dynamic and large-scale distributed systems.

--- 

### Dynamo Replication and Handling Failures: A Comprehensive Overview

**Optimistic Replication**
Dynamo uses a method called optimistic replication to ensure high availability and durability. Here's how it works:
- **Replication Factor**: Each data item is replicated on multiple nodes, where the number of replicas is defined by the replication factor.
- **Coordinator Node**: Each key is assigned to a coordinator node, which is the first node in the hash range.
- **Replication Process**: The coordinator node stores the data locally and then replicates it to its `N-1` clockwise successor nodes on the ring.
- **Asynchronous Replication**: This replication happens asynchronously in the background, supporting an eventually consistent model. This means that replicas are not guaranteed to be identical at all times.

**Consistent Hashing**
Dynamo employs consistent hashing to distribute data across nodes:
- **Data Ownership**: Each node is responsible for a specific range of data.
- **Replication**: Each data item is replicated on `N` nodes. If one node is down, other replicas can handle the queries.
- **Preference List**: This list contains the nodes responsible for storing a particular key, including extra nodes to account for failures and ensuring only distinct physical nodes are included.

**Sloppy Quorum**
Dynamo does not enforce strict quorum requirements to enhance availability:
- **Quorum Requirements**: Traditional quorum systems can become unavailable during failures. Dynamo uses a sloppy quorum instead.
- **Operation on Healthy Nodes**: Read/write operations are performed on the first `N` healthy nodes from the preference list, which might not be the first `N` nodes encountered on the hash ring.

**Example Scenario**
In a Dynamo setup with replication factor `N = 3`:
- If Server 1 is down during a write operation, the data will be stored on Server 4 instead.
- This transfer ensures that the system remains available even during temporary failures.

**Hinted Handoff**
Hinted handoff is a mechanism to handle node unavailability:
- **Temporary Storage**: When a node is unreachable, another node temporarily stores the writes.
- **Metadata Hint**: The replica contains metadata indicating the intended recipient.
- **Periodic Scans**: Nodes periodically scan their local database to check if the intended recipient has recovered.
- **Data Transfer**: Once the original node is back online, the temporarily stored data is transferred to it, and the holding node can delete the local copy.

**Conflict Resolution**
Due to the nature of sloppy quorum:
- **Divergence**: Data can diverge, with concurrent writes being accepted by non-overlapping sets of nodes.
- **Conflicts**: Multiple conflicting values for the same key can exist, leading to potential stale or conflicting reads.
- **Vector Clocks**: Dynamo uses vector clocks to resolve these conflicts, allowing the system to manage and reconcile divergent data effectively.

**Summary**
Dynamo's replication strategy ensures high availability and durability through:
- Optimistic replication with asynchronous updates.
- Consistent hashing to distribute data and handle node failures.
- Sloppy quorum to maintain operations during temporary failures.
- Hinted handoff to accept writes even when nodes are unreachable, ensuring eventual consistency.
- Conflict resolution using vector clocks to handle data divergence.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/replication.svg)

---

### Conflict Resolution in Dynamo

**Clock Skew**
**Clock skew** is the variance in the time kept by different clocks in a distributed system. Here’s how it can cause inconsistencies:

- **Single Machine**: Assumes a linear progression of time (t1 < t2), enabling straightforward versioning.
- **Distributed System**: Different machines have unsynchronized clocks, so time t on one machine doesn’t necessarily happen before time t+1 on another. This makes relying on wall clock timestamps unreliable for versioning.

**Vector Clocks**
Instead of wall clock timestamps, Dynamo uses **vector clocks** to track the causality between different versions of a data item. Here’s how they work:

1. **Structure**: A vector clock is a list of (node, counter) pairs.
2. **Versioning**: Each version of a data item is associated with a vector clock.
3. **Causality**: By comparing vector clocks, the system can determine if one version is an ancestor of another or if they are concurrent and conflicting.

**How Vector Clocks Handle Conflicts**
**Example Scenario**
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

**Conflict Resolution Process**
1. **Client-Side Reconciliation**:
   - The client receives conflicting versions and must merge them. For example, it might decide which value to keep based on application-specific logic.

2. **Semantic Reconciliation**:
   - The client merges different branches of data evolution. For example, merging shopping cart items ensures no items are lost.

3. **Truncation**:
   - Dynamo truncates vector clocks when they grow too large. This is a potential issue for maintaining eventual consistency if older vector clocks necessary for reconciliation are deleted.

**Conflict-Free Replicated Data Types (CRDTs)**
**CRDTs** are designed to resolve conflicts automatically, ensuring strong eventual consistency. Here’s how they work:

1. **Modeling Data**: Data is modeled such that concurrent changes can be applied in any order, yielding the same result.
2. **Example**: Amazon’s shopping cart:
   - Adding items A and B can be done in any order. Both additions result in a cart containing A and B.
   - Removing items is modeled as a negative add operation.

**Last-Write-Wins (LWW)**
Dynamo and systems like Apache Cassandra often use a simpler, though less reliable, conflict resolution strategy: **last-write-wins** (LWW):

1. **Wall Clock Timestamp**: Conflicts are resolved by choosing the version with the most recent timestamp.
2. **Drawbacks**:
   - LWW can lead to data loss if conflicting writes occur simultaneously.
   - It essentially discards one of the conflicting updates, akin to flipping a coin to decide which version to keep.

---

### The Life of Dynamo’s put() & get() Operations

Dynamo handles `get()` and `put()` requests through a well-defined process designed to ensure availability, durability, and consistency. Here’s a detailed look into how Dynamo manages these operations, including strategies for choosing the coordinator node, the consistency protocol, and the specifics of the `put()` and `get()` processes.

**Strategies for Choosing the Coordinator Node**
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

**Consistency Protocol**
Dynamo uses a quorum-like system for its consistency protocol, defined by parameters \(N\), \(R\), and \(W\):
- **\(N\)**: Number of replicas.
- **\(R\)**: Minimum number of nodes that must participate in a successful read.
- **\(W\)**: Minimum number of nodes that must participate in a successful write.

Common configurations include:
- \(N = 3\), \(R = 2\), \(W = 2\)
- \(N = 3\), \(R = 3\), \(W = 1\): Fast reads, slow writes, not very durable.
- \(N = 3\), \(R = 1\), \(W = 3\): Fast writes, slow reads, durable.

The latency of operations depends on the slowest replica involved. Lower values of \(R\) and \(W\) can improve latency but increase the risk of inconsistency and reduce durability.

**put() Process**
1. **Version and Vector Clock**: The coordinator generates a new data version and updates the vector clock.
2. **Local Storage**: The coordinator saves the new data locally.
3. **Replication**: The coordinator sends the write request to \(W\) highest-ranked healthy nodes from the preference list.
4. **Confirmation**: The `put()` operation is considered successful after receiving \(W\) confirmations.

**get() Process**
1. **Request Data**: The coordinator requests the data version from \(R\) highest-ranked healthy nodes from the preference list.
2. **Wait for Replies**: The coordinator waits until \(R\) replies are received.
3. **Causal Versions**: The coordinator uses vector clocks to handle causal data versions.
4. **Return Data**: All relevant data versions are returned to the caller.

**Request Handling through State Machine**
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

**Load Distribution**
To avoid uneven load distribution:
- Any of the top \(N\) nodes in the preference list can coordinate writes.
- The coordinator for a write operation is often the node that responded fastest to the preceding read operation, increasing the chances of achieving "read-your-writes" consistency.

**Summary**

Dynamo's approach to `put()` and `get()` operations ensures high availability and eventual consistency through:
- Multiple strategies for choosing the coordinator node.
- A quorum-like consistency protocol with configurable \(N\), \(R\), and \(W\) parameters.
- Detailed processes for handling writes (`put()`) and reads (`get()`).
- The use of state machines for efficient request handling and consistency maintenance.
- Optimizations for load distribution and maintaining data consistency.

Would you like for me to generate a downloadable Word document of these notes?


### Anti-Entropy Through Merkle Trees in Dynamo

**Merkle Trees:**
- **Structure**: Binary tree of hashes; each internal node is the hash of its children, each leaf node is the hash of a data portion.
- **Comparison**: 
  1. Compare root hashes of two trees.
  2. If equal, datasets are identical.
  3. If not, recursively compare child nodes.
- **Efficiency**: Minimizes data transfer by identifying and synchronizing only differing parts.
- **Independence**: Each branch can be checked independently, reducing disk reads.

**Advantages:**
- **Data Transfer Minimization**: Only differing data parts are exchanged.
- **Efficient Disk Reads**: Reduced number of disk reads during synchronization.

**Disadvantages:**
- **Recalculation**: Required when nodes join or leave, changing key ranges.
- **Resource Intensive**: Building and maintaining trees can be demanding.

**Anti-Entropy Process:**
1. **Build Merkle Trees**: Each node creates a Merkle tree for its data range.
2. **Periodic Comparison**: Nodes exchange and compare root hashes.
3. **Identify Differences**: Recursively compare differing branches.
4. **Synchronize Data**: Exchange only the differing data parts to achieve consistency.

Would you like for me to generate a downloadable Word document of these notes?

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/merkeltree.svg)



**Summary: Dynamo**
**Overview**:
- **Purpose**: Highly available key-value store developed by Amazon.
- **Design Philosophy**: Sacrifices strong consistency for high availability.

**Key Features**:
- **Peer-to-Peer System**: No leader or follower nodes; all nodes are equal.
- **Consistent Hashing**: Distributes data across nodes automatically.
- **Data Replication**: Uses sloppy quorum for fault tolerance and redundancy.
- **Conflict Resolution**: Employs Merkle trees for anti-entropy and vector clocks for reconciliation.
- **Inter-Node Communication**: Utilizes the gossip protocol.
- **Always Writeable**: Uses hinted handoff to handle writes during node failures.

**Techniques and Advantages**:

| Problem | Technique | Advantage |
|---------|-----------|-----------|
| Partitioning | Consistent Hashing | Incremental Scalability |
| High availability for writes | Vector clocks with reconciliation during reads | Decoupled version size from update rates |
| Handling temporary failures | Sloppy Quorum and Hinted Handoff | High availability and durability when replicas are unavailable |
| Recovering from permanent failures | Anti-entropy using Merkle trees | Synchronizes divergent replicas in the background |
| Membership and failure detection | Gossip protocol | Symmetry and avoidance of centralized monitoring |

**System Design Patterns**:
- **Consistent Hashing**: Distributes data efficiently.
- **Quorum**: Ensures data consistency with configurable write success criteria.
- **Gossip Protocol**: Maintains cluster state information.
- **Hinted Handoff**: Handles writes for failing nodes.
- **Read Repair**: Updates nodes with the latest data version.
- **Vector Clocks**: Reconciles concurrent updates.
- **Merkle Trees**: Resolves conflicts and ensures data consistency in the background.

Would you like for me to generate a downloadable Word document of these notes?


### Gossip Protocol in Dynamo with Example

**Gossip Protocol:**
- **Purpose**: Keeps track of the state of all nodes in a Dynamo cluster.
- **Mechanism**: Nodes periodically exchange state information with random peers.
- **Frequency**: Each node initiates a gossip round every second.
- **Benefits**: Ensures all nodes quickly learn about each other's state, minimizing data transfer and enhancing cluster coherence.

**Example:**
1. **Initial State**: Nodes A, B, and C are part of the Dynamo cluster.
2. **Gossip Round**: 
   - Node A gossips with Node B, sharing its state and the states it knows.
   - Node C gossips with Node A, learning the states of Nodes A and B.
3. **Propagation**: 
   - Node B then gossips with Node C, completing the state exchange cycle.
   - Within a few rounds, all nodes are aware of each other's states.

**Seed Nodes:**
- **Function**: Prevent logical partitions by serving as known points of contact for new nodes.
- **Configuration**: Obtained from static configuration or configuration service.
- **Role**: Ensure all nodes are aware of each other and reconcile membership changes.

**Example:**
1. **New Node Joins**:
   - Administrator adds Node D to the cluster.
   - Node D contacts a seed node (e.g., Node A) to join the ring.
2. **Seed Node Interaction**:
   - Node A gossips with Node D, updating Node D with the current cluster state.
   - Node D becomes aware of Nodes A, B, and C, avoiding logical partition.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/something.svg)
