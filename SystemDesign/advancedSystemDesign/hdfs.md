### Hadoop Distributed File System: Introduction

#### Goal
- **Objective**: Design a distributed system capable of storing huge files (terabyte and larger).
- **Requirements**: Scalability, reliability, high availability.

#### What is Hadoop Distributed File System (HDFS)?
- **Definition**: HDFS is a distributed file system built to store unstructured data.
- **Purpose**: Designed for reliable storage of huge files and to stream those files at high bandwidth to user applications.
- **Inspiration**: HDFS is a variant and simplified version of the Google File System (GFS).
- **Design Principle**: Optimized for a write-once, read-many-times pattern.

#### Background
- **Apache Hadoop**: 
  - Software framework providing distributed file storage and computing.
  - Analyzes and transforms very large data sets using the MapReduce programming model.
- **HDFS**: 
  - Default file storage system in Hadoop.
  - Designed to be distributed, scalable, and fault-tolerant.
  - Primarily caters to MapReduce paradigm needs.
- **Comparison with GFS**:
  - Both store very large files and scale to petabytes of storage.
  - Both handle batch processing on huge data sets.
  - Designed for data-intensive applications, not for end-users.
  - Not POSIX-compliant and not mountable as traditional file systems.
  - Accessed via HDFS clients or API calls from Hadoop libraries.

#### Applications Not Suitable for HDFS
1. **Low-latency data access**:
   - HDFS is optimized for high throughput, often at the expense of latency.
   - Not ideal for applications requiring low-latency data access.

2. **Handling lots of small files**:
   - NameNode, the central server in HDFS, holds all filesystem metadata in memory.
   - Memory limits on the NameNode restrict the number of files.
   - Storing millions of files is feasible; billions are beyond current hardware capabilities.

3. **Concurrent writers and arbitrary file modifications**:
   - Unlike GFS, HDFS does not support multiple concurrent writers for a single file.
   - Writes are made at the end of the file in an append-only fashion.
   - No support for modifications at arbitrary offsets in a file.

#### APIs
- **User-level APIs**: HDFS does not provide standard POSIX-like APIs.
- **File Organization**:
  - Files are organized hierarchically in directories, identified by pathnames.
  - Supports usual file system operations: creation, deletion, renaming, moving, and symbolic links.
- **Operation Mode**: All read and write operations are performed in an append-only fashion.



### High-Level Architecture of HDFS

#### Overview of HDFS Architecture
- **File Storage**:
  - Files are divided into fixed-size blocks (default size: 128 MB, configurable per file).
  - Each file consists of:
    - **File Data**: Actual data of the file.
    - **Metadata**: Information about the file blocks, their locations, total file size, etc.

- **Cluster Components**:
  - **NameNode**: Manages filesystem metadata.
  - **DataNodes**: Store the actual data.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/hdfs1.svg)

#### Key Architectural Features

- **File Blocks**:
  - All blocks of a file are the same size except the last one.
  - Large block sizes optimize the storage of extremely large files and facilitate efficient MapReduce job processing.
  - Each block has a unique 64-bit ID (BlockID).
  - Read/write operations occur at the block level.

- **DataNode Operations**:
  - DataNodes store each block in a separate file on the local file system.
  - Provide read/write access to these blocks.
  - On startup, DataNodes scan their local file system and send a BlockReport to the NameNode.

- **NameNode Operations**:
  - Maintains two on-disk data structures:
    - **FsImage**: A checkpoint of the filesystem metadata.
    - **EditLog**: Logs all filesystem metadata transactions since the last checkpoint.
  - These structures aid in recovering from failures.

- **Client Interactions**:
  - User applications interact with HDFS through its client.
  - The HDFS Client interacts with the NameNode for metadata.
  - Data transfers occur directly between the client and DataNodes.

- **High Availability**:
  - HDFS creates multiple copies of the data and distributes them across nodes in the cluster.
  - Ensures data redundancy and reliability.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/hdfs2.svg)

#### Comparison Between GFS and HDFS

| Feature | GFS | HDFS |
|---------|-----|------|
| **Storage Node** | ChunkServer | DataNode |
| **File Part** | Chunk | Block |
| **File Part Size** | Default: 64MB (adjustable) | Default: 128MB (adjustable) |
| **Metadata Checkpoint** | Checkpoint image | FsImage |
| **Write Ahead Log** | Operation log | EditLog |
| **Platform** | Linux | Cross-Platform |
| **Language** | Developed in C++ | Developed in Java |
| **Availability** | Internal to Google | Open source |
| **Monitoring** | Master receives HeartBeat from ChunkServers | NameNode receives HeartBeat from DataNodes |
| **Concurrency** | Multiple writers/readers model | Write-once, read-many model |
| **File Operations** | Append and random writes possible | Only append is possible |
| **Garbage Collection** | Deleted files are renamed for later garbage collection | Deleted files are renamed for later garbage collection |
| **Communication** | RPC over TCP | RPC over TCP |
| **Cache Management** | Clients cache metadata; ChunkServers use buffer cache | Distributed cache, user-specified paths cached in DataNode's memory |
| **Replication Strategy** | Chunk replicas spread across racks, 3 copies by default, auto re-replication | Automatic rack-aware replication, 3 copies by default, user can specify replication factor |
| **File System Namespace** | Hierarchical, identified by pathnames | Hierarchical, supports traditional file organization and third-party file systems (e.g., S3, Cloud Store) |
| **Database** | Bigtable uses GFS | HBase uses HDFS |

