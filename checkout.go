package busha_commerce_go

import (
	"errors"
	"fmt"
	"github.com/gobuffalo/uuid"
	"strings"
	"time"
)

type CheckoutType string

const (
	Donation   CheckoutType = "donation"
	FixedPrice CheckoutType = "fixed_price"
)

type CheckoutService service

type Checkout struct {
	Id            uuid.UUID    `json:"id"`
	BusinessId    string       `json:"business_id"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	CheckoutType  CheckoutType `json:"checkout_type"`
	RequestedInfo []string     `json:"requested_info"`
	LocalAmount   float64      `json:"local_amount,string"`
	LocalCurrency string       `json:"local_currency"`
	Active        bool         `json:"active"`
	CreatedAt     time.Time    `json:"created_at"`
}

type CheckoutResponse struct {
	Response
	Data Checkout `json:"data"`
}

func (s *CheckoutService) Create(req *CheckoutRequest) (*CheckoutResponse, error) {
	var resp = new(CheckoutResponse)
	err := s.client.call("POST", "/checkouts", req, &resp)
	return resp, err
}

type ListCheckoutResponse struct {
	ResponseWithPagination
	Data []*Checkout `json:"data"`
}

func (s *CheckoutService) List(params ListParameters) (*ListCheckoutResponse, error) {
	var resp = new(ListCheckoutResponse)
	err := s.client.call("GET", fmt.Sprintf("/checkouts?sort=%s&limit=%d&page=%d",
		params.Sort, params.Limit, params.Page), params, &resp)
	return resp, err
}

func (s *CheckoutService) Get(id string) (*CheckoutResponse, error) {
	var resp = new(CheckoutResponse)
	if id == "" {
		return nil, errors.New("no checkoutID provided")
	}
	err := s.client.call("GET", fmt.Sprintf("/checkouts/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *CheckoutService) Update(id string, req *CheckoutRequest) (*Response, error) {
	var resp = new(Response)
	if id == "" {
		return nil, errors.New("no checkoutID provided")
	}
	err := s.client.call("PUT", fmt.Sprintf("/checkouts/%s", strings.TrimSpace(id)), req, &resp)
	return resp, err
}

func (s *CheckoutService) ToggleStatus(id string) (*CheckoutResponse, error) {
	var resp = new(CheckoutResponse)
	if id == "" {
		return nil, errors.New("no checkoutID provided")
	}
	err := s.client.call("PATCH", fmt.Sprintf("/checkouts/%s/active", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *CheckoutService) Delete(id string) (*Response, error) {
	var resp = new(Response)
	if id == "" {
		return nil, errors.New("no checkoutID provided")
	}
	err := s.client.call("DELETE", fmt.Sprintf("/checkouts/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *CheckoutService) CreateCharge(id string, req *ChargeRequest) (*ChargeResponse, error) {
	var resp = new(ChargeResponse)
	if id == "" {
		return nil, errors.New("no checkoutID provided")
	}
	err := s.client.call("POST", fmt.Sprintf("/checkouts/%s/charge", strings.TrimSpace(id)), req, &resp)
	return resp, err
}
