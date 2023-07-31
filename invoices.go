package busha_commerce_go

import (
	"errors"
	"fmt"
	"github.com/gobuffalo/uuid"
	"strings"
	"time"
)

type InvoiceService service

type Invoice struct {
	Id            uuid.UUID  `json:"id"`
	BusinessId    uuid.UUID  `json:"business_id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	CustomerName  string     `json:"customer_name"`
	CustomerEmail string     `json:"customer_email"`
	LocalAmount   float64    `json:"local_amount,string"`
	LocalCurrency string     `json:"local_currency"`
	Status        string     `json:"status"`
	Reference     string     `json:"reference"`
	CreatedAt     time.Time  `json:"created_at"`
	DueDate       *time.Time `json:"due_date"`
}

type InvoiceResponse struct {
	Response
	Data Invoice `json:"data"`
}

func (s *InvoiceService) Create(req *InvoiceRequest) (*InvoiceResponse, error) {
	var resp = new(InvoiceResponse)
	err := s.client.call("POST", "/invoices", req, &resp)
	return resp, err
}

type ListInvoiceResponse struct {
	ResponseWithPagination
	Data []*Invoice `json:"data"`
}

func (s *InvoiceService) List(params ListParameters) (*ListPaymentLinksResponse, error) {
	var resp = new(ListPaymentLinksResponse)
	err := s.client.call("GET", fmt.Sprintf("/invoices?sort=%s&limit=%d&page=%d",
		params.Sort, params.Limit, params.Page), params, &resp)
	return resp, err
}

func (s *InvoiceService) Get(id string) (*InvoiceResponse, error) {
	var resp = new(InvoiceResponse)
	if id == "" {
		return nil, errors.New("no invoiceID provided")
	}
	err := s.client.call("GET", fmt.Sprintf("/invoices/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *InvoiceService) Void(id string) (*Response, error) {
	var resp = new(Response)
	if id == "" {
		return nil, errors.New("no invoiceID provided")
	}
	err := s.client.call("DELETE", fmt.Sprintf("/invoices/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *InvoiceService) CreateCharge(id string) (*ChargeResponse, error) {
	var resp, req = new(ChargeResponse), struct{}{}
	if id == "" {
		return nil, errors.New("no invoiceID provided")
	}
	err := s.client.call("POST", fmt.Sprintf("/invoices/%s/charge", strings.TrimSpace(id)), req, &resp)
	return resp, err
}
