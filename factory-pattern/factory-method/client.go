package main

import (
	"design-pattern/factory-pattern/factory-method/notification"
	"fmt"
)

func sendNotification(factory notification.NotifierFactory, message string) {
	notifier := factory.Create()
	fmt.Println(notifier.Send(message))
}
func main() {
	sendNotification(&notification.EmailFactory{}, "Hello via Email")
	sendNotification(&notification.SMSFactory{}, "Hello via SMS")

}
