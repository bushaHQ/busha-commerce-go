package busha_commerce_go

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventService_List(t *testing.T) {
	type args struct {
		params ListParameters
	}
	tests := []struct {
		name    string
		s       EventService
		args    args
		want    *ListEventResponse
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "List events successfully",
			s: EventService{
				client: c,
			},
			args: args{
				params: ListParameters{
					Limit: 10,
				},
			},
			want: &ListEventResponse{
				ResponseWithPagination{
					Response:   Response{Status: Success},
					Pagination: Paginator{}},
				[]*Event{},
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
