package images

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type Encoder func(io.Writer, image.Image) error

func GetJPEGEncoder() Encoder {
	return jpegEncoder(70)
}

func GetPNGEncoder() Encoder {
	return pngEncoder()
}

// JPEGEncoder returns an encoder to JPEG given the argument 'quality'
func jpegEncoder(quality int) Encoder {
	return func(w io.Writer, img image.Image) error {
		return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
	}
}

// PNGEncoder returns an encoder to PNG
func pngEncoder() Encoder {
	return func(w io.Writer, img image.Image) error {
		return png.Encode(w, img)
	}
}