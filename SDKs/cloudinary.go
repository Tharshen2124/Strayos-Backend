package SDKs

import (
    "context"
	"log"
    "github.com/cloudinary/cloudinary-go/v2"
    "github.com/cloudinary/cloudinary-go/v2/api"
    "github.com/cloudinary/cloudinary-go/v2/api/admin"
    "github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Credentials() (*cloudinary.Cloudinary, context.Context) {
    // Add your Cloudinary credentials, set configuration parameter 
    // Secure=true to return "https" URLs, and create a context
    cloudinaryInstance, _ := cloudinary.New()
    cloudinaryInstance.Config.URL.Secure = true
    ctx := context.Background()
    return cloudinaryInstance, ctx
}

func UploadImage(cloudinaryInstance *cloudinary.Cloudinary, ctx context.Context, image any, publicId string) {
	// Upload image & set the asset's public ID and allow overwriting the asset with new versions
	response, err := cloudinaryInstance.Upload.Upload(ctx, image, uploader.UploadParams{
		PublicID:      publicId,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})
	
	if err != nil {
		log.Println("error")
	}
  	log.Println("2. Upload an image\nDelivery URL:", response.SecureURL, "\n")
}

func GetAssetInfo(cloudinaryInstance *cloudinary.Cloudinary, ctx context.Context, publicId string) {
	// Get and use details of the image
	response, err := cloudinaryInstance.Admin.Asset(ctx, admin.AssetParams{PublicID: publicId})
	if err != nil {
		log.Printf("error: %v", err)
	}
	log.Println("3. Get and use details of the image\nDetailed response:\n", response, "\n")

	// Assign tags to the uploaded image based on its width. Save the response to the update in the variable 'update_resp'.
	if response.Width > 900 {
		update_resp, err := cloudinaryInstance.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: publicId,
			Tags:     []string{"large"}})
		if err != nil {
			log.Printf("error: %v", err)
		} else {
			log.Println("New tag: ", update_resp.Tags, "\n")
		}
	} else {
		update_resp, err := cloudinaryInstance.Admin.UpdateAsset(ctx, admin.UpdateAssetParams{
			PublicID: publicId,
			Tags:     []string{"small"}})
		if err != nil {
			log.Printf("error: %v", err)
		} else {
			log.Println("New tag: ", update_resp.Tags, "\n")
		}
	}
}

func TransformImage(cloudinaryInstance *cloudinary.Cloudinary, ctx context.Context, publicId string) {
    // Instantiate an object for the asset with public ID "my_image"
    qs_img, err := cloudinaryInstance.Image(publicId)
    if err != nil {
        log.Println("error")
    }

    // Add the transformation
    qs_img.Transformation = "r_max/e_sepia"

    // Generate and log the delivery URL
    new_url, err := qs_img.String()
    if err != nil {
        log.Println("error")
    } else {
        print("4. Transform the image\nTransfrmation URL: ", new_url, "\n")
    }
}