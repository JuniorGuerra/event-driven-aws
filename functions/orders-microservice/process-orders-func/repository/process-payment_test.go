package repository

import "testing"

func TestPaymentHandler_UpdatePayment(t *testing.T) {
	type args struct {
		payment ItemOrderDB
	}
	tests := []struct {
		name    string
		orders  PaymentHandler
		args    args
		wantErr bool
	}{
		{
			name:    "NotCreatedItem",
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
