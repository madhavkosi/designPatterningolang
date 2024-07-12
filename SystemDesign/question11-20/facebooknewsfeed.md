# Designing Facebook’s Newsfeed

## 1. What is Facebook’s Newsfeed?

- **Definition**: A constantly updating list of stories in the middle of Facebook’s homepage.
- **Content Types**:
  - Status updates
  - Photos
  - Videos
  - Links
  - App activity
  - 'Likes' from people, pages, and groups that a user follows
- **Purpose**: A complete scrollable version of friends' and user’s activities including photos, videos, locations, status updates, and other activities.

## 2. Requirements and Goals of the System

### Functional Requirements

1. **Newsfeed Generation**:
   - Based on posts from people, pages, and groups a user follows.
   - Users may have many friends and follow numerous pages/groups.

2. **Content Types**:
   - Images
   - Videos
   - Text

3. **Real-time Updates**:
   - New posts should append to the newsfeed as they arrive for all active users.

### Non-Functional Requirements

1. **Real-time Generation**:
   - Newsfeed must be generated in real-time with a maximum latency of 2 seconds for the end user.

2. **Post Availability**:
   - A new post should appear in the user's feed within 5 seconds of posting, assuming a new newsfeed request occurs.

## 3. Capacity Estimation and Constraints

### User and Traffic Assumptions

- **Active Users**: 300M daily
- **Average Friends per User**: 300
- **Average Pages/Groups Followed per User**: 200
- **Newsfeed Requests**: 1.5B per day (300M users * 5 requests per day)
- **Requests per Second**: ~17,500

### Storage Requirements

- **Posts per User in Memory**: 500 posts
- **Average Post Size**: 1KB
- **Storage per User**: 500KB
- **Total Storage for All Users**: 
  - 300M users * 500KB = 150TB
  - Server Capacity: 100GB per server
  - Number of Servers Needed: 150TB / 100GB = 1500 servers


### System APIs

#### getUserFeed API

```plaintext
getUserFeed(api_dev_key, user_id, since_id, count, max_id, exclude_replies)
```

- **api_dev_key** (string): The API developer key of a registered developer.
- **user_id** (number): The ID of the user for whom the system will generate the newsfeed.
- **since_id** (number): Optional; returns results with an ID higher than (that is, more recent than) the specified ID.
- **count** (number): Optional; specifies the number of feed items to try and retrieve up to a maximum of 200 per distinct request.
- **max_id** (number): Optional; returns results with an ID less than (that is, older than) or equal to the specified ID.
- **exclude_replies** (boolean): Optional; this parameter will prevent replies from appearing in the returned timeline.

**Returns**: (JSON) Returns a JSON object containing a list of feed items.

## Database Design

### Primary Objects

1. **User**: Represents a user in the system.
2. **Entity**: Represents an entity such as a page or group.
3. **FeedItem**: Represents a post or status update.

### Relationships

- **User-Entity Relationship**: 
  - A user can follow other entities and become friends with other users.
  - Both users and entities can post FeedItems.

- **FeedItem-User Relationship**:
  - Each FeedItem will have a UserID pointing to the User who created it.
  - Each FeedItem can optionally have an EntityID pointing to the page or group where that post was created.


### Database Design

| Table         | Column        | Type          | Description                                        |
|---------------|---------------|---------------|----------------------------------------------------|
| Users         | UserID        | INT           | Primary Key                                        |
|               | Name          | VARCHAR(255)  | Name of the user                                   |
|               | Email         | VARCHAR(255)  | Unique email of the user                           |
|               | Password      | VARCHAR(255)  | User's password                                    |
|               | CreatedAt     | TIMESTAMP     | Timestamp of user creation                         |
|---------------|---------------|---------------|----------------------------------------------------|
| Entities      | EntityID      | INT           | Primary Key                                        |
|               | Name          | VARCHAR(255)  | Name of the entity                                 |
|               | Type          | VARCHAR(50)   | Type of entity (e.g., Page, Group)                 |
|               | CreatedAt     | TIMESTAMP     | Timestamp of entity creation                       |
|---------------|---------------|---------------|----------------------------------------------------|
| FeedItems     | FeedItemID    | INT           | Primary Key                                        |
|               | UserID        | INT           | Foreign Key referencing Users                      |
|               | EntityID      | INT           | Foreign Key referencing Entities (optional)        |
|               | Content       | TEXT          | Content of the feed item                           |
|               | MediaURL      | VARCHAR(255)  | URL of the media associated with the feed item     |
|               | CreatedAt     | TIMESTAMP     | Timestamp of feed item creation                    |
|---------------|---------------|---------------|----------------------------------------------------|
| UserFollow    | UserFollowID  | INT           | Primary Key                                        |
|               | UserID        | INT           | Foreign Key referencing Users                      |
|               | FollowedEntityID| INT         | Foreign Key referencing Entities                   |
|               | Type          | VARCHAR(50)   | Type of followed entity (User or Entity)           |
|               | CreatedAt     | TIMESTAMP     | Timestamp of the follow relationship creation      |
|---------------|---------------|---------------|----------------------------------------------------|
| FeedMedia     | FeedMediaID   | INT           | Primary Key                                        |
|               | FeedItemID    | INT           | Foreign Key referencing FeedItems                  |
|               | MediaType     | VARCHAR(50)   | Type of media (Image, Video, etc.)                 |
|               | MediaURL      | VARCHAR(255)  | URL of the media                                   |
|               | CreatedAt     | TIMESTAMP     | Timestamp of media creation                        |

