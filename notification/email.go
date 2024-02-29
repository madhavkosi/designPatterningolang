package notification

import (
	"log"

	User "test/user"
)

type Email struct{}

func (e Email) Send(u User.User) {
	log.Printf("Sending email%v\n",u)
}
