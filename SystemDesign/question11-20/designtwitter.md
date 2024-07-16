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


![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/twitter.gif)
