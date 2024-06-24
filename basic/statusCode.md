HTTP status codes are essential communication tools between a web server and a client (usually a web browser or a crawler). They indicate the outcome of the HTTP request made by the client. Here are some important HTTP status codes and their meanings:

1. **1xx Informational:**
   - **100 Continue:** The server has received the initial part of the request and the client can continue with the rest of the request or ignore it if it's already finished.
   - **101 Switching Protocols:** The server is changing protocols per the client's request (e.g., HTTP to WebSocket).

2. **2xx Success:**
   - **200 OK:** The request has succeeded. The meaning of the success depends on the HTTP method:
     - GET: The resource has been fetched and is transmitted in the message body.
     - HEAD: The entity headers are in the message body.
     - POST: The resource describing the result of the action is transmitted in the message body.
   - **201 Created:** The request has been fulfilled, resulting in the creation of a new resource.
   - **204 No Content:** The server successfully processed the request and is not returning any content.

3. **3xx Redirection:**
   - **301 Moved Permanently:** The resource requested has been permanently moved to a new URL.
   - **302 Found (or Moved Temporarily):** The resource requested temporarily resides under a different URL.
   - **304 Not Modified:** Indicates that the resource has not been modified since the version specified by the request headers If-Modified-Since or If-None-Match.

4. **4xx Client Error:**
   - **400 Bad Request:** The server cannot process the request due to a client error (e.g., malformed request syntax, invalid request message framing).
   - **401 Unauthorized:** Similar to 403 Forbidden, but specifically for use when authentication is required and has failed or has not yet been provided.
   - **403 Forbidden:** The client does not have permission to access the resource.
   - **404 Not Found:** The server cannot find the requested resource.
   - **409 Too many Request:**  too many requests .

5. **5xx Server Error:**
   - **500 Internal Server Error:** A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.
   - **502 Bad Gateway:** The server was acting as a gateway or proxy and received an invalid response from the upstream server.
   - **503 Service Unavailable:** The server is currently unavailable (overloaded or down).

These status codes help developers and users understand what happened with a request, aiding in troubleshooting and improving the web experience.



### Step 3 - Design Deep Dive

#### Key Questions
1. **Rule Creation and Storage**: How are rate limiting rules created and where are they stored?
2. **Handling Rate-Limited Requests**: How to manage requests that exceed the rate limit?

#### Rate Limiting Rules
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

#### Handling Exceeding Requests
- **Response**: Return HTTP 429 (Too Many Requests) status code.
- **Queuing**: Optionally enqueue rate-limited requests for later processing.

#### Rate Limiter Headers
- **Informing Clients**: Use HTTP response headers to communicate rate limit status.
  - `X-Ratelimit-Remaining`: Number of remaining allowed requests.
  - `X-Ratelimit-Limit`: Maximum number of allowed requests per time window.
  - `X-Ratelimit-Retry-After`: Seconds to wait before making a new request.

#### Detailed Design (Figure 13)
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

#### Summary
- **Rules**: Defined in config files, stored on disk, cached by workers.
- **Request Handling**: Middleware checks limits, returns HTTP 429 if exceeded, uses headers to inform clients.
- **Design Flow**:
  - Client -> Middleware -> Check Redis -> Decision -> API Server or 429 Error