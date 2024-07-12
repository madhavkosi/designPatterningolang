# Designing a Yelp-like Service or Nearby Friends

## 1. Introduction

A Yelp-like service allows users to search for nearby places such as restaurants, theaters, shopping malls, etc., and to add/view reviews of these places. This service can be considered a proximity server used to discover nearby attractions.

## 2. Requirements and Goals

### Functional Requirements

1. Users can add, delete, and update information about places.
2. Users can find nearby places given their location (longitude/latitude) within a specified radius.
3. Users can add reviews/feedback about a place, including pictures, text, and a rating.

### Non-functional Requirements

1. Real-time search experience with minimal latency.
2. Support for a heavy search load, with a higher volume of search requests compared to place updates.

## 3. Scale Estimation

1. **Places**: 500M places
2. **Queries per Second (QPS)**: 100K
3. **Annual Growth**: 20% increase in places and QPS


## 4. Database Schema

**Places Table**: This table will store information about each place.

- **LocationID** (8 bytes): Uniquely identifies a location.
- **Name** (256 bytes): Name of the place.
- **Latitude** (8 bytes): Latitude of the place.
- **Longitude** (8 bytes): Longitude of the place.
- **Description** (512 bytes): Description of the place.
- **Category** (1 byte): Category of the place (e.g., coffee shop, restaurant, theater, etc.).

**Total Size**: 8 + 256 + 8 + 8 + 512 + 1 = 793 bytes

**Reviews Table**: This table will store reviews for each place.

- **LocationID** (8 bytes): Foreign key referencing the Places table.
- **ReviewID** (4 bytes): Uniquely identifies a review.
- **ReviewText** (512 bytes): Text of the review.
- **Rating** (1 byte): Rating of the place (0-10 stars).

**Total Size**: 8 + 4 + 512 + 1 = 525 bytes

**Photos Table**: This table will store photos for each place and review.

- **PhotoID** (4 bytes): Uniquely identifies a photo.
- **LocationID** (8 bytes): Foreign key referencing the Places table.
- **ReviewID** (4 bytes): Foreign key referencing the Reviews table (can be NULL for place photos).
- **PhotoURL** (256 bytes): URL of the photo.

**Total Size**: 4 + 8 + 4 + 256 = 272 bytes


## 5. Api contract

| **Function**          | **Parameters**                                                                                                                                                                       | **Returns**                                                |
|-----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|
| **search**            | api_dev_key (string), search_terms (string), user_location (string), radius_filter (number), maximum_results_to_return (number), category_filter (string), sort (number), page_token (string) | JSON with a list of places (name, address, category, rating, thumbnail) |
| **add_place**         | api_dev_key (string), name (string), latitude (number), longitude (number), description (string), category (string)                                                                   | JSON confirming place creation                             |
| **get_place_details** | api_dev_key (string), location_id (number)                                                                                                                                           | JSON with detailed information about the place             |
| **add_review**        | api_dev_key (string), location_id (number), review_text (string), rating (number), photos (array of strings)                                                                          | JSON confirming review submission                          |
| **get_reviews_for_place** | api_dev_key (string), location_id (number)                                                                                                                                           | JSON with a list of reviews for the specified place        |

### Basic System Design and Algorithm

**Overview:**
- **Goal:** Efficiently store and index datasets (places, reviews, etc.) for real-time querying.
- **Challenges:** Ensuring read efficiency, especially when location data changes infrequently.

### Storage and Indexing Methods

#### a. SQL Solution
- **Database:** MySQL
- **Structure:**
  - Each place stored in a row, identified by `LocationID`.
  - Longitude and latitude in separate columns.
  - Indexes on longitude and latitude for fast searching.
- **Query Example:**
  ```sql
  SELECT * FROM Places 
  WHERE Latitude BETWEEN X-D AND X+D 
  AND Longitude BETWEEN Y-D AND Y+D
  ```
- **Efficiency Issues:**
  - Separate indexes may return large lists, making intersections inefficient.
  - High volume of places between ranges `X-D` to `X+D` and `Y-D` to `Y+D`.

#### b. Grids
- **Concept:** Divide the map into smaller grids to group locations.
- **Advantages:** 
  - Query fewer grids for nearby places.
  - Store `GridID` with each location for faster searches.
