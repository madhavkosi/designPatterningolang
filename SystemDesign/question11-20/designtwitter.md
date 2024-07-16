# Designing Twitter-like Social Networking Service

## 1. What is Twitter?
Twitter is an online social networking service where users post and read short messages called "tweets." Registered users can post and read tweets, while non-registered users can only read them. Users access Twitter through their website interface, SMS, or mobile app.

## 2. Requirements and Goals of the System

### Functional Requirements
- Users should be able to post new tweets.
- A user should be able to follow other users.
- Users should be able to mark tweets as favorites.
- The service should be able to create and display a user's timeline consisting of top tweets from all the people the user follows.
- Tweets can contain photos and videos.

### Non-functional Requirements
- Our service needs to be highly available.
- Acceptable latency of the system is 200ms for timeline generation.
- Consistency can take a hit (in the interest of availability); if a user doesn’t see a tweet for a while, it should be fine.

### Extended Requirements
- Searching for tweets.
- Replying to a tweet.
- Trending topics – current hot topics/searches.
- Tagging other users.
- Tweet Notification.
- Who to follow? Suggestions?
- Moments.


## Capacity Estimation and Constraints

### 1. **Favorites per Day**
Given:
- 200 million daily active users (DAU)
- Each user favorites 5 tweets per day

Calculation:
200M DAU × 5 favorites/user = 1B favorites/day

### 2. **Total Tweet-Views per Day**
Assumptions:
- Each user visits their timeline twice a day
- Each user visits five other people's pages daily
- Each page shows 20 tweets

Calculation:
200M DAU × ((2 + 5) × 20 tweets) = 28B tweet-views/day

### 3. **Storage Estimates for Tweets**

#### Daily Storage
- 100 million new tweets per day
- Each tweet has 140 characters
- Each character requires 2 bytes
- Metadata per tweet: 30 bytes

Calculation:
100M tweets × (280 bytes + 30 bytes) = 30GB/day

#### Storage for Five Years
Calculation:
30GB/day × 365 days/year × 5 years = 54,750GB = 54.75TB

### 4. **Media Storage Estimates**

#### Daily Media Storage
- Every fifth tweet has a photo
- Every tenth tweet has a video
- Average photo size: 200KB
- Average video size: 2MB

Calculation:
(100M tweets ÷ 5 × 200KB) + (100M tweets ÷ 10 × 2MB) ≈ 24TB/day

### 5. **Bandwidth Estimates**

#### Ingress Bandwidth
- Daily media storage: 24TB

Calculation: 

24TB/day ÷ 86400 seconds/day ≈ 290MB/sec

#### Egress Bandwidth
- Total tweet views per day: 28B
- Each tweet's text: 280 bytes
- Every 5th tweet has a photo: 200KB each
- Every 10th tweet has a video, watched every 3rd time: 2MB each

### Calculation:

Text:
28B tweet-views × 280 bytes 86400 seconds ≈ 93MB/sec
86400 seconds
28B tweet-views × 280 bytes ≈ 93MB/sec

Photos:
28B ÷ 5 × 200KB 86400 seconds ≈ 13GB/sec
86400 seconds
28B ÷ 5 × 200KB ≈ 13GB/sec

Videos:
28B ÷ 10 ÷ 3 × 2MB 86400 seconds ≈ 22GB/sec
86400 seconds
28B ÷ 10 ÷ 3 × 2MB ≈ 22GB/sec

Total Egress:
93MB/sec + 13GB/sec + 22GB/sec ≈ 35GB/sec
93MB/sec + 13GB/sec + 22GB/sec ≈ 35GB/sec

### Summary
- **Favorites per Day:** 1 billion
- **Total Tweet-Views per Day:** 28 billion
- **Daily Tweet Storage:** 30GB
- **Five-Year Tweet Storage:** 54.75TB
- **Daily Media Storage:** 24TB
- **Ingress Bandwidth:** 290MB/sec
- **Egress Bandwidth:** 35GB/sec


## System APIs for Twitter-like Service

