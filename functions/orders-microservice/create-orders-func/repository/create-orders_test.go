package repository

import (
	"testing"
)

func TestOrderHandler_CreateOrder(t *testing.T) {
	type args struct {
		order CreateOrderDB
	}
	// para no crear una estructura muy completa, solo vamos a testear por encima...
	tests := []struct {
		name    string
		orders  OrderHandler
		args    args
		want    OrderResult
		wantErr bool
	}{
		{
			name:    "testErr",
			orders:  OrderHandler{},
			args:    args{},
			want:    OrderResult{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orders := OrderHandler{}
			_, err := orders.CreateOrder(tt.args.order)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrderHandler.CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
