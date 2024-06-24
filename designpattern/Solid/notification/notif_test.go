package notification

import (
	"testing"
)

func TestNotificationService(t *testing.T) {
	email := Email{}
	sms := Sms{}
	teams := Teams{}
	a := NewNotificationService(email)
	a.Notification.Send(User{Name: "Madhav"})
	a.Notification = Teams{}
	a.Notification.Send(User{Name: "Madhav"})
	Notifications := []Notification{sms, email, teams}
	ProcessNotification(Notifications, User{Name: "Madhav"})
}
