package busha_commerce_go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddressService_Create(t *testing.T) {
	type args struct {
		req *AddressRequest
	}
	tests := []struct {
		name    string
		s       AddressService
		args    args
		want    *AddressResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Create successful address",
			s: AddressService{
				client: c,
			},
			args: args{
				req: &AddressRequest{
					CurrencyId: "USDT",
					Chains:     []string{"ETH"},
					Label:      "Test Address",
				},
			},
			want: &AddressResponse{
				Response: Response{Status: Success, Message: "Address created successfully"},
				Data:     Address{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Create failed address (chain cannot be empty)",
			s: AddressService{
				client: c,
			},
			args: args{
				req: &AddressRequest{
					CurrencyId: "USDT",
					Chains:     []string{""},
					Label:      "Test Address2",
				},
			},
			want:    &AddressResponse{},
			wantErr: assert.Error,
		},
		{
			name: "Create failed address (invalid currency)",
			s: AddressService{
				client: c,
			},
			args: args{
				req: &AddressRequest{
					CurrencyId: "USDT0192020",
					Chains:     []string{"ETH"},
					Label:      "Test Address3",
				},
			},
			want:    &AddressResponse{},
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

func TestAddressService_List(t *testing.T) {
	type args struct {
		params ListParameters
	}
	tests := []struct {
		name    string
		s       AddressService
		args    args
		want    *ListAddressesResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "List addresses successful",
			s: AddressService{
				client: c,
			},
			args: args{
				params: ListParameters{
					Limit: 10,
				},
			},
			want: &ListAddressesResponse{
				ResponseWithPagination{
					Response:   Response{Status: Success},
					Pagination: Paginator{}},
				[]*Address{},
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

func TestAddressService_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       AddressService
		args    args
		want    *AddressResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Get successful address",
			s: AddressService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Address.Create(&AddressRequest{
						CurrencyId: "USDT",
						Chains:     []string{"ETH"},
						Label:      "Test Address",
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want: &AddressResponse{
				Response: Response{Status: Success},
				Data:     Address{},
			},
			wantErr: assert.NoError,
		},
		{
			name: "Get address that does not exist",
			s: AddressService{
				client: c,
			},
			args: args{
				id: "999",
			},
			want:    &AddressResponse{},
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
