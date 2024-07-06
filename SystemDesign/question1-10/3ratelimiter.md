### Rate Limiter

**Definition**
A rate limiter controls the rate of traffic sent by a client or a service, limiting the number of requests allowed over a specified period. Excess requests are blocked once the threshold is exceeded.

 **Examples**
- Users can write no more than 2 posts per second.
- Create a maximum of 10 accounts per day from the same IP address.
- Claim rewards no more than 5 times per week from the same device.

 **Benefits of Using an API Rate Limiter**

1. **Prevent Resource Starvation**
   - Stops Denial of Service (DoS) attacks by blocking excessive requests.
   - Example: Twitter limits the number of tweets to 300 per 3 hours; Google Docs APIs limit to 300 per user per 60 seconds.

2. **Reduce Cost**
   - Limits excess requests, reducing server load and resource usage.
   - Essential for paid third-party APIs to avoid high costs (e.g., checking credit, making payments, retrieving health records).

3. **Prevent Overloading**
   - Filters out excess requests from bots or misbehaving users.
   - Ensures servers are not overwhelmed by high traffic volumes.

 **Summary**
A rate limiter is crucial for maintaining system stability, preventing abuse, and managing costs effectively. It ensures that API usage remains within safe limits, protecting the system from excessive load and potential attacks.


### Step 1 - Understand the Problem and Establish Design Scope

**Questions and Answers**
- **Rate Limiter Type**:
  - **Focus**: Server-side API rate limiter.
- **Throttle Rules**:
  - **Flexibility**: Support different sets of throttle rules.
- **System Scale**:
  - **Requirement**: Handle a large number of requests.
- **Environment**:
  - **Setup**: Work in a distributed environment.
- **Implementation**:
  - **Decision**: Separate service or within application code.
- **User Notification**:
  - **Requirement**: Inform users when throttled.

**Requirements Summary**
- **Accurate** request limiting.
- **Low latency** to avoid slowing down HTTP response.
- **Memory efficient** usage.
- **Distributed rate limiting** across servers/processes.
- **Exception handling** for clear user notifications when throttled.
- **High fault tolerance** to avoid system-wide issues.



### Step 2 - Propose High-Level Design and Get Buy-In

**Implementation Location**
- **Client-side**: Unreliable due to potential forgery and lack of control.
- **Server-side**: More secure and controllable.

**Middleware Approach**
- **Rate Limiter Middleware**: Throttles requests before reaching API servers.
- **Example**: API allows 2 requests per second; a third request gets a 429 status code (too many requests).

**API Gateway**
- **Function**: A middleware component supporting rate limiting, SSL termination, authentication, etc.
- **Recommendation**: Use API gateway for rate limiting in microservice architecture.


### Algorithms for Rate Limiting

Rate limiting can be implemented using various algorithms, each with its own pros and cons. Understanding these at a high level helps in choosing the right one for specific use cases. Popular algorithms include:

- **Token Bucket**
- **Leaking Bucket**
- **Fixed Window Counter**
- **Sliding Window Log**
- **Sliding Window Counter**


**Token Bucket Algorithm - Short Notes**
**Overview**
- **Widely Used**: Commonly implemented by companies like Amazon and Stripe.
- **Simple and Well-Understood**: Known for its straightforward implementation.

 **How It Works**
- **Bucket Capacity**: The bucket has a predefined capacity.
- **Token Refill**: Tokens are added at preset rates periodically.
- **Overflow**: Once full, extra tokens overflow and are discarded.
- **Token Consumption**: Each request consumes one token.
  - If tokens are available, the request proceeds.
  - If tokens are insufficient, the request is dropped.

 **Parameters**
- **Bucket Size**: Maximum number of tokens the bucket can hold.
- **Refill Rate**: Number of tokens added to the bucket every second.

 **Use Cases**
- **Multiple Buckets**: Different buckets for different API endpoints or rate-limiting rules.
  - Example: Separate buckets for posting, adding friends, and liking posts.
- **IP-Based Throttling**: Each IP address has its own bucket.
- **Global Bucket**: For systems allowing a high volume of requests, a shared bucket for all requests.

**Pros**
- **Easy to Implement**: Straightforward setup.
- **Memory Efficient**: Minimal memory usage.
- **Allows Burst Traffic**: Can handle short bursts of traffic as long as tokens are available.

 **Cons**
- **Parameter Tuning**: Requires careful adjustment of bucket size and refill rate to meet needs accurately.


**Leaking Bucket Algorithm - Short Notes**
**Overview**
- **Mechanism**: Similar to the token bucket but processes requests at a fixed rate using a FIFO queue.
- **Function**: Adds requests to the queue if not full; otherwise, drops them. Processes requests at regular intervals.

**Parameters**
- **Bucket Size**: Equal to the queue size, determines how many requests can be held.
- **Outflow Rate**: Number of requests processed at a fixed rate (e.g., per second).

**Pros**
- Memory efficient due to the limited queue size.
- Suitable for use cases requiring a stable outflow rate.

**Cons**
- Traffic bursts can fill the queue with old requests, delaying recent ones.
- Tuning the parameters (bucket size and outflow rate) can be challenging.


**Fixed Window Counter Algorithm - Short Notes**
**Overview**
- **Mechanism**: Divides the timeline into fixed-sized windows, assigning a counter for each window.
- **Function**: Increments the counter with each request. Drops new requests once the counter reaches the threshold until the next window starts.

**Example**
- **Scenario**: Allows a maximum of 3 requests per second.
- **Issue**: Traffic spikes at window edges can cause more requests than the allowed quota to pass through.

**Pros**
- Memory efficient.
- Easy to understand.
- Suitable for use cases where quota resets at the end of fixed time windows.

