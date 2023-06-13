package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
)

func main() {
	stripe.Key = "sk_test_sample"

	stripeParams := &stripe.CustomerParams{
		Name:  stripe.String("sample"),
		Email: stripe.String("sample@gmail.com"),
	}

	customerID, err := customer.New(stripeParams)
	if err != nil {
		panic(err)
	}

	fmt.Println(customerID.ID)

}
