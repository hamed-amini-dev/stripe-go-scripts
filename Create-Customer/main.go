package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
)

func main() {
	stripe.Key = "sk_test_51LfLkJLdKtGMf7ZsI7yvCPZizcWoRxYSqN49SCsqzuQv3UVwu7t4hrvOS4CSakkpZPnKkV3sjICOENjK0MEZWAUT00u7sKul49"

	stripeParams := &stripe.CustomerParams{
		Name:  stripe.String("hamed"),
		Email: stripe.String("hamed@hamed.com"),
	}

	customerID, err := customer.New(stripeParams)
	if err != nil {
		panic(err)
	}

	fmt.Println(customerID.ID)
	//cus_O3WsRj7vAEkKyA

}
