package busha_commerce_go

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const Success = "success"

func TestChargeService_Create(t *testing.T) {
	type args struct {
		req *ChargeRequest
	}
	tests := []struct {
		name    string
		s       ChargeService
		args    args
		want    *ChargeResponse
		wantErr bool
	}{
		{
			name: "Create successful charge with no fixed price",
			s: ChargeService{
				client: c,
			},
			args: args{
				req: &ChargeRequest{
					FixedPrice: false,
					Meta:       json.RawMessage(`{"name":"test","email":"sarah.shaw@example.co"}`),
				},
			},
			want:    &ChargeResponse{Response{Status: Success}, Charge{}},
			wantErr: false,
		},
		{
			name: "Create fixed price charge without amount or currency",
			s: ChargeService{
				client: c,
			},
			args: args{
				req: &ChargeRequest{
					FixedPrice:    true,
					LocalAmount:   0,
					LocalCurrency: "",
					Meta:          json.RawMessage(`{"name":"test","email":"sarah.shaw@example.co"}`),
				},
			},
			want:    &ChargeResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Create(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want.Status, got.Status) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChargeService_List(t *testing.T) {
	type args struct {
		params ListParameters
	}
	tests := []struct {
		name    string
		s       ChargeService
		args    args
		want    *ListChargesResponse
		wantErr bool
	}{
		{
			name: "List charges with no filters",
			s: ChargeService{
				client: c,
			},
			args: args{
				params: ListParameters{},
			},
			want: &ListChargesResponse{ResponseWithPagination{
				Response:   Response{Status: Success},
				Pagination: Paginator{},
			}, []*Charge{}},
			wantErr: false,
		},
		{
			name: "List charges with filters",
			s: ChargeService{
				client: c,
			},
			args: args{
				params: ListParameters{
					Sort:  "asc",
					Page:  1,
					Limit: 2,
				},
			},
			want: &ListChargesResponse{ResponseWithPagination{
				Response:   Response{Status: Success},
				Pagination: Paginator{},
			}, []*Charge{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.List(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want.Response.Status, got.Response.Status) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
			if len(tt.want.Data) > 0 {
				assert.Greater(t, 0, got.Pagination.Page)
			}
		})
	}
}

func TestChargeService_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       ChargeService
		args    args
		want    *ChargeResponse
		wantErr bool
	}{
		{
			name: "Get charge with invalid id",
			s: ChargeService{
				client: c,
			},
			args: args{
				id: "999",
			},
			want:    &ChargeResponse{},
			wantErr: true,
		},
		{
			name: "Get charge with valid id",
			s: ChargeService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Charge.Create(&ChargeRequest{
						FixedPrice: false,
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want:    &ChargeResponse{Response{Status: Success}, Charge{}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "Get(%v)", tt.args.id)
		})
	}
}

func TestChargeService_Resolve(t *testing.T) {
	type args struct {
		id      string
		context string
	}
	tests := []struct {
		name    string
		s       ChargeService
		args    args
		want    *ChargeResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Resolve charge with invalid id",
			s: ChargeService{
				client: c,
			},
			args: args{
				id:      "999",
				context: "",
			},
			want:    &ChargeResponse{},
			wantErr: assert.Error,
		},
		{
			name: "Resolve charge with valid id and no context",
			s: ChargeService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Charge.Create(&ChargeRequest{
						FixedPrice: false,
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
				context: "",
			},
			want:    &ChargeResponse{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Resolve(tt.args.id, tt.args.context)
			if !tt.wantErr(t, err, fmt.Sprintf("Resolve(%v, %v)", tt.args.id, tt.args.context)) {
				return
			}
			assert.Equalf(t, tt.want, got, "Resolve(%v, %v)", tt.args.id, tt.args.context)
		})
	}
}

func TestChargeService_Cancel(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       ChargeService
		args    args
		want    *ChargeResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Cancel charge with invalid id",
			s: ChargeService{
				client: c,
			},
			args: args{
				id: "999",
			},
			want:    &ChargeResponse{},
			wantErr: assert.Error,
		},
		{
			name: "Cancel charge with valid id",
			s: ChargeService{
				client: c,
			},
			args: args{
				id: func() string {
					got, err := c.Charge.Create(&ChargeRequest{
						FixedPrice: false,
					})
					if err != nil {
						return ""
					}
					return got.Data.Id.String()
				}(),
			},
			want:    &ChargeResponse{Response{Status: Success}, Charge{}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Cancel(tt.args.id)
			if !tt.wantErr(t, err, fmt.Sprintf("Cancel(%v)", tt.args.id)) {
				return
			}
			assert.Equalf(t, tt.want.Status, got.Status, "Cancel(%v)", tt.args.id)
		})
	}
}
