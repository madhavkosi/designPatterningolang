### Designing a File Hosting Service like Google Drive

#### Understanding Google Drive

Google Drive is a cloud-based file storage and synchronization service that allows users to:
- Store documents, photos, videos, and other files.
- Access their files from any device (computer, smartphone, tablet).
- Share files with friends, family, and coworkers.

### Why Cloud Storage?

#### Popularity
- Simplifies storage and exchange of digital resources across multiple devices.
- Supports various platforms and operating systems (smartphones, tablets, etc.).
- Portable access from different geographical locations at any time.

#### Benefits

1. **Availability**:
   - Data available anywhere, anytime.
   - Access files/photos from any device.

2. **Reliability and Durability**:
   - 100% data reliability and durability.
   - Multiple copies of data stored on different geographically located servers.

3. **Scalability**:
   - Unlimited storage space, expandable as needed for a fee.


### 2. Requirements and Goals of the System

#### Top-Level Requirements

1. **File Upload and Download**:
   - Users should be able to upload and download their files/photos from any device.

2. **File Sharing**:
   - Users should be able to share files or folders with other users.

3. **Automatic Synchronization**:
   - Support automatic synchronization between devices.
   - After updating a file on one device, it should synchronize on all devices.

4. **Support for Large Files**:
   - The system should support storing large files up to a GB.

5. **ACID Properties**:
   - Ensure Atomicity, Consistency, Isolation, and Durability of all file operations.

6. **Offline Editing**:
   - Users should be able to add/delete/modify files while offline.
   - Changes should sync to remote servers and other devices when online.

#### Extended Requirements

- **Snapshotting**:
  - Support snapshotting of data so that users can revert to any version of the files.

### 3. Some Design Considerations

1. **Read and Write Volumes**:
   - Expect huge read and write volumes.
   - Read to write ratio is expected to be nearly the same.

2. **File Storage in Chunks**:
   - Internally, store files in small parts or chunks (e.g., 4MB).
   - Benefits include retrying failed operations for smaller parts, reducing the amount of data exchange by transferring updated chunks only.

3. **Data Efficiency**:
   - Remove duplicate chunks to save storage space and bandwidth usage.
   - Keep a local copy of metadata (file name, size, etc.) with the client to reduce round trips to the server.

4. **Efficient Data Transfer**:
   - For small changes, clients can intelligently upload the diffs instead of the whole chunk.


### Capacity Estimation and Constraints

#### Assumptions

1. **User Base**:
   - Total users: 500 million
   - Daily Active Users (DAU): 100 million
   - Average devices per user: 3

2. **File Statistics**:
   - Average files/photos per user: 200
   - Total files: 100 billion
   - Average file size: 100KB
   - Total storage: 10 petabytes (PB)
     ```
     100 billion files * 100KB = 10PB
     ```

3. **Active Connections**:
   - One million active connections per minute

#### Back of the Envelope Estimation

1. **User Statistics**:
   - Signed up users: 50 million
   - Daily Active Users (DAU): 10 million

2. **Storage Allocation**:
   - Free space per user: 10GB
   - Total allocated space: 500 Petabytes
     ```
     50 million users * 10GB = 500PB
     ```

3. **File Upload Statistics**:
   - Average file size: 500KB
   - Files uploaded per user per day: 2

4. **Queries Per Second (QPS)**:
   - QPS for upload API:
     ```
     10 million DAU * 2 uploads / 24 hours / 3600 seconds ≈ 240 QPS
     ```

5. **Peak QPS**:
   - Peak QPS:
     ```
     240 QPS * 2 = 480
     ```


### Step 2 - Propose High-Level Design and Get Buy-In

#### Initial Single Server Setup

1. **Components**:
   - **Web Server**: Apache server for uploading and downloading files.
   - **Database**: MySQL database to manage metadata (user data, login info, files info, etc.).
   - **Storage System**: 1TB allocated storage space for files.

2. **Directory Structure**:
   - Root directory: `drive/`
   - **Namespaces**: Directories under `drive/` for each user.
   - **File Identification**: Each file or folder is uniquely identified by combining the namespace and relative path.
     ```
     drive/
     ├── user1/
     │   ├── file1.txt
     │   └── file2.jpg
     └── user2/
         ├── file3.docx
         └── file4.png
     ```

#### APIs

1. **Upload a File**
   - **Types**:
     - Simple upload for small files.
     - Resumable upload for large files and high chances of network interruption.

   - **Resumable Upload API**:
     ```
     POST https://api.example.com/files/upload?uploadType=resumable
     ```
   - **Params**:
     - `uploadType=resumable`
     - `data`: Local file to be uploaded.

   - **Steps**:
     1. Send initial request to get the resumable URL.
     2. Upload data and monitor upload state.
     3. Resume upload if interrupted.

2. **Download a File**
   - **API**:
     ```
     GET https://api.example.com/files/download
     ```
   - **Params**:
     - `path`: Download file path.
   - **Example**:
     ```json
     {
       "path": "/recipes/soup/best_soup.txt"
     }
     ```

3. **Get File Revisions**
   - **API**:
     ```
     GET https://api.example.com/files/list_revisions
     ```
   - **Params**:
     - `path`: Path to the file for revision history.
     - `limit`: Maximum number of revisions to return.
   - **Example**:
     ```json
     {
       "path": "/recipes/soup/best_soup.txt",
       "limit": 20
     }
     ```

