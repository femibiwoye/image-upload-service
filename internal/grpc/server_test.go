package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

func TestServer_Start(t *testing.T) {
	type args struct {
		port string
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Start(tt.args.port); (err != nil) != tt.wantErr {
				t.Errorf("Server.Start() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_Upload(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *Image
	}
	tests := []struct {
		name    string
		s       *Server
		args    args
		want    *ImageUploadReply
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "upload file successfully",
			s:    &Server{uploader: testUploader{}},
			args: args{
				ctx: context.TODO(),
				in:  &Image{Name: "hello world"},
			},
			want:    &ImageUploadReply{Location: "hello world"},
			wantErr: false,
		},
		{
			name: "upload file failed",
			s:    &Server{uploader: testUploader{shouldFail: true}},
			args: args{
				ctx: context.TODO(),
				in:  &Image{Name: "hello world"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Upload(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.Upload() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testUploader struct {
	shouldFail bool
}

func (t testUploader) Upload(ctx context.Context, name string, data []byte) (string, error) {
	if t.shouldFail {
		return "", errors.New("failed")
	}

	return name, nil
}
