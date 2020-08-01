package go_img

import (
	"bytes"
	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
)

//生成并写入二维码图片
func QrcodeImage(str string, size int) (image image.Image, err error) {
	pngData := make([]byte, 0)
	if pngData, err = qrcode.Encode(str, qrcode.Highest, size); err != nil {
		return nil, errors.Wrap(err, "qrcode.Encode")
	}
	buffer := bytes.NewBuffer(pngData)
	return png.Decode(buffer)
}
