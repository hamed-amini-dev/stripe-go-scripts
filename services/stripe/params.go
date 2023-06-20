package stripe

type CreateCustomerParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AddPaymentMethodParams struct {
	CustomerID string `json:"customerID"`
	Number     string `json:"number"`
	Cvc        string `json:"cvc"`
	ExpMonth   int64  `json:"expMonth"`
	ExpYear    int64  `json:"expYear"`
}

type CreateSubscriptionParams struct {
	CustomerID      string `json:"customerID"`
	PromotionCode   string `json:"promotionCode"`
	Price           string `json:"price"`
	TrialPeriodDays int64  `json:"trialPeriodDays"`
	Qty             int64  `json:"qty"`
	Payout          bool   `json:"payout"`
}

type PaymentIntentWithConfirmRefundFlowParams struct {
	CustomerID  string `json:"customerID"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}

type PaymentIntentWithCancellFlowParams struct {
	CustomerID  string `json:"customerID"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
}
