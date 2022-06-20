package grpc

import (
	"context"
	"testing"
)

func TestClient_UploadImage(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		data []byte
	}
	var tests []struct {
		name    string
		c       *Client
		args    args
		want    string
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UploadImage(tt.args.ctx, tt.args.name, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.UploadImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.UploadImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_DialServer(t *testing.T) {
	var tests []struct {
		name    string
		c       *Client
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.DialServer(); (err != nil) != tt.wantErr {
				t.Errorf("Client.DialServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
