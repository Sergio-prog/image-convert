package convert

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"strings"

	"github.com/sunshineplan/imgconv"
)

type Image struct {
	Image *bytes.Buffer
	Format imgconv.Format
}

func ReadFile(file *multipart.FileHeader) ([]byte, error) {
	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, src)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (i *Image) WriteImageBytes(fileName string) error {
	imageBytes := i.Image.Bytes()
	err := os.WriteFile(fileName, imageBytes, 0666)
	if err != nil {
		return err
	}

	return nil
}

func ImageToBytes(imagePath string) ([]byte, error) {
	f, err := os.ReadFile(imagePath)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func ConvertImage(imageBytes []byte, convertType string) (*Image, error) {
	reader := bytes.NewReader(imageBytes)
	image, err := imgconv.Decode(reader)
	if err != nil {
		return nil, err
	}

	convertType = strings.Replace(strings.TrimSpace(strings.ToLower(convertType)), ".", "", 1)

	format, err := imgconv.FormatFromExtension(convertType)
	if err != nil {
		return nil,  err
	}

	writer := new(bytes.Buffer)
	err = imgconv.Write(writer, image, &imgconv.FormatOption{Format: format})
	if err != nil {
		return nil, err
	}

	resImage := &Image{Image: writer, Format: format}

	return resImage, nil
}
