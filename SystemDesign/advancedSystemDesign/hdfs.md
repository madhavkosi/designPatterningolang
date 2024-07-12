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