package decorator

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (EmailNotifier) Send(message string) {
	fmt.Println("Email sent: ", message)
}

type SMSNotifier struct {
	Notifier
}

func (s SMSNotifier) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("SMS sent: ", message)
}