### Observations

- **User-Entity Relation**: Stored in a separate table to manage many-to-many relationships between users and entities.
- **FeedItem-Media Relation**: Stored in a separate table to manage multiple media attachments to a feed item.


## 6. High Level System Design

### Overview

The problem can be divided into two main parts:
1. **Feed Generation**
2. **Feed Publishing**

### 1. Feed Generation

#### Steps:
1. **Retrieve Followed IDs**:
   - Retrieve IDs of all users and entities that the user (e.g., Jane) follows.
2. **Fetch Relevant Posts**:
   - Retrieve the latest, most popular, and relevant posts for those IDs.
   - These are potential posts for Jane’s newsfeed.
3. **Rank Posts**:
   - Rank these posts based on relevance to Jane.
   - This represents Jane’s current feed.
4. **Store in Cache**:
   - Store this feed in the cache.
   - Return the top posts (e.g., 20) to be rendered on Jane’s feed.
5. **Fetch More Data**:
   - When Jane reaches the end of her current feed, she can fetch the next set of posts from the server.

#### Handling New Posts:
- **Periodic Updates**:
  - Periodically (e.g., every 5 minutes), perform the above steps to rank and add newer posts to Jane’s feed.
- **User Notification**:
  - Jane can be notified of newer items in her feed that she can fetch.

### 2. Feed Publishing

#### Process:
- **Initial Load**:
  - Whenever Jane loads her newsfeed page, she requests and pulls feed items from the server.
- **Load More Data**:
  - When she reaches the end of her current feed, she can pull more data from the server.
- **New Items Notification**:
  - Either the server can notify Jane to pull new posts, or the server can push these new posts.

### System Components

1. **Web Servers**:
   - Maintain a connection with the user.
   - Transfer data between the user and the server.

2. **Application Servers**:
   - Execute workflows of storing new posts in the database servers.
   - Retrieve and push the newsfeed to the end user.

3. **Metadata Database and Cache**:
   - Store metadata about users, pages, and groups.

4. **Posts Database and Cache**:
   - Store metadata about posts and their contents.

5. **Video and Photo Storage, and Cache**:
   - Blob storage to store all media included in the posts.

6. **Newsfeed Generation Service**:
   - Gather and rank all relevant posts for a user.
   - Generate newsfeed and store in the cache.
   - Receive live updates and add newer feed items to any user’s timeline.

7. **Feed Notification Service**:
   - Notify the user that newer items are available for their newsfeed.

### High-Level Architecture Diagram

#### Description
- **User B and C are following User A.**
- **Components**:
  - Web Servers
  - Application Servers
  - Metadata Database and Cache
  - Posts Database and Cache
  - Video and Photo Storage and Cache
  - Newsfeed Generation Service
  - Feed Notification Service

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/NewsfeedHLD.gif)

### High-Level Architecture Diagram (Illustrative Description)

- **Web Servers**:
  - Interface between users and application servers.
  
- **Application Servers**:
  - Handle the logic for storing new posts, retrieving feeds, and pushing updates.

- **Metadata Database and Cache**:
  - Store user, page, and group metadata.

- **Posts Database and Cache**:
  - Store post metadata and contents.

- **Video and Photo Storage**:
  - Store media files related to posts.

