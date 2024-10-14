package main

import (
	"fmt"
	"time"
)

type User struct {
	Name      string
	Email     string
	Scheduler *MeetingScheduler
}

type Interval struct {
	StartTime time.Time
	EndTime   time.Time
}

func (i *Interval) Overlaps(other Interval) bool {
	return i.StartTime.Before(other.EndTime) && i.EndTime.After(other.StartTime)
}

type MeetingRoom struct {
	ID             int
	Capacity       int
	BookedIntervals []Interval
}

func (room *MeetingRoom) isAvailable(interval Interval) bool {
	for _, booked := range room.BookedIntervals {
		if booked.Overlaps(interval) {
			return false
		}
	}
	return true
}

func (room *MeetingRoom) bookRoom(interval Interval) bool {
	if room.isAvailable(interval) {
		room.BookedIntervals = append(room.BookedIntervals, interval)
		return true
	}
	return false
}

type Meeting struct {
	ID          int
	Participants []User
	Interval    Interval
	Room        *MeetingRoom
	Subject     string
}

type Calendar struct {
	Meetings []Meeting
}

func (c *Calendar) addMeeting(meeting Meeting) {
	c.Meetings = append(c.Meetings, meeting)
}

type MeetingScheduler struct {
	Organizer User
	Calendar  Calendar
	Rooms     []*MeetingRoom
}

func (scheduler *MeetingScheduler) checkRoomAvailability(numberOfPersons int, interval Interval) *MeetingRoom {
	for _, room := range scheduler.Rooms {
		if room.Capacity >= numberOfPersons && room.isAvailable(interval) {
			return room
		}
	}
	return nil
}

func (scheduler *MeetingScheduler) scheduleMeeting(users []User, interval Interval, subject string, numberOfPersons int) bool {
	room := scheduler.checkRoomAvailability(numberOfPersons, interval)
	if room == nil {
		return false
	}

	meeting := Meeting{
		ID:          len(scheduler.Calendar.Meetings) + 1,
		Participants: users,
		Interval:    interval,
		Room:        room,
		Subject:     subject,
	}

	scheduler.Calendar.addMeeting(meeting)
	room.bookRoom(interval)
	return true
}

func main() {
	// Create some users
	user1 := User{Name: "Alice", Email: "alice@example.com"}
	user2 := User{Name: "Bob", Email: "bob@example.com"}

	// Create some meeting rooms
	room1 := &MeetingRoom{ID: 1, Capacity: 5, BookedIntervals: []Interval{}}
	room2 := &MeetingRoom{ID: 2, Capacity: 10, BookedIntervals: []Interval{}}

	// Create a meeting scheduler
	scheduler := MeetingScheduler{
		Organizer: user1,
		Rooms:     []*MeetingRoom{room1, room2},
	}

	// Users involved in the meeting
	users := []User{user1, user2}

	// Define an interval for the meeting
	startTime := time.Now().Add(time.Hour)
	endTime := startTime.Add(time.Hour)
	interval := Interval{StartTime: startTime, EndTime: endTime}

	// Schedule the meeting
	success := scheduler.scheduleMeeting(users, interval, "Project Discussion", 4)

	if success {
		fmt.Println("Meeting scheduled successfully!")
	} else {
		fmt.Println("Failed to schedule the meeting.")
	}
}
