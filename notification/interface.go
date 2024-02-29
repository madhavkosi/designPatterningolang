package notification

import (
	User "test/user"
)

type Notification interface {
	Send(u User.User)
}
