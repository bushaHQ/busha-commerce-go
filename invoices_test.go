package busha_commerce_go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoiceService_Create(t *testing.T) {
	type args struct {
		req *InvoiceRequest
	}
	tests := []struct {
		name    string
		s       InvoiceService
		args    args
		want    *InvoiceResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Create successful invoice",
			s: InvoiceService{
				client: c,
			},
			args: args{
				req: &InvoiceRequest{
					Name:          "Payment for Development Services",
					CustomerEmail: "syz@g.com",
					LocalAmount:   5000,
					LocalCurrency: "NGN",
					CustomerName:  "Astro",
					Description:   "Test description",
					DueDate:       nil,
				},
			},
			want: &InvoiceResponse{
				Response: Response{Status: Success, Message: "Invoice created successfully"},
				Data:     Invoice{},
			},
			wantErr: assert.NoError,
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

func TestInvoiceService_CreateCharge(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       InvoiceService
		args    args
		want    *ChargeResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Create charge successful invoice",
			s: InvoiceService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Invoice.Create(&InvoiceRequest{
						Name:          "Payment for Development Services",
						CustomerEmail: "syz@g.com",
						LocalAmount:   5000,
						LocalCurrency: "NGN",
						CustomerName:  "Astro",
						Description:   "Test description",
						DueDate:       nil,
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want: &ChargeResponse{
				Response: Response{Status: Success, Message: "Charge created successfully"},
				Data:     Charge{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Create charge with invalid invoice",
			s: InvoiceService{
				client: c,
			},
			args: args{
				id: "1",
			},
			want: &ChargeResponse{
				Response: Response{},
				Data:     Charge{},
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.CreateCharge(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("CreateCharge(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "CreateCharge(%v)", tt.args.id)
		})
	}
}

func TestInvoiceService_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       InvoiceService
		args    args
		want    *InvoiceResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get successful invoice",
			s: InvoiceService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Invoice.Create(&InvoiceRequest{
						Name:          "Payment for Development Services",
						CustomerEmail: "syz@g.com",
						LocalAmount:   5000,
						LocalCurrency: "NGN",
						CustomerName:  "Astro",
						Description:   "Test description",
						DueDate:       nil,
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want: &InvoiceResponse{
				Response: Response{Status: Success, Message: "Invoice created successfully"},
				Data:     Invoice{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Get with invalid invoice",
			s: InvoiceService{
				client: c,
			},
			args: args{
				id: "1",
			},
			want: &InvoiceResponse{
				Response: Response{},
				Data:     Invoice{},
			},
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

func TestInvoiceService_List(t *testing.T) {
	type args struct {
		params ListParameters
	}
	tests := []struct {
		name    string
		s       InvoiceService
		args    args
		want    *ListInvoiceResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "List successful invoice",
			s: InvoiceService{
				client: c,
			},
			args: args{
				params: ListParameters{
					Limit: 10,
				},
			},
			want: &ListInvoiceResponse{
				ResponseWithPagination{
					Response:   Response{Status: Success},
					Pagination: Paginator{}},
				[]*Invoice{},
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

func TestInvoiceService_Void(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       InvoiceService
		args    args
		want    *Response
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Void successful invoice",
			s: InvoiceService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Invoice.Create(&InvoiceRequest{
						Name:          "Payment for Development Services",
						CustomerEmail: "syz@g.com",
						LocalAmount:   5000,
						LocalCurrency: "NGN",
						CustomerName:  "Astro",
						Description:   "Test description",
						DueDate:       nil,
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want:    &Response{Status: Success, Message: "Invoice voided successfully"},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Void(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("Void(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "Void(%v)", tt.args.id)
		})
	}
}
