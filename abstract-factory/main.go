package main

import "abstractfactory/notification"

// SMS
// Email

func main() {
	notificationFactory, err := notification.GetNotificationFactory("SMS")
	if err != nil {
		panic(err)
	}
	notification.SendNotification(notificationFactory)
}
