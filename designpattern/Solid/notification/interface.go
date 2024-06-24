package notification

type Notification interface {
	Send(User)
}
