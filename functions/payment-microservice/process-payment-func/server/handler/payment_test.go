package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestProcessPayment(t *testing.T) {
	type args struct {
		in0 context.Context
		e   events.APIGatewayV2HTTPRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayV2HTTPResponse
		wantErr bool
	}{
		{
			name: "EmptyBody",
			args: args{
				context.TODO(),
				events.APIGatewayV2HTTPRequest{
					Body: "",
				},
			},
			want: events.APIGatewayV2HTTPResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "unexpected end of JSON input",
			},
			wantErr: true,
		},
		{
			name: "StatussSuccessBadWrite",
			args: args{
				context.TODO(),
				events.APIGatewayV2HTTPRequest{
					Body: `{"order_id": "user123", "status": "sucess"}`,
				},
			},
			want: events.APIGatewayV2HTTPResponse{
				StatusCode: http.StatusInternalServerError,
			},
			wantErr: true,
		},
		{
			name: "StatussSuccessWrite",
			args: args{
				context.TODO(),
				events.APIGatewayV2HTTPRequest{
					Body: `{"order_id": "user123", "status": "success"}`,
				},
			},
			want: events.APIGatewayV2HTTPResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "unexpected end of JSON input",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProcessPayment(tt.args.in0, tt.args.e)
			if (err != nil) == tt.wantErr {
				t.Errorf("ProcessPayment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	// No mock para no crear complejidad
}
