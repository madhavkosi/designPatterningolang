package notification

import (
	"log"

	User "test/user"
)

type Sms struct{}

func (s Sms) Send(u User.User) {
	log.Printf("Sending sms%v",u)
}
