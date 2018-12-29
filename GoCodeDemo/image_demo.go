package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"image/jpeg"
	_ "image/jpeg"
)

func main() {
	reader, err := os.Open("11.jpg")
	if err != nil {
		log.Fatal("Cannot open image")
	}
	defer reader.Close()
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	bounds := m.Bounds()

	fmt.Println(bounds) // rectangel

	x := m.(interface {
		SubImage(r image.Rectangle) image.Image
	}).SubImage(image.Rect(0, 0, 400, 400))

	fmt.Println(x.Bounds())

	outputFile, err := os.Create("22.jpeg")
	defer outputFile.Close()
	if err != nil {
		log.Fatal("wrong while genrating")
	}

	// Then you can see 22.jepg on your directory
	jpeg.Encode(outputFile, x, nil)

	// startPoint := image.Point{
	// 	X: 1,
	// 	Y: 1,
	// }
	// endPoint := image.Point{
	// 	X: 3,
	// 	Y: 3,
	// }

	// rec := image.Rectangle{
	// 	Min: startPoint,
	// 	Max: endPoint,
	// }

}
