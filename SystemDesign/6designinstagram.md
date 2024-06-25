### Designing a Photo-Sharing Service like Instagram

#### Overview
Instagram is a social networking service that allows users to upload, share, and view photos and videos. Users can share content publicly or privately and follow other users. A user's News Feed consists of top photos from the people they follow.

#### Requirements and Goals

**Functional Requirements:**
- Upload, download, and view photos.
- Search based on photo/video titles.
- Follow and unfollow other users.
- Generate and display a user's News Feed with top photos from followed users.

**Non-functional Requirements:**
- **High Availability:** The service must be accessible at all times.
- **Low Latency:** News Feed generation should have an acceptable latency of 200ms.
- **Consistency vs. Availability:** Prioritize availability; slight delays in photo visibility are acceptable.
- **High Reliability:** Ensure no loss of uploaded photos or videos.

**Not in Scope:**
- Adding tags, searching by tags, commenting on photos, tagging users, and recommendation systems.

### Capacity Estimation for Instagram-like Service

#### Assumptions
- **Total Users:** 500 million
- **Active Users:** 100 million daily
- **Photos per User per Day:** 2
- **Photo Size:** 200 KB
- **Likes per Photo:** 20
- **Comments per Photo:** 5
- **Like Size:** 50 bytes
- **Comment Size:** 200 bytes

#### Estimations

1. **Photos:**
   - **Daily Uploads:** 200 million
   - **Daily Storage:** 40 TB
   - **Annual Storage:** 14.6 PB
   - **5-Year Storage:** 73 PB

2. **Likes:**
   - **Daily Likes:** 4 billion
   - **Daily Storage:** 200 GB
   - **Annual Storage:** 73 TB
   - **5-Year Storage:** 365 TB

3. **Comments:**
   - **Daily Comments:** 1 billion
   - **Daily Storage:** 200 GB
   - **Annual Storage:** 73 TB
   - **5-Year Storage:** 365 TB

4. **Bandwidth:**
   - **Daily Downloads:** 2 billion
   - **Daily Bandwidth:** 400 TB
   - **Monthly Bandwidth:** 12 PB

#### Summary
- **5-Year Storage:**
  - Photos: 73 PB
  - Likes: 365 TB
  - Comments: 365 TB
  - **Total:** 73.73 PB
- **Monthly Bandwidth:** 12 PB

This summary provides a concise view of the storage and bandwidth requirements for an Instagram-like service.

To estimate the capacity required for an Instagram-like service, we need to consider several factors including the number of users, the amount of data uploaded, storage requirements, and bandwidth needs.

#### Assumptions

- **Total Users:** 500 million
- **Active Users:** 100 million daily active users (DAUs)
- **Average Photos Uploaded per User per Day:** 2
- **Average Photo Size:** 200 KB
- **Retention Period:** Indefinite (all photos are stored permanently)

#### Estimations

1. **Photos Uploaded per Day:**
   - 100 million DAUs * 2 photos/user/day = 200 million photos/day

2. **Storage Requirement per Day:**
   - 200 million photos/day * 200 KB/photo = 40 TB/day

3. **Storage Requirement per Year:**
   - 40 TB/day * 365 days/year = 14,600 TB/year ≈ 14.6 PB/year

4. **Total Storage for 5 Years:**
   - 14.6 PB/year * 5 years = 73 PB

5. **Bandwidth Requirement:**
   - Assume each photo is downloaded 10 times on average.
   - Total downloads per day: 200 million photos/day * 10 = 2 billion downloads/day
   - Bandwidth required for downloads per day: 2 billion photos/day * 200 KB/photo = 400 TB/day

6. **Monthly Bandwidth Requirement:**
   - 400 TB/day * 30 days = 12 PB/month

#### Summary

- **Daily Uploads:** 200 million photos
- **Daily Storage:** 40 TB
- **Annual Storage:** 14.6 PB
- **5-Year Storage:** 73 PB
- **Daily Bandwidth for Downloads:** 400 TB
- **Monthly Bandwidth:** 12 PB



