### Designing Typeahead Suggestion

**1. What is Typeahead Suggestion?**

Typeahead suggestions help users search for known and frequently searched terms by predicting the query based on the characters entered in real-time. This feature aids users in constructing their search queries more effectively, guiding them rather than just speeding up the search process.

**2. Requirements and Goals of the System**

**Functional Requirements:**
- **Real-time Suggestions:** As the user types their query, the service should suggest the top 10 terms starting with the entered characters.

**Non-Functional Requirements:**
- **Low Latency:** Suggestions should appear within 200ms to ensure a seamless user experience.


### Back of the Envelope Estimation for Typeahead Suggestion

**Assumptions:**
- **Daily Active Users (DAU):** 10 million
- **Average Searches per User per Day:** 10
- **Data per Query String:** 20 bytes (ASCII encoding, 4 words per query, 5 characters per word)

**Requests per Search Query:**
- **Average Requests per Search Query:** 20
- Example of requests for "dinner":
  - search?q=d
  - search?q=di
  - search?q=din
  - search?q=dinn
  - search?q=dinne
  - search?q=dinner

**Query Per Second (QPS):**
- **QPS Calculation:**
  - 10,000,000 users * 10 queries/day * 20 characters/24 hours/3600 seconds
  - ~24,000 QPS
- **Peak QPS:** 48,000 (QPS * 2)

**Storage Estimates:**
- **New Queries Daily:** 20% of daily queries
- **New Data Added Daily:** 
  - 10,000,000 users * 10 queries/day * 20 bytes/query * 20%
  - 0.4 GB/day

### Basic System Design and Algorithm

**Problem Definition:**
- Store a large number of strings allowing users to search by prefix.
- Provide suggestions for the next terms matching a given prefix.
- Example: With terms like "cap," "cat," "captain," and "capital" in the database, typing "cap" should suggest "cap," "captain," and "capital."

**Design Considerations:**
- High efficiency and low latency in querying.
- In-memory data structure is preferred over database dependency.

**Data Structure: Trie:**
- A tree-like structure for storing phrases, where each node represents a character of the phrase.
- Example storage for "cap," "cat," "caption," "captain," and "capital":
  - Traverse from the root node through each character.
- Nodes can be merged to save space if they have only one branch.

**Case Sensitivity:**
- Assume the data is case insensitive for simplicity.

**Finding Top Suggestions:**
- Store the count of searches terminating at each node.
- To find top suggestions, traverse the sub-tree of the given prefix.
- High latency due to large trees.

**Optimization:**
- Store top suggestions at each node to speed up searches.
- Store references of terminal nodes instead of entire phrases.
- Use parent references to traverse back and find suggested terms.

**Trie Construction:**
- Build the trie bottom-up.
- Parent nodes recursively call child nodes to calculate top suggestions and counts.
- Combine child suggestions to determine parent's top suggestions.

**Updating the Trie:**
- High query volume (~60K queries per second).
- Offline updates to avoid blocking read requests.
- Log queries and track frequencies.
- Use Map-Reduce (MR) setup for periodic processing (e.g., hourly).

**Offline Update Strategies:**
1. Make a copy of the trie, update offline, then switch to the new version.
2. Primary-secondary server configuration:
   - Update secondary while primary serves traffic.
   - Switch roles after the update.

**Frequency Update for Typeahead Suggestions:**
- Update only frequency differences.
- Use Exponential Moving Average (EMA) for weighting recent data.
- Traverse back from the terminal node, updating top 10 queries at each node.

**Removing a Term from Trie:**
- Remove terms completely during regular updates.
- Use a filtering layer on each server to exclude terms before sending to users.

**Ranking Criteria for Suggestions:**
- Consider factors beyond simple count:
  - Freshness
  - User location
  - Language
  - Demographics
  - Personal search history


### Trie Data Structure and Optimization for Autocomplete

**Overview of Trie:**
- **Definition:** Trie (pronounced "try") is a tree-like data structure designed for efficient string retrieval.
- **Structure:**
  - The root represents an empty string.
  - Each node stores a character and has up to 26 children for each possible character (A-Z).
  - Nodes represent complete words or prefixes.

