package types

import "fmt"

type SMSNotification struct {
}

func (sms *SMSNotification) SendNotification() {
	fmt.Println("Sending via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}
