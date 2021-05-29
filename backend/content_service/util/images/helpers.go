package images

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"os"
	"strings"
)

// Open an image from a file
func LoadImage(filename string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}


func save(filename string, img image.Image, encoder Encoder) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return encoder(f, img)
}

// SaveImage saves an image with appropriate extension using base64string
func SaveImage(filename string, base64string string) error{
	img, imageType, err := getImageFromBase64(base64string)
	if err != nil {
		return err
	}

	var extension string
	var encoder Encoder

	switch imageType{
	case "image/jpeg", "image/jpg":
		extension = ".jpg"
		encoder = GetJPEGEncoder(EncoderSaving)
		break
	case "image/png":
		extension = ".png"
		encoder = GetPNGEncoder()
		break
	default:
		return errors.New("save: unsupported image type")
	}

	filename += extension
	err = save(filename, img, encoder)

	return err
}

// GetImageFromBase64 converts base64 string, including suffix with "data:image/..." part, to an image
func getImageFromBase64(base64string string) (image.Image, string, error){
	base64parts := strings.Split(base64string, ";base64,")
	imageType := base64parts[0][5:]
	base64string = base64parts[1]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64string))
	img, _, err := image.Decode(reader)
	if err != nil{
		return nil, "", err
	}

	return img, imageType, nil
}

// LoadImageToBase64 loads an image and convert it to base64 string.
// Supported image formats are jpeg and png.
func LoadImageToBase64(filename string) (string, error) {
	image, err := LoadImage(filename)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	filenameParts := strings.Split(filename, ".")
	mimetypeSuffix := filenameParts[len(filenameParts) - 1]

	var base64string string
	var encoder Encoder

	switch mimetypeSuffix {
	case "jpg", "jpeg":
		base64string = "data:image/jpeg;base64,"
		encoder = GetJPEGEncoder(EncoderLoading)
		break
	case "png":
		base64string = "data:image/png;base64,"
		encoder = GetPNGEncoder()
		break
	default:
		return "", errors.New("toBase64: unsupported image type")
	}

	imageBuffer := new(bytes.Buffer)
	err = encoder(imageBuffer, image)
	if err != nil{
		return "", nil
	}

	base64string += base64.StdEncoding.EncodeToString(imageBuffer.Bytes())

	return base64string, nil
}