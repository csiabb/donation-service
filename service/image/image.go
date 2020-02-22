/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"image"
	"image/png"
	"math"
	"os"

	"github.com/nfnt/resize"
)

// calculating thumbnail size
func calculateRatioFit(srcWidth, srcHeight, newDx int) (int, int) {
	ratio := math.Min(float64(newDx)/float64(srcWidth), float64(newDx)/float64(srcHeight))
	return int(math.Ceil(float64(srcWidth) * ratio)), int(math.Ceil(float64(srcHeight) * ratio))
}

// Compression ...
func (i *ImagesImpl) Compression(imageFile image.Image, newDx int) (thumbnail image.Image, err error) {
	bound := imageFile.Bounds()
	dx := bound.Max.X
	dy := bound.Max.Y

	// thumbnail size
	w, h := calculateRatioFit(dx, dy, newDx)

	// create thumbnail
	thumbnail = resize.Resize(uint(w), uint(h), imageFile, resize.Lanczos3)

	return
}

// Save ...
func (i *ImagesImpl) Save(imageFile image.Image, path string) (err error) {
	out, err := os.Create(path)
	if err != nil {
		logger.Errorf("Create file failed , path = %s", path)
		return
	}
	defer out.Close()

	// write new image to file
	err = png.Encode(out, imageFile)
	if err != nil {
		logger.Errorf("Images encoding failed , path = %s", path)
		return
	}

	return
}

// Load ...
func (i *ImagesImpl) Load(path string) (img image.Image, err error) {
	file, err := os.Open(path)
	if err != nil {
		logger.Errorf("fail to open the file , path = %s", path)
		return
	}
	defer file.Close()

	// decode png into image.Image
	img, err = png.Decode(file)
	if err != nil {
		logger.Errorf("Images decoding failed , path = %s", path)
		return
	}

	return
}
