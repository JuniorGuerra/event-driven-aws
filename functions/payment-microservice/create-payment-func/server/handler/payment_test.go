package handler

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestCreatePayment(t *testing.T) {
	type args struct {
		in0   context.Context
		event events.CloudWatchEvent
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "CreatePaymentTest",
			args:    args{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreatePayment(tt.args.in0, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("CreatePayment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