- **Query Example:**
  ```sql
  SELECT * FROM Places 
  WHERE Latitude BETWEEN X-D AND X+D 
  AND Longitude BETWEEN Y-D AND Y+D 
  AND GridID IN (GridID, GridID1, GridID2, ..., GridID8)
  ```
- **Memory Usage:**
  - Storing the index in memory using a hash table.
  - Estimated memory requirement: 4GB.

### Dynamic Size Grids

**Goal:** Efficiently manage and query places data by dynamically adjusting grid sizes based on the number of places within each grid.

#### Data Structure: QuadTree

- **QuadTree Overview:**
  - A tree structure where each node has up to four children.
  - Each node represents a geographic grid.
  - Leaf nodes contain a list of places within the grid.
  - Internal nodes store pointers to their child nodes.

- **Node Splitting:**
  - If a grid/node contains more than 500 places, it splits into four smaller grids.
  - Each new grid becomes a child node.
  - This process continues until all grids have ≤500 places.

- **Leaf Nodes:**
  - Represent grids that cannot be split further.
  - Maintain a list of places (LocationID and lat/long).

#### Memory Management

- **Memory Requirement Calculation:**
  - Each place needs 24 bytes (8 bytes for `LocationID`, 8 bytes each for `lat` and `long`).
  - For 500M places: 
    \[
    24 \times 500M = 12 \text{ GB}
    \]
  - QuadTree with 1M leaf nodes and additional internal nodes requires:
    \[
    1M \times \frac{1}{3} \times 4 \times 8 \text{ bytes} = 10 \text{ MB}
    \]
  - Total memory: 12.01GB (easily fits in modern servers).

#### Querying and Searching

- **Search Workflow:**
  - **Step 1:** Start at the root node and traverse down to the leaf node containing the user's location.
  - **Step 2:** If the leaf node has sufficient places, return them.
  - **Step 3:** If not, expand the search to neighboring grids (via parent pointers or doubly linked list) until finding enough places or reaching the maximum search radius.

- **Finding Neighboring Grids:**
  - **Method 1:** Use a doubly linked list to connect all leaf nodes.
  - **Method 2:** Use parent pointers to navigate to sibling nodes and expand the search.

#### Inserting a New Place

- **Single Server:**
  - Directly add the new place to the appropriate node in the QuadTree.
  - Split the node if it exceeds 500 places.

- **Distributed Servers:**
  - Locate the correct server/grid for the new place.
  - Insert the place into the corresponding node on that server.



### Summary
Dynamic Size Grids using a QuadTree structure provide an efficient and scalable solution for storing and querying large datasets of places. By dynamically adjusting grid sizes based on place density, the system ensures fast and memory-efficient operations, making it suitable for applications with varying geographical densities and real-time querying needs.

### Summary
- **SQL Solution:** Straightforward but potentially inefficient due to large index intersections.
- **Static Grids:** Improves efficiency by limiting search to fewer grids.
- **Dynamic Size Grids (QuadTree):** Most efficient for densely populated areas, dynamically adjusts grid sizes, and minimizes search time and memory usage.




### 7. Data Partitioning

#### Overview
- **Problem**: Huge number of places causes the index to exceed single machine’s memory and server capacity for read traffic.
- **Solution**: Partition the QuadTree to distribute load and data efficiently.

#### Solutions for Partitioning

1. **Sharding Based on Regions**
   - **Method**: Divide places into regions (e.g., zip codes). Each region is assigned to a fixed node.
   - **Storage**: Places are stored on the server corresponding to their region.
   - **Querying**: Query the server responsible for the region containing the user's location.

   **Issues**:
   - **Hot Regions**: High query traffic in a single region can slow down the server, affecting performance.
   - **Uneven Distribution**: Over time, some regions may store significantly more places than others, leading to imbalance.
   - **Solutions**:
     - **Repartition Data**: Redistribute places across servers to maintain balance.
     - **Consistent Hashing**: Dynamically assign places to servers to achieve more uniform distribution.

