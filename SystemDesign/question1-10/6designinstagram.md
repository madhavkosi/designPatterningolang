### Designing a Photo-Sharing Service like Instagram

**Overview**
Instagram is a social networking service that allows users to upload, share, and view photos and videos. Users can share content publicly or privately and follow other users. A user's News Feed consists of top photos from the people they follow.

### Requirements and Goals

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

**Assumptions**
- **Total Users:** 500 million
- **Active Users:** 100 million daily
- **Photos per User per Day:** 2
- **Photo Size:** 200 KB
- **Likes per Photo:** 20
- **Comments per Photo:** 5
- **Like Size:** 50 bytes
- **Comment Size:** 200 bytes

**Estimations**
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

**Summary**
- **5-Year Storage:**
  - Photos: 73 PB
  - Likes: 365 TB
  - Comments: 365 TB
  - **Total:** 73.73 PB
- **Monthly Bandwidth:** 12 PB





### High-Level System Design for Instagram-like Photo-Sharing Service

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/instagram1.gif)
**Scenarios to Support:**
1. **Upload Photos**
2. **View/Search Photos**
3. **Like Photos**
4. **Comment on Photos**
5. **Tag Users in Photos**
6. **Search by Tags**

**Core Components:**
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

**Data Flow**
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

**1. Users Table**
- **user_id:** INT, Primary Key, Auto Increment
- **username:** VARCHAR(50), Unique, Not Null
- **email:** VARCHAR(100), Unique, Not Null
- **password_hash:** VARCHAR(255), Not Null
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP
- **profile_picture_url:** VARCHAR(255)
- **bio:** TEXT

**2. Followers Table**
- **follower_id:** INT, Foreign Key (References Users)
- **followee_id:** INT, Foreign Key (References Users)
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP
- **Primary Key:** (follower_id, followee_id)


**3. Photos Table**
- **photo_id:** INT, Primary Key, Auto Increment
- **user_id:** INT, Foreign Key (References Users)
- **photo_url:** VARCHAR(255), Not Null
- **caption:** TEXT
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP


**4. Likes Table**
- **like_id:** INT, Primary Key, Auto Increment
- **photo_id:** INT, Foreign Key (References Photos)
- **user_id:** INT, Foreign Key (References Users)
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP

**5. Comments Table**
- **comment_id:** INT, Primary Key, Auto Increment
- **photo_id:** INT, Foreign Key (References Photos)
- **user_id:** INT, Foreign Key (References Users)
- **comment_text:** TEXT
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP


**6. Tags Table**
- **tag_id:** INT, Primary Key, Auto Increment
- **photo_id:** INT, Foreign Key (References Photos)
- **user_id:** INT, Foreign Key (References Users)
- **created_at:** TIMESTAMP, Default CURRENT_TIMESTAMP

### Storing Instagram-like Schema: SQL vs. NoSQL

**SQL (e.g., MySQL)**
- **Pros:** Structured data, ACID properties, efficient joins.
- **Cons:** Scalability challenges, performance bottlenecks with high loads.

**NoSQL (e.g., Cassandra, MongoDB)**
- **Pros:** Horizontal scalability, flexibility, fast read/write, reliability through replication.
- **Cons:** Eventual consistency, limited support for complex queries.

**Hybrid Approach**
- **Metadata in NoSQL:**
  - Key: `PhotoID`
  - Value: JSON object (PhotoLocation, UserLocation, CreationTimestamp, etc.)
- **Photos in Distributed Storage:**
  - Use HDFS or S3 for actual photo files.

**NoSQL Characteristics**
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


### Data Sharding for Metadata

**a. Partitioning based on UserID**
- **Sharding Method**: UserID % 10
- **PhotoID Generation**: Each shard uses its own auto-increment sequence, appended with ShardID for uniqueness.
- **Issues**:
  - **Hot Users**: Popular users can create load imbalance.
  - **Non-uniform Storage**: Some users have significantly more photos.
  - **Single Shard Limitation**: If one shard can't store all photos of a user, distributing photos across shards may increase latency.
  - **Shard Availability**: If a shard is down, all data for that user is unavailable, causing higher latency under high load.