**Basic Trie Operations:**
- **Insertion:** Add characters of a word sequentially from the root, creating new nodes as needed.
- **Search:** Traverse nodes according to characters of the query string.
- **Prefix Search:** Follow nodes matching the prefix, then collect all child nodes forming valid words.

**Autocomplete Using Trie:**
- **Steps:**
  1. **Find the Prefix Node:** Traverse nodes to match the prefix. (Time Complexity: \(O(p)\), where \(p\) is the length of the prefix)
  2. **Traverse Subtree:** Collect all valid children nodes forming words. (Time Complexity: \(O(c)\), where \(c\) is the number of children nodes)
  3. **Sort and Select Top \(k\):** Sort children nodes by frequency and select the top \(k\). (Time Complexity: \(O(c \log c)\))

**Optimization Strategies:**
1. **Limit Prefix Length:**
   - **Concept:** Users rarely type long queries, so limiting the prefix length (e.g., to 50 characters) reduces the prefix search complexity to \(O(1)\).
   - **Benefit:** This reduces the time complexity for finding the prefix node to a constant time.

2. **Cache Top Search Queries at Each Node:**
   - **Concept:** Store the top \(k\) most frequently used queries at each node to avoid traversing the entire trie.
   - **Benefit:** Fetching top \(k\) queries becomes a simple lookup, reducing time complexity to \(O(1)\).

**Optimized Algorithm:**
1. **Find the Prefix Node:** Time complexity \(O(1)\).
2. **Return Top \(k\) Queries:** Time complexity \(O(1)\).

**Space-Time Trade-off:**
- **Increased Space Usage:** Storing top \(k\) queries at every node requires more memory.
- **Improved Time Efficiency:** Fast response time is achieved, making the trade-off worthwhile for practical applications.

**Example:**
- **Initial Setup:**
  - Trie with queries: "tree", "try", "true", "toy", "wish", "win" and their frequencies.
  - Cached top 5 queries at each node.
- **Execution:**
  - User types "tr".
  - Prefix node "tr" is found.
  - Cached top 2 queries [true: 35, try: 29] are returned instantly.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/trie.svg)


### Scalable Data Gathering Service for Trie-based Autocomplete

In the previous design, real-time updates to the trie were inefficient and impractical due to the high volume of user queries. To address this, a scalable data gathering service is designed, focusing on aggregating data efficiently and updating the trie periodically.

**Components of the Data Gathering Service:**

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/datagatheringservice.svg)
1. **Analytics Logs:**
   - **Function:** Stores raw search query data.
   - **Characteristics:** Append-only, non-indexed logs.
   - **Example Log:**
     ```plaintext
     query     time
     tree      2019-10-01 22:01:01
     try       2019-10-01 22:01:05
     tree      2019-10-01 22:01:30
     toy       2019-10-01 22:02:22
     tree      2019-10-02 22:02:42
     try       2019-10-03 22:03:03
     ```

2. **Aggregators:**
   - **Function:** Aggregates raw data into a more usable format.
   - **Use Cases:**
     - **Real-time Applications:** Shorter aggregation intervals (e.g., every minute).
     - **Less Time-sensitive Applications:** Longer aggregation intervals (e.g., weekly).
   - **Example Aggregated Data:**
     ```plaintext
     query    time          frequency
     tree     2019-10-01    12000
     tree     2019-10-08    15000
     tree     2019-10-15    9000
     toy      2019-10-01    8500
     toy      2019-10-08    6256
     toy      2019-10-15    8866
     ```

3. **Workers:**
   - **Function:** Perform asynchronous jobs to build and update the trie data structure.
   - **Frequency:** Weekly updates to maintain a balance between data freshness and system performance.

4. **Trie Cache:**
   - **Function:** Distributed cache system that stores the trie in memory for fast access.
   - **Update Mechanism:** Takes weekly snapshots of the Trie DB.

5. **Trie DB:**
   - **Function:** Persistent storage for the trie.
   - **Storage Options:**
     - **Document Store:** Serializes the trie and stores snapshots (e.g., MongoDB).
     - **Key-Value Store:** Maps each prefix and its associated data to key-value pairs.
     - **Example Mapping:**
       ```plaintext
       Trie Node (Prefix)   <Key, Value>
       "tr"                 <"tr", {"tree": 12000, "try": 29000, "true": 35000}>
       "to"                 <"to", {"toy": 14000}>
       ```

