#### SQL Normalization and Denormalization

**SQL Normalization**

**Definition**: Normalization in SQL is a database design technique that organizes tables to reduce redundancy and dependency by dividing a database into two or more tables and defining relationships between them.

**Characteristics**:
- **Reduces Redundancy**: Avoids data duplication.
- **Improves Data Integrity**: Ensures data accuracy and consistency.
- **Database Design**: Involves creating tables and establishing relationships through primary and foreign keys.

**Example: Customer Orders Database**
- **Original Table (Before Normalization)**:

| Customer ID | Customer Name | Customer Address | Order ID | Order Date | Product |
|-------------|---------------|------------------|----------|------------|---------|
| 001         | John Doe      | 123 Apple St.    | 1001     | 2021-08-01 | Laptop  |
| 001         | John Doe      | 123 Apple St.    | 1002     | 2021-08-05 | Phone   |
| 002         | Jane Smith    | 456 Orange Ave.  | 1003     | 2021-08-03 | Tablet  |

- **After Normalization**:
  - **Customers Table (1NF, 2NF, 3NF)**:

  | Customer ID | Customer Name | Customer Address |
  |-------------|---------------|------------------|
  | 001         | John Doe      | 123 Apple St.    |
  | 002         | Jane Smith    | 456 Orange Ave.  |

  - **Orders Table (1NF, 2NF, 3NF)**:

  | Order ID | Order Date | Product | Customer ID |
  |----------|------------|---------|-------------|
  | 1001     | 2021-08-01 | Laptop  | 001         |
  | 1002     | 2021-08-05 | Phone   | 001         |
  | 1003     | 2021-08-03 | Tablet  | 002         |

**Levels (Normal Forms)**:
- **1NF**: Data is stored in atomic form with no repeating groups.
- **2NF**: Meets 1NF and has no partial dependency on any candidate key.
- **3NF**: Meets 2NF and has no transitive dependency.

**Use Cases**: Ideal for complex systems where data integrity is critical, such as financial or enterprise applications.

**SQL Denormalization**

**Definition**: Denormalization is the process of combining tables to reduce the complexity of database queries, which can introduce redundancy but may lead to improved performance by reducing the number of joins required.

**Characteristics**:
- **Increases Redundancy**: May involve some data duplication.
- **Improves Query Performance**: Reduces the complexity of queries by reducing the number of joins.
- **Data Retrieval**: Optimized for read-heavy operations.

**Denormalization Example**:
- **Denormalized Orders Table**:

| Customer ID | Customer Name | Customer Address | Order ID | Order Date | Product |
|-------------|---------------|------------------|----------|------------|---------|
| 001         | John Doe      | 123 Apple St.    | 1001     | 2021-08-01 | Laptop  |
| 001         | John Doe      | 123 Apple St.    | 1002     | 2021-08-05 | Phone   |
| 002         | Jane Smith    | 456 Orange Ave.  | 1003     | 2021-08-03 | Tablet  |

**When to Use**:
- **Read-Heavy Systems**: Where query performance is a priority.
- **Infrequent Data Changes**: When a slightly less normalized structure doesn't compromise data integrity.

**Key Differences**

**Purpose**:
- **Normalization**: Minimizes data redundancy and improves data integrity.
- **Denormalization**: Improves query performance.

**Data Redundancy**:
- **Normalization**: Reduces redundancy.
- **Denormalization**: May introduce redundancy.

**Performance**:
- **Normalization**: Can lead to more complex queries, affecting read performance.
- **Denormalization**: Improves read performance but may affect write performance due to redundancy.

**Complexity**:
- **Normalization**: Increases complexity of write operations.
- **Denormalization**: Simplifies read operations but can make write operations more complex.

**Conclusion**

- **Normalization**: Focuses on reducing redundancy and improving data integrity, but can lead to more complex queries.
- **Denormalization**: Simplifies queries but increases data redundancy and potential maintenance challenges.
- **Choice**: Depends on database system requirements, balancing read vs. write operation frequency, and prioritizing query performance vs. data integrity.



### Detailed Explanation of 1NF, 2NF, and 3NF

**Normalization** is a step-by-step process of organizing data in a database to reduce redundancy and improve data integrity. It involves dividing large tables into smaller, related tables and defining relationships between them. The process follows several normal forms, each with specific requirements. Here’s a detailed look at the first three normal forms: 1NF, 2NF, and 3NF.

### 1NF (First Normal Form)

**Definition**: A table is in the first normal form if it meets the following criteria:
- The table only contains atomic (indivisible) values.
- There are no repeating groups or arrays within rows.

