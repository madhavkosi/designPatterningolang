### Functional Requirements for Flash Sale System:

1. **Limited Stock Availability**: The system must handle a limited inventory (e.g., 1000 items) and ensure that no more items are sold once the stock is depleted.

2. **First-Come-First-Served**: The system must allocate items to users on a first-come-first-served basis, ensuring fairness and consistency.

3. **High Traffic Handling**: The system must be capable of handling millions or even billions of requests in a short period during the flash sale.

4. **Real-time Inventory Updates**: Inventory levels must be updated in real-time, and no overselling should occur.

5. **Cart Reservation**: Once an item is added to the cart, it must be reserved for a limited time (e.g., 5 minutes), after which the item is returned to the available stock if not purchased.

6. **Order Processing**: The system should allow users to complete orders via a secure checkout, integrating with multiple payment gateways.

7. **User Notifications**: Users must be notified via email/SMS upon successful order placement or stock exhaustion.

8. **Scalability**: The system must scale to accommodate millions of concurrent users, ensuring low latency for all operations (inventory checks, adding to cart, checkout).

9. **Bot Prevention**: Mechanisms to detect and block bot traffic, ensuring only legitimate users can participate in the sale.

### Non-Functional Requirements for Flash Sale System:

1. **Performance**:
   - The system should be able to handle high volumes of traffic with minimal latency (targeting response times of under 100ms for most operations).
   - Inventory updates and checkout processes should be completed in real-time, ensuring users have up-to-date information.

2. **Scalability**:
   - The system must scale horizontally to accommodate sudden spikes in traffic during a flash sale event, potentially handling billions of requests.
   - The architecture should support dynamic scaling to handle increased loads and reduced traffic outside of flash sales.

3. **Availability**:
   - The system must ensure **99.99% uptime**, especially during the flash sale event. Downtime during the sale can result in significant business losses and customer dissatisfaction.
   - Fault-tolerance mechanisms should be in place to handle failures gracefully.

4. **Reliability**:
   - All order processes and inventory updates should be reliable, ensuring consistency in stock levels and no duplicate orders or overselling.

5. **Consistency**:
   - **Strong consistency** is required for critical operations like inventory reduction during checkout to ensure no overselling of items.

6. **Security**:
   - User data (e.g., payment information, personal details) must be handled securely, following encryption standards like TLS for data in transit and AES for data at rest.
   - Implement bot protection mechanisms such as CAPTCHA and rate-limiting.

7. **Resilience**:
   - The system must handle unexpected failures gracefully, ensuring no data loss and maintaining functionality during partial outages.

8. **Graceful Degradation**:
   - In the event of high traffic or component failures, the system should degrade gracefully (e.g., slower response times but no crashes), rather than failing completely.

9. **Compliance**:
   - The system should adhere to relevant regulations such as PCI DSS (for handling payment information) and GDPR (for protecting personal data in the EU).


### Capacity Estimation for Flash Sale System:

1. **Users**:
   - Assume the system needs to handle up to **100 million users** during the flash sale, with **10 million concurrent users** actively interacting with the system at peak times.

2. **Requests per Second (RPS)**:
   - Each user may generate 5 requests (view product, add to cart, check inventory, checkout, and payment). For 10 million concurrent users, the system must handle:
     ``` 
     10,000,000 users * 5 requests = 50,000,000 requests over a 10-minute window.
     ```
     This translates to **83,333 requests per second (RPS)**.

3. **Inventory Updates**:
   - For an inventory of 1,000 units, there will be 1,000 atomic inventory updates, but the system needs to handle checking inventory availability potentially millions of times as users compete for the same items.

4. **Data Storage**:
   - If each order record is approximately 1KB (including user data, order details, and payment status), with 1 million completed orders, the system will store around **1 GB of order data**.

5. **Bandwidth**:
   - Assuming each request/response (e.g., adding to cart, inventory checks) is around 2KB, handling **83,333 RPS** results in:
     ```
     83,333 RPS * 2KB = 166,666 KB/s = ~162 MB/s of bandwidth.
     ```

6. **Latency Targets**:
   - **Sub-100ms latency** for critical operations like inventory checks and order placements to ensure a smooth user experience.

7. **Payment Processing**:
   - If 10% of users complete purchases (1 million users), and payment processing takes around 2-3 seconds per transaction, the system should handle **333 transactions per second**.


| **API Endpoint**          | **Method** | **Description**                                           | **Request Body**                                                                                                                                                                   | **Response**                                                                                                                                                                        |
|---------------------------|------------|-----------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `/add-to-cart`             | POST       | Adds an item to the user’s cart and reserves it temporarily| `{ "user_id": "string", "product_id": "string", "quantity": "integer" }`                                                                                                            | `{ "success": true, "cart_id": "string", "reservation_expiry": "timestamp" }`                                                                                                       |
| `/checkout`               | POST       | Confirms purchase and initiates payment                   | `{ "user_id": "string", "cart_id": "string", "payment_method": "string", "shipping_address": { "address_line1": "string", "city": "string", "postal_code": "string", "country": "string" } }` | `{ "order_id": "string", "payment_status": "pending | completed | failed", "inventory_status": "confirmed | failed" }`                                                            |
| `/inventory/{product_id}` | GET        | Retrieves current inventory level                         | N/A                                                                                                                                                                                | `{ "product_id": "string", "available_quantity": "integer" }`                                                                                                                       |
| `/payment`                | POST       | Processes payment and finalizes the transaction            | `{ "order_id": "string", "payment_method": "string", "amount": "float" }`                                                                                                          | `{ "success": true, "transaction_id": "string", "payment_status": "completed" }`                                                                                                    |
| `/reserve-inventory`      | POST       | Temporarily reserves inventory for checkout                | `{ "user_id": "string", "product_id": "string", "quantity": "integer" }`                                                                                                            | `{ "success": true, "reservation_expiry": "timestamp" }`                                                                                                                            |

