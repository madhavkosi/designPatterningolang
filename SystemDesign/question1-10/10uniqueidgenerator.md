## Designing a Unique ID Generator in Distributed Systems

### Step 1 - Understand the Problem and Establish Design Scope

#### Clarification Questions

1. **What are the characteristics of unique IDs?**
   - **Interviewer:** IDs must be unique and sortable.

2. **For each new record, does ID increment by 1?**
   - **Interviewer:** The ID increments by time but not necessarily only increments by 1. IDs created in the evening are larger than those created in the morning on the same day.

3. **Do IDs only contain numerical values?**
   - **Interviewer:** Yes, that is correct.

4. **What is the ID length requirement?**
   - **Interviewer:** IDs should fit into 64-bit.

5. **What is the scale of the system?**
   - **Interviewer:** The system should be able to generate 10,000 IDs per second.

#### Requirements

- IDs must be unique.
- IDs are numerical values only.
- IDs fit into 64-bit.
- IDs are ordered by date.
- Ability to generate over 10,000 unique IDs per second.

### Step 2 - Propose High-Level Design and Get Buy-In

#### Multiple Options for Unique ID Generation

1. **Multi-Master Replication**

   - **How It Works:** Uses databases’ auto_increment feature with an increment by k, where k is the number of database servers in use.
   - **Pros:** 
     - Utilizes existing database features.
   - **Cons:** 
     - Hard to scale with multiple data centers.
     - IDs do not go up with time across multiple servers.
     - Difficult to manage when a server is added or removed.

2. **Universally Unique Identifier (UUID)**

   - **How It Works:** Each web server contains an ID generator responsible for generating IDs independently.
   - **Pros:** 
     - Simple generation with no server coordination.
     - Easy to scale with web servers.
   - **Cons:** 
     - 128-bit length, which is not suitable for the 64-bit requirement.
     - IDs do not go up with time and could be non-numeric.

3. **Ticket Server**

   - **How It Works:** A centralized auto_increment feature in a single database server (Ticket Server).
   - **Pros:** 
     - Numeric IDs.
     - Easy to implement for small to medium-scale applications.
   - **Cons:** 
     - Single point of failure.
     - Scalability issues with multiple ticket servers.

4. **Twitter Snowflake Approach**

   - **How It Works:** Divides an ID into different sections (sign bit, timestamp, datacenter ID, machine ID, and sequence number).
   - **Pros:** 
     - Numeric, 64-bit IDs.
     - IDs are sortable by time.
     - High scalability and fault tolerance.
   - **Cons:** 
     - More complex to implement compared to simpler methods.

### Step 3 - Design Deep Dive

#### Twitter Snowflake Approach

We will use the Twitter Snowflake approach, which divides the ID into sections:

- **Sign Bit (1 bit):** Reserved for future use, always 0.
- **Timestamp (41 bits):** Milliseconds since a custom epoch (Twitter's epoch: Nov 04, 2010, 01:42:54 UTC).
- **Datacenter ID (5 bits):** Supports up to 32 datacenters.
- **Machine ID (5 bits):** Supports up to 32 machines per datacenter.
- **Sequence Number (12 bits):** Supports 4096 IDs per millisecond per machine.

**Example ID Layout:**
```
| 1 bit (sign) | 41 bits (timestamp) | 5 bits (datacenter ID) | 5 bits (machine ID) | 12 bits (sequence number) |
```


<p float="left">
  <img src="https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/datacenter.svg" width="800" />
</p>

#### Detailed Design

1. **Initialization:**
   - Datacenter IDs and machine IDs are assigned at startup and should remain fixed unless reviewed and changed carefully.

2. **Timestamp Handling:**
   - Timestamps ensure that IDs are sortable by time.
   - The timestamp part of the ID uses the milliseconds since the custom epoch.
   - The maximum timestamp that can be represented in 41 bits is approximately 69 years from the custom epoch.

3. **Sequence Number:**
   - Resets to 0 every millisecond.
   - Ensures that up to 4096 IDs can be generated per millisecond per machine.
### Summary

The Twitter Snowflake approach provides a robust solution for generating unique 64-bit IDs in a distributed system. By dividing the ID into a sign bit, timestamp, datacenter ID, machine ID, and sequence number, we achieve:

- Uniqueness and sortability by time.
- Numeric IDs that fit into 64-bit.
- Scalability to handle over 10,000 IDs per second.

This design ensures that IDs are unique, ordered by date, and efficiently generated across multiple datacenters and machines.


