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
