package strategy

import (
	"fmt"
)

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCard struct{}

func (CreditCard) Pay(amount float64) {
	fmt.Printf("Paid $%.2f with Credit Card\n", amount)
}

type PayPal struct{}

func (PayPal) Pay(amount float64) {
	fmt.Printf("Paid $%.2f with PayPal\n", amount)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

func (p *PaymentContext) Pay(amount float64) {
	p.strategy.Pay(amount)
}
