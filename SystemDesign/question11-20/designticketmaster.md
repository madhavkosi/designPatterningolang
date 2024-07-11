### Online Movie Ticket Booking System Design

#### 1. What is an Online Movie Ticket Booking System?
An online movie ticket booking system allows customers to purchase theater seats online. It enables customers to browse through currently playing movies and book seats anytime, anywhere.

#### 2. Requirements and Goals of the System

**Functional Requirements:**
- **City Listing:** List different cities with affiliate cinemas.
- **Movie Display:** Display movies released in the selected city.
- **Cinema and Showtimes:** Show cinemas running the selected movie and their available showtimes.
- **Seat Selection:** Allow users to choose a show at a cinema and book tickets, displaying the seating arrangement and letting them select multiple seats.
- **Seat Availability:** Distinguish between available and booked seats.
- **Seat Hold:** Allow users to hold seats for five minutes before payment.
- **Waiting System:** Enable users to wait for seats to become available, servicing them in a first-come, first-serve manner.

**Non-Functional Requirements:**
- **Concurrency:** Handle multiple booking requests for the same seat gracefully and fairly.
- **Security:** Ensure secure financial transactions and ACID-compliant database.
  
#### 3. Design Considerations
- **No User Authentication:** For simplicity, no user authentication is required.
- **All-or-Nothing Orders:** Either fulfill the entire ticket order or none.
- **Fairness:** Ensure fairness in seat booking.
- **Booking Limit:** Restrict users from booking more than ten seats at a time.
- **Scalability:** Handle traffic spikes during popular movie releases, ensuring the system is scalable and highly available.

#### 4. Capacity Estimation

**Traffic Estimates:**
- 3 billion page views per month.
- 10 million tickets sold per month.

**Storage Estimates:**
- **Cities and Cinemas:** 500 cities, 10 cinemas per city.
- **Seats and Shows:** 2000 seats per cinema, 2 shows per day.
- **Data Storage:** 50 bytes per seat booking, 50 bytes for movie and cinema information.

Calculation for daily data storage:
500 cities * 10 cinemas * 2000 seats * 2 shows * (50 + 50) bytes = 2 GB/day

To store five years of this data:
2 GB/day * 365 days/year * 5 years = 3.6 TB 


### API Definitions for Movie Ticket Booking System

#### 1. SearchMovies API
The `SearchMovies` API allows users to search for movie shows based on various filters. This API can be useful for finding specific movies or shows in a particular location and timeframe.

**Function Signature**: `SearchMovies(api_dev_key, keyword, city, lat_long, radius, start_datetime, end_datetime, postal_code, includeSpellcheck, results_per_page, sorting_order)`

**Parameters**:
- `api_dev_key` (string): The API developer key of a registered account. Used for authentication and quota management.
- `keyword` (string): Keyword to search for movies.
- `city` (string): City to filter movies.
- `lat_long` (string): Latitude and longitude to filter the search area.
- `radius` (number): Radius (in kilometers) around the specified lat_long to search within.
- `start_datetime` (string): Filter movies starting from this datetime.
- `end_datetime` (string): Filter movies ending at this datetime.
- `postal_code` (string): Postal code to filter movies.
- `includeSpellcheck` (Enum: "yes" or "no"): Include spell check suggestions if "yes".
- `results_per_page` (number): Number of results per page, maximum is 30.
- `sorting_order` (string): Sorting order of the results. Acceptable values include 'name,asc', 'name,desc', 'date,asc', 'date,desc', 'distance,asc', 'name,date,asc', 'name,date,desc', 'date,name,asc', 'date,name,desc'.

**Returns**:
A JSON object containing a list of movies and their shows. Each movie show includes details such as MovieID, ShowID, Title, Description, Duration, Genre, Language, ReleaseDate, Country, StartTime, EndTime, and a list of available seats with their types, prices, and status.

**Sample Response**:
```json
[
  {
    "MovieID": 1,
    "ShowID": 1,
    "Title": "Cars 2",
    "Description": "About cars",
    "Duration": 120,
    "Genre": "Animation",
    "Language": "English",
    "ReleaseDate": "8th Oct. 2014",
    "Country": "USA",
    "StartTime": "14:00",
    "EndTime": "16:00",
    "Seats": [
      {
        "Type": "Regular",
        "Price": 14.99,
        "Status": "Almost Full"
      },
      {
        "Type": "Premium",
        "Price": 24.99,
        "Status": "Available"
      }
    ]
  },
  {
    "MovieID": 1,
    "ShowID": 2,
    "Title": "Cars 2",
    "Description": "About cars",
    "Duration": 120,
    "Genre": "Animation",
    "Language": "English",
    "ReleaseDate": "8th Oct. 2014",
    "Country": "USA",
    "StartTime": "16:30",
    "EndTime": "18:30",
    "Seats": [
      {
        "Type": "Regular",
        "Price": 14.99,
        "Status": "Full"
      },
      {
        "Type": "Premium",
        "Price": 24.99,
        "Status": "Almost Full"
      }
    ]
  }
]
```