2. **Sharding Based on LocationID**
   - **Method**: Use a hash function to map each LocationID to a server.
   - **Storage**: Calculate the hash of each LocationID and store the place on the corresponding server.
   - **Querying**: Query all servers; each returns a set of nearby places, which are aggregated by a centralized server.

   **Considerations**:
   - **Different QuadTree Structures**: 
     - Due to uneven distribution of places, each partition may have a different QuadTree structure.
     - Ensures approximately equal number of places on all servers.
   - **Impact**: Different tree structures do not cause issues as searches involve querying all neighboring grids across all partitions.


### 8. Replication and Fault Tolerance

#### Overview
- **Purpose**: Enhance system reliability and distribute read traffic using replication.
- **Configuration**: Use a primary-secondary setup for QuadTree servers.

#### Replication Strategy

1. **Primary-Secondary Configuration**
   - **Primary Server**: Handles all write traffic.
   - **Secondary Server(s)**: Serve read traffic; receive updates from the primary.
   - **Delay**: Secondary servers may have a slight delay (a few milliseconds) in reflecting recent changes.

#### Fault Tolerance

1. **Primary Server Failure**
   - **Solution**: Secondary server takes over as the primary after a failover.
   - **Result**: Both servers maintain the same QuadTree structure.

2. **Simultaneous Primary and Secondary Server Failure**
   - **Issue**: Rebuilding the QuadTree without knowing the exact places stored on the server.
   - **Brute-force Solution**: Iterate through the database and use the hash function to filter required places.
     - **Disadvantages**: Inefficient, slow, and causes service disruption.

#### Efficient Recovery with Reverse Index

1. **Reverse Index Creation**
   - **Purpose**: Efficiently map Places to their QuadTree server.
   - **QuadTree Index Server**:
     - Holds information on which Places are stored on each QuadTree server.
     - Uses a HashMap where:
       - **Key**: QuadTree server number.
       - **Value**: HashSet of Places (including LocationID and Lat/Long).
     - **Advantages**: 
       - Quick addition/removal of Places.
       - Fast rebuild of QuadTree servers.

2. **Rebuilding a QuadTree Server**
   - **Process**: QuadTree server requests the list of Places from the QuadTree Index server.
   - **Result**: Rapid recovery and minimized service disruption.

3. **Fault Tolerance for QuadTree Index Server**
   - **Replication**: Maintain a replica of the QuadTree Index server.
   - **Recovery**: If the primary QuadTree Index server fails, it can rebuild its index by iterating through the database.

### Key Points
- **Primary-Secondary Setup**: Ensures read traffic distribution and fault tolerance.
- **Reverse Index**: Enables efficient recovery and minimizes downtime.
- **Replication**: Critical for maintaining high availability and quick recovery.


### Lecture Notes

#### 9. Cache
- **Purpose**: To efficiently manage hot Places by introducing a caching layer.
- **Solution**: Utilize an off-the-shelf caching solution like Memcache.
  - **Functionality**:
    - Stores data about hot Places.
    - Allows application servers to check the cache before querying the backend database.
  - **Scalability**:
    - Adjust the number of cache servers based on client usage patterns.
  - **Eviction Policy**:
    - **Least Recently Used (LRU)**: Suitable for the system to manage which data should be removed when the cache reaches its capacity.
    
#### 10. Load Balancing (LB)
- **Purpose**: To distribute incoming requests efficiently and manage server load.
- **Placement**: Load Balancers can be added in two key locations:
  1. **Between Clients and Application Servers**
  2. **Between Application Servers and Backend Servers**
- **Initial Approach**: Round Robin
  - **Advantages**:
    - Simple implementation.
    - Equally distributes incoming requests among backend servers.
    - Automatically removes dead servers from rotation, preventing them from receiving traffic.
  - **Disadvantages**:
    - Does not account for the actual load or performance of the servers.
    - May continue sending requests to overloaded or slow servers.
- **Advanced Approach**:
  - **Intelligent Load Balancing**:
    - Periodically queries backend servers about their current load.
    - Adjusts traffic distribution based on server load information to ensure optimal performance.

### Key Points Summary
- **Cache**: 
  - Use Memcache for hot Places.
  - Check cache before database queries.
  - LRU policy for cache eviction.
- **Load Balancing**:
  - Round Robin for initial implementation.
  - Intelligent LB for load-aware traffic distribution.