package ImageDownloader

import (
	"errors"
	"image"
	"image/color"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
)

func RemoveGlob(path string) (err error) {
	contents, err := filepath.Glob(path)
	if err != nil {
		return
	}
	for _, item := range contents {
		err = os.RemoveAll(item)
		if err != nil {
			return
		}
	}
	return
}

func CreateThumb(s string, o string, size int) error {
	// Open a test image.
	src, err := imaging.Open(s)
	if err != nil {
		return err
	}

	// Resize the cropped image to width = 200px preserving the aspect ratio.
	src = imaging.Resize(src, size, 0, imaging.Lanczos)

	// Create a new image and paste the four produced images into it.
	dst := imaging.New(size, size, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, src, image.Pt(0, 0))

	// Save the resulting image as JPEG.
	err = imaging.Save(dst, o)
	if err != nil {
		return err
	}

	return nil
}

func Exists(filePath string) bool {

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
