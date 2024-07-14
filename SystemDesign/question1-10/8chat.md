## Designing a Chat System

### Key Points

- **Diverse Functions**: A chat app serves different purposes for different users.
- **Requirement Clarity**: Crucial to define exact requirements to ensure the design meets user needs.
- **Focus Area**: Determine whether the system should focus on group chat, one-on-one chat, or both.
- **Feature Exploration**: Identify and explore the necessary features required by the target users.


## Step 1 - Understand the Problem and Establish Design Scope

### Clarification Questions

1. **Type of Chat App**:
   - Both 1 on 1 and group chat.

2. **Platform**:
   - Mobile and web app.

3. **Scale**:
   - Support 50 million daily active users (DAU).

4. **Group Chat Member Limit**:
   - Maximum of 100 people.

5. **Features**:
   - 1 on 1 chat, group chat, online indicator, text messages only.

6. **Message Size Limit**:
   - Less than 100,000 characters.

7. **Encryption**:
   - Not required initially.

8. **Chat History**:
   - Store forever.

### Design Scope

- **One-on-One Chat**: Low delivery latency.
- **Group Chat**: Max 100 people.
- **Online Presence**: Show if a user is online.
- **Multiple Device Support**: Same account can be logged in on multiple devices.
- **Push Notifications**: Notify users of new messages.
- **Scale**: Support 50 million DAU.


### High Level Design

#### Overview
- **Clients:** Can be either mobile applications or web applications.
- **Chat Service:** Acts as an intermediary to handle message transmission between clients. Clients do not communicate directly with each other.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/chat.webp)

#### Fundamental Operations of the Chat Service
1. **Receive messages from clients.**
2. **Identify the correct recipients for each message and relay the messages accordingly.**
3. **Hold messages for offline recipients until they come online.**

#### Client-Server Relationship (Refer to Figure 2)
- **Client Initiation:** When a client wants to start a chat, it connects to the chat service using one or more network protocols.
- **Importance of Network Protocols:** The choice of network protocol is crucial for the efficiency and reliability of the chat service.

#### Sending Messages (Sender Side)
- **Protocol Used:** HTTP (HyperText Transfer Protocol)
  - **Connection:** The client opens an HTTP connection with the chat service.
  - **Message Transmission:** The sender sends a message to the receiver via the chat service.
  - **Keep-Alive Header:** 
    - Maintains a persistent connection.
    - Reduces the number of TCP handshakes.
  - **Historical Context:** Many popular chat applications, like Facebook, initially used HTTP for sending messages.

#### Receiving Messages (Receiver Side)
- **Challenge:** HTTP is client-initiated, making it non-trivial to send messages from the server to the client.
- **Techniques for Simulating Server-Initiated Connections:**
  1. **Polling (Refer to Figure 3):** 
     - The client periodically asks the server if there are messages available.
     - **Drawback:** Can be costly due to the frequent queries which often result in no new messages, consuming server resources.
  2. **Long Polling (Refer to Figure 4):**
     - The client holds the connection open until new messages are available or a timeout occurs.
     - Once a message is received, the client immediately sends another request.
     - **Drawbacks:**
       - Sender and receiver may not connect to the same chat server. HTTP based servers are usually stateless. If you use round robin for load balancing, the server that receives the message might not have a long-polling connection with the client who receives the message.
       - Inefficient if the user is inactive, as it still makes periodic connections.
       - Difficulty in detecting client disconnections.
  3. **WebSocket (Refer to Figure 5):**
     - Initiated by the client, allowing bi-directional and persistent communication.
     - Begins as an HTTP connection, then upgraded to WebSocket via a handshake.
     - Uses ports 80 or 443, generally working even with firewalls in place.
     - Efficient for both sending and receiving messages (Refer to Figure 6).
     - **Advantages:**
       - Simplifies design and implementation for both client and server.
       - Persistent connections require efficient server-side connection management.

### Summary
- **Sender Side:** HTTP with keep-alive for persistent connections.
- **Receiver Side:** Techniques like polling, long polling, and WebSocket are used to efficiently handle server-initiated communication.
- **Optimal Solution:** Using WebSocket for both sending and receiving simplifies the system design, making it more efficient and straightforward to implement.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/polling.svg" width="500" />
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/longpolling.svg" width="500" /> 
   <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/websocket.svg" width="500" /> 
   <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/general.webp" width="500" /> 

</p>


### High-Level Design of the Chat System

