### Twitter Search Design

**1. Twitter Search Overview**
- A service to store and search user tweets, consisting of plain text.
- Users can update statuses (tweets) which are searchable via the system.

**2. System Requirements and Goals**
- 1.5 billion total users, 800 million daily active users.
- 400 million new tweets per day, each averaging 300 bytes.
- 500 million searches per day.
- Search queries involve multiple words with AND/OR operations.
- Efficient storage and querying of tweets are critical.

**3. Capacity Estimation and Constraints**
- **Daily Storage**: 
  - 400M tweets * 300 bytes/tweet = 120GB/day.
- **Per Second Storage**: 
  - 120GB / 24 hours / 3600 seconds ≈ 1.38MB/second.

**Key Points**
- Efficient indexing and retrieval mechanisms are essential.
- Scalability to handle high tweet and search volumes.
- Use of distributed storage systems to manage large data volumes.
- Potential use of caching and sharding for performance optimization.


### System APIs for Twitter Search

**Search API Definition**
- **Endpoint**: `search(api_dev_key, search_terms, maximum_results_to_return, sort, page_token)`

**Parameters**:
- `api_dev_key (string)`: API developer key for registered accounts, used for throttling based on quota.
- `search_terms (string)`: String containing the search terms.
- `maximum_results_to_return (number)`: Number of tweets to return.
- `sort (number)`: Optional sort mode; `0` for Latest first (default), `1` for Best matched, `2` for Most liked.
- `page_token (string)`: Token specifying a page in the result set to return.

**Returns**:
- `JSON`: Contains a list of tweets matching the search query, each entry includes user ID, user name, tweet text, tweet ID, creation time, number of likes, etc.


### High Level Design for Twitter Search

**Storage and Indexing**

- **Database**: Store all tweets in a database.
- **Indexing**: Build an index to track which words appear in which tweets.
  - **Purpose**: Quickly find tweets matching the search queries.

**Key Components**:
- **Tweet Storage**: Centralized database to store tweet data.
- **Search Index**: Inverted index mapping words to tweet IDs for efficient search.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/twittersearch.gif" width="800" />
</p>

### Detailed Component Design for Twitter Search

**1. Storage**

- **Data Partitioning**:
  - **Daily Data**: 120GB.
  - **Five Years Storage**:
    - Raw Data: 200TB.
    - With 80% Full Capacity: 250TB.
    - Including Fault Tolerance (Replication): 500TB.
  - **Modern Server Capacity**:
    - Typical Server: 4TB.
    - Required Servers: 125 servers (500TB / 4TB per server).

- **Database**:
  - **MySQL**:
    - Table: `Tweets` with columns `TweetID` (Primary Key) and `TweetText`.
    - Data Partitioning: Partition data based on `TweetID`.

- **Unique TweetIDs**:
  - **Estimated Tweets in Five Years**:
    - 400M tweets/day * 365 days/year * 5 years = 730 billion tweets.
  - **TweetID Size**:
    - Each tweet needs a unique identifier.
    - Number of Bits: log2(730 billion)≈39.4 bits
    - Rounding up, we need 40 bits.
    - 40 bits = 5 bytes.
  - **ID Generation Service**:
    - A service that generates a unique 5-byte TweetID.
    - Hash function to map TweetID to a storage server.

**2. Index**

- **Index Design**:
  - **Word Tracking**:
    - Total Words in Index: 500K (English words + common nouns).
    - Average Word Length: 5 characters.
    - Memory for Words: 2.5MB.

**Total Index Size Calculation**:

- **Memory for TweetIDs**:
  - Tweets in Two Years: 292 billion.
  - Each TweetID requires 5 bytes.
  - Memory for TweetIDs:   292 billion×5 bytes=1460 GB

- **Average Words per Tweet**:
  - On average, each tweet has 40 words.
  - Excluding common words (prepositions, articles, etc.), 15 words per tweet are indexed.
  
- **Index Entries**:
  - Each word occurrence in a tweet needs an entry in the index.
  - Each TweetID (5 bytes) is stored 15 times (once for each relevant word in the tweet).

- **Memory Requirement for Storing Index**:
  - Total memory required:
    1460GB×15=21900GB 21900GB÷1024≈21.3867TB
  - This accounts for the storage of TweetIDs across all relevant words.

- **Index Storage**:
  - The index is essentially a distributed hash table.
  - **Key**: Word.
  - **Value**: List of TweetIDs containing that word.
  - Index needs to be partitioned and distributed across multiple servers to handle memory requirements.

- **Server Requirements for Index**:
  - Assuming each server has 144GB of RAM.
  - Total Memory for Index: 21TB = 21000GB.
  - Required Servers: 
   21000/144 =~ 145.83

**Data Partitioning for Index**:

- **Sharding Based on Words**:
  - **Process**: 
    - Build the index by iterating through words in a tweet.
    - Calculate the hash of each word to determine the storage server.
  - **Advantages**:
    - Efficient querying: Query only the server holding the word.
  - **Issues**:
    - Hot words: Servers handling popular words may become overloaded.
    - Uneven Distribution: Some words may end up with a disproportionate number of TweetIDs.
  - **Solutions**:
    - Repartition data periodically.
    - Use Consistent Hashing to manage dynamic changes and ensure a more balanced distribution.

- **Sharding Based on Tweet Object**:
  - **Process**:
    - Use `TweetID` to determine the storage server.
    - Index all words of the tweet on the same server.
  - **Querying**:
    - Query all servers for tweets containing specific words.
    - Centralized server aggregates results and returns them to the user.
  - **Advantages**:
    - Uniform distribution of tweets.
    - Simpler data management.
  - **Challenges**:
    - Query latency: Increased query time due to multiple server requests.
    - Centralized Aggregation: Potential bottleneck at the aggregation server.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/twittersearch2.gif" width="800" />
</p>


### Twitter Search Design - Additional Components

**8. Cache**
- **Purpose**: Handle hot tweets efficiently.
- **Implementation**: Use Memcached to store frequently accessed tweets.
- **Operation**: Application servers check cache before querying the database.
- **Cache Eviction Policy**: Least Recently Used (LRU).
- **Scalability**: Adjust the number of cache servers based on usage patterns.

**9. Load Balancing**
- **Placement**:
  1. Between Clients and Application Servers.
  2. Between Application Servers and Backend Servers.
- **Initial Approach**: Round Robin.
  - Distributes requests equally among servers.
  - Automatically removes dead servers from rotation.
- **Advanced Approach**: Intelligent Load Balancer.
  - Monitors server load.
  - Adjusts traffic based on server capacity.

**10. Ranking**
- **Ranking Criteria**: Social graph distance, popularity, relevance.
- **Example**: Rank by popularity (likes, comments).
- **Process**:
  - Calculate a popularity score for each tweet.
  - Store the score with the index.
  - Each partition sorts results by popularity before sending to the aggregator.
  - Aggregator combines and sorts results, returning top tweets to the user.