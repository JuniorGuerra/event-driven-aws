package repository

import (
	"testing"
)

func TestPaymentHandler_UpdatePayment(t *testing.T) {
	type args struct {
		payment ItemPaymentDB
	}
	tests := []struct {
		name    string
		orders  PaymentHandler
		args    args
		wantErr bool
	}{
		{
			name:    "CantUpdate",
			orders:  PaymentHandler{},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orders := PaymentHandler{}
			if err := orders.UpdatePayment(tt.args.payment); (err != nil) != tt.wantErr {
				t.Errorf("PaymentHandler.UpdatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPaymentHandler_GetPaymentItem(t *testing.T) {
	type args struct {
		orderId string
	}
	tests := []struct {
		name      string
		orders    PaymentHandler
		args      args
		want      ItemPaymentDB
		want1     error
		wantFalse bool
	}{
		{
			name:      "CantGetItem",
			orders:    PaymentHandler{},
			args:      args{},
			want:      ItemPaymentDB{},
			want1:     nil,
			wantFalse: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orders := PaymentHandler{}
			_, _, err := orders.GetPaymentItem(tt.args.orderId)

			if (err != nil) == tt.wantFalse {
				t.Errorf("PaymentHandler.GetPaymentItem() error = %v, wantErr %v", err, tt.wantFalse)
				return
			}
		})
	}
}
