

###  What is Pastebin?
- A web service that allows users to store plain text or images online.
- Generates unique URLs to access the uploaded content.
- Used for quickly sharing data by passing the URL to others.

### Requirements and Goals of the System

#### Functional Requirements:
1. **Text Uploading**: Users can upload or paste text and receive a unique URL for access.
2. **Text Retrieval**: The system allows retrieving the text using the generated URL.
3. **Expiration**: 
   - Data and links expire after a specific timespan automatically.
   - Users can specify a custom expiration time.
4. **Custom Alias**: Users can optionally choose a custom alias for their paste URLs.

#### Non-Functional Requirements:
1. **Reliability**: 
   - Ensure no data loss; data persistence is critical.
2. **Availability**: 
   - The service should be available 24/7 to prevent downtime and ensure users can always access their pastes.
3. **Real-time Access**: 
   - Provide minimal latency for users accessing their pastes.
4. **Security**: 
   - URLs should be unique and non-predictable to prevent unauthorized access.

#### Extended Requirements:
1. **Analytics**: 
   - Track how many times a paste has been accessed.
2. **REST APIs**: 
   - Offer APIs for external services to interact with the Pastebin service.


#### Design Considerations for Pastebin Service

##### Text Size Limit
- **Limit on Text Size**: Set a maximum of 10MB per paste to prevent abuse.

##### Custom URL Size Limit
- **Size Limit on Custom URLs**: Limit custom URLs to 30-50 characters for consistency.



### Capacity Estimation and Constraints for Pastebin Service

#### Traffic Estimates
- **New Pastes per Day**: 1 million
- **Read Requests per Day**: 5 million
- **Write Requests per Second**: 
  - \( \frac{1M}{24 \times 3600} \approx 12 \) pastes/sec
- **Read Requests per Second**: 
  - \( \frac{5M}{24 \times 3600} \approx 58 \) reads/sec

#### Storage Estimates
- **Average Paste Size**: 10KB
- **Daily Storage Requirement**: 
  - \( 1M \times 10KB = 10GB \)/day
- **10-Year Storage Requirement**: 
  - \( 10GB/day \times 365 \times 10 = 36TB \)
- **Number of Pastes in 10 Years**: 
  - \( 1M \times 365 \times 10 = 3.6 \) billion pastes
- **Storage for Unique Keys**: 
  - Base64 encoding with 6 characters: \( 64^6 \approx 68.7 \) billion unique strings
  - Total size for 3.6B keys: \( 3.6B \times 6 = 22GB \)
- **Total Storage with 70% Capacity Model**: 
  - \( 36TB / 0.7 \approx 51.4TB \)

#### Bandwidth Estimates
- **Ingress (Write Requests)**:
  - \( 12 \times 10KB = 120KB/s \)
- **Egress (Read Requests)**:
  - \( 58 \times 10KB = 0.6MB/s \)

#### Memory Estimates for Caching
- **Read Requests per Day**: 5 million
- **Hot Pastes (20%)**:
  - \( 0.2 \times 5M = 1M \) requests/day
- **Memory for Caching 20% of Requests**: 
  - \( 1M \times 10KB = 10GB \)

#### Summary
- **Write Requests**: 12 pastes/sec
- **Read Requests**: 58 reads/sec
- **Daily Storage**: 10GB
- **10-Year Storage**: 51.4TB (with 70% capacity model)
- **Bandwidth**: 120KB/s ingress, 0.6MB/s egress
- **Caching Memory**: 10GB for hot pastes



### System APIs for Pastebin Service

#### 1. Create Paste API
- **Endpoint**: `POST /api/paste`
- **Parameters**:
  - `api_dev_key` (string): API key for authentication.
  - `paste_data` (string): Text content of the paste.
  - `custom_url` (string, optional): Custom URL for the paste.
  - `user_name` (string, optional): Username for URL generation.
  - `paste_name` (string, optional): Name of the paste.
  - `expire_date` (string, optional): Expiration date for the paste.
- **Returns**: URL for accessing the paste or error code.

#### 2. Retrieve Paste API
- **Endpoint**: `GET /api/paste/{api_paste_key}`
- **Parameters**:
  - `api_dev_key` (string): API key for authentication.
  - `api_paste_key` (string): Key of the paste to retrieve.
- **Returns**: Text content of the paste or error code.

#### 3. Delete Paste API
- **Endpoint**: `DELETE /api/paste/{api_paste_key}`
- **Parameters**:
  - `api_dev_key` (string): API key for authentication.
  - `api_paste_key` (string): Key of the paste to delete.
