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