- **Newsfeed Generation Service**:
  - Compile and rank posts for newsfeed generation.

- **Feed Notification Service**:
  - Notify users about new feed items.

The system architecture ensures efficient and real-time generation and delivery of personalized newsfeeds, handling large volumes of users and data effectively.


### 7. Detailed Component Design

#### a. Feed Generation
**Objective**: Generate a newsfeed with recent posts from all users and entities a user follows.

##### Query Structure
- **Query**: Fetch most recent posts from followed users/entities:
  ```sql
  (SELECT FeedItemID FROM FeedItem WHERE UserID in (
      SELECT EntityOrFriendID FROM UserFollow WHERE UserID = <current_user_id> and type = 0(user))
  )
  UNION
  (SELECT FeedItemID FROM FeedItem WHERE EntityID in (
      SELECT EntityOrFriendID FROM UserFollow WHERE UserID = <current_user_id> and type = 1(entity))
  )
  ORDER BY CreationDate DESC 
  LIMIT 100
  ```

##### Issues
1. **Performance**: Slow for users with many follows due to sorting, merging, and ranking.
2. **Latency**: High latency as timeline is generated when the user loads the page.
3. **Live Updates**: High backlog for Newsfeed Generation Service with each status update.
4. **Server Load**: Heavy load from server pushing new posts, especially for users with many followers.


## Offline Generation for Newsfeed

### Overview
The idea is to utilize dedicated servers to continuously generate users' newsfeeds and store them in memory. This approach ensures that when a user requests new posts, they are served from a pre-generated, stored location. The feeds are not compiled on-the-fly but are generated at regular intervals and served when requested.

### Key Concepts

#### Feed Generation Process
1. **Dedicated Servers**: Servers are tasked with continuously generating newsfeeds for users.
2. **Regular Updates**: Feeds are generated on a regular basis, not on-demand.
3. **Memory Storage**: Generated feeds are stored in memory for quick access.

#### Data Retrieval Process
1. **User Request**: When a user requests new posts, they are served from the pre-generated feed.
2. **Feed Update Check**: The server checks the last time the feed was generated for the user.
3. **New Data Generation**: New feed data is generated from the last generated timestamp onwards.

### Data Structures

#### Hash Table for Feed Storage
- **Key**: UserID
- **Value**: STRUCT containing feed items and the last generated timestamp.

#### STRUCT Definition
```cpp
Struct {
    LinkedHashMap<FeedItemID, FeedItem> feedItems;
    DateTime lastGenerated;
}
```

### Detailed Implementation

#### FeedItem Storage
- **LinkedHashMap or TreeMap**: Used to store FeedItemIDs, allowing easy navigation and iteration.
- **Benefits**: Enables jumping to any feed item and iterating through the map efficiently.

#### Pagination of Feed Items
1. **User's Last Seen FeedItemID**: When users want more feed items, they send the ID of the last feed item they currently see.
2. **Jump to FeedItemID**: The server jumps to the provided FeedItemID in the hash-map.
3. **Return Next Batch**: The next batch or page of feed items is returned from that point.

### Summary
- **Continuous Generation**: Newsfeeds are continuously generated and stored in memory.
- **Efficient Retrieval**: Pre-generated feeds ensure quick response times for user requests.
- **Scalable Storage**: Using data structures like LinkedHashMap or TreeMap ensures efficient feed item management and pagination.


### Feed Item Storage for User's Feed

- **Initial Decision**
  - Store 500 feed items per user.

- **Adjustments Based on Usage Patterns**
  - Monitor usage patterns to optimize storage.
  
- **Example Calculation**
  - One page of a user’s feed: 20 posts.
  - Users rarely browse more than ten pages:
    - 10 pages x 20 posts/page = 200 posts.
  - Adjust storage to 200 posts per user based on the example.

- **Handling More Posts**
  - For users who need to see more than the stored posts:
    - Query backend servers to fetch additional posts.


### Generating and Storing Newsfeeds for Users

- **Issue**
  - Many users do not log in frequently.
  - Generating and keeping newsfeeds in memory for all users is inefficient.

- **Possible Solutions**

  1. **LRU (Least Recently Used) Based Cache**
     - Remove users from memory who haven't accessed their newsfeed for a long time.
     - Simple and straightforward approach.
     - Efficient memory management.

  2. **Smarter Pre-Generation Based on User Patterns**
     - Analyze user login patterns:
       - Determine the time of day users are active.
       - Identify specific days of the week users access their newsfeed.
     - Pre-generate newsfeeds based on these patterns.
     - More efficient than generating feeds for all users indiscriminately.
     - Enhances user experience by providing timely updates.