- **Returns**: `true` if successful, otherwise `false`.



### Database Schema for Pastebin Service

#### Tables

1. **Users**

| Column           | Type         | Description                                 |
|------------------|--------------|---------------------------------------------|
| id               | UUID         | Primary key, unique identifier for each user|
| username         | VARCHAR(255) | Unique username                             |
| email            | VARCHAR(255) | User's email address                        |
| password_hash    | VARCHAR(255) | Hashed password                             |
| created_at       | TIMESTAMP    | Timestamp of when the user was created      |

**Users Table Explanation**:
- **id**: Unique identifier for each user.
- **username**: Unique username for each user.
- **email**: Email address of the user.
- **password_hash**: Hashed password for security.
- **created_at**: Timestamp indicating when the user account was created.

2. **Pastes**

| Column           | Type         | Description                                      |
|------------------|--------------|--------------------------------------------------|
| id               | UUID         | Primary key, unique identifier for each paste    |
| user_id          | UUID         | Foreign key referencing `Users(id)`              |
| url_hash         | VARCHAR(255) | Unique URL equivalent of the TinyURL             |
| content_key      | VARCHAR(255) | Reference to external object storing paste content|
| custom_url       | VARCHAR(255) | Optional custom URL                              |
| paste_name       | VARCHAR(255) | Optional name of the paste                       |
| expire_date      | TIMESTAMP    | Expiration date and time of the paste            |
| created_at       | TIMESTAMP    | Timestamp of when the paste was created          |
| access_count     | INT          | Number of times the paste has been accessed      |

**Pastes Table Explanation**:
- **id**: Unique identifier for each paste.
- **user_id**: Foreign key linking the paste to its creator.
- **url_hash**: Unique URL equivalent of the TinyURL.
- **content_key**: Reference to an external storage object for the paste content.
- **custom_url**: Optional custom URL for accessing the paste.
- **paste_name**: Optional name for the paste.
- **expire_date**: Timestamp indicating when the paste will expire.
- **created_at**: Timestamp indicating when the paste was created.
- **access_count**: Counter tracking the number of accesses.

#### Example SQL Schema

```sql
CREATE TABLE Users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Pastes (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES Users(id) ON DELETE CASCADE,
    url_hash VARCHAR(255) UNIQUE NOT NULL,
    content_key VARCHAR(255) NOT NULL,
    custom_url VARCHAR(255),
    paste_name VARCHAR(255),
    expire_date TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
);
```

**Indexes**:
- Create an index on `url_hash` in the `Pastes` table for fast retrieval.
- Create an index on `expire_date` in the `Pastes` table for efficient expiration handling.

**Indexes Example**:

```sql
CREATE INDEX idx_paste_url_hash ON Pastes(url_hash);
CREATE INDEX idx_paste_expire_date ON Pastes(expire_date);
```

This schema design provides a structured way to manage users and their pastes, while also planning for the efficient handling of external storage references for paste content.


### High-Level Design

- **Application Layer**:
  - Serves all read and write requests.
  - Communicates with the storage layer to store and retrieve data.

- **Storage Layer**:
  - **Metadata Storage**:
    - Stores metadata related to each paste and user information.
  - **Content Storage**:
    - Stores the actual paste contents in object storage.

- **Scalability**:
  - The separation of metadata and content storage allows for individual scaling of each component.



### Component Design

#### a. Application Layer
- **Handles Requests**: Processes all incoming and outgoing requests by communicating with the backend data store components.
- **Write Requests**:
  - Generates a six-letter random string as the key (unless a custom key is provided).
  - Stores the paste contents and key in the database.
  - Retries key generation in case of duplicate key insertion failure.
  - Returns an error if a provided custom key is already in use.
- **Key Generation Service (KGS)**:
  - Standalone service generating and storing unique six-letter strings in a key database (key-DB).
  - Ensures uniqueness and avoids duplication.
  - Uses two tables: one for unused keys and one for used keys.
  - Keeps keys in memory for quick access.
  - Handles failures by having a standby replica.
- **Read Requests**:
  - Contacts the datastore to retrieve paste content using the key.
  - Returns the paste content if the key is found, otherwise returns an error code.

#### b. Datastore Layer
- **Metadata Database**:
  - Stores metadata related to pastes and users.
  - Can use a relational database (e.g., MySQL) or a distributed key-value store (e.g., DynamoDB, Cassandra).
- **Object Storage**:
  - Stores actual paste contents.
  - Uses scalable solutions like Amazon S3.
  - Capacity can be increased by adding more servers when needed.


![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/pastebin.png)