#### Security
- **Authentication**: All APIs require user authentication
- **HTTPS**: Use HTTPS and SSL to protect data transfer between clients and backend servers.



### Moving Away from Single Server Setup

#### Initial Problem: Storage Space Full
- **Alert**: Only 10 MB of storage space left, preventing further file uploads .

#### Solution: Sharding
- **Sharding**: Data stored on multiple storage servers based on `user_id` .

#### Enhanced Storage: Amazon S3
- **Amazon S3**: Chosen for scalability, data availability, security, and performance.
  - **Same-region replication**: Data replicated within the same geographic area (left side of Figure 6).
  - **Cross-region replication**: Data replicated across different geographic areas (right side of Figure 6).
- **Benefits**: Ensures data redundancy and availability across multiple regions.

#### Further Improvements

1. **Load Balancer**:
   - **Function**: Distributes network traffic evenly.
   - **Benefit**: Redistributes traffic if a web server fails.

2. **Web Servers**:
   - **Scalability**: Add/remove web servers easily based on traffic load.

3. **Metadata Database**:
   - **Separation**: Move database out of the server to avoid single point of failure.
   - **Enhancements**: Set up data replication and sharding for availability and scalability.

4. **File Storage**:
   - **Usage**: Store files in Amazon S3.
   - **Replication**: Files replicated in two separate geographical regions for availability and durability.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/googledrive.webp" width="300" />
</p>


#### Sync Conflicts

- **Issue**: When two users modify the same file simultaneously, a conflict occurs (Figure 8).
- **Resolution Strategy**: 
  - First processed version wins.
  - Later version receives a conflict and the user is presented with both the local and server versions (Figure 9).
  - User can merge both files or override one version with the other.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/gd2.svg" width="400" />
    <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/gd3.svg" width="400" />
</p>


### High-Level Design Components (Figure 10)

1. **User**:
   - Accesses the application via browser or mobile app.

2. **Block Servers**:
   - **Function**: Upload blocks to cloud storage (S3).
   - **Mechanism**: Split files into blocks (max 4MB each).
   - **Hashing**: Each block is given a unique hash value stored in the metadata database.
   - **Storage**: Blocks treated as independent objects, stored in cloud storage (S3).
   - **Reconstruction**: To reconstruct a file, blocks are joined in a specific order.

3. **Cloud Storage**:
   - **Storage**: Store file blocks in cloud storage.

4. **Cold Storage**:
   - **Usage**: Store inactive data not accessed for a long time.

5. **Load Balancer**:
   - **Function**: Evenly distribute requests among API servers.

6. **API Servers**:
   - **Responsibilities**: Handle user authentication, profile management, and file metadata updates.

7. **Metadata Database**:
   - **Storage**: Store metadata of users, files, blocks, versions, etc.
   - **Separation**: Files stored in the cloud; only metadata stored in the database.

8. **Metadata Cache**:
   - **Function**: Cache metadata for fast retrieval.

9. **Notification Service**:
   - **Function**: Notify clients of file changes to sync updates.
   - **Mechanism**: Publisher/subscriber system for event notifications.
   - **Real-Time Updates**: Sends push notifications to clients for real-time synchronization.
   - **Event Handling**: Handles events such as file addition, modification, and deletion.
   - **Scalability**: Designed to handle a large number of events and subscribers efficiently.

10. **Offline Backup Queue**:
    - **Function**: Store changes for offline clients to sync when they are online.
    - **Mechanism**: Queues changes and ensures they are applied once the client reconnects.

#### Summary
- **Initial Setup**: Single server with Apache, MySQL, and local storage.
- **Scaling**: Implemented sharding to handle data growth.
- **Enhanced Storage**: Moved file storage to Amazon S3 for reliability and scalability.
- **Further Improvements**: Added load balancer, decoupled components, and ensured replication for high availability.
- **Sync Conflict Resolution**: Strategy to handle file modifications by multiple users.


<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/gd4.webp" width="400" />
</p>

### Step 3 - Design deep dive


## Block Servers

For large files that are updated regularly, sending the whole file on each update consumes a lot of bandwidth. Two optimizations are proposed to minimize the amount of network traffic being transmitted:

1. **Delta Sync**: When a file is modified, only the modified blocks are synced instead of the whole file using a sync algorithm.
2. **Compression**: Applying compression on blocks can significantly reduce the data size. Blocks are compressed using compression algorithms depending on file types. For example, gzip and bzip2 are used to compress text files. Different compression algorithms are needed to compress images and videos.

### Block Servers in Action

Block servers handle the heavy lifting work for uploading files. They process files passed from clients by:

1. **Splitting a File into Blocks**: A file is divided into smaller, manageable blocks.
2. **Compressing Each Block**: Each block is compressed using appropriate compression algorithms.
3. **Encrypting Each Block**: For security, each block is encrypted before it is sent to cloud storage.
4. **Uploading Blocks**: The blocks are then uploaded to the cloud storage.
<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/gd5.webp" width="400" />
   <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/gd6.webp" width="400" />
</p>
### Delta Sync

Delta sync ensures that only modified blocks are transferred to cloud storage. For instance, if a file has several blocks and only "block 2" and "block 5" are changed, then only these two blocks are uploaded to the cloud storage, reducing the amount of data transmitted and saving network traffic.

### Benefits of Block Servers

By implementing delta sync and compression, block servers significantly save network traffic and improve efficiency in file uploads and updates.