#### Overview
- **WebSocket Protocol:** Used for bidirectional communication between client and server. Other features (e.g., sign up, login, user profile) can use the traditional request/response method over HTTP.
- **System Components:** The system is divided into three major categories: stateless services, stateful services, and third-party integration.

#### Components of the System

##### Stateless Services
- **Purpose:** Manage login, signup, user profile, etc.
- **Characteristics:** Public-facing request/response services, can be monolithic or microservices.
- **Load Balancer:** Routes requests to the correct services based on request paths.
- **Service Discovery:** Provides clients with a list of DNS hostnames of chat servers they can connect to.

##### Stateful Service
- **Chat Service:** 
  - **Nature:** Stateful, as each client maintains a persistent connection to a chat server.
  - **Connection:** Clients do not switch servers unless necessary, avoiding server overloading through service discovery.
  - **Function:** Facilitates real-time messaging.

##### Third-Party Integration
- **Push Notification:**
  - **Importance:** Notifies users of new messages even when the app is not running.
  - **Integration:** Crucial for user engagement and timely updates.

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/dg.webp)

#### Scalability
- **Initial Design:** 
  - Starting with a single server design is acceptable as a conceptual starting point.
  - **Limitation:** Single server design is impractical for large scale due to single points of failure.
- **Scalable Design:**
  - Multiple servers to handle different components (chat servers, presence servers, API servers, notification servers).
  - **Presence Servers:** Manage online/offline status of users.
  - **API Servers:** Handle user login, signup, profile changes, etc.
  - **Notification Servers:** Send push notifications.
  - **Key-Value Store:** Stores chat history, allowing users to retrieve past messages.

#### Storage
- **Data Layer:** Critical to the system, requiring the correct type of database.
- **Data Types:**
  1. **Generic Data:** User profiles, settings, friends list.
     - **Storage:** Relational databases (robust, reliable).
     - **Techniques:** Replication and sharding for scalability and availability.
  2. **Chat History Data:** Unique to chat systems, large volume.
     - **Access Pattern:** Recent chats accessed frequently; old chats accessed infrequently but require random access support.
     - **Read/Write Ratio:** Approximately 1:1 for one-on-one chat apps.
- **Storage System Recommendation:**
  - **Key-Value Stores:** Preferred due to easy horizontal scaling, low latency access, and proven reliability in other chat applications.
  - **Examples:** 
    - Facebook Messenger uses HBase.
    - Discord uses Cassandra.

### High-Level Design Diagram (Refer to Figure 8)
- **Client:** Maintains a persistent WebSocket connection to a chat server.
- **Chat Servers:** Handle real-time message sending and receiving.
- **Presence Servers:** Manage user online/offline status.
- **API Servers:** Handle all HTTP-based requests (login, signup, profile management).
- **Notification Servers:** Send push notifications.
- **Key-Value Store:** Stores chat history for retrieval by clients.

By designing the system with these components and considerations, the chat application can efficiently handle real-time messaging, user management, and scalability, ensuring a robust and reliable user experience.
![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/highlevelDiagram.webp)


### Data Models for Chat System

#### Key-Value Stores
- **Storage Layer:** Using key-value stores for the storage layer.
- **Focus:** The most important data is message data.

#### Message Table for 1-on-1 Chat (Refer to Figure 9)
- **Primary Key:** `message_id`
  - **Purpose:** Determines the sequence of messages.
  - **Reason:** Relying on `created_at` is unreliable as two messages can be created simultaneously.

#### Message Table for Group Chat (Refer to Figure 10)
- **Composite Primary Key:** `(channel_id, message_id)`
  - **channel_id:** Partition key, representing a group or channel.
  - **message_id:** Ensures the sequence of messages within the channel.

#### Message ID Generation
- **Requirements:**
  1. **Uniqueness:** IDs must be unique.
  2. **Sortable by Time:** New rows should have higher IDs than old ones.

- **Approaches:**
  1. **Auto Increment (MySQL):** 
     - Not suitable for NoSQL databases as they usually do not provide this feature.
  2. **Global 64-bit Sequence Number Generator (e.g., Snowflake):**
     - Ensures globally unique and sortable IDs.
     - Discussed in detail in the “Design a unique ID generator in a distributed system” chapter.
  3. **Local Sequence Number Generator:**
     - IDs are unique within a group.
     - Suitable for maintaining message sequence within a one-on-one or group channel.
     - Easier to implement compared to a global ID generator.

### Example Data Models

#### 1-on-1 Chat Message Table

