package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"path"
	"sobel/internal/img"
	"sobel/internal/worker"
)

var (
	imagePath  string
	numWorkers int
)

func main() {
	flag.StringVar(&imagePath, "image", "", "The path to the image")
	flag.IntVar(&numWorkers, "numWorkers", 0, "Count of workers (goroutines)")
	flag.Parse()

	if imagePath == "" {
		log.Fatalln("the --image value is not set")
	} else if numWorkers == 0 {
		log.Fatalln("the --numWorkers value is not set")
	}

	srcImage, err := img.Read(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	bounds := srcImage.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

	outputDir, outputFile := path.Split(imagePath)
	outputFile = fmt.Sprintf("sobel-%s", outputFile)
	outputPath := path.Join(outputDir, outputFile)

	imgData := img.ImageData{
		SrcImage: srcImage,
		DstImage: dstImage,
		Width:    width,
		Height:   height,
	}

	worker.RunWorkerPool(numWorkers, &imgData)

	fmt.Println(outputPath)
	err = img.Write(imgData.DstImage, outputPath)
	if err != nil {
		log.Fatal(err)
	}

}