| **Function**                   | **Parameters**                                                                                                    | **Returns**                                                 |
|--------------------------------|------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------|
| **tweet**                      | api_dev_key (string), tweet_data (string), tweet_location (string, optional), user_location (string, optional), media_ids (number[], optional) | (string): URL to access the tweet if successful, otherwise an appropriate HTTP error is returned. |
| **follow_user**                | api_dev_key (string), user_id (number), follow_user_id (number)                                                   | (string): Success message if successful, otherwise an appropriate HTTP error is returned.         |
| **favorite_tweet**             | api_dev_key (string), user_id (number), tweet_id (number)                                                         | (string): Success message if successful, otherwise an appropriate HTTP error is returned.         |
| **get_timeline**               | api_dev_key (string), user_id (number), page (number), count (number)                                             | (dict): A dictionary containing the timeline tweets if successful, otherwise an appropriate HTTP error is returned. |
| **search_tweets**              | api_dev_key (string), query (string), page (number), count (number)                                               | (dict): A dictionary containing the search results if successful, otherwise an appropriate HTTP error is returned. |
| **reply_to_tweet**             | api_dev_key (string), tweet_id (number), reply_data (string), reply_location (string, optional), user_location (string, optional), media_ids (number[], optional) | (string): URL to access the reply if successful, otherwise an appropriate HTTP error is returned. |
| **tag_user_in_tweet**          | api_dev_key (string), tweet_id (number), tagged_user_ids (number[])                                               | (string): Success message if successful, otherwise an appropriate HTTP error is returned.         |


### 5. High Level System Design

- **System Requirements:**
  - **Write Load:** 
    - Daily: 100 million tweets
    - Per second: 1150 tweets
  - **Read Load:** 
    - Daily: 28 billion tweets
    - Per second: 325,000 tweets
  - Read-heavy system
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/twitter.gif)

- **Components:**
  - **Application Servers:**
    - Multiple servers to handle the high volume of requests
    - Load balancers for traffic distribution
  - **Backend Database:**
    - Efficient database capable of handling high write and read loads
  - **File Storage:**
    - Storage solution for photos and videos

- **Expected Traffic Patterns:**
  - Average per second:
    - 1160 new tweets
    - 325,000 read requests
  - Peak times:
    - Several thousand write requests per second
    - Approximately 1 million read requests per second

- **Key Considerations:**
  - Uneven traffic distribution throughout the day
  - Design must accommodate peak traffic loads


## Database Schema

We need to store data about users, their tweets, their favorite tweets, and people they follow. Here is a high-level schema for our Twitter-like service.

### Users Table

| **Column**       | **Type**     | **Description**                        |
|------------------|--------------|----------------------------------------|
| `user_id`        | BIGINT       | Primary Key, unique identifier for user|
| `username`       | VARCHAR(50)  | Unique username                        |
| `email`          | VARCHAR(100) | Unique email address                   |
| `password`       | VARCHAR(255) | Hashed password                        |
| `created_at`     | TIMESTAMP    | Account creation time                  |
| `profile_info`   | TEXT         | User profile information (bio, etc.)   |

### Tweets Table

| **Column**       | **Type**     | **Description**                        |
|------------------|--------------|----------------------------------------|
| `tweet_id`       | BIGINT       | Primary Key, unique identifier for tweet|
| `user_id`        | BIGINT       | Foreign Key, identifier for user       |
| `content`        | TEXT         | Tweet content (up to 280 characters)   |
| `media_ids`      | TEXT         | Comma-separated list of media IDs      |
| `tweet_location` | POINT        | Optional location (longitude, latitude)|
| `created_at`     | TIMESTAMP    | Tweet creation time                    |

### Favorites Table

| **Column**       | **Type**     | **Description**                        |
|------------------|--------------|----------------------------------------|
| `user_id`        | BIGINT       | Foreign Key, identifier for user       |
| `tweet_id`       | BIGINT       | Foreign Key, identifier for tweet      |
| `favorited_at`   | TIMESTAMP    | Time when tweet was favorited          |

### Follows Table

| **Column**       | **Type**     | **Description**                        |
|------------------|--------------|----------------------------------------|
| `follower_id`    | BIGINT       | Foreign Key, identifier for follower   |
| `followee_id`    | BIGINT       | Foreign Key, identifier for followee   |
| `followed_at`    | TIMESTAMP    | Time when follow action was made       |

### Media Table

| **Column**       | **Type**     | **Description**                        |
|------------------|--------------|----------------------------------------|
| `media_id`       | BIGINT       | Primary Key, unique identifier for media|
| `user_id`        | BIGINT       | Foreign Key, identifier for user       |
| `media_type`     | VARCHAR(50)  | Type of media (photo, video)           |
| `media_url`      | VARCHAR(255) | URL to access the media                |
| `uploaded_at`    | TIMESTAMP    | Time when media was uploaded           |

### Diagram