#### 2. ReserveSeats API
The `ReserveSeats` API allows users to reserve seats for a specific movie show.

**Function Signature**: `ReserveSeats(api_dev_key, session_id, movie_id, show_id, seats_to_reserve[])`

**Parameters**:
- `api_dev_key` (string): The API developer key of a registered account. Used for authentication and quota management.
- `session_id` (string): User's session ID to track the reservation. The reservation will be removed when the session expires.
- `movie_id` (string): ID of the movie to reserve.
- `show_id` (string): ID of the show to reserve.
- `seats_to_reserve` (array of numbers): An array containing seat IDs to reserve.

**Returns**:
A JSON object indicating the status of the reservation. Possible statuses include:
1. "Reservation Successful"
2. "Reservation Failed - Show Full"
3. "Reservation Failed - Retry, as other users are holding reserved seats"

**Sample Response**:
```json
{
  "status": "Reservation Successful"
}
```


### Database Design for Movie Ticket Booking System

| Entity  | Description                                          | Relationships                                       |
|---------|------------------------------------------------------|-----------------------------------------------------|
| **City** | Each city can have multiple cinemas.                 | One-to-Many with Cinema                             |
| **Cinema** | Each cinema belongs to one city.<br>Each cinema can have multiple halls. | Many-to-One with City<br>One-to-Many with Hall |
| **Hall** | Each hall belongs to one cinema.<br>Each hall can host multiple shows. | Many-to-One with Cinema<br>One-to-Many with Show |
| **Movie** | Each movie can have multiple shows.                  | One-to-Many with Show                               |
| **Show** | Each show belongs to one hall and is for one specific movie.<br>Each show can have multiple bookings. | Many-to-One with Hall<br>Many-to-One with Movie<br>One-to-Many with Booking |
| **User** | Each user can have multiple bookings.                 | One-to-Many with Booking                            |
| **Booking** | Each booking belongs to one show and one user.     | Many-to-One with Show<br>Many-to-One with User      |
| **Seat** | Each seat belongs to a hall.<br>Each seat can be associated with multiple bookings over time. | Many-to-One with Hall                               |

This table outlines the entities, their descriptions, and the relationships they have with other entities in the movie ticket booking system.

#### Entity-Relationship Diagram (ERD)

```plaintext
City --< Cinema --< Hall --< Show --< Booking >-- User
                      |                 |
                    Seat              Movie
```
### Database Design for Movie Ticket Booking System

| Table       | Columns                                                        |
|-------------|----------------------------------------------------------------|
| **City**    | CityID (PK), Name, State, Country, PostalCode                  |
| **Cinema**  | CinemaID (PK), Name, CityID (FK)                               |
| **Hall**    | HallID (PK), Name, CinemaID (FK)                               |
| **Seat**    | SeatID (PK), HallID (FK), SeatNumber, Type                     |
| **Movie**   | MovieID (PK), Title, Description, Duration, Genre, Language, ReleaseDate, Country |
| **Show**    | ShowID (PK), MovieID (FK), HallID (FK), StartTime, EndTime     |
| **User**    | UserID (PK), Name, Email, PhoneNumber                          |
| **Booking** | BookingID (PK), ShowID (FK), UserID (FK), SeatID (FK), BookingTime, Status |


![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/databasedeisgn.svg)


### Detailed Component Design

#### Ticket Booking Workflow

1. **User Searches for a Movie**: User inputs movie search criteria.
2. **User Selects a Movie**: User chooses a movie from the search results.
3. **Available Shows Displayed**: User views showtimes for the selected movie.
4. **User Selects a Show**: User picks a specific showtime.
5. **User Selects Number of Seats**: User specifies the number of seats to reserve.
6. **Seat Availability Check**:
   - **If Seats Available**: User is shown a theater map to select seats.
   - **If Seats Not Available**: User is taken to step 8.
7. **Seat Reservation Attempt**:
   - **Successful Reservation**: Seats are reserved, and user proceeds to payment.
   - **Failed Reservation**: User sees an error message or is taken back to the theater map to choose different seats.
8. **Waiting for Seats**:
   - **Seats Become Available**: User is notified to select seats.
   - **Seats Not Available**: User is shown an error message or taken back to the movie search page.
   - **Session Timeout**: After one hour, user is taken back to the movie search page.
9. **Payment Process**:
   - **Successful Payment**: Booking is completed.
   - **Failed Payment**: Reserved seats are released after five minutes.

### ActiveReservationsService

#### Overview:
- **Purpose**: Manage reservations of a 'show' in memory and database.
- **Memory Storage**: Utilizes a data structure similar to Linked HashMap or TreeMap.
- **Database Storage**: Reservations stored in the 'Booking' table.

