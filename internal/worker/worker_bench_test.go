package worker

import (
	"image"
	"log"
	"sobel/internal/img"
	"strconv"
	"testing"
)

// Helper function to load a test image from a file
func loadTestImage(imagePath string) (*img.ImageData, error) {
	srcImage, err := img.Read(imagePath)
	if err != nil {
		return nil, err
	}

	bounds := srcImage.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

	return &img.ImageData{
		SrcImage: srcImage,
		DstImage: dstImage,
		Width:    width,
		Height:   height,
	}, nil
}

// Benchmark for the worker pool with a real photo
func BenchmarkRunWorkerPool(b *testing.B) {
	// Define the path to the test image
	imagePath := "../../images/image.jpg" // Update the path to your test image

	// Load the test image
	imgData, err := loadTestImage(imagePath)
	if err != nil {
		log.Fatalf("Failed to load test image: %v", err)
	}

	// List of worker counts to benchmark
	workerCounts := []int{1, 2, 4, 8, 10, 16, 32, 64, 128, 256, 512, 1024}

	for _, numWorkers := range workerCounts {
		b.Run("Workers_"+strconv.Itoa(numWorkers), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Create a fresh destination image for each iteration
				b.StopTimer()
				imgData.DstImage = image.NewRGBA(image.Rect(0, 0, imgData.Width, imgData.Height))
				b.StartTimer()

				// Run the worker pool
				RunWorkerPool(numWorkers, imgData)
			}
		})
	}
}
