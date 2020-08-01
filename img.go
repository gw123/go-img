package go_img

import (
	"bytes"
	"github.com/pkg/errors"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	ModeVertical   = "vertical"
	ModeHorizontal = "horizontal"
)

func LoadImageFromFile(filePath string) (image.Image, error) {
	imageFile, err := os.Open(filePath)
	if err != nil {
		return nil, errors.Wrap(err, "os.Open")
	}

	img, err := png.Decode(imageFile)
	if err != nil {
		return nil, errors.Wrap(err, "png.Decode")
	}
	return img, nil
}

//
func LoadImageFromUrl(url string) (image.Image, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "http.Get")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("LoadImageFromUrl resp.statusCode != 200")
	}
	//invalid JPEG format: missing SOI marker 解决read 之后读指针后移动问题
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("ReadAll")
	}

	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		if img, err = jpeg.Decode(bytes.NewReader(data)); err != nil {
			return nil, errors.Wrap(err, "jpeg.Decode")
		} else {
			return img, nil
		}
	}
	return img, nil
}

func SaveImageToFile(img image.Image, filePath string) error {
	imageFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	return png.Encode(imageFile, img)
}

/***
  垂直拼接两个图片
*/
func CombineV(upImg image.Image, downImg image.Image) (image.Image, error) {
	upRect := upImg.Bounds()
	downRect := downImg.Bounds()
	outHeight := upRect.Max.Y + downRect.Max.Y
	outWeight := 0
	if upRect.Max.X > downRect.Max.X {
		outWeight = upRect.Max.X
	} else {
		outWeight = downRect.Max.X
	}
	outRect := image.Rect(0, 0, outWeight, outHeight)
	outImg := image.NewRGBA(outRect)
	draw.Draw(outImg, upRect, upImg, image.Pt(0, 0), draw.Src)
	draw.Draw(outImg, downRect.Add(image.Pt(0, upRect.Max.Y)), downImg, image.Pt(0, 0), draw.Src)
	return outImg, nil
}
