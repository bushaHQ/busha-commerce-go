package busha_commerce_go

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCheckoutService_Create(t *testing.T) {
	type args struct {
		req *CheckoutRequest
	}
	tests := []struct {
		name    string
		s       CheckoutService
		args    args
		want    *CheckoutResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Create successful checkout",
			s: CheckoutService{
				client: c,
			},
			args: args{
				req: &CheckoutRequest{
					Name:          "Astro checkout",
					Description:   "Testing the checkout for my iphone 14",
					CheckoutType:  FixedPrice,
					RequestedInfo: []string{"name", "email", "phone"},
					LocalAmount:   5000,
					LocalCurrency: "NGN",
				},
			},
			want: &CheckoutResponse{
				Response: Response{Status: Success, Message: "Checkout created successfully"},
				Data:     Checkout{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Create failed checkout (amount must not be attached for donations)",
			s: CheckoutService{
				client: c,
			},
			args: args{
				req: &CheckoutRequest{
					Name:          "Astro NGO",
					Description:   "Raising money to buy a dog",
					CheckoutType:  Donation,
					RequestedInfo: []string{"name", "email", "phone"},
					LocalAmount:   5000,
					LocalCurrency: "NGN",
				},
			},
			want:    &CheckoutResponse{},
			wantErr: assert.Error,
		},
		{
			name: "Create failed checkout (invalid currency)",
			s: CheckoutService{
				client: c,
			},
			args: args{
				req: &CheckoutRequest{
					Name:          "Astro NGO",
					Description:   "Raising money to buy a dog",
					CheckoutType:  FixedPrice,
					RequestedInfo: []string{"name", "email", "phone"},
					LocalAmount:   5000,
					LocalCurrency: "ASTROPCOIN",
				},
			},
			want:    &CheckoutResponse{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Create(tt.args.req)
			if !tt.wantErr(t, err, fmt.Sprintf("Create(%v)", tt.args.req)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "Create(%v)", tt.args.req)
		})
	}
}

func TestCheckoutService_CreateCharge(t *testing.T) {
	type args struct {
		id  string
		req *ChargeRequest
	}
	tests := []struct {
		name    string
		s       CheckoutService
		args    args
		want    *ChargeResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Create charge successful",
			s: CheckoutService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Checkout.Create(&CheckoutRequest{
						Name:          "Astro checkout",
						Description:   "Testing the checkout for my iphone 14",
						CheckoutType:  FixedPrice,
						RequestedInfo: []string{"name", "email", "phone"},
						LocalAmount:   5000,
						LocalCurrency: "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
				req: &ChargeRequest{
					FixedPrice:    true,
					LocalAmount:   5000,
					Meta:          json.RawMessage(`{"name": "Astro Boy", "email": "x@y.com", "phone": "0987654321"}`),
					LocalCurrency: "NGN",
				},
			},
			want: &ChargeResponse{
				Response: Response{Status: Success, Message: "Charge created successfully"},
				Data:     Charge{},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateCharge(tt.args.id, tt.args.req)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateCharge(%v, %v)", tt.args.id, tt.args.req)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "CreateCharge(%v, %v)", tt.args.id, tt.args.req)
			assert.Equalf(t, strings.ToLower(tt.want.Message), strings.ToLower(got.Message), "CreateCharge(%v, %v)", tt.args.id, tt.args.req)
		})
	}
}

func TestCheckoutService_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       CheckoutService
		args    args
		want    *Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Delete successful",
			s: CheckoutService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Checkout.Create(&CheckoutRequest{
						Name:          "Astro checkout",
						Description:   "Testing the checkout for my iphone 14",
						CheckoutType:  FixedPrice,
						RequestedInfo: []string{"name", "email", "phone"},
						LocalAmount:   5000,
						LocalCurrency: "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want:    &Response{Status: Success, Message: "Checkout deleted successfully"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Delete(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("Delete(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "Delete(%v)", tt.args.id)
		})
	}
}

func TestCheckoutService_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       CheckoutService
		args    args
		want    *CheckoutResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get successful checkout",
			s: CheckoutService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Checkout.Create(&CheckoutRequest{
						Name:          "Astro checkout",
						Description:   "Testing the checkout for my iphone 14",
						CheckoutType:  FixedPrice,
						RequestedInfo: []string{"name", "email", "phone"},
						LocalAmount:   5000,
						LocalCurrency: "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want: &CheckoutResponse{
				Response: Response{Status: Success},
				Data:     Checkout{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Get checkout that does not exist",
			s: CheckoutService{
				client: c,
			},
			args: args{
				id: "999",
			},
			want:    &CheckoutResponse{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Get(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("Get(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "Get(%v)", tt.args.id)
		})
	}
}

func TestCheckoutService_List(t *testing.T) {
	type args struct {
		params ListParameters
	}
	tests := []struct {
		name    string
		s       CheckoutService
		args    args
		want    *ListCheckoutResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "List checkout successful",
			s: CheckoutService{
				client: c,
			},
			args: args{
				params: ListParameters{
					Limit: 10,
				},
			},
			want: &ListCheckoutResponse{
				ResponseWithPagination{
					Response:   Response{Status: Success},
					Pagination: Paginator{}},
				[]*Checkout{},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.List(tt.args.params)
			if !tt.wantErr(t, err, fmt.Sprintf("List(%v)", tt.args.params)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "List(%v)", tt.args.params)
		})
	}
}

func TestCheckoutService_ToggleStatus(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       CheckoutService
		args    args
		want    *CheckoutResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Toggle status successful",
			s: CheckoutService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Checkout.Create(&CheckoutRequest{
						Name:          "Astro checkout",
						Description:   "Testing the checkout for my iphone 14",
						CheckoutType:  FixedPrice,
						RequestedInfo: []string{"name", "email", "phone"},
						LocalAmount:   5000,
						LocalCurrency: "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want: &CheckoutResponse{
				Response: Response{Status: Success},
				Data:     Checkout{},
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToggleStatus(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("ToggleStatus(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "ToggleStatus(%v)", tt.args.id)
		})
	}
}

func TestCheckoutService_Update(t *testing.T) {
	type args struct {
		id  string
		req *CheckoutRequest
	}
	tests := []struct {
		name    string
		s       CheckoutService
		args    args
		want    *Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Update successful",
			s: CheckoutService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Checkout.Create(&CheckoutRequest{
						Name:          "Astro checkout",
						Description:   "Testing the checkout for my iphone 14",
						CheckoutType:  FixedPrice,
						RequestedInfo: []string{"name", "email", "phone"},
						LocalAmount:   5000,
						LocalCurrency: "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
				req: &CheckoutRequest{
					Name:          "Check",
					Description:   "Testing the checkout for my iphone 14",
					CheckoutType:  FixedPrice,
					RequestedInfo: []string{"name", "email", "phone"},
					LocalAmount:   5000,
					LocalCurrency: "NGN",
				},
			},
			want:    &Response{Status: Success},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Update(tt.args.id, tt.args.req)
			if !tt.wantErr(t, err, fmt.Sprintf("Update(%v, %v)", tt.args.id, tt.args.req)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "Update(%v, %v)", tt.args.id, tt.args.req)
		})
	}
}
