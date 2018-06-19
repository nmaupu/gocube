package image

import (
	"github.com/fogleman/gg"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Trim image from white color
func TrimImageWhite(ctx *gg.Context) (*gg.Context, error) {
	img := ctx.Image()
	xs := make([]int, 0)
	ys := make([]int, 0)

	b := img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			px := img.At(x, y)
			r, g, b, _ := px.RGBA()
			if r != 65535 && g != 65535 && b != 65535 {
				xs = append(xs, x)
				ys = append(ys, y)
			}
		}
	}

	xL, _ := getMin(xs, ys)
	xR, _ := getMax(xs, ys)
	yU, _ := getMin(ys, xs)
	yD, _ := getMax(ys, xs)

	tmpDir, err := ioutil.TempDir("", "gocube")
	defer os.RemoveAll(tmpDir) // clean up
	if err != nil {
		return nil, err
	}
	tmpFile := filepath.Join(tmpDir, "out.png")
	fso, err := os.Create(tmpFile)
	if err != nil {
		return nil, err
	}
	defer fso.Close()

	croppedimg := img.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(xL, yU, xR, yD))

	png.Encode(fso, croppedimg)
	retImg, err := gg.LoadPNG(tmpFile)
	if err != nil {
		return nil, err
	}

	return gg.NewContextForImage(retImg), nil
}

func getMin(xs, ys []int) (int, int) {
	min := 100000000
	y := 0
	for k, v := range xs {
		if v < min {
			min = v
			y = ys[k]
		}
	}

	return min, y
}
func getMax(xs, ys []int) (int, int) {
	max := 0
	y := 0
	for k, v := range xs {
		if v > max {
			max = v
			y = ys[k]
		}
	}

	return max, y
}
