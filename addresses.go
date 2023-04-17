package busha_commerce_go

import (
	"errors"
	"fmt"
	"github.com/gobuffalo/uuid"
	"strings"
	"time"
)

type AddressService service

type Address struct {
	Id         uuid.UUID `json:"id"`
	BusinessID uuid.UUID `json:"business_id"`
	CurrencyId string    `json:"currency_id"`
	Chain      string    `json:"chain"`
	Address    string    `json:"address"`
	Memo       string    `json:"memo"`
	Label      string    `json:"label"`
	CreatedAt  time.Time `json:"created_at"`
}

type AddressResponse struct {
	Response
	Data Address `json:"data"`
}

type ListAddressesResponse struct {
	ResponseWithPagination
	Data []*Address `json:"data"`
}

func (s *AddressService) Create(req *AddressRequest) (*AddressResponse, error) {
	var resp = new(AddressResponse)
	err := s.client.call("POST", "/addresses", req, &resp)
	return resp, err
}

func (s *AddressService) List(params ListParameters) (*ListAddressesResponse, error) {
	var resp = new(ListAddressesResponse)
	err := s.client.call("GET", fmt.Sprintf("/addresses?sort=%s&limit=%d&page=%d&currency=%s",
		params.Sort, params.Limit, params.Page, params.Currency), params, &resp)
	return resp, err
}

func (s *AddressService) Get(id string) (*AddressResponse, error) {
	var resp = new(AddressResponse)
	if id == "" {
		return nil, errors.New("no addressID provided")
	}
	err := s.client.call("GET", fmt.Sprintf("/addresses/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}