**Characteristics**:
- **Atomicity**: Each column contains unique and indivisible values.
- **No Repeating Groups**: Each column should have a single value for each row, ensuring there are no sets of columns representing the same type of data (e.g., phone1, phone2).

**Example**:
Original table (non-1NF):

| Order ID | Customer | Items             |
|----------|----------|-------------------|
| 1001     | John Doe | Laptop, Mouse     |
| 1002     | Jane Doe | Phone, Headphones |

Normalized to 1NF:

| Order ID | Customer | Item      |
|----------|----------|-----------|
| 1001     | John Doe | Laptop    |
| 1001     | John Doe | Mouse     |
| 1002     | Jane Doe | Phone     |
| 1002     | Jane Doe | Headphones|

### 2NF (Second Normal Form)

**Definition**: A table is in the second normal form if:
- It is already in 1NF.
- All non-key attributes are fully functionally dependent on the entire primary key.

**Characteristics**:
- **No Partial Dependency**: No non-key attribute should depend on a part of the composite primary key.

**Example**:
Original table (1NF, but not 2NF):

| Order ID | Product ID | Product Name | Customer ID | Customer Name |
|----------|------------|--------------|-------------|---------------|
| 1001     | 201        | Laptop       | 001         | John Doe      |
| 1002     | 202        | Phone        | 002         | Jane Doe      |

Here, `Product Name` depends only on `Product ID`, and `Customer Name` depends only on `Customer ID`, which are parts of the composite primary key (`Order ID`, `Product ID`, `Customer ID`).

Normalized to 2NF:

**Products Table**:

| Product ID | Product Name |
|------------|--------------|
| 201        | Laptop       |
| 202        | Phone        |

**Customers Table**:

| Customer ID | Customer Name |
|-------------|---------------|
| 001         | John Doe      |
| 002         | Jane Doe      |

**Orders Table**:

| Order ID | Product ID | Customer ID |
|----------|------------|-------------|
| 1001     | 201        | 001         |
| 1002     | 202        | 002         |

### 3NF (Third Normal Form)

**Definition**: A table is in the third normal form if:
- It is already in 2NF.
- All the attributes are functionally dependent only on the primary key, not on any other non-key attributes (no transitive dependency).

**Characteristics**:
- **No Transitive Dependency**: Non-key attributes should not depend on other non-key attributes.

**Example**:
Original table (2NF, but not 3NF):

| Order ID | Customer ID | Customer Name | Customer Address |
|----------|-------------|---------------|------------------|
| 1001     | 001         | John Doe      | 123 Apple St.    |
| 1002     | 002         | Jane Doe      | 456 Orange Ave.  |

Here, `Customer Name` and `Customer Address` depend on `Customer ID`, which is a non-key attribute.

Normalized to 3NF:

**Customers Table**:

| Customer ID | Customer Name | Customer Address |
|-------------|---------------|------------------|
| 001         | John Doe      | 123 Apple St.    |
| 002         | Jane Doe      | 456 Orange Ave.  |

**Orders Table**:

| Order ID | Customer ID |
|----------|-------------|
| 1001     | 001         |
| 1002     | 002         |

### Summary

**1NF**:
- Ensures each column contains atomic values.
- Eliminates repeating groups.

**2NF**:
- Meets all requirements of 1NF.
- Eliminates partial dependency, ensuring that non-key attributes depend on the whole primary key.

**3NF**:
- Meets all requirements of 2NF.
- Eliminates transitive dependency, ensuring that non-key attributes depend only on the primary key.

By following these normalization steps, databases can be designed to reduce redundancy, improve data integrity, and enhance the efficiency of data retrieval and manipulation.


### Types of Relationships in General Database Design

| Relationship Type | Description | Example |
|-------------------|-------------|---------|
| **One-to-One (1:1)** | Each row in Table A is linked to one and only one row in Table B, and vice versa. | Person and Passport tables where each person has one passport and each passport belongs to one person. |
| **One-to-Many (1:N)** | Each row in Table A can be linked to multiple rows in Table B, but each row in Table B is linked to one and only one row in Table A. | Customer and Order tables where each customer can place multiple orders, but each order is placed by one customer. |
| **Many-to-One (N:1)** | Each row in Table B can be linked to multiple rows in Table A. | Book and Library tables where each library can have multiple books, but each book is located in one library. |
| **Many-to-Many (N:N)** | Each row in Table A can be linked to multiple rows in Table B, and each row in Table B can be linked to multiple rows in Table A. | Student and Course tables where each student can enroll in multiple courses, and each course can have multiple students. |



