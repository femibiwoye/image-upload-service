package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/femibiwoye/image-upload-service/internal/imageuploader"
	"google.golang.org/grpc"
)

var uploader imageuploader.ImageUploader

type Server struct {
	srv *grpc.Server
	UnimplementedImageUploaderServer
	uploader imageuploader.ImageUploader
}

// NewServer creates a new server
func NewServer() *Server {
	// Create a new gRPC server
	return &Server{
		srv:      grpc.NewServer(),
		uploader: imageuploader.NewS3Uploader(),
	}
}

func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		return err
	}

	// Register server method (actions the server will do)
	RegisterImageUploaderServer(s.srv, s)

	log.Printf("server listening at %v", lis.Addr())

	return s.srv.Serve(lis)
}

// This is the method function that call the upload interface in the imageuploader package
func (s *Server) Upload(ctx context.Context, in *Image) (*ImageUploadReply, error) {
	loc, err := s.uploader.Upload(ctx, in.Name, in.Body)

	if err != nil {
		return nil, err
	}

	return &ImageUploadReply{Location: loc}, nil
}