**Cons**
- Traffic spikes at window edges can exceed the allowed request quota.


**Sliding Window Log Algorithm - Short Notes**
**Overview**
- **Mechanism**: Keeps track of request timestamps, usually in a cache like Redis.
- **Function**:
  - **New Request**: Remove outdated timestamps (older than the current time window).
  - Add the new request's timestamp to the log.
  - If the log size is within the allowed limit, the request is accepted; otherwise, it is rejected.

**Example**
- **Scenario**: Allows 2 requests per minute.
- **Operations**:
  - **1:00:01**: Log is empty, request allowed, timestamp added. **Log Update: [1:00:01].**
  - **1:00:30**: Log size is 2 after adding, request allowed.**Log Update: [1:00:01, 1:00:30].**
  - **1:00:50**: Log size exceeds limit (3), request rejected.**Log Update:[1:00:01, 1:00:30, 1:00:50].**
  - **1:01:40**: Outdated timestamps removed, log size is 2, request allowed.**Log Update: [1:00:50, 1:01:40].**

**Pros**
- Accurate rate limiting within any rolling window.
- Prevents excess requests at window edges.

**Cons**
- High memory usage as timestamps are stored even for rejected requests.


**Sliding Window Counter Algorithm - Notes**
**Overview**
- **Hybrid Approach**: Combines the fixed window counter and sliding window log to provide a more accurate rate limiting mechanism.

**How It Works**
- **Calculation**: Uses the request counts from the current and previous windows to compute the effective request count for the rolling window.
- **Formula**: 
  \[ \text{Effective Request Count} = \text{Requests in Current Window} + (\text{Requests in Previous Window} \times \text{Overlap Percentage}) \]
- **Example**:
  - **Allowed Requests**: 7 per minute.
  - **Previous Minute Requests**: 5.
  - **Current Minute Requests**: 3.
  - **Overlap Percentage**: 30% into the current minute.
  - **Calculation**: 
    \[ 3 + (5 \times 0.7) = 6.5 \]
    - Rounded down to 6.
  - **Decision**: Current request allowed since 6 < 7.

**Pros**
- **Smooths Traffic Spikes**: Averages request load over recent windows, reducing the impact of traffic bursts.
- **Memory Efficient**: Requires less memory compared to logging every request.

**Cons**
- **Approximation**: Assumes an even distribution of requests within the windows, which may not always be accurate.
- **Limited Lookback**: Effective primarily for not-so-strict lookback windows.

**Conclusion**
- The sliding window counter algorithm provides a balanced and memory-efficient rate limiting solution by interpolating between fixed windows, offering a more accurate approximation of the request rate while smoothing out traffic spikes.


### High-Level Architecture for Rate Limiting

**Basic Concept**
- **Counter**: Tracks the number of requests sent by a user, IP address, etc. If the counter exceeds the limit, the request is disallowed.

**Counter Storage**
- **In-Memory Cache**: Preferred over a database due to faster access and support for time-based expiration. Redis is a popular choice.
- **Redis Commands**:
  - **INCR**: Increments the stored counter by 1.
  - **EXPIRE**: Sets a timeout for the counter, automatically deleting it upon expiration.

**Architecture Workflow**

1. **Client Request**: Sent to rate limiting middleware.
2. **Counter Check**: Middleware fetches counter from Redis.
3. **Decision**:
   - **Limit Reached**: Request rejected.
   - **Limit Not Reached**: Request forwarded to API servers, counter incremented in Redis.


**Summary**
The high-level architecture for rate limiting involves using an in-memory cache like Redis to store request counters. The rate limiting middleware checks these counters before allowing requests to proceed to the API servers, ensuring efficient and effective rate limiting.

### Step 3 - Design Deep Dive

**Key Questions**
1. **Rule Creation and Storage**: How are rate limiting rules created and where are they stored?
2. **Handling Rate-Limited Requests**: How to manage requests that exceed the rate limit?

**Rate Limiting Rules**
- **Examples**:
  - **Messaging Domain**:
    ```yaml
    domain: messaging
    descriptors:
      - key: message_type
        value: marketing
        rate_limit:
          unit: day
          requests_per_unit: 5
    ```
  - **Auth Domain**:
    ```yaml
    domain: auth
    descriptors:
      - key: auth_type
        value: login
        rate_limit:
          unit: minute
          requests_per_unit: 5
    ```
- **Storage**: Rules are written in configuration files and saved on disk.

**Handling Exceeding Requests**
- **Response**: Return HTTP 429 (Too Many Requests) status code.
- **Queuing**: Optionally enqueue rate-limited requests for later processing.

**Rate Limiter Headers**
- **Informing Clients**: Use HTTP response headers to communicate rate limit status.
  - `X-Ratelimit-Remaining`: Number of remaining allowed requests.
  - `X-Ratelimit-Limit`: Maximum number of allowed requests per time window.
  - `X-Ratelimit-Retry-After`: Seconds to wait before making a new request.

**Detailed Design (Figure 13)**
1. **Rule Storage and Loading**:
   - Rules are stored on disk.
   - Workers pull and cache rules from disk.
2. **Request Processing**:
   - Client sends a request to the server.
   - Request is sent to the rate limiter middleware.
   - Middleware loads rules from cache, checks counters and timestamps in Redis.
   - **Decision**:
     - **Not Rate Limited**: Request forwarded to API servers.
     - **Rate Limited**: Return HTTP 429 error, optionally enqueue the request.

**Summary**
- **Rules**: Defined in config files, stored on disk, cached by workers.
- **Request Handling**: Middleware checks limits, returns HTTP 429 if exceeded, uses headers to inform clients.
- **Design Flow**:
  - Client -> Middleware -> Check Redis -> Decision -> API Server or 429 Error

  
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/rateLimiter.png)