package notification

import "fmt"

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) string {
	return fmt.Sprintf("Sending SMS: %s", message)
}

type SMSFactory struct{}

func (f *SMSFactory) Create() Notifier {
	return &SMSNotifier{}
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) string {
	return fmt.Sprintf("Sending email: %s", message)
}

type EmailFactory struct{}

func (f *EmailFactory) Create() Notifier {
	return &EmailNotifier{}
}
