package busha_commerce_go

import (
	"errors"
	"fmt"
	"github.com/gobuffalo/uuid"
	"strings"
	"time"
)

type PaymentLinkType string

const (
	Donation   PaymentLinkType = "donation"
	FixedPrice PaymentLinkType = "fixed_price"
)

type PaymentLinkService service

type PaymentLink struct {
	Id              uuid.UUID       `json:"id"`
	BusinessId      string          `json:"business_id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	PaymentLinkType PaymentLinkType `json:"payment_link_type"`
	RequestedInfo   []string        `json:"requested_info"`
	LocalAmount     float64         `json:"local_amount,string"`
	LocalCurrency   string          `json:"local_currency"`
	Active          bool            `json:"active"`
	CreatedAt       time.Time       `json:"created_at"`
}

type PaymentLinkResponse struct {
	Response
	Data PaymentLink `json:"data"`
}

func (s *PaymentLinkService) Create(req *PaymentLinkRequest) (*PaymentLinkResponse, error) {
	var resp = new(PaymentLinkResponse)
	err := s.client.call("POST", "/payment_links", req, &resp)
	return resp, err
}

type ListPaymentLinksResponse struct {
	ResponseWithPagination
	Data []*PaymentLink `json:"data"`
}

func (s *PaymentLinkService) List(params ListParameters) (*ListPaymentLinksResponse, error) {
	var resp = new(ListPaymentLinksResponse)
	err := s.client.call("GET", fmt.Sprintf("/payment_links?sort=%s&limit=%d&page=%d",
		params.Sort, params.Limit, params.Page), params, &resp)
	return resp, err
}

func (s *PaymentLinkService) Get(id string) (*PaymentLinkResponse, error) {
	var resp = new(PaymentLinkResponse)
	if id == "" {
		return nil, errors.New("no payment link ID provided")
	}
	err := s.client.call("GET", fmt.Sprintf("/payment_links/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *PaymentLinkService) Update(id string, req *PaymentLinkRequest) (*Response, error) {
	var resp = new(Response)
	if id == "" {
		return nil, errors.New("no payment link ID provided")
	}
	err := s.client.call("PUT", fmt.Sprintf("/payment_links/%s", strings.TrimSpace(id)), req, &resp)
	return resp, err
}

func (s *PaymentLinkService) ToggleStatus(id string) (*PaymentLinkResponse, error) {
	var resp = new(PaymentLinkResponse)
	if id == "" {
		return nil, errors.New("no payment link ID provided")
	}
	err := s.client.call("PATCH", fmt.Sprintf("/payment_links/%s/active", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *PaymentLinkService) Delete(id string) (*Response, error) {
	var resp = new(Response)
	if id == "" {
		return nil, errors.New("no payment link ID provided")
	}
	err := s.client.call("DELETE", fmt.Sprintf("/payment_links/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *PaymentLinkService) CreateCharge(id string, req *ChargeRequest) (*ChargeResponse, error) {
	var resp = new(ChargeResponse)
	if id == "" {
		return nil, errors.New("no payment link ID provided")
	}
	err := s.client.call("POST", fmt.Sprintf("/payment_links/%s/charge", strings.TrimSpace(id)), req, &resp)
	return resp, err
}
