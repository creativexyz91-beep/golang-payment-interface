package main

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(amount float64) error
}

type CreditCard struct {
	Holder string
}

type Stripe struct {
	Email string
}

type PayPal struct {
	Email string
}

func (c CreditCard) Pay(amount float64) error {
	fmt.Printf("EUR %.2f paid using Credit Card (%s)\n",
		amount, c.Holder)
	return nil
}

func (s Stripe) Pay(amount float64) error {
	fmt.Printf("EUR %.2f paid using Stripe (%s)\n",
		amount, s.Email)
	return nil
}

func (p PayPal) Pay(amount float64) error {
	fmt.Printf("EUR %.2f paid using PayPal (%s)\n",
		amount, p.Email)
	return nil
}

func ProcessPayment(p Payment, amount float64) error {

	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	return p.Pay(amount)
}

func main() {

	card := CreditCard{
		Holder: "Alexander",
	}

	stripe := Stripe{
		Email: "alexander@example.com",
	}

	paypal := PayPal{
		Email: "alexander@example.com",
	}

	fmt.Println("===== PAYMENT SYSTEM =====")

	err := ProcessPayment(card, 2500)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ProcessPayment(stripe, 1500)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ProcessPayment(paypal, 25)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = ProcessPayment(card, -100)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
