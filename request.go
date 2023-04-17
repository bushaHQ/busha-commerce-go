package busha_commerce_go

import (
	"encoding/json"
	"time"
)

type ListParameters struct {
	//Sort is the number of entries you want per page.
	Sort string `url:"sort,omitempty"`
	//Page is the page-number you want to get entries for
	Page int64 `url:"page,omitempty"`
	//Limit is the number of entries you want per page.
	Limit int64 `url:"limit,omitempty"`
	//Currency is the currency you want to get entries for
	Currency string `url:"currency,omitempty"`
}

type ChargeRequest struct {
	//FixedPrice has the value true if the charge price is fixed
	//or the value false if the charge price is not fixed
	FixedPrice bool `json:"fixed_price"`
	//LocalAmount amount in the currency to be charged
	LocalAmount int `json:"local_amount,string,omitempty"`
	//LocalAmount currency of the charge i.e, NGN
	LocalCurrency string `json:"local_currency,omitempty"`
	//Reference (optional) could be passed to create a charge with a
	//custom reference, should be between 5 and 100 characters.
	Reference *string `json:"reference,omitempty"`
	//Meta set of key:value that you can add to a charge
	//this can be useful for storing additional information to be retrieved later on a charge
	//i.e {"name": "Sarah Shaw", "email": "iL6jP@example.com"}
	Meta json.RawMessage `json:"meta,omitempty"`
	//SuccessURL (optional) is the URL that will be customer will be directed to once the charge is successful.
	SuccessRedirectURL *string `json:"success_redirect_url,omitempty"`
	//CancelURL (optional) is the URL that will be customer will be directed to once the charge is cancelled.
	CancelRedirectURL *string `json:"cancel_redirect_url,omitempty"`
}

type CheckoutRequest struct {
	//Name is the name of the checkout
	Name string `json:"name"`
	//Description is the description of the checkout
	Description string `json:"description"`
	//CheckoutType is the type of the checkout i.e. donation or fixed_price
	CheckoutType CheckoutType `json:"checkout_type"`
	//RequestedInfo is the requested information you'd want from the customer
	//i.e. ["name", "email","phone_number]
	RequestedInfo []string `json:"requested_info"`
	//LocalAmount amount in the currency to be charged
	LocalAmount float64 `json:"local_amount,string"`
	//LocalCurrency currency of the charge i.e, NGN
	LocalCurrency string `json:"local_currency"`
}

type InvoiceRequest struct {
	//Name is the name of the invoice
	Name string `json:"name"`
	//CustomerEmail is the email of the customer
	CustomerEmail string `json:"customer_email"`
	//LocalAmount amount in the currency to be charged
	LocalAmount float64 `json:"local_amount,string"`
	//LocalCurrency currency of the charge i.e, NGN
	LocalCurrency string `json:"local_currency"`
	//CustomerName is the name of the customer
	CustomerName string `json:"customer_name"`
	//Description is the description of the invoice
	Description string `json:"description"`
	//DueDate is the date by which the invoice will be void
	DueDate *time.Time `json:"due_date"`
}

type AddressRequest struct {
	//CurrencyID is the currency of the address i.e, BTC, ETH
	CurrencyId string `json:"currency_id"`
	//Chains is the list of chains that the address belongs to i.e []string{"TRX", "ETH"}
	Chains []string `json:"chains"`
	//Label is the description you want to give to the address
	Label string `json:"label,omitempty"`
}
