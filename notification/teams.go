package notification

import (
	"log"

	User "test/user"
)

type Teams struct{}

func (e Teams) Send(u User.User) {
	log.Printf("Sending on Teams%v\n", u)
}
