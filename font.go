package go_img

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/pkg/errors"
	"golang.org/x/image/font"
	"image"
	"io/ioutil"
)

var (
	simFontFile = "resources/fonts/simsun.ttf"
	fontType    *truetype.Font
	fontDpi     float64
	fontSize    float64
	fontHinting string
)

func SetFontType(t *truetype.Font) {
	fontType = t
}

func SetFontDpi(d float64) {
	fontDpi = d
}

func SetFontSize(s float64) {
	fontSize = s
}

func SetFontHinting(h string) {
	fontHinting = h
}

func LoadFont(fontPath string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(simFontFile)
	if err != nil {
		return nil, err
	}
	fontType, err = freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}
	return fontType, nil
}

func DrawString(canvas *image.RGBA, x, y int, text string) error {
	fg := image.Black
	c := freetype.NewContext()
	if fontDpi != 0 {
		c.SetDPI(fontDpi)
	}

	if fontType != nil {
		c.SetFont(fontType)
	}

	if fontSize != 0 {
		c.SetFontSize(fontSize)
	}

	if fontType != nil {
		c.SetFont(fontType)
	}

	c.SetClip(canvas.Bounds())
	c.SetDst(canvas)
	c.SetSrc(fg)

	switch fontHinting {
	default:
		c.SetHinting(font.HintingNone)
	case "full":
		c.SetHinting(font.HintingFull)
	}
	// Draw the text.
	pt := freetype.Pt(x, y)
	_, err := c.DrawString(text, pt)
	if err != nil {
		return errors.Wrap(err, "c.DrawString")
	}
	return nil
}
