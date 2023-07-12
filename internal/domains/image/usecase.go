package image

import (
	"Datearn/pkg/advancedError"
	"Datearn/pkg/userHandler"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"time"
)

const imagePath = "../../../static/images/"

// ImageUseCase is a struct which satisfies user use case interface functionalities
type ImageUseCase struct {
	sv        ImageServiceInterface
	decoderFn func(ctx context.Context, token string) (uint, error)
}

// NewImageUseCase and return it
func NewImageUseCase(sv ImageServiceInterface, decoderFn func(ctx context.Context, token string) (uint, error)) ImageUseCaseInterface {
	if decoderFn == nil {
		decoderFn = userHandler.ExtractUserIdFromToken
	}

	return &ImageUseCase{
		sv:        sv,
		decoderFn: decoderFn,
	}
}

func (i *ImageUseCase) GetAllImage(ctx context.Context, token string, request *ImageGetRequest) (*[]Image, error) {
	_, err := i.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}
	return i.sv.GetAllImage(request.Id, request.ImageableId, request.Name, request.ImageableType, request.Limit, request.Offset)
}

func (i *ImageUseCase) CreateImage(ctx context.Context, token string, request *ImageCreateRequest) (*Image, error) {
	_, err := i.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	_, path := createFile(request.Name, request.Img)

	return i.sv.CreateImage(request.ImageableId, request.Name, path, request.ImageableType)

}

func (i *ImageUseCase) UpdateImage(ctx context.Context, token string, request *ImageUpdateRequest) (*Image, error) {
	//TODO implement me
	panic("implement me")
}

func (i *ImageUseCase) DeleteImage(ctx context.Context, token string, request *ImageDeleteRequest) (*Image, error) {
	_, err := i.decoderFn(ctx, token)
	if err != nil {
		return nil, advancedError.New(err, "Decoding token failed")
	}

	return i.sv.DeleteImage(&request.Id)
}

func createFile(name string, img *image.RGBA) (string, string) {

	pathDir, err := createDirectoryFromDate()

	if err != nil {
		panic(err)
	}

	// Create a file to save the image to.
	file, err := os.Create(imagePath + pathDir + name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the image as jpeg and write it to the file.
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		panic(err)
	}

	return filepath.Base(file.Name()), filepath.Join("../../static/images/", file.Name())
}

func createDirectoryFromDate() (string, error) {
	// Get the current date.
	t := time.Now()

	// Format the date as year/month/day.
	date := t.Format("2006/01/02")

	// Print the date.
	//fmt.Println(date)

	// Create the directories with the date as the path.
	err := os.MkdirAll(imagePath+date, 0755)
	if err != nil {
		// Print the error if any.
		fmt.Println(err)
		return "", err
	}

	// Print a success message.
	//fmt.Println("Directories have been created successfully!")
	return date + "/", nil
}