```plaintext
+-----------------+     +-----------------+     +-----------------+     +-----------------+
|     Users       |     |     Tweets      |     |    Favorites    |     |     Follows     |
+-----------------+     +-----------------+     +-----------------+     +-----------------+
| user_id (PK)    |<----| tweet_id (PK)   |<----| user_id (FK)    |<----| follower_id (FK)|
| username        |     | user_id (FK)    |     | tweet_id (FK)   |     | followee_id (FK)|
| email           |     | content         |     | favorited_at    |     | followed_at     |
| password        |     | media_ids       |     +-----------------+     +-----------------+
| created_at      |     | tweet_location  |     
| profile_info    |     | created_at      |     
+-----------------+     +-----------------+
         ^                     ^
         |                     |
         |                     |
         +---------------------+
                 |
         +-----------------+
         |     Media       |
         +-----------------+
         | media_id (PK)   |
         | user_id (FK)    |
         | media_type      |
         | media_url       |
         | uploaded_at     |
         +-----------------+
```

This schema is designed to handle the high volume of tweets and user interactions efficiently, ensuring that both reads and writes are optimized for performance and scalability.


## Data Sharding

Given the high volume of tweets and read requests, we need to distribute our data onto multiple machines for efficient read/write operations. Below are the various sharding strategies and their pros and cons.

### Sharding Strategies

1. **Sharding Based on UserID**

   - **Approach:** Store all data of a user on one server. Use a hash function to map UserID to a database server.
   - **Pros:**
     - Efficient for querying data of a single user.
   - **Cons:**
     - **Hot Users:** High load on the server holding data of popular users.
     - **Data Imbalance:** Over time, some users may accumulate more data than others, leading to uneven distribution.
     - **Solution:** Repartition/redistribute data or use consistent hashing.

2. **Sharding Based on TweetID**

   - **Approach:** Use a hash function to map each TweetID to a random server.
   - **Pros:**
     - Solves the problem of hot users.
   - **Cons:**
     - **High Latency:** Requires querying all database partitions to find tweets of a user.
     - **Solution:** Use caching to store hot tweets in front of the database servers.

3. **Sharding Based on Tweet Creation Time**

   - **Approach:** Store tweets based on their creation time.
   - **Pros:**
     - Efficient for fetching top tweets quickly.
   - **Cons:**
     - **Uneven Load Distribution:** New tweets will concentrate on one server, leaving others underutilized.
     - **Solution:** Combine with TweetID-based sharding for balanced load.

### Combined Sharding Strategy: TweetID and Creation Time

- **Approach:** Use TweetID to reflect creation time. Make each TweetID universally unique with a timestamp.
- **TweetID Structure:**
  - **Epoch Time (31 bits):** To store number of seconds for the next 50 years.
  - **Auto-incrementing Sequence (17 bits):** To store up to 130K new tweets per second.

- **TweetID Example:**
  - Epoch seconds: `1483228800`
  - TweetID: `1483228800 000001`, `1483228800 000002`, ...

- **Benefits:**
  - **Efficient Writes:** Reduced write latency by avoiding secondary indexes.
  - **Efficient Reads:** Faster reads due to epoch time in primary key.
  - **Scalability:** Can store tweets for the next 100 years with millisecond granularity.

### Implementation

- **Auto-incrementing Sequence:**
  - Use two database servers for generating keys.
  - One server generates even-numbered keys; the other generates odd-numbered keys.

### Data Sharding Diagram

```plaintext
               +---------------------+
               |      Load Balancer  |
               +---------+-----------+
                         |
         +---------------+------------------+
         |               |                  |
+--------+-----+ +-------+--------+ +-------+--------+
| App Server 1 | | App Server 2   | | App Server 3   |  ... (Horizontally scalable)
+--------------+ +----------------+ +----------------+
         |               |                  |
+--------+-----+ +-------+--------+ +-------+--------+
| Write DB      | | Read DB        | | Cache Layer   |
+--------------+ +----------------+ +----------------+
                         |
               +---------+-----------+
               |     File Storage    |
               +---------------------+
```

### High-Level Flow

1. **Writing a Tweet:**
   - Generate TweetID with epoch time and auto-incrementing sequence.
   - Use hash function to determine shard.
   - Store tweet in the appropriate shard.

2. **Reading a Tweet:**
   - Use hash function to determine shard based on TweetID.
   - Fetch tweet from the appropriate shard.

3. **Generating Timeline:**
   - Find all users a user follows.
   - Query relevant shards for tweets.
   - Aggregate and sort results to generate the timeline.

This combined sharding strategy allows efficient handling of both read and write operations, ensuring scalability and fault tolerance while managing high traffic loads effectively.