This format provides a clear view of the APIs, their methods, request body, and the expected responses.

### Database Design for Flash Sale System:

The database design will focus on ensuring **consistency**, **high performance**, and **scalability**. We’ll use **PostgreSQL** or **Cassandra** for persistent data storage and **Redis** for real-time inventory management and caching.

#### 1. **Table: Inventory**
   - **Purpose**: Stores the product inventory details, including stock levels.
   - **Columns**:
     | Column Name       | Data Type      | Description                              |
     |-------------------|----------------|------------------------------------------|
     | `product_id`       | VARCHAR(255)   | Unique identifier for the product        |
     | `product_name`     | VARCHAR(255)   | Name of the product                      |
     | `available_quantity` | INT            | Number of units available                |
     | `price`            | DECIMAL(10, 2) | Price of the product                     |
     | `last_updated`     | TIMESTAMP      | Timestamp of the last update             |

   ```sql
   CREATE TABLE Inventory (
       product_id VARCHAR(255) PRIMARY KEY,
       product_name VARCHAR(255),
       available_quantity INT,
       price DECIMAL(10, 2),
       last_updated TIMESTAMP
   );
   ```

#### 2. **Table: Orders**
   - **Purpose**: Stores confirmed orders with associated user and product information.
   - **Columns**:
     | Column Name      | Data Type      | Description                              |
     |------------------|----------------|------------------------------------------|
     | `order_id`       | VARCHAR(255)   | Unique identifier for the order          |
     | `user_id`        | VARCHAR(255)   | Identifier for the user placing the order|
     | `product_id`     | VARCHAR(255)   | Identifier of the product ordered        |
     | `quantity`       | INT            | Number of units ordered                  |
     | `order_status`   | ENUM(pending, completed, failed) | Current status of the order   |
     | `payment_status` | ENUM(pending, paid, failed) | Payment status of the order   |
     | `created_at`     | TIMESTAMP      | Timestamp when the order was placed      |

   ```sql
   CREATE TABLE Orders (
       order_id VARCHAR(255) PRIMARY KEY,
       user_id VARCHAR(255),
       product_id VARCHAR(255),
       quantity INT,
       order_status ENUM('pending', 'completed', 'failed'),
       payment_status ENUM('pending', 'paid', 'failed'),
       created_at TIMESTAMP,
       FOREIGN KEY (product_id) REFERENCES Inventory(product_id)
   );
   ```

#### 3. **Table: Cart**
   - **Purpose**: Temporarily stores items added to the user's cart along with an expiration for the cart reservation.
   - **Columns**:
     | Column Name     | Data Type      | Description                              |
     |-----------------|----------------|------------------------------------------|
     | `cart_id`       | VARCHAR(255)   | Unique identifier for the cart           |
     | `user_id`       | VARCHAR(255)   | Identifier for the user                  |
     | `product_id`    | VARCHAR(255)   | Identifier for the product in the cart   |
     | `quantity`      | INT            | Number of units in the cart              |
     | `cart_expiry`   | TIMESTAMP      | Expiration time for the cart reservation |
   
   ```sql
   CREATE TABLE Cart (
       cart_id VARCHAR(255) PRIMARY KEY,
       user_id VARCHAR(255),
       product_id VARCHAR(255),
       quantity INT,
       cart_expiry TIMESTAMP,
       FOREIGN KEY (user_id) REFERENCES Users(user_id),
       FOREIGN KEY (product_id) REFERENCES Inventory(product_id)
   );
   ```

#### 4. **Table: Payment**
   - **Purpose**: Stores payment information linked to orders.
   - **Columns**:
     | Column Name       | Data Type      | Description                              |
     |-------------------|----------------|------------------------------------------|
     | `payment_id`       | VARCHAR(255)   | Unique identifier for the payment        |
     | `order_id`         | VARCHAR(255)   | The order associated with the payment    |
     | `payment_method`   | VARCHAR(50)    | Payment method used (credit card, PayPal)|
     | `amount`           | DECIMAL(10, 2) | Amount paid                              |
     | `payment_status`   | ENUM(pending, completed, failed) | Payment status |
     | `created_at`       | TIMESTAMP      | Timestamp when payment was processed     |

   ```sql
   CREATE TABLE Payment (
       payment_id VARCHAR(255) PRIMARY KEY,
       order_id VARCHAR(255),
       payment_method VARCHAR(50),
       amount DECIMAL(10, 2),
       payment_status ENUM('pending', 'completed', 'failed'),
       created_at TIMESTAMP,
       FOREIGN KEY (order_id) REFERENCES Orders(order_id)
   );
   ```

---

### Redis Integration:
- **Purpose**: To store frequently accessed data such as real-time inventory levels and user sessions for faster read/write operations.
   - **Inventory Cache**:
     - Key: `inventory:product_id`
     - Value: Available stock level for the product.

   - **Cart Data**:
     - Key: `cart:user_id`
     - Value: The list of items in the user’s cart along with expiration time.

---
