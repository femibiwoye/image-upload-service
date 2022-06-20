package imageuploader

import "context"

/**
 * ImageUploader is the interface that defines the methods that the image uploader service will use.
 * The image uploader service will use the implementation of this interface to upload images to the cloud.
 */
type ImageUploader interface {
	Upload(ctx context.Context, name string, data []byte) (string, error)
}
