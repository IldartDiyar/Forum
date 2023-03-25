package service

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
)

func toBase64(b []byte, format string) string {
	return "data:image/" + format + ";base64," + base64.StdEncoding.EncodeToString(b)
}

func resize(file []byte, width, height int) (string, error) {
	img, format, err := image.Decode(bytes.NewReader(file))
	if err != nil {
		return "", err
	}
	rect := img.Bounds()
	newWidth := rect.Dx() * width / rect.Dx()
	newHeight := rect.Dy() * height / rect.Dy()

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			oldX := x * rect.Dx() / newWidth
			oldY := y * rect.Dy() / newHeight
			newImg.Set(x, y, img.At(oldX, oldY))
		}
	}

	var b bytes.Buffer

	switch format {
	case "jpeg":
		err = jpeg.Encode(&b, newImg, nil)
	case "png":
		err = png.Encode(&b, newImg)
	default:
		err = errors.New("Invalid image format")
	}

	return toBase64(b.Bytes(), format), nil
}
