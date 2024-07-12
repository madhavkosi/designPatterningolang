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

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/NewsfeedHLD.gifas)

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