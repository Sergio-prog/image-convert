package main

import (
	"convert/api"
	// "log"
	// "convert/convert"
	// "log"
)

func main() {
	// image, err := convert.ImageToBytes("tests\\vs.webp")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(image)

	// format := "png"
	// convertedImage, err := convert.ConvertImage(image, format)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = convertedImage.WriteImageBytes("ConvertedImage" + "." + format)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	api.Run()
}

