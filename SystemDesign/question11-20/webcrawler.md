### Designing a Web Crawler

**1. What is a Web Crawler?**
- **Definition**: A software program that browses the World Wide Web methodically and automatically.
- **Functions**:
  - Collect documents by recursively fetching links.
  - Create search engine indexes.
  - Test web pages and links for valid syntax and structure.
  - Monitor site changes.
  - Maintain mirror sites.
  - Search for copyright infringements.
  - Build special-purpose indexes (e.g., multimedia content).

**2. Requirements and Goals of the System**
- **Scalability**: 
  - Crawl the entire Web.
  - Fetch hundreds of millions of web documents.
- **Extensibility**:
  - Modular design for adding new functionality.
  - Support for future document types.
  
  **3. Some Design Considerations**
- **Content Type**:
  - Initially focus on HTML pages.
  - Extensible to support other media types (images, videos, etc.).
  - Separate parsing modules for different media types.
- **Protocols**:
  - Initial focus on HTTP.
  - Extensible to support other protocols (FTP, etc.).
- **Scale**:
  - Target to crawl one billion websites.
  - Assume an upper bound of 15 billion different web pages.
- **Robots Exclusion Protocol**:
  - Implement the Robots Exclusion Protocol.
  - Fetch and respect `robots.txt` before downloading any real content.


### Capacity Estimation and Constraints

**Crawling Rate**:
- **Target**: Crawl 15 billion pages in 4 weeks.
- **Required Rate**: 

```
15 billion pages / (4 weeks * 7 days/week * 86400 seconds/day) ≈ 6200 pages/second
```


Certainly! Here's the content with both LaTeX and plain text formats:

---

**Storage Requirements**:
- **Average Page Size**: 100KB (HTML text).
- **Metadata Size**: 500 bytes per page.
- **Total Storage**:

```
Total Storage: 
15 billion pages * (100 KB + 500 bytes) ≈ 1.5 petabytes
```

---

**Capacity Model**:
- **70% Capacity Utilization**: Avoid exceeding 70% of total storage capacity.
- **Total Storage Needed**:
 

```
Total Storage Needed: 
1.5 petabytes / 0.7 ≈ 2.14 petabytes
```

---

By including both LaTeX and plain text formats, you ensure clarity for different viewing environments.
  ### High Level Design for Web Crawler

**Basic Algorithm**:
1. **Seed URLs**: Start with a list of seed URLs.
2. **Steps**:
   - **Pick URL**: Select a URL from the unvisited URL list.
   - **Resolve IP**: Determine the IP address of its host-name.
   - **Establish Connection**: Connect to the host to download the document.
   - **Parse Document**: Extract new URLs from the document.
   - **Add URLs**: Add new URLs to the unvisited URL list.
   - **Process Document**: Store or index the downloaded document.
   - **Repeat**: Go back to step 1.

**Crawling Strategy**:
- **Breadth-First Search (BFS)**:
  - Commonly used.
  - Ensures broad coverage of the web.
- **Depth-First Search (DFS)**:
  - Used in specific situations to save connection overhead.
  - Efficient for crawling all URLs within a single website after establishing a connection.

**Path-Ascending Crawling**:
- **Purpose**: Discover isolated resources without inbound links.
- **Method**:
  - Ascend to every path in the URL.
  - Example: For `http://foo.com/a/b/page.html`, attempt to crawl:
    - `/a/b/`
    - `/a/`
    - `/`
- **Benefits**: Helps find resources that regular crawling might miss.

### Summary
A web crawler should start with a list of seed URLs and use BFS to explore the web broadly. In some cases, DFS can be used to optimize connections. Path-ascending crawling is an additional strategy to find isolated resources by ascending paths in the URL.


### Difficulties in Implementing an Efficient Web Crawler

**Key Challenges**:

1. **Large Volume of Web Pages**:
   - **Issue**: The web is vast, and crawlers can only download a fraction at a time.
   - **Solution**: Prioritize downloads intelligently to ensure the most relevant or important pages are fetched first.

2. **Rate of Change on Web Pages**:
   - **Issue**: Web pages change frequently, meaning a page might change or a new page might be added while the crawler is still processing the site.
   - **Solution**: Implement strategies for frequent revisits and updates to ensure the freshness of the indexed content.

**Essential Components of a Basic Web Crawler**:

1. **URL Frontier**:
   - **Function**: Store and manage the list of URLs to be downloaded.
   - **Prioritization**: Determine the order in which URLs should be crawled based on factors like relevance, importance, and freshness.

2. **HTML Fetcher**:
   - **Function**: Retrieve web pages from servers using HTTP requests.

3. **Extractor**:
   - **Function**: Extract links from HTML documents to discover new URLs for crawling.

4. **Duplicate Eliminator**:
   - **Function**: Ensure that the same content is not downloaded and processed multiple times.
   - **Method**: Use hashing or checksums to detect duplicate content.

5. **Datastore**:
   - **Function**: Store retrieved pages, URLs, and metadata for indexing and further processing.


<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/crawler2.gif" width="800" />
</p>


### Detailed Component Design
<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/crawler3.gif" width="800" />
</p>

#### Distributed URL Frontier

**Overview**
- **URL Frontier:** The data structure holding URLs yet to be downloaded.
- **Breadth-First Traversal:** Initial pages (seed set) are traversed using a FIFO (First-In, First-Out) queue.

**Distributed URL Frontier**
- **Distribution Across Servers:**
  - URLs are distributed to multiple servers.
  - Each server has multiple worker threads for crawling tasks.
  - URLs are assigned to servers via a hash function mapping each URL to a specific server.

