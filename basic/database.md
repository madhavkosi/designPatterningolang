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