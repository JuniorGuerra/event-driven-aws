package repository

import "testing"

func TestPaymentHandler_CreatePayment(t *testing.T) {
	type args struct {
		payment CreatePaymentDB
	}
	tests := []struct {
		name    string
		orders  PaymentHandler
		args    args
		wantErr bool
	}{
		{
			name:    "Cantput",
			orders:  PaymentHandler{},
			args:    args{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orders := PaymentHandler{}
			if err := orders.CreatePayment(tt.args.payment); (err != nil) != tt.wantErr {
				t.Errorf("PaymentHandler.CreatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
