package notification

import (
	"log"
)

type Sms struct{}

func (s Sms) Send(u User) {
	log.Printf("Sending sms%v", u)
}
