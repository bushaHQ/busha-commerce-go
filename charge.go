package busha_commerce_go

import (
	"errors"
	"fmt"
	"github.com/gobuffalo/uuid"
	"strings"
	"time"
)

type ChargeService service

type Charge struct {
	Id               uuid.UUID               `json:"id"`
	BusinessId       uuid.UUID               `json:"business_id"`
	BusinessName     string                  `json:"business_name"`
	Reference        string                  `json:"reference"`
	HostedUrl        string                  `json:"hosted_url"`
	PriceFixed       bool                    `json:"price_fixed"`
	Meta             map[string]interface{}  `json:"meta"`
	ExpiresAt        time.Time               `json:"expires_at"`
	Timeline         []ChargeTimeline        `json:"timeline"`
	SupportedAssets  []ChargeSupportedAssets `json:"supported_assets"`
	PaymentThreshold PaymentThreshold        `json:"payment_threshold"`
	Payments         []ChargePayment         `json:"payments"`
	Pricing          []ChargePricing         `json:"pricing"`
	Addresses        []ChargeAddress         `json:"addresses"`
	CallbackUrl      string                  `json:"callback_url"`
	LocalAmount      float64                 `json:"local_amount,string"`
	LocalCurrency    string                  `json:"local_currency"`
}

type ChargeAddress struct {
	CurrencyId string `json:"currency_id"`
	Chain      string `json:"chain"`
	Address    string `json:"address"`
	Memo       string `json:"memo"`
	Label      string `json:"label"`
}

type ChargePricing struct {
	CurrencyId string  `json:"currency_id"`
	Amount     float64 `json:"amount,string"`
	Rate       float64 `json:"rate,string"`
	IsLocal    bool    `json:"is_local"`
}

type ChargePayment struct {
	Chain            string             `json:"chain"`
	LocalAmount      float64            `json:"local_amount,string"`
	LocalCurrency    string             `json:"local_currency"`
	Amount           float64            `json:"amount,string"`
	Currency         string             `json:"currency"`
	TransactionId    string             `json:"transaction_id"`
	TransactionHash  string             `json:"transaction_hash"`
	Reference        string             `json:"reference"`
	Status           string             `json:"status"`
	Traded           bool               `json:"traded"`
	Address          string             `json:"address"`
	Confirmation     int                `json:"confirmation"`
	PaymentThreshold []PaymentThreshold `json:"payment_threshold"`
	BlockUrl         string             `json:"block_url"`
	Internal         bool               `json:"internal"`
}

type ChargeTimeline struct {
	Status    string    `json:"status"`
	Context   string    `json:"context"`
	CreatedAt time.Time `json:"created_at"`
}

type PaymentThreshold struct {
	OverpaymentAbsoluteThreshold  float64 `json:"overpayment_absolute_threshold,string"`
	OverpaymentRelativeThreshold  float64 `json:"overpayment_relative_threshold,string"`
	UnderpaymentAbsoluteThreshold float64 `json:"underpayment_absolute_threshold,string"`
	UnderpaymentRelativeThreshold float64 `json:"underpayment_relative_threshold,string"`
}

type ChargeSupportedAssets struct {
	CurrencyId string   `json:"currency_id"`
	Chains     []string `json:"chains"`
	Name       string   `json:"name"`
}

type ChargeResponse struct {
	Response
	Data Charge `json:"data"`
}

type ListChargesResponse struct {
	ResponseWithPagination
	Data []*Charge `json:"data"`
}

func (s *ChargeService) Create(req *ChargeRequest) (*ChargeResponse, error) {
	var resp = new(ChargeResponse)
	err := s.client.call("POST", "/charges", req, &resp)
	return resp, err
}

func (s *ChargeService) List(params ListParameters) (*ListChargesResponse, error) {
	var resp = new(ListChargesResponse)
	err := s.client.call("GET", fmt.Sprintf("/charges?sort=%s&limit=%d&page=%d",
		params.Sort, params.Limit, params.Page), params, &resp)
	return resp, err
}

func (s *ChargeService) Get(id string) (*ChargeResponse, error) {
	var resp = new(ChargeResponse)
	if id == "" {
		return nil, errors.New("no chargeID provided")
	}
	err := s.client.call("GET", fmt.Sprintf("/charges/%s", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}

func (s *ChargeService) Resolve(id, context string) (*ChargeResponse, error) {
	req := struct {
		Context string `json:"context"`
	}{
		Context: context,
	}
	var resp = new(ChargeResponse)
	if id == "" {
		return nil, errors.New("please provide a chargeID")
	}
	err := s.client.call("POST", fmt.Sprintf("/charges/%s/resolve", strings.TrimSpace(id)), req, &resp)
	return resp, err
}

func (s *ChargeService) Cancel(id string) (*ChargeResponse, error) {
	var resp = new(ChargeResponse)
	if id == "" {
		return resp, errors.New("please provide a chargeID")
	}
	err := s.client.call("PUT", fmt.Sprintf("/charges/%s/cancel", strings.TrimSpace(id)), nil, &resp)
	return resp, err
}