### High-Level System Design for Instagram-like Photo-Sharing Service

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/instagram1.gif)
#### Scenarios to Support:
1. **Upload Photos**
2. **View/Search Photos**
3. **Like Photos**
4. **Comment on Photos**
5. **Tag Users in Photos**
6. **Search by Tags**

#### Core Components:

1. **Client Application:**
   - **Mobile Apps/Web Interface:** For user interactions including uploading, viewing, liking, commenting, tagging, and searching photos.

2. **API Gateway:**
   - **Load Balancer:** Distributes incoming requests to appropriate servers to ensure even load distribution and high availability.

3. **Application Servers:**
   - **User Service:** Manages user profiles, authentication, and authorization.
   - **Photo Service:** Handles photo uploads, storage, retrieval, and metadata management.
   - **Follow Service:** Manages user follow relationships.
   - **Search Service:** Enables searching of photos by titles and tags.
   - **News Feed Service:** Generates and serves the News Feed for users.
   - **Like Service:** Manages likes on photos.
   - **Comment Service:** Manages comments on photos.
   - **Tag Service:** Manages tagging of users in photos.

4. **Storage:**
   - **Object Storage Servers:** Store the actual photo files (e.g., Amazon S3, Google Cloud Storage).
   - **Database Servers:** Store metadata about photos, likes, comments, and tags (e.g., SQL or NoSQL databases like MySQL, Cassandra).

5. **Cache:**
   - **In-Memory Cache:** Stores frequently accessed data to reduce latency (e.g., Redis, Memcached).
   - **CDN (Content Delivery Network):** Distributes photo content to users globally, reducing load on origin servers and improving access speed.

6. **Backend Services:**
   - **Worker Servers:** Handle asynchronous tasks such as processing uploaded photos (e.g., resizing, thumbnail generation), updating caches, and managing notifications.
   - **Queue System:** Manages task queues for background processing (e.g., RabbitMQ, Apache Kafka).

7. **Analytics and Monitoring:**
   - **Logging Service:** Collects and stores logs for system monitoring and debugging.
   - **Analytics Service:** Collects and analyzes user interaction data for insights and performance improvements.
   - **Monitoring and Alerts:** Monitors system health and triggers alerts for any issues (e.g., Prometheus, Grafana).

#### Data Flow

1. **Upload Photos:**
   - User uploads a photo through the client application.
   - The photo is sent to the API Gateway.
   - The load balancer routes the upload request to the Photo Service on the application servers.
   - Photo Service stores the photo in Object Storage and saves metadata in the Database.
   - Photo metadata is cached for quick access and CDN is updated for global distribution.

2. **View/Search Photos:**
   - User requests to view or search photos through the client application.
   - The request is sent to the API Gateway.
   - The load balancer routes the request to the appropriate service (Photo Service/Search Service) on the application servers.
   - Service fetches data from the cache or Database and returns the photo URLs and metadata to the client.
   - Photos are served from the CDN for fast access.

3. **Like Photos:**
   - User likes a photo through the client application.
   - The like request is sent to the API Gateway.
   - The load balancer routes the request to the Like Service.
   - Like Service updates the like count in the Database and possibly in the cache.
   - Notification services can update the photo owner's notification feed.

4. **Comment on Photos:**
   - User comments on a photo through the client application.
   - The comment request is sent to the API Gateway.
   - The load balancer routes the request to the Comment Service.
   - Comment Service stores the comment in the Database and updates any relevant caches.
   - Notification services can update the photo owner's notification feed.

5. **Tag Users in Photos:**
   - User tags another user in a photo through the client application.
   - The tag request is sent to the API Gateway.
   - The load balancer routes the request to the Tag Service.
   - Tag Service updates the tags in the Database.
   - Notification services can update tagged users' notification feeds.

6. **Search by Tags:**
   - User searches photos by tags through the client application.
   - The search request is sent to the API Gateway.
   - The load balancer routes the request to the Search Service.
   - Search Service queries the Database for photos with the specified tags and returns the results to the client.

### Database Schema for Instagram-like Service

