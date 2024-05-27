package worker

import (
	"sobel/internal/img"
	"sobel/internal/sobel"
	"sync"
)

type SobelTask struct {
	StartRow int
	EndRow   int
}

func worker(tasks chan SobelTask, results chan struct{}, imgData *img.ImageData) {
	for task := range tasks {
		sobel.ApplySobel(imgData, task.StartRow, task.EndRow)
		results <- struct{}{}
	}
}

func RunWorkerPool(numWorkers int, imgData *img.ImageData) {
	tasks := make(chan SobelTask, numWorkers)
	results := make(chan struct{}, numWorkers)

	wg := sync.WaitGroup{}

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			worker(tasks, results, imgData)
			wg.Done()
		}()
	}

	rowsPerWorker := imgData.Height / numWorkers
	for i := 0; i < numWorkers; i++ {
		startRow := i * rowsPerWorker
		endRow := startRow + rowsPerWorker
		if i == numWorkers-1 {
			endRow = imgData.Height
		}
		tasks <- SobelTask{StartRow: startRow, EndRow: endRow}
	}
	close(tasks)

	for i := 0; i < numWorkers; i++ {
		<-results
	}
	close(results)

	wg.Wait()
}
