package imageuploader

import (
	"context"
	"testing"
)

func Test_extensionFromContentType(t *testing.T) {
	type args struct {
		contentType string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"png file", args{contentType: "image/png"}, ".png"},
		{"jpg file", args{contentType: "image/jpg"}, ".jpg"},
		{"gif file", args{contentType: "image/gif"}, ".gif"},
		{"invalid mime type", args{contentType: "image-gif"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extensionFromContentType(tt.args.contentType); got != tt.want {
				t.Errorf("extensionFromContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidImageType(t *testing.T) {
	type args struct {
		mimeType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"png file", args{mimeType: "image/png"}, true},
		{"jpg file", args{mimeType: "image/jpg"}, true},
		{"jpeg file", args{mimeType: "image/jpeg"}, true},
		{"gif file", args{mimeType: "image/gif"}, true},
		{"webp file", args{mimeType: "image/webp"}, false},
		{"mp3 file", args{mimeType: "audio/mp3"}, false},
		{"mp4 file", args{mimeType: "video/mp4"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidImageType(tt.args.mimeType); got != tt.want {
				t.Errorf("isValidImageType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_prefixKey(t *testing.T) {
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"prefix present", args{key: "world", prefix: "hello"}, "hello/world"},
		{"trailing slash prefix present", args{key: "world", prefix: "hello/"}, "hello/world"},
		{"no prefix", args{key: "world", prefix: ""}, "world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixKey(tt.args.prefix, tt.args.key); got != tt.want {
				t.Errorf("prefixKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_awsS3_Upload(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		data []byte
	}
	var tests []struct {
		name         string
		as3          *awsS3
		args         args
		wantLocation string
		wantErr      bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLocation, err := tt.as3.Upload(tt.args.ctx, tt.args.name, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("awsS3.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLocation != tt.wantLocation {
				t.Errorf("awsS3.Upload() = %v, want %v", gotLocation, tt.wantLocation)
			}
		})
	}
}
