package busha_commerce_go

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gobuffalo/uuid"
	"strings"
	"time"
)

type EventService service

type Event struct {
	Id         string    `json:"id"`
	BusinessId string    `json:"business_id,omitempty"`
	Resource   string    `json:"resource"`
	CreatedAt  time.Time `json:"created_at"`
	Data       EventData `json:"data"`
}

type EventResponse struct {
	Response
	Data Event `json:"data"`
}

type EventData struct {
	Id            uuid.UUID        `json:"id"`
	BusinessId    uuid.UUID        `json:"business_id"`
	Reference     string           `json:"reference"`
	HostedUrl     string           `json:"hosted_url"`
	PriceFixed    bool             `json:"price_fixed"`
	LocalCurrency string           `json:"local_currency"`
	CallBackUrl   string           `json:"callback_url,omitempty"`
	Meta          json.RawMessage  `json:"meta"`
	ExpiresAt     time.Time        `json:"expires_at"`
	CreatedAt     time.Time        `json:"created_at"`
	Type          *string          `json:"type,omitempty"`
	ResourceID    *uuid.UUID       `json:"resource_id,omitempty"`
	TimeLine      []ChargeTimeline `json:"timeline"`
	Addresses     []ChargeAddress  `json:"addresses"`
	Payment       []ChargePayment  `json:"payments"`
	Pricing       []ChargePricing  `json:"pricing"`
}

type ListEventResponse struct {
	ResponseWithPagination
	Data []*Event `json:"events"`
}

func (s *EventService) List(params ListParameters) (*ListEventResponse, error) {
	var resp = new(ListEventResponse)
	err := s.client.call("GET", fmt.Sprintf("/events?sort=%s&limit=%d&page=%d",
		params.Sort, params.Limit, params.Page), params, &resp)
	return resp, err
}

func (s *EventService) Get(id string) (*EventResponse, error) {
	var resp = new(EventResponse)
	if id == "" {
		return nil, errors.New("no eventID provided")
	}
	err := s.client.call("GET", fmt.Sprintf("/events/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}