**Optimized Algorithm for Top K Queries:**

1. **Find the Prefix Node:** Time complexity \(O(1)\) due to limited prefix length.
2. **Return Top K Queries:** Time complexity \(O(1)\) with cached results.

**Design Considerations:**

- **Data Freshness vs. System Performance:** Real-time updates are avoided to prevent slowing down the query service.
- **Use Case Variability:** Adjust the aggregation interval based on the application's real-time requirements.
- **Space-Time Trade-off:** Caching top queries at each node increases memory usage but significantly reduces query response time.

By leveraging these components and optimizations, the data gathering service can handle high query volumes efficiently while providing fast and relevant autocomplete suggestions. This design balances the need for up-to-date data with system performance, ensuring scalability and responsiveness.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/trie.svg)


### Optimized Query Service for Autocomplete

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/searchQuery.webp)


In the improved design of the query service, several optimizations ensure efficient and fast retrieval of autocomplete suggestions. Here is an overview of the optimized query service architecture and its components:

**Architecture Overview**

1. **Search Query Handling:**
   - **Step 1:** A search query is sent to the **load balancer**.
   - **Step 2:** The load balancer routes the request to **API servers**.
   - **Step 3:** API servers fetch trie data from the **Trie Cache** and construct autocomplete suggestions for the client.
   - **Step 4:** If the data is not found in the Trie Cache (cache miss), the system replenishes the cache to handle subsequent requests more efficiently.

2. **Cache Management:**
   - **Trie Cache:** Stores the trie structure in memory for quick access. Cache replenishment ensures that cache misses are minimized and that the trie is always up-to-date.

**Optimizations for Lightning-Fast Query Service**

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/ajax.svg)

1. **AJAX Requests:**
   - **Benefit:** For web applications, using AJAX requests allows sending and receiving autocomplete results without refreshing the entire web page. This reduces latency and improves user experience.

2. **Browser Caching:**
   - **Mechanism:** Autocomplete suggestions can be cached in the browser to serve subsequent requests directly from the cache.
   - **Example (Google Search):** Google caches autocomplete results in the browser for one hour. The response header includes `cache-control: private, max-age=3600`, indicating that the cache is valid for 3600 seconds and is intended for a single user.

3. **Data Sampling:**
   - **Purpose:** To reduce the processing power and storage required for logging every search query, data sampling can be implemented.
   - **Implementation:** Only 1 out of every N requests is logged by the system, significantly reducing the load on logging infrastructure.

**Detailed Example and Figures**

- **Figure 11: Improved Query Service Design:**
  - **Load Balancer:** Distributes incoming queries evenly across API servers.
  - **API Servers:** Handle requests by fetching data from the Trie Cache.
  - **Trie Cache:** Maintains frequently accessed trie data in memory for fast retrieval.

- **Figure 12: Browser Caching Example:**
  - **Cache-Control Header:** Example from Google search engine showing `cache-control: private, max-age=3600`, indicating a cache duration of one hour for autocomplete suggestions.

By implementing these optimizations, the query service can achieve high-speed performance, ensuring that autocomplete suggestions are provided to users promptly and efficiently. This design leverages caching at both the server and browser levels, as well as efficient data handling strategies to maintain scalability and performance.



### Trie Operations in the Autocomplete System

Trie is an essential component of the autocomplete system, enabling efficient string retrieval. Here, we detail the create, update, and delete operations for maintaining the trie structure.

**Create Operation**

**Purpose:** To build the trie using aggregated data from the Analytics Log/DB.

**Process:**
1. **Data Aggregation:** Workers aggregate search query data over a specified period (e.g., weekly).
2. **Trie Construction:** Using the aggregated data, workers construct the trie by inserting each query and its frequency into the trie structure.
3. **Storage:** The newly constructed trie is stored in the Trie Cache and Trie DB for fast access and persistence.

**Update Operation**

There are two primary methods for updating the trie:

**Option 1: Weekly Update**
- **Description:** A new trie is created weekly using the latest aggregated data and replaces the old trie.
- **Advantages:** Ensures the trie is updated comprehensively and consistently.
- **Process:**
  1. **Data Aggregation:** Aggregate data weekly.
  2. **Trie Construction:** Build a new trie from the aggregated data.
  3. **Replacement:** Replace the old trie with the newly constructed trie in both Trie Cache and Trie DB.