#### Summary of HDFS High-Level Architecture
- **Block-based Storage**: Files broken into 128 MB blocks.
- **Cluster Roles**: NameNode (metadata) and DataNodes (actual data).
- **File Operations**: Read/write operations at the block level, append-only writes.
- **High Availability**: Multiple data copies across the cluster.
- **Client Interactions**: Metadata from NameNode, data transfers directly with DataNodes.
- **Comparative Features**: Similarities with GFS, but with distinct differences in terms of implementation, replication strategies, and file operations.


### Deep Dive into HDFS Design Components

#### Cluster Topology
- **Data Center Configuration**:
  - Many racks of servers connected via switches.
  - Typical configuration: 30 to 40 servers per rack.
  - Each rack has a dedicated gigabit switch connecting its servers.
  - Uplink from each rack switch to a core switch/router, bandwidth shared by many racks.

- **Server Configuration in HDFS**:
  - Each server is mapped to a particular rack.
  - Network distance between servers is measured in hops:
    - One hop corresponds to one link in the topology.
    - Tree-style topology is assumed by Hadoop.
    - Distance calculation: Sum of distances to the closest common ancestor.
  
- **Example**:
  - Distance between Node 1 and itself: 0 hops.
  - Distance between Node 1 and Node 2: 2 hops.
  - Distance between Node 3 and Node 4: 4 hops.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/hdfs3.svg)

#### Rack-Aware Replication
- **Replica Placement**:
  - Critical for HDFS reliability and performance.
  - **Replication Factor**: Typically set to 3.
  - Placement Strategy:
    - First replica on the same node as the client writing the block (if the client is within the cluster).
    - Second replica on a different rack from the first (off-rack replica).
    - Third replica on another random node on the same rack as the second.
    - Additional replicas on random nodes, avoiding excessive replicas on the same rack.
  
- **Failure Tolerance**:
  - Designed to tolerate node and rack failures.
  - Example: If an entire rack goes offline, the block can still be located on a different rack.

- **Policy Summary**:
  - No DataNode contains more than one replica of any block.
  - No rack contains more than two replicas of the same block, given enough racks.
  - Intentional tradeoff: Slows write operations but enhances reliability and performance.

#### Synchronization Semantics
- **Early Versions**:
  - Followed strict immutable semantics: Files once written could not be re-opened for writes.
  - Files could be deleted but not modified.
  
- **Current Versions**:
  - Support append operations.
  - Existing binary data once written cannot be modified in place.
  - Design aligns with common MapReduce workloads (write-once, read-many pattern).

- **MapReduce**:
  - Reducers write independent files to HDFS as output.
  - Focus on fast read access for multiple clients.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/hdfs4.svg)

#### HDFS Consistency Model
- **Strong Consistency**:
  - Each data block is replicated to multiple nodes.
  - Write declared successful only when all replicas are written successfully.
  - Ensures all clients see a consistent view of the file.

- **Single Writer Policy**:
  - No multiple concurrent writers for a single HDFS file.
  - Simplifies implementation of strong consistency.

---

### Key Points Summary
- **Cluster Topology**: Servers mapped to racks, network distance measured in hops, tree-style topology.
- **Rack-Aware Replication**: Strategy for replica placement enhances reliability and performance, tolerates node and rack failures.
- **Synchronization Semantics**: Supports append operations, aligns with write-once, read-many pattern common in MapReduce workloads.
- **Consistency Model**: Strong consistency ensured by successful replication to multiple nodes, single writer policy simplifies consistency implementation.


### HDFS Read Process

#### Overview
The HDFS read process involves a client requesting data from the HDFS by communicating with the NameNode and DataNodes. The following steps outline the detailed process:

1. **Initiating the Read Request**:
   - The client initiates a read request by calling the `open()` method of the `DistributedFileSystem` object.
   - The client specifies:
     - File name
     - Start offset
     - Read range length

2. **Calculating Blocks to be Read**:
   - The `DistributedFileSystem` object determines the blocks needed based on the given offset and range length.
   - It requests the block locations from the NameNode.

3. **NameNode's Role**:
   - The NameNode provides metadata for all block locations.
   - It returns a list of blocks and their replicas' locations to the client.
   - Closest replica to the client is preferred:
     - **Same Node**: Preferred if the block is on the same node as the client.
     - **Same Rack**: Preferred if the block is in the same rack.
     - **Off-Rack**: Preferred if the block is off-rack when the other two options are not available.

4. **Reading Data from DataNodes**:
   - The client calls the `read()` method of `FSDataInputStream`.
   - The `FSDataInputStream` establishes a connection with the closest DataNode with the first block of the file.
   - Data is read as streams and passed to the requesting application immediately, without waiting for the entire block to transfer.
   - Once all data of a block is read, the connection closes, and the process repeats for the next block until all required blocks are read.

5. **Completing the Read Operation**:
   - After all required blocks are read, the client calls the `close()` method of the input stream object to complete the operation.

#### Short Circuit Read
- **Definition**: When the client and data are on the same machine, HDFS can bypass the DataNode and read the file directly.
- **Benefits**:
  - Reduces overhead.
  - Minimizes processing resource usage.
  - Enhances efficiency.

### Summary
- **Step-by-Step Process**:
  - Client initiates read request (`open()` method).
  - `DistributedFileSystem` calculates needed blocks and requests locations from NameNode.
  - NameNode provides block locations, preferring the closest replicas.
  - `FSDataInputStream` handles connections and streams data to the client.
  - Client closes the input stream after reading all blocks (`close()` method).

- **Efficiency**:
  - Short circuit read optimizes read operations by bypassing DataNode when the client and data are on the same node.

These steps outline the efficient process of reading data in HDFS, ensuring high performance and reliability. 

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/hdfs5.svg)
