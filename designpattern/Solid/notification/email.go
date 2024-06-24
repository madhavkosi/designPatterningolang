package notification

import (
	"log"
)

type Email struct{}

func (e Email) Send(u User) {
	log.Printf("Sending email%v\n", u)
}