**Politeness Requirements**
- **Preventing Server Overload:**
  - Avoid downloading many pages from a single server rapidly.
  - Ensure no multiple machines connect to the same web server concurrently.

**Implementation of Politeness Constraints**
- **FIFO Sub-Queues:**
  - Each server hosts multiple distinct FIFO sub-queues.
  - Each worker thread operates from its own sub-queue.
- **URL Placement in Sub-Queues:**
  - New URLs are placed in sub-queues based on their canonical hostname.
  - The hash function maps each hostname to a specific thread number.
  - Ensures:
    - Only one worker thread downloads from a specific web server at a time.
    - FIFO queue usage prevents overloading any single web server.

**Key Points**
- **URL Frontier:** Central to managing URLs yet to be crawled.
- **Breadth-First Traversal:** Simplified by FIFO queues.
- **Distribution Strategy:**
  - URLs are hashed and distributed to servers and threads.
  - Ensures balanced load and adherence to politeness policies.
- **Politeness Constraints:**
  - Critical to avoid overloading servers.
  - Managed by structured sub-queues and thread assignments.


**Size and Storage of URL Frontier**
- **URL Frontier Size:**
  - Estimated to be in the hundreds of millions of URLs.
- **Storage Solution:**
  - URLs must be stored on disk due to their large volume.

**Implementation of Queues**
- **Separate Buffers for Enqueueing and Dequeuing:**
  - **Enqueue Buffer:**
    - Collects URLs to be added to the frontier.
    - Once filled, dumps the URLs to the disk.
  - **Dequeue Buffer:**
    - Maintains a cache of URLs that need to be visited next.
    - Periodically reads from the disk to replenish the buffer.

**Key Points**
- **Disk Storage Necessity:**
  - Essential for handling the vast number of URLs.
- **Buffer System:**
  - Ensures efficient management and retrieval of URLs.
  - Separate buffers optimize enqueue and dequeue operations.


#### Fetcher Module 

- **Purpose**: Downloads documents from given URLs using network protocols (e.g., HTTP).
- **Robot.txt Compliance**: Webmasters use robot.txt to restrict crawler access to specific site parts.
- **Caching**: To prevent downloading robot.txt with each request, maintain a fixed-sized cache.
- **Cache Function**: Maps host-names to their robot's exclusion rules.


#### Document Input Stream (DIS) 

- **Purpose**: Allows the same document to be processed by multiple modules without repeated downloads.
- **Caching**: Caches documents locally to avoid multiple downloads.
- **Functionality**: 
  - Provides methods to re-read documents.
  - Caches small documents (≤ 64 KB) in memory.
  - Temporarily writes larger documents to a backing file.
- **Worker Threads**: 
  - Each worker thread has an associated DIS.
  - DIS is reused for different documents.
- **Process Flow**:
  - Worker extracts a URL from the frontier.
  - URL is passed to the relevant protocol module.
  - Protocol module initializes the DIS from the network connection with the document’s contents.
  - DIS is then passed to all relevant processing modules.


#### Document Dedupe Test Notes

- **Purpose**: Prevents processing the same document multiple times by removing duplicates.
- **Issue**: Web documents often available under different URLs or mirrored on various servers, causing redundant downloads.
- **Solution**:
  - Calculate a 64-bit checksum (using MD5 or SHA) for every processed document.
  - Store checksums in a database.
  - For each new document, compare its checksum against stored checksums to check for duplicates.
- **Checksum Store Size**:
  - 15 billion distinct web pages.
  - 15B * 8 bytes = 120 GB.
  - Fits into modern server memory, but may require an alternative approach if memory is insufficient.
- **Alternative Approach**:
  - Use a smaller LRU-based cache on each server.
  - Backed by persistent storage.
  - Dedupe test first checks if the checksum is in the cache.
  - If not, checks persistent storage.
  - If found, the document is ignored.
  - If not found, the checksum is added to both the cache and persistent storage.

#### URL Filters Notes

- **Purpose**: Provides control over which URLs are downloaded.
- **Function**: 
  - Blocks websites to be ignored by the crawler.
  - Worker thread consults user-supplied URL filter before adding URLs to the frontier.
  - Filters can restrict URLs by domain, prefix, or protocol type.

#### Domain Name Resolution Notes

- **Purpose**: Maps Web server hostnames to IP addresses using DNS.
- **Issue**: DNS name resolution can be a bottleneck due to the high number of URLs.
- **Solution**:
  - Cache DNS results by building a local DNS server.
  - Reduces repeated DNS requests.

#### URL Dedupe Test Notes

- **Purpose**: Prevents downloading and processing the same document multiple times.
- **Process**:
  - Store all seen URLs in canonical form using fixed-sized checksums in a database.
  - In-memory cache of popular URLs on each host to reduce database operations.
  - High in-memory hit rate due to common URLs.
- **Storage Requirement**:
  - 15 billion distinct URLs.
  - 15B * 8 bytes = 120 GB.
- **Bloom Filters**:
  - Probabilistic data structure for set membership testing.
  - Reduces space but may yield false positives.
  - False positives cause URLs to be ignored, preventing document download.
  - Larger bit vectors can reduce false positive rates.

#### Checkpointing Notes

- **Purpose**: Guards against failures during long web crawls.
- **Function**:
  - Regular snapshots of crawler state written to disk.
  - Allows interrupted or aborted crawls to restart from the latest checkpoint.

