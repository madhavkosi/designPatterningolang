package notification

import User "test/user"

type NotificationService struct {
	Notification Notification
}

func NewNotificationService(n Notification) *NotificationService {
	return &NotificationService{Notification: n}
}
func ProcessNotification(Notifications []Notification, u User.User) {
	for _, val := range Notifications {
		val.Send(u)
	}
}
