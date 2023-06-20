package stripe

import (
	sStripe "github.com/hamed-amini-dev/stripe-go-scripts/services/stripe"
)

type Option func(*iStripe) error

// Init option service for getting service business logic account layer

func InitOptionService(service sStripe.IStripe) Option {
	return func(is *iStripe) error {
		is.service = service
		return nil
	}
}
