package main

import "github.com/femibiwoye/image-upload-service/internal/cli"

/**
 * This is the main function of the image-uploader-cli package. It is the entry point of the image-uploader-cli application.
 * The image-uploader-cli application is a command line interface for the image-upload-service application. It is used to upload images to the cloud.
 * The image-upload-service application is a gRPC server that accepts image upload requests and returns the location of the uploaded image.
 */
func main() {
	cli.Execute()
}
