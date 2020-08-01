package go_img

import (
	"image"
	"testing"
)

const ImagePath = "./resources/asset/"

func getDefaultImg() (image.Image, error) {
	filePath := "./resources/asset/qr.png"
	return LoadImageFromFile(filePath)
}

func handleErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestCombineV(t *testing.T) {
	img, err := getDefaultImg()
	if err != nil {
		t.Error(err)
		return
	}

	img2, err := LoadImageFromFile(ImagePath + "qr.png")
	handleErr(t, err)
	img3, err := LoadImageFromFile(ImagePath + "qr.png")
	handleErr(t, err)

	img4, err := LoadImageFromUrl("https://static.tukuppt.com/common/image/collect.png")
	handleErr(t, err)
	img5, err := LoadImageFromUrl("https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=862160853,3554709241&fm=26&gp=0.jpg")
	handleErr(t, err)

	type args struct {
		upImg   image.Image
		downImg image.Image
	}
	tests := []struct {
		name    string
		args    args
		want    image.Image
		wantErr bool
	}{
		{
			name: "testCombineV-1",
			args: args{
				upImg:   img,
				downImg: img,
			},
			wantErr: false,
		},
		{
			name: "testCombineV-2",
			args: args{
				upImg:   img3,
				downImg: img2,
			},
			wantErr: false,
		},
		{
			name: "testCombineV-3",
			args: args{
				upImg:   img4,
				downImg: img5,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CombineV(tt.args.upImg, tt.args.downImg)
			if (err != nil) != tt.wantErr {
				t.Errorf("CombineV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			err = SaveImageToFile(got, "./resources/asset/"+tt.name+".png")
			if err != nil {
				t.Errorf("SaveImageToFile Err: %v", err)
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("CombineV() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
