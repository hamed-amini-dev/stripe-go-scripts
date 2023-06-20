package routes

/*
A route is a section of Express code that associates an HTTP verb ( GET , POST , PUT , DELETE , etc.), a URL path/pattern
  - Name
  - Method
  - Path
  - Handler
*/

import (
	"net/http"

	"github.com/hamed-amini-dev/stripe-go-scripts/internal/handler/stripe"
	lhttp "github.com/hamed-amini-dev/stripe-go-scripts/internal/http"
)

// Define all end point and assign to the handling function
func StripeRoutes(th stripe.IStripe) []Route {
	return []Route{
		{
			Name:    "create customer",
			Method:  http.MethodPost,
			Path:    "/customer",
			Handler: lhttp.DefaultHTTPHandler(th.CreateCustomer),
		},
		{
			Name:    "add payment method",
			Method:  http.MethodPost,
			Path:    "/paymentmethods",
			Handler: lhttp.DefaultHTTPHandler(th.AddPaymentMethod),
		},
		{
			Name:    "create subscription",
			Method:  http.MethodPost,
			Path:    "/newsubscription",
			Handler: lhttp.DefaultHTTPHandler(th.CreateSubscription),
		},
		{
			Name:    "payment intent with confirm refund flow",
			Method:  http.MethodPost,
			Path:    "/pi_with_confirm_refund_flow",
			Handler: lhttp.DefaultHTTPHandler(th.PaymentIntentWithConfirmRefundFlow),
		},
		{
			Name:    "payment intent with cancell flow",
			Method:  http.MethodPost,
			Path:    "/pi_with_cancel_flow",
			Handler: lhttp.DefaultHTTPHandler(th.PaymentIntentWithCancelFlow),
		},
	}
}
