package busha_commerce_go

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	defaultHTTPTimeout = 20 * time.Second
	baseURL            = "https://api.commerce.busha.co"
	userAgent          = "Busha/Commerce-SDK"
	liveKeyPrefix      = "live_"
	testKeyPrefix      = "test_"
)

type service struct {
	client *Client
}

type Client struct {
	base      service
	client    *http.Client
	secretKey string
	userAgent string
	baseURL   *url.URL

	LogDebug bool
	Log      Logger

	Charge   *ChargeService
	Checkout *CheckoutService
	Invoice  *InvoiceService
	Event    *EventService
	Address  *AddressService
}

type Logger interface {
	Printf(format string, v ...interface{})
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseWithPagination struct {
	Response
	Pagination Paginator `json:"pagination"`
}

type Paginator struct {
	Page               int `json:"page"`
	PerPage            int `json:"per_page"`
	Offset             int `json:"offset"`
	TotalEntriesSize   int `json:"total_entries_size"`
	CurrentEntriesSize int `json:"current_entries_size"`
	TotalPages         int `json:"total_pages"`
}

func New(key string, httpClient *http.Client) (*Client, error) {
	if err := validateKey(key); err != nil {
		return nil, err
	}

	if httpClient == nil {
		httpClient = &http.Client{Timeout: defaultHTTPTimeout}
	}

	u, _ := url.Parse(baseURL)
	c := &Client{
		client:    httpClient,
		secretKey: key,
		userAgent: userAgent,
		baseURL:   u,
		LogDebug:  false,
		Log:       log.New(os.Stderr, "", log.LstdFlags),
	}

	c.base.client = c
	c.Charge = (*ChargeService)(&c.base)
	c.Checkout = (*CheckoutService)(&c.base)
	c.Invoice = (*InvoiceService)(&c.base)
	c.Event = (*EventService)(&c.base)
	c.Address = (*AddressService)(&c.base)

	return c, nil
}

func (c *Client) call(method, path string, reqBody, response interface{}) (err error) {
	buffer := bytes.NewBuffer([]byte{})
	if method == http.MethodPost || method == http.MethodPut {
		if err = json.NewEncoder(buffer).Encode(reqBody); err != nil {
			return err
		}
	}
	u, _ := c.baseURL.Parse(path)
	req, err := http.NewRequest(method, u.String(), buffer)

	if err != nil {
		return err
	}

	req.Header.Set("X-BC-API-KEY", c.secretKey)
	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Set("Content-Type", "application/json")

	if c.LogDebug {
		c.Log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
		c.Log.Printf("%s request data %v\n", req.Method, reqBody)
	}

	req = req.WithContext(context.TODO())

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		e := ErrResponse{}
		err = json.NewDecoder(resp.Body).Decode(&e)

		if err != nil {
			return err
		}

		return e
	}

	err = json.NewDecoder(resp.Body).Decode(response)
	return
}

func (c *Client) SetDebug(debug bool) {
	c.LogDebug = debug
}

func validateKey(key string) error {
	if len(key) == 0 {
		return errors.New("commerce Secret Key cannot be empty")
	}

	if !strings.HasPrefix(key, liveKeyPrefix) && !strings.HasPrefix(key, testKeyPrefix) {
		return errors.New("Commerce Secret Key  is not valid")
	}
	return nil
}

func mustHaveTestKeyEnv() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file for test : %v", err)
	}
	key := os.Getenv("COMMERCE_KEY")
	if err := validateKey(key); err != nil {
		panic(err)
	}
	return key
}
