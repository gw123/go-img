package go_img

import (
	"image/png"
	"os"
	"testing"
)

func TestQrcodeImage(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Qrcode",
			args:    args{content: "hello world"},
			wantErr: false,
		},
	}

	filePath := "./resources/test/"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotImage, err := QrcodeImage(tt.args.content, 200)
			if (err != nil) != tt.wantErr {
				t.Errorf("QrcodeImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			file, err := os.Create(filePath + tt.name + ".png")
			if err != nil {
				t.Error(err)
				return
			}
			defer file.Close()
			err = png.Encode(file, gotImage)
			if err != nil {
				t.Error(err)
				return
			}

		})
	}
}
