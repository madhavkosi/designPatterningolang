package notification

import (
	"log"
)

type Teams struct{}

func (e Teams) Send(u User) {
	log.Printf("Sending on Teams%v\n", u)
}
