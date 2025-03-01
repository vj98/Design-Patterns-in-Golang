package main

import (
	"fmt"

	"github.com/vj98/decorator"
	"github.com/vj98/factory"
	"github.com/vj98/observer"
	"github.com/vj98/singleton"
	"github.com/vj98/strategy"
)

func main() {
	fmt.Println("Running Go Design Patterns")

	// Singleton Test
	s1 := singleton.GetInstance()
	s2 := singleton.GetInstance()
	fmt.Println("Singleton same instance:", s1 == s2)

	// Factory Test
	shape := factory.ShapeFactory("circle")
	shape.Draw()

	// Strategy Test
	payment := strategy.PaymentContext{}
	payment.SetStrategy(strategy.CreditCard{})
	payment.Pay(100.0)

	// Observer Test
	pub := observer.Publisher{}
	sub1 := &observer.Subscriber{Name: "Alice"}
	sub2 := &observer.Subscriber{Name: "Bob"}
	pub.Subscribe(sub1)
	pub.Subscribe(sub2)
	pub.Notify("New Event!")

	// Decorator Test
	email := decorator.EmailNotifier{}
	sms := decorator.SMSNotifier{Notifier: email}
	sms.Send("Hello!")
}
