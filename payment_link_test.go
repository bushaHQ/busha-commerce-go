package busha_commerce_go

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPaymentLinkService_Create(t *testing.T) {
	type args struct {
		req *PaymentLinkRequest
	}
	tests := []struct {
		name    string
		s       PaymentLinkService
		args    args
		want    *PaymentLinkResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Create successful payment link",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				req: &PaymentLinkRequest{
					Name:            "Astro payment link",
					Description:     "Testing the payment link for my iphone 14",
					PaymentLinkType: FixedPrice,
					RequestedInfo:   []string{"name", "email", "phone"},
					LocalAmount:     5000,
					LocalCurrency:   "NGN",
				},
			},
			want: &PaymentLinkResponse{
				Response: Response{Status: Success, Message: "PaymentLink created successfully"},
				Data:     PaymentLink{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Create failed payment link (amount must not be attached for donations)",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				req: &PaymentLinkRequest{
					Name:            "Astro NGO",
					Description:     "Raising money to buy a dog",
					PaymentLinkType: Donation,
					RequestedInfo:   []string{"name", "email", "phone"},
					LocalAmount:     5000,
					LocalCurrency:   "NGN",
				},
			},
			want:    &PaymentLinkResponse{},
			wantErr: assert.Error,
		},
		{
			name: "Create failed payment link (invalid currency)",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				req: &PaymentLinkRequest{
					Name:            "Astro NGO",
					Description:     "Raising money to buy a dog",
					PaymentLinkType: FixedPrice,
					RequestedInfo:   []string{"name", "email", "phone"},
					LocalAmount:     5000,
					LocalCurrency:   "ASTROPCOIN",
				},
			},
			want:    &PaymentLinkResponse{},
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

func TestPaymentLinkService_CreateCharge(t *testing.T) {
	type args struct {
		id  string
		req *ChargeRequest
	}
	tests := []struct {
		name    string
		s       PaymentLinkService
		args    args
		want    *ChargeResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Create charge successful",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.PaymentLink.Create(&PaymentLinkRequest{
						Name:            "Astro payment link",
						Description:     "Testing the payment link for my iphone 14",
						PaymentLinkType: FixedPrice,
						RequestedInfo:   []string{"name", "email", "phone"},
						LocalAmount:     5000,
						LocalCurrency:   "NGN",
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

func TestPaymentLinkService_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       PaymentLinkService
		args    args
		want    *Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Delete successful",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.PaymentLink.Create(&PaymentLinkRequest{
						Name:            "Astro payment link",
						Description:     "Testing the payment link for my iphone 14",
						PaymentLinkType: FixedPrice,
						RequestedInfo:   []string{"name", "email", "phone"},
						LocalAmount:     5000,
						LocalCurrency:   "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want:    &Response{Status: Success, Message: "Payment link deleted successfully"},
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

func TestPaymentLinkService_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       PaymentLinkService
		args    args
		want    *PaymentLinkResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get successful payment link",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.PaymentLink.Create(&PaymentLinkRequest{
						Name:            "Astro payment link",
						Description:     "Testing the payment link for my iphone 14",
						PaymentLinkType: FixedPrice,
						RequestedInfo:   []string{"name", "email", "phone"},
						LocalAmount:     5000,
						LocalCurrency:   "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want: &PaymentLinkResponse{
				Response: Response{Status: Success},
				Data:     PaymentLink{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Get payment link that does not exist",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				id: "999",
			},
			want:    &PaymentLinkResponse{},
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

func TestPaymentLinkService_List(t *testing.T) {
	type args struct {
		params ListParameters
	}
	tests := []struct {
		name    string
		s       PaymentLinkService
		args    args
		want    *ListPaymentLinksResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "List payment links successful",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				params: ListParameters{
					Limit: 10,
				},
			},
			want: &ListPaymentLinksResponse{
				ResponseWithPagination{
					Response:   Response{Status: Success},
					Pagination: Paginator{}},
				[]*PaymentLink{},
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

func TestPaymentLinkService_ToggleStatus(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       PaymentLinkService
		args    args
		want    *PaymentLinkResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Toggle status successful",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.PaymentLink.Create(&PaymentLinkRequest{
						Name:            "Astro payment link",
						Description:     "Testing the payment link for my iphone 14",
						PaymentLinkType: FixedPrice,
						RequestedInfo:   []string{"name", "email", "phone"},
						LocalAmount:     5000,
						LocalCurrency:   "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want: &PaymentLinkResponse{
				Response: Response{Status: Success},
				Data:     PaymentLink{},
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

func TestPaymentLinkService_Update(t *testing.T) {
	type args struct {
		id  string
		req *PaymentLinkRequest
	}
	tests := []struct {
		name    string
		s       PaymentLinkService
		args    args
		want    *Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Update successful",
			s: PaymentLinkService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.PaymentLink.Create(&PaymentLinkRequest{
						Name:            "Astro payment link",
						Description:     "Testing the payment link for my iphone 14",
						PaymentLinkType: FixedPrice,
						RequestedInfo:   []string{"name", "email", "phone"},
						LocalAmount:     5000,
						LocalCurrency:   "NGN",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
				req: &PaymentLinkRequest{
					Name:            "Check",
					Description:     "Testing the payment link for my iphone 14",
					PaymentLinkType: FixedPrice,
					RequestedInfo:   []string{"name", "email", "phone"},
					LocalAmount:     5000,
					LocalCurrency:   "NGN",
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
