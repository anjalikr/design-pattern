package main

import (
	"errors"
	"fmt"
	"log"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCard struct{}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card\n", amount)
}

type PayPal struct{}

func (p PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal", amount)
}

func NewPaymentFactory(method string) (PaymentMethod, error) {
	switch method {
	case "credit":
		return &CreditCard{}, nil
	case "paypal":
		return &PayPal{}, nil
	default:
		return nil, errors.New("unsupported payment method")
	}
}

func main() {
	method, err := NewPaymentFactory("credit")
	if err != nil {
		log.Fatal(err)
	}

	result := method.Pay(15.5)
	fmt.Println(result) //Paid 15.50 using Credit Card
}
