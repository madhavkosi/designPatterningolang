## Network Essentials
### HTTP vs. HTTPS

**HTTP (Hypertext Transfer Protocol)**
- **Characteristics**:
  - **No Encryption**: Data is transmitted in plain text, can be intercepted and read.
  - **Default Port**: Operates over port 80.
  - **Vulnerabilities**: Susceptible to man-in-the-middle attacks and eavesdropping.
- **Use Cases**:
  - Browsing simple websites with non-sensitive data, like informational blogs.
  - Historically standard for all web communications, but declining due to security concerns.

**HTTPS (Hypertext Transfer Protocol Secure)**
- **Characteristics**:
  - **Encryption**: Data is encrypted using SSL (Secure Sockets Layer) or TLS (Transport Layer Security), preventing easy interception and reading by attackers.
  - **Default Port**: Operates over port 443.
  - **Security**: Provides authentication of the accessed website and ensures the privacy and integrity of the data exchanged.
  - **Trustworthiness**: More trusted by users; browsers often mark HTTP sites as 'Not Secure'.
- **Use Cases**:
  - Ideal for transactions involving personal, financial, or sensitive data.
  - Recommended for all types of websites to ensure secure communication.

**Key Differences**
- **Security**: 
  - **HTTP**: No encryption, not secure.
  - **HTTPS**: Encrypted, secure.
- **Performance**: 
  - **HTTP**: No encryption overhead.
  - **HTTPS**: Slightly more server load due to encryption/decryption, minimized by modern hardware/software.
- **SEO Ranking**: 
  - **HTTP**: Less preferred by search engines.
  - **HTTPS**: Preferred by search engines, improves ranking.
- **Certificate Requirement**: 
  - **HTTP**: No certificate needed.
  - **HTTPS**: Requires an SSL/TLS certificate from a Certificate Authority (CA).

**Summary**
- **HTTP**: Suitable for non-sensitive, informational content.
- **HTTPS**: Essential for secure, private communication and trusted interactions, recommended for all websites. 

The widespread adoption of HTTPS is aimed at creating a more secure and trustworthy internet.



### TCP (Transmission Control Protocol)
- **Definition**: A connection-oriented protocol that ensures reliable, ordered, and error-checked delivery of a stream of bytes between applications.
- **Characteristics**:
  - **Reliability**: Ensures accurate and in-order delivery of data. Retransmits lost or corrupted packets.
  - **Connection-Oriented**: Establishes a connection between sender and receiver before transmitting data.
  - **Flow Control**: Manages the rate of data transmission to prevent network congestion.
  - **Congestion Control**: Adjusts the transmission rate based on network traffic conditions.
  - **Acknowledgements and Retransmissions**: Uses acknowledgments to confirm receipt of data and retransmits if necessary.
- **Use Cases**:
  - Applications where reliability and order are critical.
  - Examples: Web browsing (HTTP/HTTPS), email (SMTP, POP3, IMAP), file transfers (FTP).

### UDP (User Datagram Protocol)
- **Definition**: A connectionless protocol that sends messages, called datagrams, without establishing a prior connection and without guaranteeing reliability or order.
- **Characteristics**:
  - **Low Overhead**: Does not establish a connection, leading to lower overhead and latency.
  - **Unreliable Delivery**: Does not guarantee message delivery, order, or error checking.
  - **Speed**: Faster than TCP due to its simplicity and lack of retransmission mechanisms.
  - **No Congestion Control**: Does not reduce transmission rates under network congestion.
- **Use Cases**:
  - Applications that require speed and can tolerate some loss of data.
  - Examples: Streaming video or audio, online gaming, VoIP (Voice over Internet Protocol).

### Key Differences
- **Reliability**:
  - **TCP**: Reliable, ordered, and error-checked delivery.
  - **UDP**: Unreliable, unordered delivery without error checking.
- **Connection**:
  - **TCP**: Connection-oriented, requiring a connection before data transfer.
  - **UDP**: Connectionless, sending data without a prior connection.
- **Speed and Overhead**:
  - **TCP**: Slower with higher overhead due to handshaking, acknowledgements, and congestion control.
  - **UDP**: Faster with minimal overhead, suitable for real-time applications.
- **Data Integrity**:
  - **TCP**: High data integrity, essential for applications like file transfers and web browsing.
  - **UDP**: Lower data integrity, acceptable for applications like streaming where perfect accuracy is less critical.
- **Use Case Suitability**:
  - **TCP**: Used when data accuracy and reliability are more critical than speed.
  - **UDP**: Used when speed is more critical than reliability, and some data loss is acceptable.

### Summary
- **TCP**: Best for applications needing reliable and accurate data transmission (e.g., web browsing, email, file transfers).
- **UDP**: Ideal for applications needing speed where some data loss is tolerable (e.g., streaming, online gaming, VoIP).



### URL vs. URI vs. URN

**URL (Uniform Resource Locator)**
- **Definition**: A specific type of URI that identifies a resource and provides a method to locate it.
- **Components**: Protocol (HTTP, HTTPS, FTP), domain name, path, optional query parameters, and fragment identifier.
- **Example**: `https://www.example.com/path?query=term#section`
- **Key Characteristics**:
  - Specifies how to access the resource (protocol).
  - Includes the location of the resource (web address).

**URI (Uniform Resource Identifier)**
- **Definition**: A generic term for identifying a resource by location, name, or both.
- **Scope**: Encompasses both URLs and URNs.
- **Example**: `https://www.example.com` (URL) and `urn:isbn:0451450523` (URN) are both URIs.
- **Key Characteristics**:
  - General concept that can be a locator (URL), a name (URN), or both.

**URN (Uniform Resource Name)**
- **Definition**: A type of URI that names a resource without specifying how to locate it.
- **Example**: `urn:isbn:0451450523` (identifies a book by ISBN).
- **Key Characteristics**:
  - Provides a unique and persistent identifier.
  - Does not specify location or access method.

### Summary of Differences
- **URL**: Specifies both identity and location of a resource (How and Where).
- **URI**: A broad term covering URLs (identifying and locating) and URNs (just identifying).
- **URN**: Focuses on uniquely identifying a resource without specifying location or access method.

### Practical Usage
- **URL**: Mostly used for web browsing.
- **URI and URN**: Used in contexts like software development, digital libraries, and systems requiring unique and persistent identification. 