### Feed Publishing Options

#### Fanout Methods
1. **Pull Model (Fan-out-on-load)**
   - **Description**: Users pull recent feed data from the server as needed.
   - **Advantages**: 
     - Reduces immediate load on the server when new posts are made.
   - **Disadvantages**:
     - New data might not be shown until users pull it.
     - Hard to determine optimal pull frequency, leading to potential resource wastage.

2. **Push Model (Fan-out-on-write)**
   - **Description**: New posts are immediately pushed to all followers.
   - **Advantages**: 
     - Immediate availability of new posts.
     - Reduces the need for users to fetch feeds individually.
   - **Disadvantages**:
     - High load on the server, especially for users with many followers (e.g., celebrities).
     - Requires maintaining Long Poll requests for updates.

3. **Hybrid Model**
   - **Description**: Combination of push and pull models.
   - **Strategies**:
     - Push posts only from users with few followers.
     - For users with many followers (e.g., celebrities), let followers pull updates.
     - Push to only online friends of the user.
     - Use a combination of 'push to notify' and 'pull for serving' end-users.
   - **Advantages**:
     - Balances server load and user experience.
     - More efficient use of resources compared to purely push or pull models.

#### Feed Items per Request
- **Maximum Limit**: Set a maximum limit (e.g., 20 items) per request.
- **Customizable**: Allow clients to specify the number of items to fetch based on the device used (mobile vs. desktop).

#### Notifications for New Posts
- **General Users**: Notify users of new posts available.
- **Mobile Users**: Avoid pushing data to save bandwidth.
  - Use "Pull to Refresh" mechanism on mobile devices to allow users to fetch new posts when they choose.


### Feed Ranking

#### Basic Ranking Approach
- **Creation Time**: The simplest method is to rank posts by their creation time.

#### Advanced Ranking Approach
- **Key Signals**: Identify signals that make a post important.
  - **Examples of Signals**:
    - Number of likes.
    - Number of comments.
    - Number of shares.
    - Time of the update.
    - Presence of images or videos.

#### Calculating Ranking Score
- **Feature Selection**: Choose relevant features for ranking.
- **Score Calculation**: Combine selected features to calculate a final ranking score.
  - **Simple Ranking System**:
    - Calculate score using a weighted sum of features.
  - **Advanced Ranking System**:
    - Continuously evaluate and adjust the ranking algorithm.
    - Focus on improving user stickiness, retention, and ads revenue.

#### Continuous Improvement
- **Evaluation**: Regularly assess if the ranking system improves user engagement and platform metrics.
- **Adjustment**: Make data-driven adjustments to the ranking algorithm based on performance evaluations.

### Summary
- **Initial Ranking**: Start with simple ranking by creation time.
- **Signal-Based Ranking**: Incorporate key signals like likes, comments, shares, etc., to calculate a more sophisticated ranking score.
- **Continuous Evaluation**: Regularly evaluate and adjust the ranking system to enhance user engagement and platform performance.


### Data Partitioning

#### Sharding Posts and Metadata
- **Objective**: Efficiently distribute data to handle high read/write loads.
- **Approach**: Follow a similar design as used in 'Designing Twitter'.

#### Sharding Feed Data
- **Partitioning Based on UserID**: 
  - Store all data of a user on one server.
  - Use a hash function with UserID to map users to specific cache servers.

- **Implementation Details**:
  - **Hash Function**: Map UserID to a cache server where user’s feed objects are stored.
  - **Data Limit**: Ensure storage of no more than 500 FeedItemIDs per user to fit data on a single server.

- **Query Efficiency**: 
  - Always query one server to get a user's feed.

- **Scalability and Replication**:
  - Use **Consistent Hashing** for future growth and replication.
    - **Consistent Hashing**: Helps in distributing data evenly across servers and facilitates easy addition/removal of servers.

### Summary
- **Posts and Metadata**: Distribute using sharding strategies similar to those used in large-scale systems like Twitter.
- **Feed Data**: Partition based on UserID with a hash function mapping users to specific servers, ensuring efficient querying and handling up to 500 FeedItemIDs per user.
- **Future Proofing**: Implement consistent hashing for scalable and reliable data distribution.