**Option 2: Direct Node Update**
- **Description:** Individual trie nodes are updated directly. This method is less preferred due to its complexity and potential slowness but can be acceptable for small trie sizes.
- **Advantages:** Allows for real-time updates of specific nodes.
- **Process:**
  1. **Node Update:** Update the specific node with new data.
  2. **Ancestor Update:** Update all ancestor nodes up to the root to reflect changes in the top queries of the updated node.
  3. **Example:** If the query "beer" is updated from a frequency of 10 to 30, the node and all its ancestors are updated to reflect this change (as shown in Figure 13).

**Delete Operation**

**Purpose:** To remove unwanted autocomplete suggestions (e.g., hateful, violent, sexually explicit, or dangerous content).

**Process:**
1. **Filter Layer:** A filter layer is added in front of the Trie Cache to intercept and filter out unwanted suggestions based on predefined rules (as shown in Figure 14).
2. **Asynchronous Removal:** Unwanted suggestions are marked for deletion and removed asynchronously from the database to ensure the correct data set is used for future trie updates.
3. **Physical Deletion:** The unwanted suggestions are physically removed from the database during the next update cycle.
4. **Example:** If a suggestion like "violent" needs to be removed, it is first filtered out at the cache level, and then removed from the database asynchronously.

![alt text](https://raw.githubusercontent.com/madhavkosi/designPatterningolang/main/SystemDesign/image%20folder/filter.webp)


### Scalable Storage for Large Trie Structures

When the trie grows too large for a single server, sharding is necessary to distribute the load.

**Sharding Strategy**

**First-Level Sharding:**
- **Method:** Based on the first character of the query.
  - **Example:**
    - Two servers: 'a' to 'm' on one, 'n' to 'z' on the other.
    - Up to 26 servers for 'a' to 'z'.
- **Issue:** Imbalance due to uneven distribution of queries (e.g., 'c' vs. 'x').

**Advanced Sharding:**
- **Method:** Use second or third characters for further splits.
  - **Example:** 'a' can be split into 'aa-ag', 'ah-an', 'ao-au', 'av-az'.
- **Smarter Sharding:** Analyze historical data for balanced distribution.
  - **Example:** Separate shards for 's' and 'u' to 'z'.

**Implementation**

1. **Shard Map Manager:** 
   - Manages lookup database for shard assignment based on historical data.
2. **Dynamic Sharding:** 
   - Periodically adjust shard boundaries for balance.
3. **Data Storage:**
   - **Trie Cache:** Distributed in-memory storage.
   - **Trie DB:** Persistent storage, distributed by sharding logic.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/shard-map-manager.svg)
**Query Handling Flow**

1. Query sent to load balancer.
2. Load balancer routes to appropriate API server.
3. API server checks Shard Map Manager for correct shard.
4. Query directed to correct server, fetching data from Trie Cache or Trie DB.

**Conclusion**

Advanced sharding and a Shard Map Manager ensure even data distribution and scalable storage for trie-based autocomplete systems, maintaining high performance as the system grows.


**Extending Trie-based Autocomplete to Multiple Languages and Real-Time Queries**

**Supporting Multiple Languages:**
- **Unicode Characters:** Use Unicode characters in trie nodes to support all global writing systems.

**Country-Specific Top Queries:**
- **Separate Tries:** Build different tries for each country.
- **CDN Storage:** Store these tries in CDNs to improve response time.

**Supporting Trending (Real-Time) Queries:**
- **Challenge:** Real-time events can make certain queries suddenly popular, which weekly updates cannot handle.
- **Ideas for Real-Time Support:**
  1. **Sharding:** Reduce the working dataset by sharding.
  2. **Ranking Model:** Adjust the ranking model to give more weight to recent queries.
  3. **Stream Processing:** Handle continuous data streams using technologies like Apache Hadoop, Apache Spark Streaming, Apache Storm, and Apache Kafka.

**Conclusion:**
Extending the system for multiple languages and real-time queries involves using Unicode, country-specific tries, CDN storage, and stream processing technologies to handle continuous data and trending search queries effectively.