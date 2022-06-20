package imageuploader

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/viper"
)

var ErrInvalidFileType = errors.New("invalid file type")

var validImageTypes = []string{
	"image/jpg",
	"image/png",
	"image/jpeg",
	"image/gif",
}

type awsS3 struct {
	client   *s3.S3
	uploader *s3manager.Uploader
}

var _ ImageUploader = (*awsS3)(nil)

func NewS3Uploader() *awsS3 {
	sess := session.Must(session.NewSession())
	svc := s3.New(sess)

	return &awsS3{
		client:   svc,
		uploader: s3manager.NewUploaderWithClient(svc),
	}
}

func (as3 *awsS3) Upload(ctx context.Context, name string, data []byte) (location string, err error) {
	contentType := http.DetectContentType(data)

	ss := strings.Split(name, ".")
	name = strings.Join(ss[:len(ss)-1], "")

	if !isValidImageType(contentType) {
		return "", ErrInvalidFileType
	}

	ext := extensionFromContentType(contentType)
	key := prefixKey(viper.GetString("EndpointPrefix"), fmt.Sprintf("%d/%s", time.Now().Unix(), name+ext))
	params := &s3manager.UploadInput{
		Bucket:      aws.String(viper.GetString("BucketName")),
		Key:         aws.String(key),
		Body:        bytes.NewBuffer(data),
		ACL:         aws.String("public-read"),
		ContentType: aws.String(contentType),
		//ContentEncoding: aws.String("base64"),
	}

	out, err := as3.uploader.Upload(params)

	if err != nil {
		log.Printf("Failed to upload file: %v", err)
		return "", err
	}

	log.Printf("Successfully uploaded file to %s", out.Location)

	return out.Location, nil
}

func prefixKey(prefix, key string) string {
	if prefix == "" {
		return key
	}

	return strings.TrimSuffix(prefix, "/") + "/" + key
}

func extensionFromContentType(contentType string) string {
	ss := strings.Split(contentType, "/")

	if len(ss) != 2 {
		return ""
	}

	return "." + ss[1]
}

func isValidImageType(mimeType string) bool {
	var isValid bool

	for _, v := range validImageTypes {
		if mimeType == v {
			isValid = true
			break
		}
	}

	return isValid
}
