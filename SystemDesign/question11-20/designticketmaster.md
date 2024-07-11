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

#### Tables and Relationships

1. **City**
   - `CityID` (PK)
   - `Name`
   - `State`
   - `Country`
   - `PostalCode`

2. **Cinema**
   - `CinemaID` (PK)
   - `Name`
   - `CityID` (FK)

3. **Hall**
   - `HallID` (PK)
   - `Name`
   - `CinemaID` (FK)

4. **Seat**
   - `SeatID` (PK)
   - `HallID` (FK)
   - `SeatNumber`
   - `Type`

5. **Movie**
   - `MovieID` (PK)
   - `Title`
   - `Description`
   - `Duration`
   - `Genre`
   - `Language`
   - `ReleaseDate`
   - `Country`

6. **Show**
   - `ShowID` (PK)
   - `MovieID` (FK)
   - `HallID` (FK)
   - `StartTime`
   - `EndTime`

7. **User**
   - `UserID` (PK)
   - `Name`
   - `Email`
   - `PhoneNumber`

8. **Booking**
   - `BookingID` (PK)
   - `ShowID` (FK)
   - `UserID` (FK)
   - `SeatID` (FK)
   - `BookingTime`
   - `Status`

![alt text](https://github.com/madhavkosi/designPatterningolang/blob/main/SystemDesign/image%20folder/databasedeisgn.svg)