#### 1. Users Table
- **user_id:** INT, Primary Key, Auto Increment
- **username:** VARCHAR(50), Unique, Not Null
- **email:** VARCHAR(100), Unique, Not Null
- **password_hash:** VARCHAR(255), Not Null
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP
- **profile_picture_url:** VARCHAR(255)
- **bio:** TEXT

#### 2. Followers Table
- **follower_id:** INT, Foreign Key (References Users)
- **followee_id:** INT, Foreign Key (References Users)
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP
- **Primary Key:** (follower_id, followee_id)


#### 3. Photos Table
- **photo_id:** INT, Primary Key, Auto Increment
- **user_id:** INT, Foreign Key (References Users)
- **photo_url:** VARCHAR(255), Not Null
- **caption:** TEXT
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP


#### 4. Likes Table
- **like_id:** INT, Primary Key, Auto Increment
- **photo_id:** INT, Foreign Key (References Photos)
- **user_id:** INT, Foreign Key (References Users)
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP

#### 5. Comments Table
- **comment_id:** INT, Primary Key, Auto Increment
- **photo_id:** INT, Foreign Key (References Photos)
- **user_id:** INT, Foreign Key (References Users)
- **comment_text:** TEXT
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP


#### 6. Tags Table
- **tag_id:** INT, Primary Key, Auto Increment
- **photo_id:** INT, Foreign Key (References Photos)
- **user_id:** INT, Foreign Key (References Users)
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP

### Storing Instagram-like Schema: SQL vs. NoSQL

#### SQL (e.g., MySQL)
- **Pros:** Structured data, ACID properties, efficient joins.
- **Cons:** Scalability challenges, performance bottlenecks with high loads.

#### NoSQL (e.g., Cassandra, MongoDB)
- **Pros:** Horizontal scalability, flexibility, fast read/write, reliability through replication.
- **Cons:** Eventual consistency, limited support for complex queries.

#### Hybrid Approach
- **Metadata in NoSQL:**
  - Key: `PhotoID`
  - Value: JSON object (PhotoLocation, UserLocation, CreationTimestamp, etc.)
- **Photos in Distributed Storage:**
  - Use HDFS or S3 for actual photo files.

#### NoSQL Characteristics
- **Reliability:** Multiple replicas for data availability.
- **Consistency:** Eventual consistency; support for undeleting.

### Summary
- **SQL:** For relational data needing joins and strong consistency.
- **NoSQL:** For scalable, high-throughput storage of photo metadata.
- **Photos:** Stored in distributed file storage (HDFS, S3).

Combining SQL and NoSQL leverages their strengths, ensuring scalability and performance for an Instagram-like service.


### Component Design

1. **Speed of Operations**:
   - **Photo Uploads (Writes)**: These operations are generally slow because they involve writing data to the disk.
   - **Reads**: These operations are faster, especially when served from cache.

2. **Connection Limitations**:
   - **Upload Bottleneck**: Uploading users can consume all available connections due to the slow nature of write operations.
   - **Read Availability**: When the system is busy with uploads, read requests may be delayed or unable to be served.

3. **Web Server Connection Limits**:
   - **Maximum Connections**: A typical web server may have a maximum of 500 concurrent connections.
   - **Concurrent Limits**: This means no more than 500 concurrent uploads or reads can occur simultaneously.

4. **System Design Considerations**:
   - **Separate Services**: To handle the bottleneck, separate the read and write operations into different services.
     - **Dedicated Servers**: Use dedicated servers for reads and different servers for writes.
   - **Scalability and Optimization**: 
     - **Independent Scaling**: By separating read and write services, each can be scaled independently.
     - **Optimization**: This separation allows for targeted optimization of each operation.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/instagram2.gif)



### Redundancy and High Availability

1. **File Redundancy**:
   - Store multiple copies of each file on different servers.
   - Ensures file retrieval if one server fails.

2. **Service Redundancy**:
   - Run multiple replicas of services.
   - Keeps the system available even if some services fail.

3. **Failover Mechanism**:
   - Secondary services can take over if primary fails.
   - Ensures continuous service availability.

4. **High Availability**:
   - Redundancy removes single points of failure.
   - Provides backup functionality in crises.
   - Failover can be automatic or manual.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/instagram3.gif)
