/*
Copyright Lingzhu Ltd. 2020 All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package image

import (
	"errors"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/csiabb/donation-service/common/log"
	"github.com/csiabb/donation-service/common/rest"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/skip2/go-qrcode"
	"golang.org/x/image/math/fixed"
)

var (
	logger = log.MustGetLogger("components/image")
)

// Client image services client
type Client struct {
	ImageConfig *Config
	Bg          *image.NRGBA
	Font        *freetype.Context
	FontType    *truetype.Font
}

// Init initializes a new background image
func (c *Client) Init() error {
	imgFile, err := os.Open(c.ImageConfig.BackgroundPath)
	if err != nil {
		logger.Errorf("failed to read bg path %s: %s", c.ImageConfig.BackgroundPath, err)
		return err
	}
	defer imgFile.Close()

	pngImg, err := png.Decode(imgFile)
	if err != nil {
		logger.Errorf("failed to bg decode : %s", err)
		return err
	}

	bg := image.NewNRGBA(pngImg.Bounds())
	for y := 0; y < bg.Bounds().Dy(); y++ {
		for x := 0; x < bg.Bounds().Dx(); x++ {
			bg.Set(x, y, pngImg.At(x, y))
		}
	}
	c.Bg = bg

	fontBytes, err := ioutil.ReadFile(c.ImageConfig.FontPath)
	if err != nil {
		logger.Errorf("failed to read font path %s: %s", c.ImageConfig.FontPath, err)
		return err
	}

	fontType, err := freetype.ParseFont(fontBytes)
	if err != nil {
		logger.Errorf("failed to parse font : %s", err)
		return err
	}
	c.FontType = fontType

	font := freetype.NewContext()
	font.SetDPI(72)
	font.SetFont(fontType)
	font.SetClip(bg.Bounds())
	font.SetDst(bg)

	c.Font = font

	return nil
}

// CreateQrCode  create a new qr code
func (c *Client) CreateQrCode(content string) (img image.Image, err error) {
	var qrCode *qrcode.QRCode

	qrCode, err = qrcode.New(content, qrcode.Highest)

	if err != nil {
		return nil, errors.New("qr code creation failed")
	}

	qrCode.DisableBorder = true
	qrCode.BackgroundColor = rest.Color3
	qrCode.ForegroundColor = rest.Color2

	img = qrCode.Image(rest.QrCodeSize)

	return img, nil
}

// DrawText define string drawing
func (c *Client) DrawText(fontColor color.Color, str string, pt fixed.Point26_6, size float64) error {
	c.Font.SetFontSize(size)
	c.Font.SetSrc(image.NewUniform(fontColor))
	_, err := c.Font.DrawString(str, pt)
	return err
}

// SlipString handles line breaks of strings
func (c *Client) SlipString(content string, fontSize float64, textWidth int) []string {
	runes := []rune(content)
	opts := truetype.Options{
		Size: fontSize,
	}
	face := truetype.NewFace(c.FontType, &opts)

	var text string
	var lines []string
	var length fixed.Int26_6
	for j := 0; j < len(runes); j++ {
		faceWidth, _ := face.GlyphAdvance(runes[j])
		length += faceWidth
		if length.Ceil() > textWidth {
			lines = append(lines, text)
			text = ""
			length = 0
		} else {
			text += string(runes[j])
		}
	}

	lines = append(lines, text)
	return lines
}

// CreateDonationImage create new image of donation items
func (c *Client) CreateDonationImage(content []string, isShare bool) (*image.NRGBA, error) {
	// init bg image
	err := c.Init()
	if err != nil {
		return nil, err
	}

	// create donation image
	var index = 0
	for i := 0; i < len(rest.Title); i++ {
		err = c.DrawText(rest.Color1, rest.Title[i], freetype.Pt(c.Bg.Bounds().Dx()-rest.SubWidth,
			c.Bg.Bounds().Dy()-(rest.SubHeight-index*48)), 28)
		lines := c.SlipString(content[i], 28, 375)
		for k := 0; k < len(lines); k++ {
			err = c.DrawText(rest.Color2, lines[k], freetype.Pt(c.Bg.Bounds().Dx()-rest.SubWidth+len(rest.Title[i])/3*28,
				c.Bg.Bounds().Dy()-(rest.SubHeight-index*48)), 28)
			index++
		}
	}

	// create qr image
	if isShare {
		var qrCodeImg image.Image
		qrCodeImg, err = c.CreateQrCode("http://www.baidu.com")
		draw.Draw(c.Bg, qrCodeImg.Bounds().Add(image.Pt(100, 710)), qrCodeImg, image.Point{X: 0, Y: 0}, draw.Over)
		err = c.DrawText(rest.Color2, rest.QRContent, freetype.Pt(235, 828), 22)
	}

	return c.Bg, err
}
