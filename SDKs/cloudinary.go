package SDKs

import (
	"context"
	"fmt"
	"log"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func GeneratePrefixedUniqueID() string {
    id := uuid.New()
    return fmt.Sprintf("img_%s", id.String()[:8])
}

func Credentials() (*cloudinary.Cloudinary, context.Context) {
    cloudinaryInstance, _ := cloudinary.New()
    cloudinaryInstance.Config.URL.Secure = true
    ctx := context.Background()
    return cloudinaryInstance, ctx
}

func UploadImage(cloudinaryInstance *cloudinary.Cloudinary, ctx context.Context, image any, publicId string) {
	_, err := cloudinaryInstance.Upload.Upload(ctx, image, uploader.UploadParams{
		PublicID:      publicId,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})
	
	if err != nil {
		log.Println("error")
	} else {
		log.Println("Image successfully uploaded")
	}
}

func GetTransformedImage(cloudinaryInstance *cloudinary.Cloudinary, ctx context.Context, publicId string) string {
    image, err := cloudinaryInstance.Image(publicId)
    if err != nil {
        log.Println("error")
    }

    image.Transformation = "c_fill,g_auto,h_500,w_500"
    transformedImageURL, err := image.String()
    if err != nil {
        log.Println("error")
    }

	return transformedImageURL
}