# golang-payment-interface
A payment processing simulation built with Go, demonstrating interfaces, structs, methods, polymorphism, and error handling using Credit Card, Stripe, and PayPal payment methods.
# Go Payment Interface System

A beginner-friendly payment processing simulation developed using **Golang** to demonstrate the practical implementation of interfaces and polymorphism.

## Overview

This project demonstrates how Go interfaces allow multiple payment methods to share a common behavior through a single payment processing function.

It simulates three payment methods: Credit Card, Stripe, and PayPal.

## Features

- Multiple payment methods
- Interface implementation
- Structs and methods
- Polymorphism
- Error handling
- Payment amount validation
- Reusable payment processing function

## Technologies Used

- Go (Golang)
- fmt package
- errors package

## Payment Methods

1. Credit Card
2. Stripe
3. PayPal

## Example Output

```text
===== PAYMENT SYSTEM =====

EUR 2500.00 paid using Credit Card (Alexander)
EUR 1500.00 paid using Stripe (alexander@example.com)
EUR 25.00 paid using PayPal (alexander@example.com)

Error: amount must be greater than zero
```

## How to Run

Make sure Go is installed on your computer.

Clone the repository:

```bash
git clone YOUR_REPOSITORY_URL
```

Navigate to the project directory:

```bash
cd golang-payment-interface
```

Run the program:

```bash
go run main.go
```

## Concepts Learned

- Go Interfaces
- Structs
- Methods
- Implicit Interface Implementation
- Polymorphism
- Functions
- Error Handling

## Learning Purpose

This project was created to strengthen my understanding of Go programming and explore how interfaces help build reusable, flexible, and maintainable software.

## Disclaimer

This is an educational payment simulation. It does not connect to actual payment gateways or process real financial transactions.
