package notification

import (
	. "abstractfactory/notification/types"
	"errors"
	"fmt"
)

type notificationType string

const (
	SMSNotifactionType    notificationType = "SMS"
	EmailNotificationType notificationType = "Email"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

func GetNotificationFactory(notificationType notificationType) (INotificationFactory, error) {
	switch notificationType {
	case SMSNotifactionType:
		return &SMSNotification{}, nil
	case EmailNotificationType:
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
