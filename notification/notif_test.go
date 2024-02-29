package notification

import (
	"test/user"
	"testing"
)

func TestNotificationService(t *testing.T) {
	email := Email{}
	sms := Sms{}
	teams := Teams{}
	a := NewNotificationService(email)
	a.Notification.Send(user.User{Name: "Madhav"})
	a.Notification = Teams{}
	a.Notification.Send(user.User{Name: "Madhav"})
	Notifications := []Notification{sms, email, teams}
	ProcessNotification(Notifications, user.User{Name: "Madhav"})
}
