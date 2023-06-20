package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
)

func CreateCustomer(name, email string) (string, error) {
	stripeParams := &stripe.CustomerParams{
		Name:  stripe.String(name),
		Email: stripe.String(email),
	}

	customerID, err := customer.New(stripeParams)
	if err != nil {
		return "", err
	}

	return customerID.ID, nil

}
