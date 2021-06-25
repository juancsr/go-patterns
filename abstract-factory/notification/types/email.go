package types

import "fmt"

type EmailNotification struct {
}

func (*EmailNotification) SendNotification() {
	fmt.Println("Sending via Email")
}

func (*EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "Protonmail"
}
