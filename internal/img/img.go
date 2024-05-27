package img

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/pkg/errors"
)

type ImageData struct {
	SrcImage image.Image
	DstImage *image.RGBA
	Height   int
	Width    int
}

func Read(imgPath string) (image.Image, error) {
	fmt.Printf("imgPath: %v\n", imgPath)
	file, err := os.Open(imgPath)
	if err != nil {
		return nil, errors.Wrap(err, "img.Read error")
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func Write(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return errors.Wrap(err, "img.Write error")
	}
	defer file.Close()

	return jpeg.Encode(file, img, nil)
}
