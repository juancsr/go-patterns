package notification

import (
	. "abstractfactory/notification/types"
	"errors"
	"fmt"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	switch notificationType {
	case "SMS":
		return &SMSNotification{}, nil
	case "Email":
		return &EmailNotification{}, nil
	default:
		return nil, errors.New("Unsuported notification type")
	}
}

func SendNotification(factory INotificationFactory) {
	factory.SendNotification()
}

func GetMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}
