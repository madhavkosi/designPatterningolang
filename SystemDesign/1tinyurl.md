
---

# Table of Contents


   - [Overview](#overview)
   - [Requirements](#requirements)
   - [Capacity Estimation and Constraints](#capacity-estimation-and-constraints)
   - [API Definition](#api-definition)
   - [Database Schema](#database-schema)
   - [Basic System Design and Algorithm for URL Shortening Service](#basic-system-design-and-algorithm-for-url-shortening-service)


---

# URL Shortening System

## Overview

**URL Shortening**

- **Definition:** URL shortening creates shorter aliases for long URLs, known as "short links."
- **Purpose:** Short links redirect users to the original URL, saving space and reducing errors in typing.
- **Example:** Original URL: `https://example.com/very/long/path/to/some/page.html`
  Shortened URL: `https://tinyurl.com/y7fxyz`

**Advantage:** URL shortening is a practical tool for simplifying, tracking, and managing links effectively across various digital platforms.

## Requirements

**Functional Requirements:**

1. **Generate Short Link:** 
   - Given a URL, generate a unique alias (short link) that is easily manageable.
   - Short link should be short enough for practical use in various applications.

2. **Redirect Functionality:** 
   - Ensure that when users access a short link, they are redirected to the original URL.

3. **Custom Short Links:** 
   - Provide an option for users to specify a custom short link for their URL.

4. **Link Expiration:** 
   - Implement a default expiration time for links.
   - Allow users to specify custom expiration times for their links.

**Non-Functional Requirements:**

1. **High Availability:** 
   - Ensure the system is highly available to prevent downtime and ensure continuous URL redirection service.
   
2. **Real-Time Redirection:** 
   - Ensure minimal latency in redirecting users from short links to original URLs.

3. **Security - Non-guessable Short Links:** 
   - Shortened links should not be predictable to prevent unauthorized access.

**Extended Requirements:**

1. **Analytics:** 
   - Track and provide analytics on short link usage, including the number of times a redirection occurs.
   
2. **REST API Access:** 
   - Provide REST APIs for seamless integration with other services, allowing them to generate short links programmatically.

## Capacity Estimation and Constraints

### Traffic Estimates

Our system anticipates the following traffic patterns:

- **URL Shortenings**: Approximately 500 million new URLs per month, translating to about 16.7 million new URLs per day on average.
- **Redirections**: Given a 100:1 read-to-write ratio, expect around 50 billion redirections monthly, or approximately 1.67 billion redirections daily.

### Queries Per Second (QPS)

Based on the corrected traffic estimates:

- **New URLs per Second**: Approximately 200 URLs/s on average (500 million URLs per month / (30 days * 24 hours * 3600 seconds)).
- **Redirections per Second**: Around 20k URLs/s on average (50 billion redirections per month / (30 days * 24 hours * 3600 seconds)).

### Storage Estimates
 
To accommodate our data retention policy of storing each URL shortening request and associated link for 5 years:

- **Total Storage Needed**: Estimated at 15 TB for 30 billion objects, assuming each object averages 500 bytes.

### Bandwidth Estimates

- **Incoming Data**: About 100 KB/s for new URL requests.
- **Outgoing Data**: Roughly 10 MB/s for redirection responses.

### Memory Estimates

- **Memory for Cache**: Approximately 170 GB, based on caching 20% of the 1.67 billion daily redirections.

## API Definition

1. **CreateURL(api_dev_key, original_url, custom_alias=None, user_name=None, expire_date=None)**
   - Creates a shortened URL with optional custom alias and expiration date.
   - Requires `api_dev_key` for authentication.
   - Returns shortened URL or error code.

2. **GetURL(url_key)**
   - Retrieves the original URL associated with `url_key`.
   - Returns original URL.

3. **DeleteURL(api_dev_key, url_key)**
   - Deletes the shortened URL identified by `url_key`.
   - Requires `api_dev_key` for authentication.
   - Returns 'URL Removed' if successful.

### Redirect Types

- **301 Redirect:** Indicates a permanent move of the requested URL to the long URL. Browser caches the redirection, so subsequent requests skip the URL shortening service and go directly to the long URL server. Ideal for reducing server load but limits analytics tracking.

- **302 Redirect:** Signals a temporary move of the URL to the long URL. Each request goes through the URL shortening service, allowing for detailed analytics on click rates and traffic sources. Useful when analytics are crucial but may result in higher server load.

### Abuse Prevention

- **API Key Limits:** Each `api_dev_key` has quotas for URL creations and redirections per time period.
  
- **Time-Based Quotas:** Limits are set per developer key (e.g., daily, weekly) to prevent misuse.

- **Monitoring:** Continuous monitoring detects and responds to unusual activity or abuse.

## Database Schema

### URLMappings Collection (NoSQL Schema)

- **\_id**: ObjectId (automatically generated unique identifier)
- **short_url**: String (indexed, unique) - Shortened URL key
- **original_url**: String - Original URL
- **custom_alias**: String (optional, indexed, unique) - Custom alias chosen for the short URL
- **created_at**: Date - Timestamp of when the short URL was created
- **expires_at**: Date (optional) - Timestamp indicating when the short URL expires
- **user_id**: ObjectId - Reference to the user who created the short URL

### Users Collection (NoSQL Schema)

- **\_id**: ObjectId (automatically generated unique identifier)
- **user_id**: String (indexed, unique) - User ID
- **username**: String - Username of the user who created the short URL
- **email**: String (optional) - Email address of the user (if applicable)
- **created_at**: Date - Timestamp of when the user account was created

### Analytics Collection (NoSQL Schema)

- **\_id**: ObjectId (automatically generated unique identifier)
- **short_url**: String - Shortened URL key (indexed)
- **timestamp**: Date - Timestamp of when the redirection or click occurred
- **referrer**: String - Referrer URL or source of the click
- **user_agent**: String - User agent information of the requester
- **ip_address**: String - IP address of the requester
- **country**: String - Country of the requester (derived from IP address geolocation)
- **device_type**: String - Type of device used (desktop, mobile, tablet, etc.)
### Database Choice: 
- **NoSQL Database:** Due to the need to handle billions of small records without relational dependencies and the service's read-heavy nature.
- **Scalability:** NoSQL databases like MongoDB or Cassandra are preferred for their ability to scale horizontally.

### Considerations:
- **Data Partitioning:** Distribute data across multiple nodes to manage large volumes efficiently.
- **Performance:** Optimize for read operations given the service's read-heavy workload.


## Basic System Design and Algorithm for URL Shortening Service


### Solution 1: Encoding Actual URL
- **Generate a Unique Hash**: Compute a hash (e.g., MD5, SHA256) for the URL.
- **Encode the Hash**: Convert the hash to a displayable format using base36 ([a-z, 0-9]) or base62 ([A-Z, a-z, 0-9]) encoding. Optionally, base64 encoding can be used.
- **Short Key Length**: 
  - A 6-character key in base64 offers 64^6 (~68.7 billion) possible combinations.
  - An 8-character key in base64 offers 64^8 (~281 trillion) combinations.
- **Hash Truncation**: For example, MD5 generates a 128-bit hash. After base64 encoding, it exceeds the desired key length, so truncate to the first 6 or 8 characters.
- **Duplication and URL-encoding Issues**:
  1. Identical URLs produce the same shortened URL.
  2. URL-encoded variations of the same URL need handling.
- **Workarounds**:
  - Append an increasing sequence number to ensure uniqueness.
  - Append user ID for uniqueness, especially for non-logged-in users.

### Solution 2: Generating Keys Offline
- **Key Generation Service (KGS)**: 
  - Pre-generate random six-letter strings and store them in a key database (key-DB).
  - KGS ensures all keys are unique.
- **Concurrency Handling**:
  - Keys are marked as used in the database immediately upon assignment to avoid duplication.
  - Use two tables: one for unused keys and one for used keys.
- **Efficiency**:
  - KGS keeps some keys in memory for quick access.
  - In case of KGS failure, a replica can take over.
  - Application servers can cache keys to improve performance, with an acceptable risk of losing some keys if a server dies.
- **Key-DB Size**:
  - With base64, 68.7 billion unique keys require approximately 412 GB of storage (6 characters per key).
- **Key Lookup**:
  - Perform a lookup in the database to retrieve the full URL.
  - If the key exists, issue an "HTTP 302 Redirect" to the original URL.
  - If the key is not found, return an "HTTP 404 Not Found" or redirect to the homepage.
- **Custom Aliases**:
  - Support user-defined custom aliases with a reasonable length limit (e.g., 16 characters).

### Data Partitioning and Replication

To scale out a database for storing billions of URLs, an effective partitioning scheme is essential. Here are two primary approaches to partitioning data:

#### a. Range-Based Partitioning
- **Concept**: Data is divided into partitions based on the range of values. For example, URLs are stored in different partitions based on the first letter of their hash key.
  - URLs starting with 'A' are stored in one partition, 'B' in another, and so on.
  - Less frequently occurring letters can be combined into a single partition.
- **Advantages**: 
  - Predictable storage and retrieval, as the partition location is determined by the value range.
- **Disadvantages**: 
  - Unbalanced Load: Certain partitions can become overloaded if the data distribution is uneven (e.g., many URLs starting with 'E').
  - Static Scheme: Requires careful planning and might need rebalancing as data distribution changes.

#### b. Hash-Based Partitioning
- **Concept**: Data is partitioned based on the hash value of the key (e.g., URL). The hash function maps each key to a number that corresponds to a partition.
  - Example: Hash function maps keys to a number between 1 and 256, with each number representing a specific partition.
- **Advantages**:
  - Random Distribution: Helps in evenly distributing the data across partitions.
  - Scalability: New partitions can be added easily by extending the hash range.
- **Disadvantages**:
  - Overloaded Partitions: Some partitions might still become overloaded due to uneven hash value distribution.
  - Complexity: Requires a good hash function to ensure even distribution and might need rehashing if partitions are added or removed.

#### Consistent Hashing
- **Concept**: An advanced form of hash-based partitioning designed to solve the issues of traditional hashing schemes.
  - Distributes both data and servers on a ring of hash values.
  - Data is assigned to the nearest server in the clockwise direction.
- **Advantages**:
  - Minimal Disruption: When servers are added or removed, only a small fraction of data needs to be moved.
  - Even Load Distribution: Virtual nodes can be used to ensure balanced partitions.
  - Flexibility: Easily handles the dynamic nature of distributed systems.

## Replication
- **Purpose**: Ensures data availability and reliability by storing copies of data on multiple servers.
- **Approach**:
  - **Master-Slave Replication**: One server (master) handles writes and propagates changes to other servers (slaves) for read operations.
  - **Peer-to-Peer Replication**: All servers are equal and can handle both read and write operations, propagating changes to each other.
- **Benefits**:
  - High Availability: Data remains accessible even if some servers fail.
  - Load Balancing: Read operations can be distributed across multiple replicas.
  - Fault Tolerance: Data loss risk is minimized through redundancy.


## Cache

Implementing caching for frequently accessed URLs can significantly enhance the performance and efficiency of a URL shortening service. Here's an overview:

### Caching Mechanism
- **Purpose**: Reduce latency and backend load by storing frequently accessed URLs and their respective hashes in memory.
- **Implementation**: Use an off-the-shelf caching solution (e.g., Redis, Memcached) to store full URLs and their hashes.
  - Application servers first check the cache before querying the backend storage.
  
### Cache Size
- **Estimation**: Start with caching 20% of daily traffic.
  - Example Calculation: For 20% of daily traffic, 170GB of cache memory is needed.
  - Modern servers typically have 256GB of memory, allowing all cache to fit into a single machine.
  - Alternatively, use multiple smaller servers to store hot URLs.

### Cache Eviction Policy
- **Policy Choice**: Least Recently Used (LRU) policy is suitable for this system.
  - LRU discards the least recently used URL first when the cache is full.

### Cache Replication

- **Purpose**: Distribute load and enhance reliability.
- **Mechanism**: Replicate caching servers to balance the load.
  - Upon a cache miss, servers query the backend database, update the cache, and propagate the new entry to all cache replicas.
  - Each replica updates its cache by adding the new entry or ignoring it if it already exists.

## Load Balancer (LB)

**Purpose**: Distribute incoming requests to ensure efficient resource use and high availability.

### Load Balancer Placement

1. **Between Clients and Application Servers**:
   - Balances client requests across application servers.
   - Prevents any single server from becoming a bottleneck.

2. **Between Application Servers and Database Servers**:
   - Distributes database queries among database servers.
   - Enhances database performance and availability.

3. **Between Application Servers and Cache Servers**:
   - Balances cache requests across cache servers.
   - Ensures faster data retrieval and prevents overloads.

### Load Balancing Strategies

1. **Round Robin Load Balancing**:
   - **Method**: Cyclically distributes requests among servers.
   - **Pros**: Simple, no overhead, handles server failures.
   - **Cons**: Ignores server load, may overload slow servers.

2. **Intelligent Load Balancing**:
   - **Method**: Adjusts traffic based on server load.
   - **Pros**: Efficient resource use, prevents overload.
   - **Cons**: Complex, requires continuous monitoring.

## Purging or DB Cleanup

**Expiration and Purging**:
- **Expiration Policy**: Links have a user-specified or default expiration time (e.g., two years).
- **Lazy Cleanup**: Remove expired links slowly to avoid database pressure.
- **On Access**: Delete the link and return an error if a user accesses an expired link.

**Cleanup Service**:
- **Periodic Execution**: Runs periodically during low traffic times to remove expired links from storage and cache.
- **Lightweight Operation**: Designed to minimize impact on system performance.

**Key Reuse**:
- After removing an expired link, the key can be returned to the key-DB for reuse.

**Considerations for Inactive Links**:
- **Storage Cost**: Given the decreasing cost of storage, links might be kept indefinitely.
- **Policy Decision**: Evaluate whether to remove links not visited in a long time (e.g., six months).

## Telemetry

**Purpose**: Collect and store usage statistics for short URLs to gain insights into user behavior and system performance.

### Key Statistics to Track

- **Usage Count**: Number of times a short URL has been accessed.
- **User Location**: Country of the visitor.
- **Access Details**: 
  - Date and time of access.
  - Referring web page.
  - Browser and platform used for access.

### Storage and Handling

- **Database Design**:
  - Store statistics in a separate database or table to avoid performance issues on the main URL table.
  - Use batch processing to update statistics periodically instead of real-time updates.
  
- **Concurrency Handling**:
  - **Queue System**: Implement a queue to handle high concurrency, where access logs are written to a queue and processed in batches.
  - **Asynchronous Updates**: Use asynchronous processes to update the statistics, reducing the load on the main database during peak traffic.

### Implementation Strategies

- **Event Logging**: Log each access event with the relevant details (e.g., country, timestamp, referrer, browser) to a separate logging system or database.
- **Analytics Tools**: Integrate with analytics tools (e.g., Google Analytics) to track and visualize usage patterns.
- **Data Aggregation**: Periodically aggregate data to generate summary statistics, reducing the volume of data stored and processed.