### Optimistic vs. Pessimistic Locking

In a system where multiple users might try to access and modify the same data simultaneously (such as a seat booking system in a movie theater), you need strategies to ensure **data consistency** and avoid conflicts. Optimistic and Pessimistic Locking are two common strategies used in concurrency control to handle these scenarios.

### 1. **Pessimistic Locking**

Pessimistic locking assumes that **conflicts are likely** to happen and takes steps to prevent them before any changes are made to the data.

#### How Pessimistic Locking Works:
- When a user (or process) accesses a piece of data (e.g., a seat in a movie theater), the system **locks** the data.
- This lock ensures that **no other user can access or modify the data** until the current user has finished the transaction (e.g., booking the seat).
- Once the user finishes the transaction (booking the seat), the lock is released, and other users can then access or modify the data.

#### Example in a Movie Booking System:
- **User A** selects seat A1 to book.
- The system places a **lock** on seat A1, ensuring that **User B** (or any other user) cannot book seat A1 while **User A** is in the process of completing the booking.
- If **User B** tries to book seat A1 during this time, they will either be blocked or receive a message saying the seat is already being processed by another user.
- Once **User A** completes the booking (or cancels), the system **releases the lock** on seat A1, allowing other users to book it.

#### Benefits:
- **Data integrity** is guaranteed because only one user can access the data at a time.
- Suitable for systems where there are high chances of conflict (e.g., a seat might be in high demand during peak hours).

#### Drawbacks:
- **Reduced performance**: Pessimistic locking can lead to bottlenecks, especially in a high-concurrency system, because users might be blocked waiting for a lock to be released.
- **Deadlocks**: If multiple users hold locks on different resources and wait for each other to release their locks, it can cause deadlocks.

### 2. **Optimistic Locking**

Optimistic locking, on the other hand, assumes that **conflicts are rare** and does not lock the data when it's read. Instead, it detects conflicts only when the data is updated.

#### How Optimistic Locking Works:
- When a user reads a piece of data (e.g., seat A1), no lock is placed.
- The system keeps track of the **version** of the data (or a timestamp) when it is read.
- When the user attempts to update the data (e.g., book the seat), the system checks whether the **data version has changed** since it was read.
    - If the version is the same, the update proceeds.
    - If the version has changed (indicating another user has modified the data), the update is **rejected**, and the user is prompted to retry the transaction.

