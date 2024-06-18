package notification

import User "test/user"

type NotificationService struct {
	Notification Notification
}

func NewNotificationService(n Notification) *NotificationService {
	return &NotificationService{Notification: n}
}
func ProcessNotification(Notifications []Notification, u User.User) {
	for _, n := range Notifications {
		n.Send(u)
	}
}
