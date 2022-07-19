package main

import "abstractfactory/notification"

// SMS
// Email

func main() {
	notificationFactory, err := notification.GetNotificationFactory(notification.SMSNotifactionType)
	if err != nil {
		panic(err)
	}
	notification.SendNotification(notificationFactory)
}