#### Example in a Movie Booking System:
- **User A** selects seat A1, and the system provides them with the current version of the seat (say, version 1).
- **User B** selects seat A1 shortly afterward, also receiving version 1.
- **User A** completes the booking and updates the seat status to "Booked." The system then increments the version of seat A1 to version 2.
- **User B** now attempts to book seat A1. However, since the version has changed from 1 to 2 (due to **User A's** booking), the system **rejects the update** from **User B** and informs them that the seat is no longer available.
- **User B** must refresh their view and choose a different seat.

#### Benefits:
- **Better performance**: Since no locks are used when reading the data, multiple users can access the data simultaneously without blocking each other.
- **Scalability**: Optimistic locking is well-suited for systems with many concurrent users because conflicts are assumed to be rare, and users are not blocked from accessing data.

#### Drawbacks:
- **Conflict handling**: Users might occasionally encounter errors if another user has updated the data before them, requiring them to retry the transaction.
- **Increased complexity**: The system needs to manage data versions or timestamps to detect conflicts.

### 3. **When to Use Each Locking Mechanism**

#### Use **Pessimistic Locking** When:
- **High contention**: When there is a high likelihood that multiple users will try to modify the same data at the same time. For example, during a flash sale for movie tickets when many users are attempting to book seats simultaneously.
- **Critical data**: If losing or corrupting data due to concurrent modifications is unacceptable, such as in financial systems or when booking highly valuable resources (e.g., VIP tickets).
- **Short transactions**: Pessimistic locking is more suitable when transactions are short, minimizing the time that data is locked.

#### Use **Optimistic Locking** When:
- **Low contention**: If it's rare for multiple users to modify the same data at the same time, optimistic locking allows for better performance.
- **Scalability**: In systems with many concurrent users but low conflict rates, such as in distributed applications or large-scale web services.
- **Read-heavy systems**: If the system is primarily read-heavy and conflicts during data writes are rare, optimistic locking allows many users to read data without delays caused by locks.

### 4. **How to Implement Each Locking Mechanism**

#### Implementing **Pessimistic Locking**:
Most relational databases support pessimistic locking natively. You can use database transactions to lock rows when they are read:

- **For SQL databases**: Use a `SELECT FOR UPDATE` statement to lock the rows being accessed. For example:

    ```sql
    SELECT * FROM seats WHERE seat_id = 'A1' FOR UPDATE;
    ```

    This locks the row corresponding to seat A1, preventing other users from modifying it until the transaction is complete.

- **In NoSQL databases**: Some NoSQL databases (like MongoDB) offer locking mechanisms at the document level, though implementation might vary.

#### Implementing **Optimistic Locking**:
Optimistic locking can be implemented by adding a **version number** or **timestamp** to each row or record in the database. Every time a record is updated, the version number is incremented or the timestamp is updated.

- When a user reads a record, the version number or timestamp is also fetched.
- When the user updates the record, the system checks if the version number or timestamp has changed since the record was last read.

For example:

1. **Select seat with version number**:

    ```sql
    SELECT seat_id, status, version FROM seats WHERE seat_id = 'A1';
    ```

2. **Update seat status with version check**:

    ```sql
    UPDATE seats
    SET status = 'Booked', version = version + 1
    WHERE seat_id = 'A1' AND version = 1;
    ```

   If the `version` is still 1 (meaning no other user has updated it), the update proceeds. If another user has already updated the record (changing the version to 2), this query will fail, and the user will be prompted to try again.

### 5. **Conclusion**

Both **optimistic** and **pessimistic locking** provide ways to handle concurrent access to data, but they are suited for different scenarios:

- **Pessimistic locking** is better for situations where conflicts are likely and data consistency is critical, but it comes with the cost of reduced system throughput and the potential for deadlocks.
- **Optimistic locking** works well when conflicts are rare, providing better performance and scalability but requiring some extra handling for retrying transactions.

In a movie booking system, **optimistic locking** is often preferred, as conflicts (multiple users booking the same seat) are rare in most cases, and it allows the system to scale and perform better under heavy load. However, for highly contested events, such as the opening of bookings for a blockbuster movie, **pessimistic locking** might be considered to ensure absolute data integrity.


### Serializable Lock vs. Optimistic and Pessimistic Locking - Short Notes

**1. Definitions**  
- **Serializable Lock:**  
  - Highest isolation level, ensures transactions appear sequential.  
  - Uses strict locks or MVCC.  
  - Goal: Perfect consistency.
  
- **Optimistic Locking:**  
  - No locks during read, checks for conflicts during write.  
  - Goal: High concurrency, detects conflicts at write time.

- **Pessimistic Locking:**  
  - Locks data as soon as it's read to prevent conflicts.  
  - Goal: Prevent conflicts by blocking access during a transaction.

**2. How They Work**  
- **Serializable:** Locks rows/tables or uses MVCC to ensure no conflicts and full isolation.
- **Optimistic:** No locks during reads, uses versioning/timestamps to detect changes at write time.
- **Pessimistic:** Locks data immediately upon read and holds until transaction ends.

**3. When to Use**  
- **Serializable:** Critical systems like financial apps, where consistency > performance.
- **Optimistic:** High concurrency, low conflict systems (e.g., web apps).
- **Pessimistic:** High conflict systems like booking and reservation systems.

**4. Performance and Trade-offs**  
- **Serializable:**  
  - **Performance:** Low due to high locking overhead, but highest consistency.  
  - **Trade-off:** Perfect consistency, low scalability.
  
- **Optimistic:**  
  - **Performance:** High concurrency, low read overhead.  
  - **Trade-off:** Possible transaction retries, adds complexity.
  
- **Pessimistic:**  
  - **Performance:** Lower concurrency due to locks, potential for deadlocks.  
  - **Trade-off:** Ensures data integrity but reduces throughput.

**5. Examples**  
- **Serializable:** Financial transactions (bank account transfers).  
- **Optimistic:** E-commerce systems (product purchases).  
- **Pessimistic:** Movie ticket booking systems (seat reservations).

**Comparison Summary:**
| Characteristic      | Serializable       | Optimistic          | Pessimistic       |
|---------------------|--------------------|---------------------|-------------------|
| Conflict Handling    | Serializes all     | Conflicts checked on write | Conflicts blocked at read |
| Performance Impact   | High (low concurrency) | Low (high concurrency) | Medium to High (due to blocking) |
| Concurrency          | Low                | High                | Low to Medium     |
| Data Integrity       | Highest            | High, with retries   | High              |
| Use Case             | Financial systems  | Web apps, distributed systems | Reservations, auctions |