**b. Partitioning based on PhotoID**
- **Sharding Method**: PhotoID % 10
- **PhotoID Generation**: 
  - **Single DB**: A dedicated instance generates auto-increment IDs, which may become a single point of failure.
  - **Dual DB**: Two databases generate even and odd IDs to avoid single point of failure.
  - **Load Balancer**: Round-robin between databases for key generation.
- **Key Generation Scripts**:
  - **Server1**:
    - auto-increment-increment = 2
    - auto-increment-offset = 1
  - **Server2**:
    - auto-increment-increment = 2
    - auto-increment-offset = 2
- **Scalability**: Implement a 'key' generation scheme similar to TinyURL.

**Planning for Future Growth**
- **Logical Partitions**: Create multiple logical partitions to accommodate data growth.
- **Physical Database Servers**: Multiple logical partitions can reside on a single physical server initially.
- **Migration**: Move logical partitions to different servers as needed.
- **Config File**: Map logical partitions to database servers for easy migration and updates.


###  Ranking and News Feed Generation

**Fetching Top Photos**
- **Process**:
  1. Get the list of people the user follows.
  2. Fetch metadata of each user's latest 100 photos.
  3. Submit photos to the ranking algorithm.
  4. Return the top 100 photos based on recency, likeness, etc.
- **Issue**: Higher latency due to multiple queries and sorting/merging/ranking operations.

**Pre-generating the News Feed**
- **Method**:
  1. Dedicated servers continuously generate and store users' News Feeds in the 'UserNewsFeed' table.
  2. Query this table to get the latest photos for the News Feed.
  3. Generate new News Feed data from the last update time for each user.
- **Benefits**: Reduces latency by pre-generating and storing the feed.

**Approaches for Sending News Feed Contents**
1. **Pull**:
   - **Method**: Clients pull the News Feed from the server at regular intervals or manually.
   - **Problems**:
     - New data might not be shown until a pull request is made.
     - Pull requests may often return no new data.

2. **Push**:
   - **Method**: Servers push new data to users as soon as it's available.
   - **Requirement**: Users maintain a Long Poll request with the server.
   - **Problems**: High push frequency for users following many people or celebrities.

3. **Hybrid**:
   - **Approach**: 
     - High-follower users use a pull-based model.
     - Few-follower users receive push updates.
   - **Alternative**:
     - Push updates to all users at a limited frequency.
     - High-update users pull data regularly.



### Cache and Load Balancing

**Massive-Scale Photo Delivery System**
- **Global Users**: Serve a large number of geographically distributed users.
- **Content Delivery**: Push content closer to the user using geographically distributed photo cache servers and CDNs.

**Caching Strategy**
- **Metadata Caching**:
  - **Cache Hot Rows**: Use Memcache to store frequently accessed database rows.
  - **Application Servers**: Check Memcache before querying the database.
  - **Eviction Policy**: Use Least Recently Used (LRU) to discard the least recently accessed rows first.

**Intelligent Caching**
- **Eighty-Twenty Rule**:
  - **Traffic Distribution**: 20% of daily read volume generates 80% of the traffic.
  - **Popular Photos**: Certain photos are significantly more popular.
  - **Caching Strategy**: Cache 20% of the daily read volume of photos and metadata to maximize efficiency.

**Implementation Example**
1. **Photo Delivery**:
   - Use CDNs to store and deliver photos closer to users.
   - Geographically distributed cache servers reduce latency and improve load times.

2. **Metadata Caching with Memcache**:
   - **Cache Check**: Before hitting the database, application servers check Memcache.
   - **Cache Hit**: If data is found in Memcache, serve it directly.
   - **Cache Miss**: If data is not found, query the database and update Memcache.

3. **LRU Eviction Policy**:
   - Maintain a priority list of cached items based on access time.
   - Discard the least recently accessed items when the cache is full.

4. **Applying the Eighty-Twenty Rule**:
   - Identify the 20% of photos that generate 80% of the traffic.
   - Prioritize caching these popular photos and their metadata.

**Benefits**
- **Reduced Latency**: Serving data from cache or nearby CDN reduces response time.
- **Load Balancing**: Distributes load across multiple cache servers, reducing the burden on the central database.
- **Efficiency**: Intelligent caching maximizes the use of cache space, focusing on the most frequently accessed data.

By implementing these strategies, the service can efficiently handle a large volume of global users, providing fast and reliable access to photos and metadata.