#### Key Components:
1. **In-Memory Data Structure**:
   - **Type**: Linked HashMap.
   - **Function**: Allows direct access to any reservation to remove it when booking is complete.
   - **Expiry Management**: Head always points to the oldest reservation for expiry management.

2. **HashTable**:
   - **Key**: ShowID.
   - **Value**: Linked HashMap containing BookingID and creation Timestamp.

3. **Database**:
   - **Table**: Booking.
   - **Fields**:
     - **Timestamp**: Stores the expiry time.
     - **Status**:
       - Reserved (1)
       - Booked (2)
       - Expired (3)
   - **Operations**:
     - Update status to Booked (2) upon completion.
     - Remove or mark as Expired (3) when reservation expires.

#### Workflow:
1. **Reservation Management**:
   - **Creation**: New reservations added to Linked HashMap and database.
   - **Completion**: Updates status to Booked (2) in the database and removes from Linked HashMap.
   - **Expiry**: Removes expired reservations from Linked HashMap and either deletes or marks as Expired (3) in the database.

2. **Payment Processing**:
   - **Integration**: Works with an external financial service to process payments upon booking completion or reservation expiry.

3. **WaitingUsersService**:
   - **Notification**: Signals WaitingUsersService to serve waiting customers when a reservation is completed or expired.


### WaitingUsersService

#### Overview:
- **Purpose**: Manage waiting users of a 'show' in memory and ensure fair service based on waiting time.
- **Memory Storage**: Utilizes a data structure similar to Linked HashMap or TreeMap.
- **Database Storage**: Not explicitly mentioned, but implied to work in conjunction with ActiveReservationsService.

#### Key Components:
1. **In-Memory Data Structure**:
   - **Type**: Linked HashMap.
   - **Function**: Allows direct access to any waiting user to remove them upon cancellation.
   - **Fair Service**: Head always points to the longest waiting user for fair first-come-first-serve handling when seats become available.

2. **HashTable**:
   - **Key**: ShowID.
   - **Value**: Linked HashMap containing UserIDs and their wait-start-time.

#### Workflow:
1. **Waiting List Management**:
   - **Addition**: New waiting users added to the Linked HashMap and HashTable.
   - **Cancellation**: Users can cancel their request, and they are removed from the Linked HashMap.
   - **Notification**: When seats become available, the head of the Linked HashMap (longest waiting user) is served first.

2. **Client Updates**:
   - **Long Polling**: Clients use long polling to stay updated on their reservation status.
   - **Notification**: Server uses long polling requests to notify users when seats become available.

3. **Reservation Expiration**:
   - **Tracking**: ActiveReservationsService tracks the expiry of active reservations based on reservation time.
   - **Client Timer**: Clients see a timer for expiration time, potentially out of sync with the server.
   - **Buffer**: A five-second buffer is added on the server to prevent timing out on the client side before the server, ensuring a seamless experience and successful purchase.

#### Summary:
- **WaitingUsersService** maintains a fair and efficient system for managing waiting users using an in-memory Linked HashMap.
- Ensures that the longest waiting users are served first when seats become available.
- Integrates with client-side long polling for real-time updates.
- Collaborates with ActiveReservationsService for managing reservation expirations, incorporating a buffer to ensure a smooth user experience.


### Lecture Notes: Concurrency Handling in SQL Databases

#### Key Topic: Concurrency
- **Objective**: Prevent multiple users from booking the same seat simultaneously.

#### Handling Concurrency Using Transactions
- **SQL Transactions**:
  - Ensure operations are completed successfully without interference.
  - Use transactions to manage concurrent seat bookings.

#### Transaction Isolation Levels
- **Serializable Isolation Level**:
  - Highest level of isolation.
  - Prevents Dirty Reads, Nonrepeatable Reads, and Phantom Reads.
  - Ensures that once rows are read in a transaction, they are locked for writing.

#### Sample SQL Code
```sql
SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;

BEGIN TRANSACTION;

    -- Intent: Reserve seats 54, 55, 56 for ShowID 99
    SELECT * FROM Show_Seat 
    WHERE ShowID=99 AND ShowSeatID IN (54, 55, 56) AND Status=0; -- Check if seats are free

    -- If the above returns three rows, proceed with update
    -- Otherwise, return failure

    UPDATE Show_Seat ...
    UPDATE Booking ...

COMMIT TRANSACTION;
```


#### Key Points
- **Serializable Isolation**:
  - Ensures safety from concurrent modifications.
  - Locks rows for writing once read within a transaction.
  
- **Locking Mechanism**:
  - Within a transaction, reading rows places a write lock on them.
  - Prevents updates by other transactions.

#### Post-Transaction Process
- **ActiveReservationService**:
  - Once the transaction is successful, track the reservation using `ActiveReservationService`.

#### Summary
- Use transactions with the `Serializable` isolation level to handle concurrency.
- Ensure seat availability before updating the booking records.
- Lock rows during transactions to prevent concurrent modifications.
- Track successful reservations with an active service.