| Field       | Type       | Description                        |
|-------------|------------|------------------------------------|
| message_id  | UUID       | Primary key, unique message ID     |
| sender_id   | UUID       | ID of the sender                   |
| receiver_id | UUID       | ID of the receiver                 |
| content     | Text       | Message content                    |
| created_at  | Timestamp  | Timestamp of message creation      |
| status      | String     | Status of the message (e.g., sent, delivered, read) |

#### Group Chat Message Table

| Field       | Type       | Description                        |
|-------------|------------|------------------------------------|
| channel_id  | UUID       | Partition key, ID of the channel   |
| message_id  | UUID       | Composite primary key, unique message ID within the channel |
| sender_id   | UUID       | ID of the sender                   |
| content     | Text       | Message content                    |
| created_at  | Timestamp  | Timestamp of message creation      |
| status      | String     | Status of the message (e.g., sent, delivered, read) |

### Summary
- **Data Storage:** Use key-value stores for efficient horizontal scaling and low latency access.
- **Message Tables:** Separate tables for 1-on-1 and group chats, each with appropriate primary keys.
- **ID Generation:** Implement either a global or local sequence number generator to ensure unique and time-sortable message IDs. Local sequence number generators are easier to implement and sufficient for maintaining order within specific channels.


### Design Deep Dive for Chat System

#### Service Discovery
The primary role of service discovery is to recommend the best chat server for a client based on criteria like geographical location, server capacity, etc. Apache Zookeeper is a popular solution for this purpose.

**How Service Discovery Works (Refer to Figure 11):**
1. User A logs into the app.
2. The load balancer routes the login request to API servers.
3. After authentication, service discovery finds the best chat server for User A (e.g., Server 2) and returns the server info to User A.
4. User A connects to Chat Server 2 via WebSocket.

#### Messaging Flows

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/chatflow.webp" width="500" />
</p>
**1-on-1 Chat Flow (Refer to Figure 12):**
1. User A sends a message to Chat Server 1.
2. Chat Server 1 obtains a message ID from the ID generator.
3. Chat Server 1 sends the message to the message sync queue.
4. The message is stored in a key-value store.
5. If User B is online, the message is forwarded to Chat Server 2 (where User B is connected).
6. If User B is offline, a push notification is sent.
7. Chat Server 2 forwards the message to User B via a persistent WebSocket connection.

**Message Synchronization Across Multiple Devices (Refer to Figure 13):**
1. User A logs in with two devices (phone and laptop), each establishing a WebSocket connection with Chat Server 1.
2. Each device maintains a `cur_max_message_id` to track the latest message.
3. New messages are identified if:
   - The recipient ID matches the logged-in user ID.
   - The message ID in the key-value store is greater than `cur_max_message_id`.
4. Each device fetches new messages from the key-value store based on these conditions.

**Small Group Chat Flow (Refer to Figures 14 and 15):**
1. User A sends a message in a group chat.
2. The message is copied to each group member’s message sync queue (inbox).
3. On the recipient side, each user has an inbox containing messages from different senders.
<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/smallgp.webp" width="500" />
</p>

#### Online Presence

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/userlogin.webp" width="500" />
</p>

**User Login (Refer to Figure 16):**
1. User logs in and establishes a WebSocket connection with the chat server.
2. User’s online status and `last_active_at` timestamp are saved in the key-value store.
3. Presence indicator shows the user as online.

**User Logout (Refer to Figure 17):**
1. User logs out.
2. Online status is updated to offline in the key-value store.
3. Presence indicator shows the user as offline.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/userlogout.webp" width="500" />
</p>

**User Disconnection (Refer to Figure 18):**
1. Persistent connection is lost if the user disconnects from the internet.
2. To avoid frequent status changes, a heartbeat mechanism is used.
3. Client sends a heartbeat event to presence servers periodically (e.g., every 5 seconds).
4. If no heartbeat is received within a certain time (e.g., 30 seconds), the user is marked offline.

<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/heartbeat.webp" width="500" />
</p>

**Online Status Fanout (Refer to Figure 19):**
1. Presence servers use a publish-subscribe model for online status updates.
2. Each friend pair maintains a channel.
3. When User A’s status changes, the event is published to the relevant channels (e.g., A-B, A-C, A-D).
4. Friends (User B, C, D) subscribed to these channels receive the status updates via WebSocket.


This design is efficient for small groups (e.g., WeChat limits groups to 500 members). For larger groups, fetching online status only when entering a group or manually refreshing the friend list can mitigate performance bottlenecks.